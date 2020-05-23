package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetMonsterDetails(t *testing.T) {
	checkKeyValuePair := func(t *testing.T, gotKey string, wanted string) {
		t.Helper()
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		detailHTML := "<h1>foo</h1><br><b>XP</b> 4,800<br>CN Large dragon<br><b>Init</b> +5;<br><b>Fort</b> +11,<b>Ref</b> +7,<b>Will</b> +9;"
		mockHTML := "<table><tbody><tr><td><span>" +
			detailHTML +
			"</span></td></tr></tbody></table>"

		httpmock.RegisterResponder("GET", "https://www.aonprd.com/MonsterDisplay.aspx?ItemName=Aashaq%27s+Wyvern",
			httpmock.NewStringResponder(200, mockHTML))

		details := GetMonsterDetails("Aashaq's Wyvern")

		if details[gotKey] != wanted {
			t.Errorf("got %v wanted %v", details[gotKey], wanted)
		}
	}

	t.Run("it returns XP", func(t *testing.T) {
		checkKeyValuePair(t, "XP", "4,800")
	})

	t.Run("it returns Init", func(t *testing.T) {
		checkKeyValuePair(t, "Init", "+5;")
	})

	t.Run("it returns Fort", func(t *testing.T) {
		checkKeyValuePair(t, "Fort", "+11,")
	})

	t.Run("it returns Ref", func(t *testing.T) {
		checkKeyValuePair(t, "Ref", "+7,")
	})

	t.Run("it returns Will", func(t *testing.T) {
		checkKeyValuePair(t, "Will", "+9;")
	})
}
