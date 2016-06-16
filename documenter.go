package main

// This is a utility program used to pull down HTML Allegro documentation, parse it, and insert it
// into the Go source code.

import (
	"bytes"
	"golang.org/x/net/html"
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const ROOT_URL string = "http://alleg.sourceforge.net/a5docs/"

var VERSION string

var verbose bool
var locate string

// Usage() prints out the program's usage.
func Usage() {
	fmt.Println("usage: documenter [-v] <allegro-version>")
	os.Exit(0)
}

// Get() takes a sub-URL, fetches the appropriate page from the Allegro docs,
// parses it, and returns the result as an HTML node.
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

// Attr() returns the value of the HTML node's attribute with the given name.
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

// HasClass() checks of the HTML node's "class" attribute contains a
// certain class value.
func HasClass(node *html.Node, class string) bool {
	classes := strings.Split(Attr(node, "class"), " ")
	for _, c := range classes {
		if c == class {
			return true
		}
	}
	return false
}

// Find() takes an HTML node and a filter function and returns that node's
// first child for which the filter returns true.
func Find(node *html.Node, cond func(*html.Node) bool) *html.Node {
	var f func(*html.Node) *html.Node
	f = func(child *html.Node) *html.Node {
		for child != nil {
			if cond(child) {
				return child
			}
			if child.FirstChild != nil {
				if result := f(child.FirstChild); result != nil {
					return result
				}
			}
			child = child.NextSibling
		}
		return nil
	}
	return f(node.FirstChild)
}

// Filter function for the content div.
func IsContentDiv(node *html.Node) bool {
	return node.Data == "div" && HasClass(node, "content")
}

// Filter function for the table of contents div.
func IsTOCDiv(node *html.Node) bool {
	return node.Data == "div" && Attr(node, "id") == "TOC"
}

// GetListContents() takes a <ul> HTML node and returns a channel of nodes that
// represent all of that list's values.
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

// GetText() takes an HTML node and returns the text of it and all of its
// subchildren.
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

// GetSource() walks the allegro packages and returns a map from file path
// to linked list of lines in that source file.
func GetSource() map[string]*list.List {
	m := make(map[string]*list.List)
	err := filepath.Walk("allegro/", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			data, err2 := ioutil.ReadFile(path)
			if err2 != nil {
				return err2
			}
			l := list.New()
			for _, line := range strings.Split(string(data), "\n") {
				l.PushBack(line)
			}
			m[path] = l
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return m
}

// TrimTo() trims a string down to a maximum length, breaking off at the
// last space to occur before the limit.
func TrimTo(str string, length int) (string, string) {
	if len(str) <= length {
		return str, ""
	}
	i := strings.LastIndex(str[:length], " ")
	return str[:i], str[i+1:]
}

// Parse() fetches a page from a table of contents link, pulls out all
// function documentation, and updates the source file implementing that
// function with it. This doesn't write anything to disk, just modifies it
// in memory to be written out later.
func Parse(href string, sources map[string]*list.List, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	page := Get(href)
	content := Find(page, IsContentDiv)
	toc := Find(content, IsTOCDiv)
	if toc == nil {
		return
	}
	ch := make(chan string)
	go func() {
		var f func(*html.Node)
		f = func(node *html.Node) {
			for item := range GetListContents(node) {
				a := Attr(item, "href")
				if a != "" && a[0] == '#' {
					ch <- a[1:]
				}
				next := item.NextSibling
				if next != nil && next.Type != html.TextNode && next.Data == "ul" {
					f(item.NextSibling)
				}
			}
		}
		f(toc.FirstChild.NextSibling)
		close(ch)
	}()
	for id := range ch {
		if id == locate {
			fmt.Println("looking for " + locate)
		}
		cgo := "C." + id + "("
		node := Find(content, func(child *html.Node) bool {
			return Attr(child, "id") == id
		})
		for node.Data != "p" {
			node = node.NextSibling
		}
		found := false
		for _, lines := range sources {
			var e *list.Element
			for e = lines.Front(); e != nil; e = e.Next() {
				if bytes.Contains([]byte(e.Value.(string)), []byte(cgo)) {
					break
				}
			}
			if e == nil {
				continue
			}
			found = true
			for !strings.HasPrefix(e.Value.(string), "func ") {
				e = e.Prev()
			}
			for strings.HasPrefix(e.Prev().Value.(string), "//") {
				lines.Remove(e.Prev())
			}
			text := GetText(node)
			if id == locate {
				fmt.Printf("%s: %s\n", locate, text)
				fmt.Printf("implemented at: %s\n", e.Value.(string))
			}
			for {
				commentLine, rest := TrimTo(text, 77)
				lines.InsertBefore("// " + commentLine, e)
				if rest == "" {
					break
				} else {
					text = rest
				}
			}
		}
		if id == locate && !found {
			fmt.Printf("'%s' not found in package sources\n", locate)
		}
	}
}

func main() {
	flag.BoolVar(&verbose, "v", false, "-v")
	flag.StringVar(&locate, "debug", "", "-debug <function>")
	flag.Parse()
	if flag.NArg() == 0 {
		Usage()
	}
	VERSION = flag.Arg(0)
	index := Get("index.html")
	content := Find(index, IsContentDiv)
	apiContent := Find(content, func(child *html.Node) bool {
		return child.Data == "h1" && Attr(child, "id") == "api"
	})
	addonContent := Find(content, func(child *html.Node) bool {
		return child.Data == "h1" && Attr(child, "id") == "addons"
	})
	sources := GetSource()
	var wg sync.WaitGroup
	for link := range GetListContents(apiContent.NextSibling.NextSibling) {
		href := Attr(link, "href")
		wg.Add(1)
		go Parse(href, sources, &wg)
	}
	for link := range GetListContents(addonContent.NextSibling.NextSibling) {
		href := Attr(link, "href")
		wg.Add(1)
		go Parse(href, sources, &wg)
	}
	wg.Wait()
	var buf bytes.Buffer
	for path, lines := range sources {
		for e := lines.Front(); e != nil; e = e.Next() {
			buf.WriteString(e.Value.(string) + "\n")
		}
		err := ioutil.WriteFile(path, buf.Bytes(), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		buf.Reset()
		if verbose {
			fmt.Println("wrote to " + path)
		}
	}
}
