package vowels

var vowels = []rune{
	'a',
	'e',
	'i',
	'o',
	'u',
	'w',
	'y',
}

func IsVowel(toCheck rune) bool {
	for _, vowel := range vowels {
		if toCheck == vowel {
			return true
		}
	}
	return false
}
