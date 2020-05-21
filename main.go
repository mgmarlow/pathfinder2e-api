package main

import (
	"fmt"
)

func main() {
	monsterNames := GetMonsterNames()
	fmt.Println("Following ", monsterNames[0])
	GetMonsterDetails(monsterNames[0])

	// for _, link := range monsterLinks {
	// 	fmt.Println("Following ", link)
	// 	monsterDetails := GetMonsterDetails(link)
	// }
}
