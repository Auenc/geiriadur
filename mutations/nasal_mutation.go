package mutations

import (
	"errors"
	"sort"
	"strings"

	"github.com/auenc/geiriadur/alphabet"
	"golang.org/x/exp/maps"
)

var NasalMutationLetters = map[string]string{
	"p": "mh",
	"t": "nh",
	"c": "ngh",
	"b": "m",
	"d": "n",
	"g": "ng",
}

func canNasalMutate(firstLetter string) bool {
	for original := range NasalMutationLetters {
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

	newLetter := NasalMutationLetters[firstLetter]
	newWord := strings.Replace(word, firstLetter, newLetter, 1)

	return Mutation(newWord), nil
}

func nasalLettersToSortedTupleArray() [][]string {
	mutationLetters := [][]string{}
	keys := maps.Keys(NasalMutationLetters)
	sort.Strings(keys)
	for _, letter := range keys {
		mutationLetters = append(mutationLetters, []string{letter, NasalMutationLetters[letter]})
	}
	return mutationLetters
}
