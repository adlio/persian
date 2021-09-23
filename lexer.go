package persian

import (
	"unicode/utf8"
)

type lexer struct {
	word  string
	pos   int
	width int
}

// next consumes the next UTF-8 rune in the word
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.word) {
		l.width = 0
		return 0
	}
	r, l.width = utf8.DecodeRuneInString(l.word[l.pos:])
	l.pos += l.width
	return r
}

// peek retrieves the next rune in the word but does not consume it
func (l *lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) nextRomanized() string {
	n := l.next()
	p := l.peek()

	if n == 0 {
		return ""
	}

	var romanized string
	switch Consonant(n) {
	case Aleph:
		romanized = "a"
	case Ayin:
		romanized = "a"
	case Bet:
		romanized = "b"
	case Dalet:
		romanized = "d"
	case Gaf:
		romanized = "g"
	case Heth:
		romanized = "h"
	case Lamedh:
		romanized = "l"
	case Mem:
		romanized = "m"
	case Sad:
		romanized = "s"
	case Seen:
		romanized = "s"
	case Yodh:
		romanized = "i"
	}

	if Diacritic(p) == Shadda {
		l.next() // consume the input
		romanized = romanized + romanized
	}

	switch {
	case Consonant(n) == Gaf && Consonant(p) == Lamedh:
		romanized = romanized + "o"
	case Consonant(n) == Heth && Consonant(p) == Mem:
		romanized = romanized + "a"
	case Consonant(p) == Yodh:
		// nothing
	case IsConsonant(n) && IsConsonant(p):
		romanized = "n=" + string(n) + "p=" + string(p)
		// romanized = romanized + "a"
	}

	if romanized == "" {
		romanized = string(n)
	}

	return romanized
}

func IsConsonant(r rune) bool {
	switch Consonant(r) {
	case Aleph:
	case Bet:
		fallthrough
	case Taw:
		fallthrough
	case Heth:
		fallthrough
	case Dalet:
		fallthrough
	case Seen:
		fallthrough
	case Sheen:
		fallthrough
	case Sad:
		fallthrough
	case Dad:
		fallthrough
	case Teth:
		fallthrough
	case Za:
		fallthrough
	case Ayin:
		fallthrough
	case Ghayn:
		fallthrough
	case PeSemitic:
		fallthrough
	case PePersian:
		fallthrough
	case Qoph:
		fallthrough
	case Kaph:
		fallthrough
	case Gaf:
		fallthrough
	case Lamedh:
	case Mem:
		fallthrough
	case Nun:
		fallthrough
	case He:
		fallthrough
	case Waw:
		fallthrough
	case Yodh:
		return true
	}
	return false
}