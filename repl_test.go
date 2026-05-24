package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hola soy Alex   , ",
			expected: []string{"hola", "soy", "alex", ","},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice against the expected slice
		// if they dont match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("expected %v, got %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %v, got %v", expectedWord, word)
			}
		}
	}
}
