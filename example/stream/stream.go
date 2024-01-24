package main

import (
	"fmt"
	"github.com/poplknight/helper/stream"
)

func main() {
	arr := stream.NewSlice([]int{1, 1, 2, 3, 2, 3, 4, 5}...).
		Peek(Print[int]("Origin")).
		Sort(func(i int, j int) bool {
			return i > j
		}).
		Peek(Print[int]("Sort")).
		Distinct(stream.NewHash(func(t int) int {
			return t
		})).
		Peek(Print[int]("Distinct")).
		Filter(func(i int) bool {
			return i%2 == 1
		}).
		Peek(Print[int]("Filter")).ToSlice()
	fmt.Println("\nResult:", arr)
}

func Print[T any](name string) func(T) {
	fmt.Printf("\n%s:", name)
	return func(i T) {
		fmt.Printf(" %v", i)
	}
}
