package persian

func Consonant(roman string) RuneProcessor {
	return func(l *lexer) string {
		return string(l.next())
	}
}

var (
	Aleph          = Consonant("a")
	AlephWithMadda = Consonant("a")
)

type RuneProcessor func(l *lexer) string

func DefaultProcessor(l *lexer) string {
	// By default we consume the next character
	return string(l.next())
}

var procs = map[rune]RuneProcessor{
	0x0622: AlephWithMadda,
	0x0627: Aleph,
}

/*
// Persian alphabet consonants
const (
	// What Unicode character is this: https://www.babelstone.co.uk/Unicode/whatisit.html
	// Romanization table information: https://en.wikipedia.org/wiki/Romanization_of_Persian#Comparison_table
	Bet        Consonant = 0x0628 // B as in Bob
	Taw        Consonant = 0x0629 // No pronunciation?
	Jeem       Consonant = 0x062C // J as in Jam
	Heth       Consonant = 0x062D // H as in Holday
	Dalet      Consonant = 0x062F // D as in Dave
	Ghain      Consonant = 0x063A // G as in Gah
	Reh        Consonant = 0x0631 // R as in Rabbit
	Zain       Consonant = 0x0632 // Z as in Zero
	Seen       Consonant = 0x0633 // S as in Sam
	Sheen      Consonant = 0x0634 // Sh as in Sheep (Sheen)
	Sad        Consonant = 0x0635 // S as in Sam
	Dad        Consonant = 0x0636 // Z as in zero
	Teth       Consonant = 0x0637 // T as in Tank
	Za         Consonant = 0x0638 // Z as in zero
	Ayin       Consonant = 0x0639 // - as in Uh-Oh
	Ghayn      Consonant = 0x063A // French R
	PeSemitic  Consonant = 0x0641 // F as in Fred
	PePersian  Consonant = 0x067E // P as in Pet
	Qoph       Consonant = 0x0642 // French R
	Kaf        Consonant = 0x0643 // Not listed here: https://en.wikipedia.org/wiki/Romanization_of_Persian#Fingilish
	KafRing    Consonant = 0x06AB //
	Keheh      Consonant = 0x06A9 // C as in Card { kaf mashkula }
	Gaf        Consonant = 0x06AF // G as in Go
	YehSemitic Consonant = 0x064A // Yeah
	Lamedh     Consonant = 0x0644 // L as in Lamp
	Qaf        Consonant = 0x0642 // Q
	Mem        Consonant = 0x0645 // M as in Michael
	Nun        Consonant = 0x0646 // N as in Name
	He         Consonant = 0x0647 // H as in Hot
	Waw        Consonant = 0x0648 // V as in Vision
	Tcheh      Consonant = 0x0686 // Ch
	YehPersian Consonant = 0x06CC // Farsi Yeh, Y as in Yale
)

// IsConsonant returns true if the provided rune is a Persian consonant
func IsConsonant(r rune) bool {
	switch Consonant(r) {
	case Aleph, Bet, Taw, Heth, Dalet, Seen, Sheen, Sad, Dad, Teth, Za:
		return true
	case Ayin, Ghayn, PeSemitic, PePersian, Qoph:
		fallthrough
	case Keheh, Gaf, Lamedh, Mem, Nun, He, Waw, YehSemitic, YehPersian:
		return true
	}
	return false
}
*/
