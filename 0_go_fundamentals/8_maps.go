package gofundamentals

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// `a` is a "raw" string, at \ is jest a \
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Code that runs before main func
// - var expressions
// - init function

func Map() {
	// mapDemo()
	file, err := os.Open("go_fundamentals/sherlock.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	freq := make(map[string]int) // word -> count
	s := bufio.NewScanner(file)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}
	}
	if e := s.Err(); e != nil {
		fmt.Println(e)
	}
	// fmt.Println("Map - ", freq)

	top := topN(freq, 10)
	fmt.Println(top)
}

// returns top N common words in freq
func topN(freq map[string]int, n int) []string {
	wordSlice := slices.Collect(maps.Keys(freq))
	sort.Slice(wordSlice, func(i, j int) bool {
		wi, wj := wordSlice[i], wordSlice[j]
		return freq[wi] > freq[wj]
	})

	n = min(n, len(wordSlice))
	return wordSlice[:n]
}

func mapDemo() {
	mapOp := map[string]string{
		"Superman":     "A",
		"Wonder women": "B",
		"Batman":       "C",
	}

	for k, v := range mapOp {
		fmt.Println("Key - ", k, ",Value - ", v)
	}

	n := mapOp["Batman"]
	m := mapOp["Aquaman"]
	fmt.Println(n)
	// Non existent key would be zero value
	fmt.Printf("%q\n", m)
	o, ok := mapOp["Aquaman"]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(o)
	}

	delete(mapOp, "Wonder women")
	fmt.Printf("%v\n", mapOp)
}
