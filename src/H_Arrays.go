package src

import "fmt"

func H() {
	//first way of declaring array
	var a [5]int
	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	//second way of declaring and assign array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//two dimension arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d ", twoD)

}
