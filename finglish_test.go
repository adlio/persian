package persian

import "testing"

func TestToFinglishBlank(t *testing.T) {
	emptyStrings := []string{"", " ", "     "}
	for _, str := range emptyStrings {
		result := ToFinglish(str)
		if ToFinglish(str) != "" {
			t.Errorf("Blank conversion failed. Expected '%s' => ''. Got '%s'.", str, result)
		}
	}
}

func TestNames(t *testing.T) {
	assertMatch(t, "عبّاس", "abbas")
	assertMatch(t, "گل", "gol")
	assertMatch(t, "حمید", "hamid")
	assertMatch(t, "عبدالله", "abdullah")
	assertMatch(t, "سکینه", "sakineh")
	assertMatch(t, "سارا", "sara")
	assertMatch(t, "سمیرا", "samira")
	assertMatch(t, "ساناز", "sanaz")
	assertMatch(t, "صدیقی", "siddiqi")
	assertMatch(t, "شهرزاد", "shahrazad")
	assertMatch(t, "ګل آغا شيرزی", "gul agha sherzai")
}
