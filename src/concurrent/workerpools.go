package concurrent

import (
	"fmt"
	"time"
)

func workerpool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker: ", id, " started job ", j)
		time.Sleep(time.Second)

		fmt.Println("worker: ", id, " finished job ", j)
		results <- j * 2
	}
}

func WorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workerpool(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		/*
				not only collect the resusults
				but also ensure the worker goroutines have finished
			An alternative way to wait for multiple goroutines is to use a WaitGroup
		*/
		<-results
	}
}
