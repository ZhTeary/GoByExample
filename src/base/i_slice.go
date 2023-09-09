package base

import (
	"fmt"
	//"slices" //go 1.21
)

func I() {
	//declare and assign a slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("lenL ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	//t := []string{"g", "h", "i"}
	//if slices.Equal(t, t2) {
	//	fmt.Println("t == t2")
	//}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
