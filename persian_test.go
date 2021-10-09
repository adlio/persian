package persian

import (
	"strings"
	"testing"

	"golang.org/x/text/unicode/runenames"
)

func assertMatch(t *testing.T, original string, expected string) {
	actual := ToFinglish(original)
	if expected != actual {
		runes := []rune(original)
		explain := make([]string, len(runes))
		for i, r := range runes {
			explain[i] = strings.TrimPrefix(runenames.Name(r), "ARABIC LETTER ")
		}
		t.Errorf("Expected '%s' => '%s'. Got '%s' [%s]", original, expected, actual, strings.Join(explain, ", "))
	}
}
