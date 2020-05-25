package api

import (
	"log"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

// Monster is an entity in the Pathfinder 2e Bestiary.
type Monster struct {
	ID           int
	Name         string
	XP           int
	SR           int
	Initiative   string
	Senses       string
	AC           string
	HP           string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Bab          string
	CMB          string
	CMD          int
}

// NewMonster creates a monster from a list of key-value pairs.
// These pairs should come from GetMonsterDetails.
func NewMonster(name string, pairs map[string]string) (*Monster, error) {
	var err error
	monster := new(Monster)

	monster.Name = name
	monster.Initiative = pairs["Init"]
	monster.Senses = pairs["Senses"]
	monster.HP = pairs["hp"]
	monster.AC = pairs["AC"]
	monster.Bab = pairs["Base Atk"]
	monster.CMB = pairs["CMB"]

	monster.XP, err = strconv.Atoi(strings.Replace(pairs["XP"], ",", "", -1))
	monster.CMD, err = strconv.Atoi(pairs["CMD"])
	monster.SR, err = strconv.Atoi(pairs["SR"])
	monster.Strength, err = strconv.Atoi(pairs["Str"])
	monster.Dexterity, err = strconv.Atoi(pairs["Dex"])
	monster.Constitution, err = strconv.Atoi(pairs["Con"])
	monster.Intelligence, err = strconv.Atoi(pairs["Int"])
	monster.Wisdom, err = strconv.Atoi(pairs["Wis"])
	monster.Charisma, err = strconv.Atoi(pairs["Cha"])

	if err != nil {
		return nil, err
	}

	return monster, nil
}

// CopyMonsters adds monsters to the database.
func CopyMonsters(monsters []*Monster) error {
	txn, err := Db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("beasts", "name", "xp", "initiative", "ac", "hp"))

	if err != nil {
		return err
	}

	for _, monster := range monsters {
		_, err = stmt.Exec(
			monster.Name,
			monster.XP,
			monster.Initiative,
			monster.AC,
			monster.HP)

		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}
