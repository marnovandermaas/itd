package translit

import (
	"strings"
)

// Transliterate runs the given maps on s and returns the result
func Transliterate(s string, useMaps ...string) string {
	// Create variable to store modified string
	out := s
	// If custom map exists
	if custom, ok := Transliterators["custom"]; ok {
		// Perform transliteration with it
		out = custom.Transliterate(out)
	}
	// For every map to use
	for _, useMap := range useMaps {
		// If custom, skip
		if useMap == "custom" {
			continue
		}
		// Get requested map
		transliterator, ok := Transliterators[useMap]
		if !ok {
			continue
		}
		transliterator.Init()
		// Perform transliteration
		out = transliterator.Transliterate(out)
	}
	// Return result
	return out
}

// Transliterator is implemented by anything with a
// Transliterate method, which performs transliteration
// and returns the resulting string.
type Transliterator interface {
	Transliterate(string) string
	Init()
}

// Map implements Transliterator using a slice where
// every even element is a key and every odd one is a value
// which replaces the key.
type Map []string

func (mt Map) Transliterate(s string) string {
	return strings.NewReplacer(mt...).Replace(s)
}

func (Map) Init() {}

// Transliterators stores transliterator implementations for each supported language.
// Some of these were sourced from https://codeberg.org/Freeyourgadget/Gadgetbridge
var Transliterators = map[string]Transliterator{
	"eASCII": Map{
		"œ", "oe",
		"ª", "a",
		"°", "o",
		"«", `"`,
		"»", `"`,
	},
	"Scandinavian": Map{
		"Æ", "Ae",
		"æ", "ae",
		"Ø", "Oe",
		"ø", "oe",
		"Å", "Aa",
		"å", "aa",
	},
	"German": Map{
		"ä", "ae",
		"ö", "oe",
		"ü", "ue",
		"Ä", "Ae",
		"Ö", "Oe",
		"Ü", "Ue",
		"ß", "ss",
		"ẞ", "SS",
	},
	"Hebrew": Map{
		"א", "a",
		"ב", "b",
		"ג", "g",
		"ד", "d",
		"ה", "h",
		"ו", "u",
		"ז", "z",
		"ח", "kh",
		"ט", "t",
		"י", "y",
		"כ", "c",
		"ל", "l",
		"מ", "m",
		"נ", "n",
		"ס", "s",
		"ע", "'",
		"פ", "p",
		"צ", "ts",
		"ק", "k",
		"ר", "r",
		"ש", "sh",
		"ת", "th",
		"ף", "f",
		"ץ", "ts",
		"ך", "ch",
		"ם", "m",
		"ן", "n",
	},
	"Greek": Map{
		"α", "a",
		"ά", "a",
		"β", "v",
		"γ", "g",
		"δ", "d",
		"ε", "e",
		"έ", "e",
		"ζ", "z",
		"η", "i",
		"ή", "i",
		"θ", "th",
		"ι", "i",
		"ί", "i",
		"ϊ", "i",
		"ΐ", "i",
		"κ", "k",
		"λ", "l",
		"μ", "m",
		"ν", "n",
		"ξ", "ks",
		"ο", "o",
		"ό", "o",
		"π", "p",
		"ρ", "r",
		"σ", "s",
		"ς", "s",
		"τ", "t",
		"υ", "y",
		"ύ", "y",
		"ϋ", "y",
		"ΰ", "y",
		"φ", "f",
		"χ", "ch",
		"ψ", "ps",
		"ω", "o",
		"ώ", "o",
		"Α", "A",
		"Ά", "A",
		"Β", "B",
		"Γ", "G",
		"Δ", "D",
		"Ε", "E",
		"Έ", "E",
		"Ζ", "Z",
		"Η", "I",
		"Ή", "I",
		"Θ", "Th",
		"Ι", "I",
		"Ί", "I",
		"Ϊ", "I",
		"Κ", "K",
		"Λ", "L",
		"Μ", "M",
		"Ν", "N",
		"Ξ", "Ks",
		"Ο", "O",
		"Ό", "O",
		"Π", "P",
		"Ρ", "R",
		"Σ", "S",
		"Τ", "T",
		"Υ", "Y",
		"Ύ", "Y",
		"Ϋ", "Y",
		"Φ", "F",
		"Χ", "Ch",
		"Ψ", "Ps",
		"Ω", "O",
		"Ώ", "O",
	},
	"Russian": Map{
		"Ё", "Йo",
		"ё", "йo",
	},
	"Ukranian": Map{
		"ґ", "gh",
		"є", "je",
		"і", "i",
		"ї", "ji",
		"Ґ", "Gh",
		"Є", "Je",
		"І", "I",
		"Ї", "JI",
	},
	"Arabic": Map{
		"ا", "a",
		"ب", "b",
		"ت", "t",
		"ث", "th",
		"ج", "j",
		"ح", "7",
		"خ", "5",
		"د", "d",
		"ذ", "th",
		"ر", "r",
		"ز", "z",
		"س", "s",
		"ش", "sh",
		"ص", "9",
		"ض", "9'",
		"ط", "6",
		"ظ", "6'",
		"ع", "3",
		"غ", "3'",
		"ف", "f",
		"ق", "q",
		"ك", "k",
		"ل", "l",
		"م", "m",
		"ن", "n",
		"ه", "h",
		"و", "w",
		"ي", "y",
		"ى", "a",
		"ﺓ", "",
		"آ", "2",
		"ئ", "2",
		"إ", "2",
		"ؤ", "2",
		"أ", "2",
		"ء", "2",
		"٠", "0",
		"١", "1",
		"٢", "2",
		"٣", "3",
		"٤", "4",
		"٥", "5",
		"٦", "6",
		"٧", "7",
		"٨", "8",
		"٩", "9",
	},
	"Farsi": Map{
		"پ", "p",
		"چ", "ch",
		"ژ", "zh",
		"ک", "k",
		"گ", "g",
		"ی", "y",
		"\u200c", " ",
		"؟", "?",
		"٪", "%",
		"؛", ";",
		"،", ":",
		"۱", "1",
		"۲", "2",
		"۳", "3",
		"۴", "4",
		"۵", "5",
		"۶", "6",
		"۷", "7",
		"۸", "8",
		"۹", "9",
		"۰", "0",
		"»", "<",
		"«", ">",
		"ِ", "e",
		"َ", "a",
		"ُ", "o",
		"ّ", "",
	},
	"Polish": Map{
		"Ł", "L",
		"ł", "l",
	},
	"Lithuanian": Map{
		"ą", "a",
		"č", "c",
		"ę", "e",
		"ė", "e",
		"į", "i",
		"š", "s",
		"ų", "u",
		"ū", "u",
		"ž", "z",
	},
	"Estonian": Map{
		"ä", "a",
		"Ä", "A",
		"ö", "o",
		"õ", "o",
		"Ö", "O",
		"Õ", "O",
		"ü", "u",
		"Ü", "U",
	},
	"Icelandic": Map{
		"Þ", "Th",
		"þ", "th",
		"Ð", "D",
		"ð", "d",
	},
	"Czeck": Map{
		"ř", "r",
		"ě", "e",
		"ý", "y",
		"á", "a",
		"í", "i",
		"é", "e",
		"ó", "o",
		"ú", "u",
		"ů", "u",
		"ď", "d",
		"ť", "t",
		"ň", "n",
	},
	"French": Map{
		"à", "a",
		"â", "a",
		"é", "e",
		"è", "e",
		"ê", "e",
		"ë", "e",
		"ù", "u",
		"ü", "u",
		"ÿ", "y",
		"ç", "c",
	},
	"Emoji": Map{
		"😂", ":')",
		"😊", ":)",
		"😃", ":)",
		"😩", "-_-",
		"😏", ":‑J",
		"💜", "<3",
		"💖", "<3",
		"💗", "<3",
		"❤️", "<3",
		"💕", "<3",
		"💞", "<3",
		"💘", "<3",
		"💓", "<3",
		"💚", "<3",
		"💙", "<3",
		"💔", "</3",
		"😱", "D:",
		"😮", ":O",
		"😝", ":P",
		"😍", ":x",
		"😢", ":(",
		"💯", ":100:",
		"🔥", ":fire:",
		"😉", ";)",
		"😴", ":zzz:",
		"💤", ":zzz:",
	},
	"Korean":   &KoreanTranslit{},
	"Chinese":  &ChineseTranslit{},
	"Armenian": &ArmenianTranslit{},
}
