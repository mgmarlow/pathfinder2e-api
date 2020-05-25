package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mgmarlow/pathfinder2e-api/pkg/api"
	"github.com/mgmarlow/pathfinder2e-api/pkg/scraper"
)

func getMonster(name string) (*api.Monster, error) {
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
	monsters := make([]*api.Monster, len(monsterNames))

	for i, name := range monsterNames {
		fmt.Println("Scraping", name, "...")

		monster, err := getMonster(name)
		if err != nil {
			fmt.Println("Could not create", name)
			continue
		}

		monsters[i] = monster
	}

	fmt.Println("Done scraping.")

	fmt.Println("Upserting...")

	err := api.CopyMonsters(monsters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All done.")
}
