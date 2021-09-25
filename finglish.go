package persian

import (
	"bufio"
	"strings"
)

type Consonant rune

const (
	// What Unicode character is this: https://www.babelstone.co.uk/Unicode/whatisit.html
	// Romanization table information: https://en.wikipedia.org/wiki/Romanization_of_Persian#Comparison_table
	Aleph     Consonant = 0x0627
	Bet       Consonant = 0x0628 // B as in Bob
	Taw       Consonant = 0x0629 // No pronunciation?
	Heth      Consonant = 0x062D // H as in Holday
	Dalet     Consonant = 0x062F // D as in Dave
	Seen      Consonant = 0x0633 // S as in Sam
	Sheen     Consonant = 0x0634 // Sh as in Sheep (Sheen)
	Sad       Consonant = 0x0635 // S as in Sam
	Dad       Consonant = 0x0636 // Z as in zero
	Teth      Consonant = 0x0637 // T as in Tank
	Za        Consonant = 0x0638 // Z as in zero
	Ayin      Consonant = 0x0639 // - as in Uh-Oh
	Ghayn     Consonant = 0x063A // French R
	PeSemitic Consonant = 0x0641 // F as in Fred
	PePersian Consonant = 0x067E // P as in Pet
	Qoph      Consonant = 0x0642 // French R
	Kaf       rune      = 0x0643 // Not listed here: https://en.wikipedia.org/wiki/Romanization_of_Persian#Fingilish
	Keheh     Consonant = 0x06A9 // C as in Card { kaf mashkula }
	Gaf       Consonant = 0x06AF // G as in Go
	Lamedh    Consonant = 0x0644 // L as in Lamp
	Mem       Consonant = 0x0645 // M as in Michael
	Nun       Consonant = 0x0646 // N as in Name
	He        Consonant = 0x0647 // H as in Hot
	Waw       Consonant = 0x0648 // V as in Vision
	Yodh      Consonant = 0x06CC // Y as in Yale
)

type Diacritic rune

const (
	Shadda Diacritic = 0x0651 // Doubles the prior consonant
)

/*
var lookup = map[string][]string{
	"تَ": {"a", "a"},
	"تِ": {"e", "e"},
	"تُ": {"o", "o"},
	"ﻁْ": {"."},
	"ا":  {"a", "a"},
	"آ":  {"a", "a"},
	"ئ":  {"a"},
	"ء":  {"a"},
	"ب":  {"b"},
	"پ":  {"p"},
	"ت":  {"t"},
	"ث":  {"s"},
	"ج":  {"j"},
	"چ":  {"ch"},
	"ح":  {"h"},
	"خ":  {"kh"},
	"د":  {"d"},
	"ذ":  {"z"},
	"ر":  {"r"},
	"ز":  {"z"},
	"ژ":  {"zh"},
	"س":  {"s"},
	"ش":  {"sh"},
	"ص":  {"s"},
	"ض":  {"z"},
	"ط":  {"t"},
	"ظ":  {"z"},
	"ع":  {"", "a"},
	"غ":  {"gh"},
	"ف":  {"f"},
	"ق":  {"gh"},
	"ک":  {"k"},
	"گ":  {"g"},
	"ل":  {"l"},
	"م":  {"m"},
	"ن":  {"n"},
	"و":  {"v", "o"},
	"ه":  {"h"},
	"ی":  {"y", "i"},
	"ي":  {"y", "i"},
}
*/

func ToFinglish(input string) string {
	var b strings.Builder
	words := bufio.NewScanner(strings.NewReader(input))
	words.Split(bufio.ScanWords)
	wordCount := 0
	for words.Scan() {
		if wordCount > 0 {
			// Add leading space before all but the first word
			b.WriteString(" ")
		}

		lex := lexer{word: words.Text()}
		str := lex.nextRomanized()
		for str != "" {
			b.WriteString(str)
			str = lex.nextRomanized()
		}

		wordCount++
	}
	return b.String()
}
