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
			input:    "hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ThIs is A test  ",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "YOU CAN DO IT",
			expected: []string{"you", "can", "do", "it"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, got %d, input: %q", len(c.expected), len(actual), c.input)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatch at index %d: expected %q, got %q, input: %q", i, c.expected[i], actual[i], c.input)
			}
		}
	}
}
