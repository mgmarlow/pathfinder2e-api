package main

import (
	"fmt"
	"log"
	"net/http"

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

type Monster struct{}

func getMonsterDetails(url string) Monster {
	return Monster{}
}

func getMonsterLinks() []string {
	doc := scrape("https://www.aonprd.com/Monsters.aspx?Letter=All")

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
		fmt.Println("Following ", link)
		// monsterDetails := getMonsterDetails("https://www.aonprd.com/" + link)
	}
}
