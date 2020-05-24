package scraper

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetMonsterNames(t *testing.T) {
	t.Run("it should call aonprd GET monsters", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		url := "https://www.aonprd.com/Monsters.aspx?Letter=All"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, ""))

		GetMonsterNames()

		if httpmock.GetTotalCallCount() != 1 {
			t.Errorf("expected %v to be called once", url)
		}
	})

	t.Run("it should return monster names from table", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		url := "https://www.aonprd.com/Monsters.aspx?Letter=All"
		mockHTML := "<div id=\"main\"><table><tbody>" +
			"<tr><td><a href=\"MonsterDisplay.aspx?ItemName=Foo Bar\">Foo Bar</a></td><td>8</td></tr>" +
			"<tr><td><a href=\"MonsterDisplay.aspx?ItemName=BazBar Foo\">BazBar Foo</a></td><td>20</td></tr>" +
			"</tbody></table></div>"
		httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, mockHTML))

		names := GetMonsterNames()

		if len(names) != 2 {
			t.Errorf("got %v wanted %v", len(names), 2)
		}

		if names[0] != "Foo Bar" {
			t.Errorf("got %v wanted %v", names[0], "Foo Bar")
		}

		if names[1] != "BazBar Foo" {
			t.Errorf("got %v wanted %v", names[1], "BazBar Foo")
		}
	})
}

func TestGetMonsterDetails(t *testing.T) {
	checkKeyValuePair := func(t *testing.T, gotKey string, wanted string) {
		t.Helper()
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		detailHTML := "<h1>foo</h1><br><b>XP</b> 4,800<br>CN Large dragon<br><b>Init</b> +5; <br><b>Fort</b> +11, <b>Ref</b> +7, <b>Will</b> +9;"
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
		checkKeyValuePair(t, "Init", "+5")
	})

	t.Run("it returns Fort", func(t *testing.T) {
		checkKeyValuePair(t, "Fort", "+11")
	})

	t.Run("it returns Ref", func(t *testing.T) {
		checkKeyValuePair(t, "Ref", "+7")
	})

	t.Run("it returns Will", func(t *testing.T) {
		checkKeyValuePair(t, "Will", "+9")
	})
}
