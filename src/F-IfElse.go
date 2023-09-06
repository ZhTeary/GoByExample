package src

import "fmt"

func F() {
	// if
	if 8%4 == 0 {
		fmt.Println("8 is divisble by 4")
	}
	//if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//if-elseif-else
	if num := 9; num < 0 {
		fmt.Println(num, "is negeative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
