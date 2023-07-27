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

	ysgol := "ysgol"

	tests := []test{
		{
			input: ysgol,
			want: &Mutations{
				Word:       "ysgol",
				HProthesis: Mutation("h" + ysgol),
			},
		},
	}

	for i, tt := range tests {
		result, err := Mutate(tt.input)
		if err != tt.expectedError {
			t.Fatalf("test[%d] - expected error was not '%s'. Got got '%s'", i, tt.expectedError.Error(), err.Error())
		}

		if result == nil && tt.want != nil {
			t.Fatalf("test[%d] - result is nil but the desired mutation is not", i)
		}

		// TODO check the result
		assert.EqualExportedValues(t, *tt.want, *result)
	}
}
