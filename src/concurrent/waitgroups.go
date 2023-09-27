package concurrent

import (
	"fmt"
	"sync"
	"time"
)

/*
	To wait for multiple goroutines to finish
*/

func workerwg(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func WG() {
	/*
		If a WaitGroup is explicitly passed into functions, it should be done by pointer
	*/
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		/*
			Avoid re-use of the same i value in each goroutine closure.
			https://golang.org/doc/faq#closures_and_goroutines
		*/
		go func(u int) {
			defer wg.Done()
			workerwg(u)
		}(i)
		/* easier way
		i := i
		go func() {
			i := i
			defer wg.Done()
			workerwg(i)
		}()

		*/
	}

	wg.Wait()

}
