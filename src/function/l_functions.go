package function

import "fmt"

func L() {
	res := plus(1, 2)
	fmt.Println("res:", res)
}

func plus(a int, b int) int {
	return a + b
}
