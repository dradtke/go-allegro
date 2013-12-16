// This is a utility program used to pull down HTML Allegro documentation, parse it, and insert it
// into the Go source code.
package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const ROOT_URL string = "http://alleg.sourceforge.net/a5docs/"
var VERSION string

func Usage() {
	fmt.Println("usage: documenter <allegro-version>")
	os.Exit(0)
}

func Get(suburl string) *html.Node {
	url := ROOT_URL + VERSION + "/" + suburl
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(fmt.Errorf("received status code %d when fetching '%s'", resp.StatusCode, url))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	node, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	return node
}

func Attr(node *html.Node, name string) string {
	if node == nil || node.Attr == nil {
		return ""
	}
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func HasClass(node *html.Node, class string) bool {
	classes := strings.Split(Attr(node, "class"), " ")
	for _, c := range classes {
		if c == class {
			return true
		}
	}
	return false
}

func Filter(node *html.Node, cond func(*html.Node) bool) <-chan *html.Node {
	ch := make(chan *html.Node)
	go func() {
		var f func(*html.Node)
		f = func(child *html.Node) {
			for child != nil {
				if cond(child) {
					ch <- child
				}
				if child.FirstChild != nil {
					f(child.FirstChild)
				}
				child = child.NextSibling
			}
		}
		f(node.FirstChild)
		close(ch)
	}()
	return ch
}

func IsContentDiv(node *html.Node) bool {
	return node.Data == "div" && HasClass(node, "content")
}

func GetListContents(node *html.Node) <-chan *html.Node {
	if node.Data != "ul" {
		log.Fatalf("'%s' is not a <ul> node", node.Data)
	}
	ch := make(chan *html.Node)
	go func() {
		li := node.FirstChild
		for li != nil {
			if li.FirstChild != nil {
				ch <- li.FirstChild
			}
			li = li.NextSibling
		}
		close(ch)
	}()
	return ch
}

func GetText(node *html.Node) string {
	var buf bytes.Buffer
	node = node.FirstChild
	for node != nil {
		if node.Type == html.TextNode {
			buf.WriteString(node.Data)
		} else {
			buf.WriteString(GetText(node))
		}
		node = node.NextSibling
	}
	return buf.String()
}

func Parse(href string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	page := Get(href)
	content :=  <-Filter(page, func(child *html.Node) bool {
		return child.Data == "div" && HasClass(child, "content")
	})
	toc := <-Filter(content, func(child *html.Node) bool {
		return child.Data == "div" && Attr(child, "id") == "TOC"
	})
	if toc == nil {
		// the main addon has no table of contents, so just return
		return
	}
	ch := make(chan string)
	go func() {
		for item := range GetListContents(toc.FirstChild.NextSibling) {
			a := Attr(item, "href")
			if a != "" && a[0] == '#' {
				ch <- a[1:]
			}
		}
		close(ch)
	}()
	for id := range ch {
		node := <-Filter(content, func(child *html.Node) bool {
			return child.Data == "h1" && Attr(child, "id") == id
		})
		fmt.Print(id + ": ")
		for node.Data != "p" {
			node = node.NextSibling
		}
		fmt.Println(GetText(node))
		// TODO: find the implementation of the Allegro function denoted by id,
		// go to the line right above the function declaration, and insert the
		// value of GetText(node)
	}
}

func main() {
	if len(os.Args) == 1 {
		Usage()
	}
	VERSION = os.Args[1]
	index := Get("index.html")
	content := <-Filter(index, IsContentDiv)
	apiContent := <-Filter(content, func(child *html.Node) bool {
		return child.Data == "h1" && Attr(child, "id") == "api"
	})
	addonContent := <-Filter(content, func(child *html.Node) bool {
		return child.Data == "h1" && Attr(child, "id") == "addons"
	})
	var wg sync.WaitGroup
	for link := range GetListContents(apiContent.NextSibling.NextSibling) {
		href := Attr(link, "href")
		wg.Add(1)
		go Parse(href, &wg)
	}
	for link := range GetListContents(addonContent.NextSibling.NextSibling) {
		href := Attr(link, "href")
		wg.Add(1)
		go Parse(href, &wg)
	}
	wg.Wait()
}
