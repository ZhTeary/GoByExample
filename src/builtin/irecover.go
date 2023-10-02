package builtin

import "fmt"

func mayPanic() {
	panic("a problem")
}
func IRecover() {
	defer func() {
		fmt.Println("Recovered.Error:")
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
