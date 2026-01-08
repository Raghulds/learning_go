package main

import (
	gofundamentals "github.com/gofundamentals/go_fundamentals"
)

func main() {
	// 0
	// gofundamentals.Banner("Go", 6)
	// gofundamentals.Banner("Go", 7)
	// gofundamentals.Banner("Go!", 7)
	// gofundamentals.Banner("Go🐱", 7)

	// 1
	// gofundamentals.CallGithub()

	// 2
	// _, err := os.Create("server.pid")
	// if err != nil {
	// 	fmt.Errorf("Err while creeating file - %s", err)
	// }
	// os.WriteFile("server.pid", []byte{byte('7')}, os.ModeAppend)
	// gofundamentals.KillServer()

	// 3
	// gofundamentals.GetFileSign()

	// 3.1
	// gofundamentals.ReadFromStr("Hello, World!")
	// gofundamentals.WriteToBuffer("Hello, World!")
	// gofundamentals.ReadWithCustomReader()
	// gofundamentals.WriteWithCustomWriter()

	// 4
	// gofundamentals.SliceFundamentals()
	// unsorted := []float64{2, 1, 3, 4}
	// fmt.Println("unsorted array ", unsorted)
	// fmt.Println("unsorted array became sorted", unsorted, "it shouldn't have - Median for the array -> ", gofundamentals.GetMedian(unsorted))

	// 5
	// game.Game()

	// 6
	// gofundamentals.Generics()

	// 7
	// gofundamentals.Panic()

	// 8
	// gofundamentals.Map()

	// 9
	gofundamentals.Concurrency()

}
