package base

import (
	"fmt"
	"time"
)

func G() {
	//switch
	i := 2
	fmt.Print("Wtite ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//switch-condition-default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is a weekday")
	}

	//switch-default
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon ")
	}

	//switch in function
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I`m a bool")
		case int:
			fmt.Println("I`m an int")
		default:
			fmt.Printf("Don`t know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hi")

}
