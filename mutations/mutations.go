package mutations

import (
	"errors"

	"github.com/auenc/geiriadur/alphabet"
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

	hProthesis, err := hProthesisMutate(word)
	if err != nil {
		return nil, err
	}
	m.HProthesis = hProthesis

	softMutation, err := softMutate(word)
	if err != nil {
		return nil, err
	}
	m.Soft = softMutation

	nasalMutation, err := nasalMutate(word)
	if err != nil {
		return nil, err
	}
	m.Nasal = nasalMutation

	aspirateMutation, err := aspirateMutate(word)
	if err != nil {
		return nil, err
	}
	m.Aspirate = aspirateMutation

	if m.Aspirate == "" && m.HProthesis == "" && m.Nasal == "" && m.Soft == "" {
		return nil, nil
	}

	return &m, nil
}

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
