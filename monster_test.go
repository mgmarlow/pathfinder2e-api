package main

import (
	"testing"
)

func TestNewMonster(t *testing.T) {
	name := "Foo Monster"
	pairs := map[string]string{
		"XP":       "4,800",
		"Init":     "+5",
		"Senses":   "foo bar baz",
		"hp":       "103 (9d12+45)",
		"AC":       "20, touch 10",
		"Base Atk": "+9",
		"CMB":      "+15",
		"CMD":      "20",
		"SR":       "19",
		"Str":      "12",
		"Dex":      "13",
		"Con":      "14",
		"Int":      "15",
		"Wis":      "16",
		"Cha":      "17",
	}

	monster, _ := NewMonster(name, pairs)

	test := func(t *testing.T, property string, got interface{}, wanted interface{}) {
		t.Helper()
		expectation := "it should have " + property

		t.Run(expectation, func(*testing.T) {
			if got != wanted {
				t.Errorf("got %v wanted %v", got, wanted)
			}
		})
	}

	test(t, "XP", monster.XP, 4800)
	test(t, "Initiative", monster.Initiative, "+5")
	test(t, "Senses", monster.Senses, "foo bar baz")
	test(t, "hp", monster.HP, "103 (9d12+45)")
	test(t, "AC", monster.AC, "20, touch 10")
	test(t, "Base Atk", monster.Bab, "+9")
	test(t, "CMB", monster.CMB, "+15")
	test(t, "CMD", monster.CMD, 20)
	test(t, "SR", monster.SR, 19)
	test(t, "Str", monster.Strength, 12)
	test(t, "Dex", monster.Dexterity, 13)
	test(t, "Con", monster.Constitution, 14)
	test(t, "Int", monster.Intelligence, 15)
	test(t, "Wis", monster.Wisdom, 16)
	test(t, "Cha", monster.Charisma, 17)
}
