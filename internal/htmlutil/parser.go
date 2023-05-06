package htmlutil

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	Title       string
	Desc        string
	URL         string
	Content     string
	Author      string
	PublishTime string // not work.
}

func Parse(url string) (*Article, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed to request %s, err: %v\n", url, err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("status code error: %d %s, %s", res.StatusCode, res.Status, url)
		return nil, errors.New("http response status error")
	}

	// 加载 HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("failed to load html body, err: %s, url: %s", err, url)
		return nil, err
	}

	art := &Article{URL: url}
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:title") {
			art.Title, _ = s.Attr("content")
		}
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:description") {
			art.Desc, _ = s.Attr("content")
		}
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:article:author") {
			art.Author, _ = s.Attr("content")
		}
	})

	// not work.
	/*doc.Find("#publish_time").Each(func(_ int, s *goquery.Selection) {
		art.PublishTime = s.Text()
	})*/
	doc.Find("#js_content").Each(func(_ int, s *goquery.Selection) {
		art.Content = trimEmpty(s.Text())
	})

	return art, nil
}

func trimEmpty(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.Join(strings.Fields(s), "")
}

func (a Article) String() string {
	return fmt.Sprintf(
		"Title: %s\n"+
			"Desc: %s\n"+
			"URL: %s\n"+
			"Author: %s\n"+
			"Content: %s\n",
		a.Title, a.Desc, a.URL, a.Author, a.Content)
}
