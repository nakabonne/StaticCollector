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

// Fetch フェッチする
func Fetch(u string, depth int) (urls []string, html string, err error) {
	baseUrl, err := url.Parse(u)
	if err != nil {
		return
	}
	// html取得------------------------------
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	if depth == 1 {
		html, err = getStaticFile(*resp)
	}
	defer resp.Body.Close()

	// スクレイピング------------------------------------
	resp2, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	defer resp2.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp2.Body)
	if err != nil {
		return
	}
	title, err := getTitle(doc)
	//rank := 1 // 仮
	//fmt.Println("titleは", title)
	page := &models.Pages{
		ID:        bson.NewObjectId(),
		Title:     title,
		URL:       u,
		HTML:      html,
		Rank:      1,
		TargetDay: time.Now(),
	}
	fmt.Println("構造体は", page)

	// 1回目のみなので関数分ける-----------------------------------------------------------------
	urls = make([]string, 0)
	doc.Find(".r").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqUrl, err := baseUrl.Parse(href)
				if err == nil {
					urls = append(urls, reqUrl.String())
				}
			}
		})
	})
	// -------------------------------------------------------------------------

	return
}

func getStaticFile(res http.Response) (file string, err error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(body)
	file = buf.String()
	return
}

func getTitle(doc *goquery.Document) (title string, err error) {
	doc.Find("title").Each(func(_ int, srg *goquery.Selection) {
		title = srg.Text()
	})

	return
}
