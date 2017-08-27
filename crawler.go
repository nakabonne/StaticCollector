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
	page *models.Pages
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

func (c *crawler) collectHTML() {
	wc := 0 // ワーカーの数
	urlMap := make(map[string]bool, 100)
	done := false
	for !done {
		select {

		case res := <-c.res:
			if res.err == nil {
				fmt.Println("構造体は", res.page)
			} else {
				fmt.Fprintf(os.Stderr, "Error %s\n%v\n", res.url, res.err)
			}

		case req := <-c.req:
			if urlMap[req.url] {
				// 取得済み
				break
			}
			urlMap[req.url] = true
			wc++
			baseURL, err := url.Parse(req.url)
			if err != nil {
				log.Fatal("エラー", err)
			}
			switch req.depth {
			case 0:
				break
			case 1:
				go getPage(baseURL, c)
			default:
				go createRequest(baseURL, req.depth, c)
			}

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
func createRequest(u *url.URL, depth int, c *crawler) {
	defer func() { c.quit <- 0 }()

	urls := make([]string, 0)
	doc, err := getDoc(u)
	doc.Find(".r").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqURL, err := u.Parse(href)
				if err == nil {
					urls = append(urls, reqURL.String())
				}
			}
		})
	})

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

func getPage(u *url.URL, c *crawler) {
	defer func() { c.quit <- 0 }()

	doc, err := getDoc(u)

	var html string
	var title string
	html, err = getHTML(u.String())
	if err != nil {
		return
	}
	title, err = getTitle(doc)
	if err != nil {
		return
	}
	url := u.String()
	page := &models.Pages{
		ID:        bson.NewObjectId(),
		Title:     title,
		URL:       url,
		HTML:      html,
		Rank:      1,
		TargetDay: time.Now(),
	}
	// 結果送信
	c.res <- &respons{
		page: page,
		err:  err,
	}
}

func getDoc(u *url.URL) (doc *goquery.Document, err error) {
	if err != nil {
		return
	}

	res, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
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
