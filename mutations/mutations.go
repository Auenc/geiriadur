package mutations

import (
	"errors"

	"github.com/auenc/geiriadur/vowels"
)

type Mutation string
type Mutations struct {
	Word       string
	Soft       Mutation
	Nasal      Mutation
	Aspirate   Mutation
	HProthesis Mutation
}

func Mutate(word string) (*Mutations, error) {
	if word == "" {
		return nil, errors.New("word cannot be empty")
	}
	m := Mutations{Word: word}

	m.HProthesis = hProthesisMutate(word)

	return &m, nil
}

func hProthesisMutate(word string) Mutation {
	if word == "" {
		return ""
	}

	firstLetter := word[0]
	if vowels.IsVowel(rune(firstLetter)) {
		return Mutation("h" + word)
	}

	return ""
}
