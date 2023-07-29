package mutations

import (
	"errors"

	"github.com/auenc/geiriadur/alphabet"
)

func hProthesisMutate(word string) (Mutation, error) {
	if word == "" {
		return "", errors.New("cannot mutate an empty string")
	}

	firstLetter, err := alphabet.GetFirstLetter(word)
	if err != nil {
		return "", err
	}
	if alphabet.IsVowel(firstLetter) {
		return Mutation("h" + word), nil
	}

	return "", nil
}
