package src

import "fmt"

func O() {
	i := 0
	fmt.Println("before closure", i)
	FF := func() int {
		i++
		return i
	}
	fmt.Println("after closure", i)

	FF()
	fmt.Println("after closure", i)

	FF()

	fmt.Println("after closure", i)
}
