package persian

import (
	"unicode/utf8"
)

type lexer struct {
	word  string
	pos   int
	width int
}

func (l *lexer) prior() (r rune) {
	l.backup()
	return l.next()
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
	case Heth, He:
		romanized = "h"
	case Jeem:
		romanized = "j"
	case Keheh:
		romanized = "k"
	case Tcheh:
		romanized = "ch"
	case Lamedh:
		romanized = "l"
	case Mem:
		romanized = "m"
	case Nun:
		romanized = "n"
	case Qaf:
		romanized = "q"
	case Reh:
		romanized = "r"
	case Sad:
		romanized = "s"
	case Seen:
		romanized = "s"
	case Sheen:
		romanized = "sh"
	case Yeh:
		romanized = "i"
	case Zain:
		romanized = "z"
	}

	if Diacritic(p) == Shadda {
		// The Shadda diacritic doubles the prior consonant
		l.next() // consume the input
		romanized = romanized + romanized
	}

	switch {
	case Consonant(n) == Gaf && Consonant(p) == Lamedh:
		romanized = romanized + "o"
	case Consonant(n) == Aleph && Consonant(p) == Lamedh:
		romanized = "u"
	case Consonant(n) == Lamedh && Consonant(p) == He:
		romanized = romanized + "a"
	case Consonant(n) == Heth && Consonant(p) == Mem:
		romanized = romanized + "a"
	case Consonant(p) == Yeh:
		// nothing
	case Consonant(n) == Seen && Consonant(p) == Keheh:
		romanized = romanized + "a"
	case Consonant(n) == Nun && Consonant(p) == He:
		romanized = romanized + "e"
	}

	if romanized == "" {
		romanized = string(n)
	}

	return romanized
}

func IsConsonant(r rune) bool {
	switch Consonant(r) {
	case Aleph, Bet, Taw, Heth, Dalet, Seen, Sheen, Sad, Dad, Teth, Za:
		return true
	case Ayin, Ghayn, PeSemitic, PePersian, Qoph:
		fallthrough
	case Keheh, Gaf, Lamedh, Mem, Nun, He, Waw, Yeh:
		return true
	}
	return false
}
