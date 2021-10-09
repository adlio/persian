package persian

// ToFinglish performs transliteration of the provided string from
// Persian to English
func ToFinglish(input string) string {
	lex := lexer{content: input}
	return lex.run()
}

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


    this.vowels = ['ﺕَ'.slice(-1), 'ﺕِ'.slice(-1), 'ﺕُ'.slice(-1)];

    this.words['ﺕَ'.slice(-1)] = ['a', 'a'];
    this.words['تِ'.slice(-1)] = ['e', 'e'];
    this.words['تُ'.slice(-1)] = ['e', 'e'];
    this.words['ﻁْ'.slice(-1)] = ['.'];
}
*/
