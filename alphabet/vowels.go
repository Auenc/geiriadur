package alphabet

var vowels = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
	"w",
	"y",
}

func IsVowel(toCheck string) bool {
	for _, vowel := range vowels {
		if toCheck == vowel {
			return true
		}
	}
	return false
}
