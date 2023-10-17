package builtin

import (
	"fmt"
	"strconv"
)

func Parsing() {
	/*
		32 for float32, or 64 for float64.
		When bitSize=32, the result still has type float64,
		but it will be convertible to float32 without changing its value
	*/
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("010", 8, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
