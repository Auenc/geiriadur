package mutations

import (
	"errors"
	"strings"

	"github.com/auenc/geiriadur/alphabet"
)

var SoftMutationLetters = map[string]string{
	"p":  "b",
	"t":  "d",
	"c":  "g",
	"b":  "f",
	"d":  "dd",
	"g":  "",
	"m":  "f",
	"ll": "l",
	"rh": "r",
}

func canSoftMutate(firstLetter string) bool {
	for original, _ := range SoftMutationLetters {
		if original == firstLetter {
			return true
		}
	}
	return false
}

func softMutate(word string) (Mutation, error) {
	if word == "" {
		return "", errors.New("cannot mutate an empty string")
	}

	firstLetter, err := alphabet.GetFirstLetter(word)
	if err != nil {
		return "", err
	}

	canMutate := canSoftMutate(firstLetter)
	if !canMutate {
		return "", nil
	}

	newLetter := SoftMutationLetters[firstLetter]
	newWord := strings.Replace(word, firstLetter, newLetter, 1)

	return Mutation(newWord), nil
}
