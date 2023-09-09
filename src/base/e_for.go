package base

import "fmt"

func E() {
	// 3 ways of for-loog
	i := 1

	// 1st way
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// 2nd
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 3rd
	for {
		fmt.Println("for-loop")
		break
	}
}
