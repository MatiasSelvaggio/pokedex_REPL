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
			input:    "  hEllo  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "charmanDer bulbasur squirte pikacHu",
			expected: []string{"charmander", "bulbasur", "squirte", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("not the same length %d vs %d", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("not the same word %s vs %s", word, expectedWord)
				t.Fail()
			}
		}
	}
}
