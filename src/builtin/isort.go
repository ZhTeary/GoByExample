package builtin

import (
	"fmt"
	"slices"
)

// go 1.21
func Isort() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strs:   ", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
