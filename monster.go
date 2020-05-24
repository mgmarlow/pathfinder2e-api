package main

import (
	"strconv"
	"strings"
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

// Create adds a monster to the database.
func (m *Monster) Create() error {
	statement := "insert into beasts (name, xp) values ($1, $2) returning id;"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(&m.Name, &m.XP).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}
