package alphabet

var Vowels = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
	"w",
	"y",
}

func IsVowel(toCheck string) bool {
	for _, vowel := range Vowels {
		if toCheck == vowel {
			return true
		}
	}
	return false
}
