package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getMonsterLinks() []string {
	url := "https://www.aonprd.com/Monsters.aspx?Letter=All"

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

	return doc.Find("#main table td a").Map(func(i int, s *goquery.Selection) string {
		link, exists := s.Attr("href")

		if exists {
			return link
		}

		return ""
	})
}

func main() {
	monsterLinks := getMonsterLinks()
	for _, link := range monsterLinks {
		fmt.Println(link)
	}
}
