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
	testCases := map[string]string{
		"عبّاس":   "abbas",
		"گل":      "gol",
		"حمید":    "hamid",
		"عبدالله": "abdullah",
		"سکینه":   "sakineh",
		"سارا":    "sara",
		"سمیرا":   "samira",
		"ساناز":   "sanaz",
		"صدیقی":   "siddiqi",
		"شهرزاد":  "shahrazad",
		"ګل آغا شيرزی": "gul agha sherzai",
	}
	for src, expected := range testCases {
		actual := ToFinglish(src)
		if actual != expected {
			t.Errorf("Expect '%s' => '%s'. Got '%s'.", src, expected, actual)
		}
	}
}
