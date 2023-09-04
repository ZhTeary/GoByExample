package src

import (
	"fmt"
	"math"
)

const s string = "constant"

func D() {
	fmt.Println(s)
	const n = 500000

	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(n)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))
	fmt.Println(5e3)
}
