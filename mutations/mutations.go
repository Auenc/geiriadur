package mutations

import (
	"errors"
	"fmt"
	"sort"

	"github.com/auenc/geiriadur/util"
	"golang.org/x/exp/maps"
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

func MutationLettersAsFlatArray(mutation string) ([][]string, error) {
	var flat [][]string
	switch mutation {
	case "soft":
		flat = softLettersToSortedTupleArray()
	case "nasal":
		flat = nasalLettersToSortedTupleArray()
	case "aspirate":
		flat = aspirateLettersToSortedTupleArray()
	case "all":
		flat = collatedMapToFlatArray()
	default:
		return flat, fmt.Errorf("unknown mutation type %s", mutation)
	}

	return flat, nil
}

func collatedMapToFlatArray() [][]string {
	collated := CollatedMutationMap()
	mutationLetters := [][]string{}
	keys := maps.Keys(collated)
	sort.Strings(keys)
	for _, letter := range keys {
		letterMutations := []string{letter}
		letterMutations = append(letterMutations, collated[letter]...)
		mutationLetters = append(mutationLetters, letterMutations)
	}

	return mutationLetters
}

func CollatedMutationMap() map[string][]string {
	collated := map[string][]string{}

	util.AppendMap(collated, SoftMutationLetters)
	util.AppendMap(collated, NasalMutationLetters)
	util.AppendMap(collated, AspirateMutationLetters)

	return collated
}
