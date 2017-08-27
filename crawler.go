package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type crawler struct {
	res  chan *respons
	req  chan *request
	quit chan int
}
type respons struct {
	url string
	err interface{}
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
				fmt.Printf("%s\n", res.url)
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
	fmt.Println(urlMap)
}

func Crawl(url string, depth int, c *crawler) {
	defer func() { c.quit <- 0 }()

	// WebページからURLを取得
	urls, err := Fetch(url)

	// 結果送信
	c.res <- &respons{
		url: url,
		err: err,
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

func Fetch(u string) (urls []string, err error) {
	baseUrl, err := url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 取得したhtmlを文字列で確認したい時はこれ
	//body, err := ioutil.ReadAll(resp.Body)
	//buf := bytes.NewBuffer(body)
	//html := buf.String()
	//fmt.Println(html)
	// ---------------

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

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

	return
}
