package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"webCrawler/models"

	"gopkg.in/mgo.v2/bson"

	"github.com/PuerkitoBio/goquery"
)

type crawler struct {
	res  chan *respons
	req  chan *request
	quit chan int
}
type respons struct {
	url  string
	html string
	err  interface{}
}
type request struct {
	url   string
	depth int
}

func newCrawler() *crawler {
	return &crawler{
		res:  make(chan *respons),
		req:  make(chan *request),
		quit: make(chan int),
	}
}

func (c *crawler) execute() {
	wc := 0 // ワーカーの数
	urlMap := make(map[string]bool, 100)
	done := false
	for !done {
		select {
		case res := <-c.res:
			if res.err == nil {
				//fmt.Printf("%s\n", res.url)
				//fmt.Println("htmlは")
				//fmt.Printf("%s\n", res.html)
			} else {
				fmt.Fprintf(os.Stderr, "Error %s\n%v\n", res.url, res.err)
			}
		case req := <-c.req:
			if req.depth == 0 {
				break
			}

			if urlMap[req.url] {
				// 取得済み
				break
			}
			urlMap[req.url] = true

			wc++
			go Crawl(req.url, req.depth, c)
		case <-c.quit:
			wc--
			if wc == 0 {
				done = true
			}
		}
	}
	log.Println("スクレイピング完了")
}

// Crawl クロールする
func Crawl(url string, depth int, c *crawler) {
	defer func() { c.quit <- 0 }()

	// WebページからURLを取得
	urls, html, err := Fetch(url, depth)

	// 結果送信
	c.res <- &respons{
		url:  url,
		html: html,
		err:  err,
	}

	if err == nil {
		for _, url := range urls {
			// 新しいリクエスト送信
			c.req <- &request{
				url:   url,
				depth: depth - 1,
			}
		}
	}
}

// Fetch ②pages構造体を返す関数にして、1回目は別の関数を実行するようにする
func Fetch(u string, depth int) (urls []string, html string, err error) {
	baseURL, err := url.Parse(u)
	if err != nil {
		return
	}
	// 検索結果ページの場合のみhtml取得
	if depth == 1 {
		html, err = getHTML(baseURL.String())
		if err != nil {
			return
		}
	}

	// スクレイピング------------------------------------
	res, err := http.Get(baseURL.String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	title, err := getTitle(doc)
	page := &models.Pages{
		ID:        bson.NewObjectId(),
		Title:     title,
		URL:       u,
		HTML:      html,
		Rank:      1,
		TargetDay: time.Now(),
	}
	fmt.Println("構造体は", page)

	// 次ページのurlも取得したい場合
	if depth != 1 {
		urls = make([]string, 0)
		doc.Find(".r").Each(func(_ int, srg *goquery.Selection) {
			srg.Find("a").Each(func(_ int, s *goquery.Selection) {
				href, exists := s.Attr("href")
				if exists {
					reqURL, err := baseURL.Parse(href)
					if err == nil {
						urls = append(urls, reqURL.String())
					}
				}
			})
		})
	}

	return
}

func getHTML(url string) (html string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(body)
	html = buf.String()
	return
}

func getTitle(doc *goquery.Document) (title string, err error) {
	doc.Find("title").Each(func(_ int, srg *goquery.Selection) {
		title = srg.Text()
	})

	return
}
