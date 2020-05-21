package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

func GetMonsterDetails(name string) Monster {
	monsterDetailLink := "https://www.aonprd.com/MonsterDisplay.aspx?ItemName=" + url.QueryEscape(name)
	doc := scrape(monsterDetailLink)

	doc.Find("table span").Each(func(i int, s *goquery.Selection) {
		// TODO:
		// s.Find("b").Map(func(i int, s *goquery.Selection) string {
		// 	//
		// })
		// fmt.Println(s.Text())
	})

	return Monster{}
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
