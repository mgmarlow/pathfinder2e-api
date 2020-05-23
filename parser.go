package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// TODO: Don't kill program on failure, log and move on instead
func scrape(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func GetMonsterDetails(name string) map[string]string {
	monsterDetailLink := "https://www.aonprd.com/MonsterDisplay.aspx?ItemName=" + url.QueryEscape(name)
	doc := scrape(monsterDetailLink)

	m := make(map[string]string)

	var keys []string
	var values []string

	doc.Find("table span").Each(func(i int, s *goquery.Selection) {
		rawHTML, err := s.Html()
		if err != nil {
			fmt.Println("Error retrieving HTML for: " + monsterDetailLink)
			return
		}

		doc, err := html.Parse(strings.NewReader(rawHTML))
		if err != nil {
			fmt.Println("Error parsing HTML for: " + monsterDetailLink)
			return
		}

		// Key-value pairs look like the following, making them difficult to parse.
		// Note that values are not nested in any containing DOM node.
		//
		// <b>title</b>
		// value
		// <br>
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.TextNode {
				if n.Parent.Data == "b" {
					keys = append(keys, n.Data)
				}

				// TODO: need a way of collecting node values
				// if n.Parent.Data == "a" { }

				if n.Parent.Data == "body" && n.Data[0] == ' ' {
					values = append(values, n.Data[1:])
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}

		f(doc)

	})

	for i, key := range keys {
		m[key] = values[i]
	}

	return m
}

func GetMonsterNames() []string {
	doc := scrape("https://www.aonprd.com/Monsters.aspx?Letter=All")

	return doc.Find("#main table td a").Map(func(i int, s *goquery.Selection) string {
		link, exists := s.Attr("href")

		if exists {
			return strings.Split(link, "ItemName=")[1]
		}

		return ""
	})
}
