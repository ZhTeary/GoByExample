package builtin

import (
	"cmp"
	"fmt"
	"slices"
)

func SortbyFunc() {
	//demo1
	fruits := []string{"peach", "banana", "kiwi"}

	//这里的sortFunc()和cmp.Compare是一个combine(combo)
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)

	//demo2
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		Person{"Jax", 37},
		Person{"TJ", 25},
		Person{"Alex", 72},
	}
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
