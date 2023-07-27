package mutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutate(t *testing.T) {
	type test struct {
		input         string
		want          *Mutations
		expectedError error
	}

	tests := []test{
		{
			input: "ysgol",
			want: &Mutations{
				Word:       "ysgol",
				HProthesis: Mutation("h" + "ysgol"),
			},
		},
		{
			input: "chwarae",
			want:  nil,
		},
		{
			input: "mynd",
			want: &Mutations{
				Word: "mynd",
				Soft: Mutation("fynd"),
			},
		},
		{
			input: "bangor",
			want: &Mutations{
				Word:  "bangor",
				Soft:  Mutation("fangor"),
				Nasal: Mutation("mangor"),
			},
		},
		{
			input: "caerdydd",
			want: &Mutations{
				Word:     "caerdydd",
				Soft:     Mutation("gaerdydd"),
				Nasal:    Mutation("nghaerdydd"),
				Aspirate: Mutation("chaerdydd"),
			},
		},
	}

	for _, tt := range tests {
		result, err := Mutate(tt.input)
		assert.Equal(t, err, tt.expectedError)
		assert.Equal(t, tt.want, result)
	}
}
