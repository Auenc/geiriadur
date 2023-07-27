package alphabet

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDoubleLetter(t *testing.T) {
	type test struct {
		input         string
		want          bool
		expectedError error
	}

	tests := []test{
		{
			input:         "",
			want:          false,
			expectedError: errors.New("cannot check if an empty string is a double letter"),
		},
		{
			input: "a",
			want:  false,
		},
		{
			input: "ll",
			want:  true,
		},
		{
			input: "dd",
			want:  true,
		},
		{
			input: "rh",
			want:  true,
		},
		{
			input: "ch",
			want:  true,
		},
		{
			input: "ph",
			want:  true,
		},
		{
			input: "th",
			want:  true,
		},
		{
			input: "ff",
			want:  true,
		},
		{
			input: "ng",
			want:  true,
		},
	}

	for _, tt := range tests {
		result, err := IsDoubleLetter(tt.input)
		assert.Equal(t, err, tt.expectedError)
		assert.Equal(t, result, tt.want)
	}
}

func TestGetFirstLetter(t *testing.T) {
	type test struct {
		input         string
		want          string
		expectedError error
	}
	tests := []test{
		{
			input:         "",
			want:          "",
			expectedError: errors.New("cannot get first letter of an empty string"),
		},
		{
			input: "chwarae",
			want:  "ch",
		},
		{
			input: "cig",
			want:  "c",
		},
	}

	for _, tt := range tests {
		result, err := GetFirstLetter(tt.input)
		assert.Equal(t, err, tt.expectedError)
		assert.Equal(t, result, tt.want)
	}
}
