package gofundamentals

import (
	"errors"
	"fmt"
	"time"
)

func Generics() {
	fmt.Println(Ret(7))
	fmt.Println(Ret(4.3))
	fmt.Println(Ret(time.February))

	m, _ := NewMatrix[int](3, 10)
	fmt.Println("m - ", m)
	fmt.Println(m.GetCell(2, 8))
	fmt.Println(Max([]int{4, 5, 1}))
	fmt.Println(Max([]float64{4, 5.2, 1}))
	fmt.Println(Max[int](nil))
}

/*
func RetInt(i int64) int64 {
	if i < 0 {
		return 0
	}

	return i
}

func RetFloat(i float64) float64 {
	if i < 0 {
		return 0
	}

	return i
}
*/

type Number interface {
	~int | ~float64
}

// T is a type constraint, not a new type
func Ret[T Number](data T) T {
	if data < 0 {
		return 0
	}

	return data
}

type Matrix[T Number] struct {
	Row  int
	Col  int
	data []T
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows < 0 || cols < 0 {
		return nil, fmt.Errorf("Bad dimensions: %d %d", rows, cols)
	}

	m := Matrix[T]{
		Row:  rows,
		Col:  cols,
		data: make([]T, rows*cols),
	}

	return &m, nil
}

func (m *Matrix[T]) GetCell(row, col int) (T, error) {
	var zero T
	if row < 0 || row > m.Row || col < 0 || col > m.Col {
		return zero, fmt.Errorf("Not a valid cell for the matrix")
	}

	i := (row * m.Col) + col
	return m.data[i], nil
}

func Max[T Number](values []T) (T, error) {
	if len(values) == 0 {
		return 0, errors.New("Max of empty slice")
	}
	max := values[0]

	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}
