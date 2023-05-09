package htmlutil

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type Article struct {
	Title       string
	Desc        string
	URL         string
	Content     string
	Author      string
	PublishTime string // not work.
}

var wxHTTPClient *http.Client

func init() {
	wxHTTPClient = newHTTPClient()
}

func Parse(url string) (*Article, error) {

	res, err := wxHTTPClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to request %s", url))
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request url %s, response status not 200.", url))
	}

	// 加载 HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "load html body")
	}

	art := &Article{URL: url}
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		var isExist bool
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:title") {
			art.Title, isExist = s.Attr("content")
			if !isExist {
				logx.Errorf("url %s doesn't contain content mark.")
			}
		}
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:description") {
			art.Desc, isExist = s.Attr("content")
			if !isExist {
				logx.Errorf("url %s doesn't contain description mark.")
			}
		}
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:article:author") {
			art.Author, isExist = s.Attr("content")
			if !isExist {
				logx.Errorf("url %s doesn't contain author mark.")
			}
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

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 3 * time.Second, // 发送完request后等待服务端响应时间
			IdleConnTimeout:       60 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second, // 连接建立超时
				KeepAlive: 5 * time.Second, // 客户端发送心跳间隔
			}).DialContext,
			DisableKeepAlives: false,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // self-signed certificate
			},
		},
		Timeout: 10 * time.Second, // 请求整体完成超时时间（截止到内核接收完响应体，不含用户程序读取Body的延时）
	}
}
