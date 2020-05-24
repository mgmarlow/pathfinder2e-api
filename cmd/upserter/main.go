package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mgmarlow/pathfinder2e-api/pkg/api"
	"github.com/mgmarlow/pathfinder2e-api/pkg/scraper"
)

func createMonster(name string) (*api.Monster, error) {
	details := scraper.GetMonsterDetails(name)
	monster, err := api.NewMonster(name, details)

	if err != nil {
		fmt.Println("error creating monster ", name)
		log.Println(err)
		return nil, err
	}

	return monster, nil
}

func main() {
	monsterNames := scraper.GetMonsterNames()
	// TODO: goroutine this
	fmt.Println("Scraping ", monsterNames[0])
	monster, err := createMonster(monsterNames[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Upserting ", monsterNames[0])
	err = monster.Create()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All done.")

	// for _, link := range monsterLinks {
	// 	fmt.Println("Following ", link)
	// 	monsterDetails := GetMonsterDetails(link)
	// }
}
