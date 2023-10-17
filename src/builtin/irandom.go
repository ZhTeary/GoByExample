package builtin

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {
	fmt.Println(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64())       // [0.0 , 1.0)
	fmt.Println(rand.Float64()*5 + 5) // [5.0 , 10.0)
	fmt.Println(rand.Float64()*5 + 5) // [5.0 , 10.0)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))

	fmt.Println("===========")
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100))
	fmt.Println(r3.Intn(100))

}
