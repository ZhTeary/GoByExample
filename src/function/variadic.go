package function

import "fmt"

func N(nums ...int) {
	fmt.Println(nums, "")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total", total)
}
