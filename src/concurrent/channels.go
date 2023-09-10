package concurrent

import "fmt"

func Z() {
	messages := make(chan string)
	go func() {
		messages <- "string"
	}()

	msg := <-messages
	fmt.Println(msg)
}
