package main

import (
	"fmt"
)

func main() {
	monsterLinks := GetMonsterLinks()
	for _, link := range monsterLinks {
		fmt.Println("Following ", link)
		// monsterDetails := getMonsterDetails("https://www.aonprd.com/" + link)
	}
}
