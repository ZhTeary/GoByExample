package function

import (
	"errors"
	"fmt"
)

func f1(args int) (int, error) {
	if args == 42 {
		return -1, errors.New("can't work with 42")
	}
	return args + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "can't work with it."}
	}
	return arg + 3, nil
}

func X() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, ok := f2(i); ok != nil {
			fmt.Println("f1 failed", ok)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
