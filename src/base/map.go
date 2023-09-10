package base

import (
	"fmt"
)

func J() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k1"]
	fmt.Println("v1:", v3)

	fmt.Println("len:", len(m))

	//map delete
	delete(m, "k2")
	fmt.Println("m deleted k2:", m)

	//clear go1.21
	//clear(m)

	if _, prs := m["k2"]; !prs {
		fmt.Println("prs:", prs)
	} else {
		fmt.Println("True")

	}

	//n:=map[string]int{"foo":1,"bar":2}
	//fmt.Println("map",n)
	//
	//n2:=map[string]int{"foo":1,"bar":2}
	//maps package in go 1.21
	//if maps.Equal(n,n2){
	//	fmt.Print("n==n2")
	//}

}
