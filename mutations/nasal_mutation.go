package mutations

import (
	"errors"
	"strings"

	"github.com/auenc/geiriadur/alphabet"
)

var NasalMurationLetters = map[string]string{
	"p": "mh",
	"t": "nh",
	"c": "ngh",
	"b": "m",
	"d": "n",
	"g": "ng",
}

func canNasalMutate(firstLetter string) bool {
	for original := range NasalMurationLetters {
		if original == firstLetter {
			return true
		}
	}
	return false
}

func nasalMutate(word string) (Mutation, error) {
	if word == "" {
		return "", errors.New("cannot mutate an empty string")
	}

	firstLetter, err := alphabet.GetFirstLetter(word)
	if err != nil {
		return "", err
	}

	canMutate := canNasalMutate(firstLetter)
	if !canMutate {
		return "", nil
	}

	newLetter := NasalMurationLetters[firstLetter]
	newWord := strings.Replace(word, firstLetter, newLetter, 1)

	return Mutation(newWord), nil
}
