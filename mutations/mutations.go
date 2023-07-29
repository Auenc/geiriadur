package mutations

import (
	"errors"
	"fmt"
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

	return &m, nil
}

func MutationLettersAsTuples(mutation string) ([][]string, error) {
	var tuples [][]string
	switch mutation {
	case "soft":
		tuples = softLettersToSortedTupleArray()
	case "nasal":
		tuples = nasalLettersToSortedTupleArray()
	case "aspirate":
		tuples = aspirateLettersToSortedTupleArray()
	default:
		return tuples, fmt.Errorf("unknown mutation type %s", mutation)
	}

	return tuples, nil
}
