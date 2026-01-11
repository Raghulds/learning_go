package nlp

import (
	"os"
	"slices"
	"testing"
)

func TestCustomTokenize(t *testing.T) {
	text := "Who's on first?"
	tokens := CustomTokenize(text)
	expected := []string{"who", "on", "first"}
	if !slices.Equal(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
}

func TestCustomTokenizeTable(t *testing.T) {
	var cases = []struct {
		text   string
		tokens []string
	}{
		{"Who's on first", []string{"who", "on", "first"}},
		{"Who's on Sec", []string{"who", "on", "sec"}},
		{"", nil},
	}

	for _, tc := range cases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := CustomTokenize(tc.text)
			if !slices.Equal(tc.tokens, tokens) {
				t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
			}
		})
	}
}

/* Selecting tests
- "-run" flag: regexp
- build tags (//go:build comment)
- environment variables
*/

// In jenkins, use BUILD_NUMBER
var inCI = os.Getenv("CI") != ""

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("no in CI")
	}
}
