package concurrent

import "fmt"

func RangeChan() {
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
