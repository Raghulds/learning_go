package gofundamentals

import (
	"fmt"
	"sort"
)

/*
Arrays are of fixed size. Length is part of their type
To increase length, new array should be created.

SLICE comes in!
s := []string{"a", "b", "c"}

Internally it's - *unsafe pointer, len 3, cap 3

slice can be nil if it's backing array is not set
nil slice VS empty slice

Cons of slices: Updating any element in slice updates the element in main array

Purpose of cap: to tell append when it should allocate a new array for the slice.
*/

func SliceFundamentals() {
	cart := []string{"apples", "bananas", "oranges"}
	fmt.Println("length - ", len(cart))
	fmt.Println("cart[1] - ", cart[1])

	for i := range cart {
		fmt.Println(i)
	}

	for i, v := range cart {
		fmt.Println(i, " has ", v)
	}

	for _, v := range cart {
		fmt.Println(v)
	}

	cart = append(cart, "grapes")
	fmt.Println("Appended fruit to cart - ", cart)

	slicedFruits := cart[:3]
	fmt.Println("sliced fruits till [:3]: ", slicedFruits)

	slicedFruits = append(slicedFruits, "pears")
	fmt.Println("Appended to sliced fruits: ", slicedFruits)

	fmt.Println("cart again - ", cart)
	fmt.Println("Why does the cart slice have the fruit appended to the sliced fruits!?")
	fmt.Println("Lets see how slice append works..")

	var sl []int
	for i := range 10_000 {
		sl = appendInt(sl, i)
	}

	fmt.Println("Concat func without using a for loop")

	type PlayerScore struct {
		Name  string
		Score float64
	}

	playerScores := []PlayerScore{
		{"Player A", 10_000},
		{"Player B", 300},
	}

	// Value semantics in for loop
	for _, p := range playerScores {
		p.Score += 100
	}
	fmt.Println("Array elem update with actual variable - ", playerScores)

	// Use index to access the element or use Pointer receiver in for loop
	for i := range playerScores {
		playerScores[i].Score += 100
	}
	fmt.Println("Array elem update with index - ", playerScores)

	fmt.Println("Slice without backing array will be 'nil'")
	var s1 []string
	s2 := []string{}
	s3 := make([]string, 0)
	fmt.Println(s1, len(s1), cap(s1), s1 == nil)
	fmt.Println(s2, len(s2), cap(s2), s2 == nil)
	fmt.Println(s3, len(s3), cap(s3), s3 == nil)
	fmt.Println("------------------------------------------")

}

// slice struct = unsafe pointer, len, cap
// Re-allocates the memory when append exceeds cap and return as slice. Old memory slice would be garbage collected
func appendInt(slice []int, val int) []int {
	i := len(slice)
	if len(slice) == cap(slice) {
		sizeToCreate := 2 * (len(slice) + 1)
		fmt.Println(cap(slice), " -> ", sizeToCreate)
		newSlice := make([]int, sizeToCreate)
		copy(newSlice, slice)
		slice = newSlice[:len(slice)]
	}

	slice = slice[:len(slice)+1]
	slice[i] = val
	return slice
}

func Concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}

func GetMedian(values []float64) float64 {
	copyValues := make([]float64, len(values))
	copy(copyValues, values)

	// values is a copy of the main slice but still points to the same underlying array
	sort.Float64s(copyValues) // This modifies the sort order of the array

	mid := len(copyValues) / 2
	if len(copyValues)%2 == 1 {
		return copyValues[mid]
	}
	return (copyValues[mid-1] + copyValues[mid]) / 2
}

/*

Backing Array: The underlying array that a slice header points to.
Multiple slices can sharethe same backing array, and changes in one slice might affect others if
they share the samebacking array and their ranges overlap.

*/
