package persian

import (
	"bufio"
	"strings"
)

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

		chars := bufio.NewScanner(strings.NewReader(words.Text()))
		chars.Split(bufio.ScanRunes)
		charCount := 0
		for chars.Scan() {
			char := chars.Text()
			if resultChars, found := lookup[char]; found {
				b.WriteString(resultChars[0])
			}
			charCount++
		}

		wordCount++
	}
	return b.String()
}
