package persian


// Diacritic defines a diacritic character in Persian unicode
type Diacritic rune

// All known Diacritics
const (
	Fatha  Diacritic = 0x064E // Vowel point for long a
	Damma  Diacritic = 0x064F // In the Arabic script, the vowel point for u, appearing as a small curl placed above a letter ( ـُ ) and designating a short u /u/. If the Arabic letter و‎ (wāw) immediately follows, it indicates a long ū /uː/
	Kasra  Diacritic = 0x0650 // In Persian represents the vowel /e/ and is only used in front of ـی‎ to represent a diphthong, so unmarked دی‎ is /di/ but دِی‎ with the diacritic is /dej/. This usage is different from Arabic.
	Shadda Diacritic = 0x0651 // Doubles the prior consonant
	Sukun  Diacritic = 0x0652 // Indicates the absence of a vowel
)

// IsDiacritic returns true if the provided rune is a diacritic in
// Persian
func IsDiacritic(r rune) bool {
	switch Diacritic(r) {
	case Fatha, Damma, Kasra, Shadda, Sukun:
		return true
	default:
		return false
	}
}
