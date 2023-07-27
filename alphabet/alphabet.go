package alphabet

import "errors"

var Alphabet []string = []string{
	"a",
	"b",
	"c",
	"ch",
	"d",
	"dd",
	"e",
	"f",
	"ff",
	"g",
	"ng",
	"h",
	"i",
	"j",
	"l",
	"ll",
	"m",
	"n",
	"o",
	"p",
	"ph",
	"r",
	"rh",
	"s",
	"t",
	"th",
	"u",
	"w",
	"y",
}

var DoubleLetters []string = []string{
	"ch",
	"dd",
	"ff",
	"ng",
	"ll",
	"ph",
	"rh",
	"th",
}

// Getting the first letter is not as simple as getting the first character.
// In order to support double letters (and the fact that most people don't use the unicode double letters)
// we have to check the first two letters
func GetFirstLetter(word string) (string, error) {
	if word == "" {
		return "", errors.New("cannot get first letter of an empty string")
	}
	firstTwoCharacters := word[0:2]
	isDoubleLetter, err := IsDoubleLetter(firstTwoCharacters)
	if err != nil {
		return "", err
	}
	if isDoubleLetter {
		return firstTwoCharacters, nil
	}
	return string(word[0]), nil
}

func IsDoubleLetter(characters string) (bool, error) {
	if characters == "" {
		return false, errors.New("cannot check if an empty string is a double letter")
	}
	for _, letter := range DoubleLetters {
		if characters == letter {
			return true, nil
		}
	}
	return false, nil
}
