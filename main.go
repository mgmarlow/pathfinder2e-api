package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var Db *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PASS"))

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func createMonster(name string) (*Monster, error) {
	details := GetMonsterDetails(name)
	monster, err := NewMonster(name, details)

	if err != nil {
		fmt.Println("error creating monster ", name)
		log.Println(err)
		return nil, err
	}

	return monster, nil
}

func main() {
	monsterNames := GetMonsterNames()
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
