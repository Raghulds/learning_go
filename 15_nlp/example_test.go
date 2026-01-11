package nlp_test

import (
	"fmt"

	"github.com/Raghulds/nlp"
)

func ExampleCustomTokenize() {
	tokens := nlp.CustomTokenize("Who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who on first]
}
