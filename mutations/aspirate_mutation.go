package mutations

import (
	"errors"
	"sort"
	"strings"

	"github.com/auenc/geiriadur/alphabet"
	"golang.org/x/exp/maps"
)

var AspirateMutationLetters = map[string]string{
	"p": "ph",
	"t": "th",
	"c": "ch",
}

func canAspirateMutate(firstLetter string) bool {
	for original := range AspirateMutationLetters {
		if original == firstLetter {
			return true
		}
	}
	return false
}

func aspirateMutate(word string) (Mutation, error) {
	if word == "" {
		return "", errors.New("cannot mutate an empty string")
	}

	firstLetter, err := alphabet.GetFirstLetter(word)
	if err != nil {
		return "", err
	}

	canMutate := canAspirateMutate(firstLetter)

	if !canMutate {
		return "", nil
	}

	newLetter := AspirateMutationLetters[firstLetter]
	newWord := strings.Replace(word, firstLetter, newLetter, 1)

	return Mutation(newWord), nil
}

func aspirateLettersToSortedTupleArray() [][]string {
	mutationLetters := [][]string{}
	keys := maps.Keys(AspirateMutationLetters)
	sort.Strings(keys)
	for _, letter := range keys {
		mutationLetters = append(mutationLetters, []string{letter, NasalMutationLetters[letter]})
	}
	return mutationLetters
}
