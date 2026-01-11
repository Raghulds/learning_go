package nlp

import (
	"os"
	"slices"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestCustomTokenize(t *testing.T) {
	text := "Who's on first?"
	tokens := CustomTokenize(text)
	expected := []string{"who", "on", "first"}
	require.Equal(t, expected, tokens)
	// if !slices.Equal(expected, tokens) {
	// 	t.Fatalf("expected %#v, got %#v", expected, tokens)
	// }
}

type tokCase struct {
	Text   string
	Tokens []string
	Name   string
}

func loadTokenizeCases(t *testing.T) []tokCase {
	file, err := os.Open("testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close()

	var data struct {
		Cases []tokCase `toml:"case"`
	}
	dec := toml.NewDecoder(file)
	_, err = dec.Decode(&data)
	require.NoError(t, err)
	return data.Cases
}

func TestCustomTokenizeTable(t *testing.T) {
	// var cases = []struct {
	// 	text   string
	// 	tokens []string
	// }{
	// 	{"Who's on first", []string{"who", "on", "first"}},
	// 	{"Who's on Sec", []string{"who", "on", "sec"}},
	// 	{"", nil},
	// }
	cases := loadTokenizeCases(t)

	for _, tc := range cases {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := CustomTokenize(tc.Text)
			if !slices.Equal(tc.Tokens, tokens) {
				t.Fatalf("expected %#v, got %#v", tc.Tokens, tokens)
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
