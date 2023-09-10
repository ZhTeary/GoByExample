package concurrent

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	/*
		If you removed the <- done line from this program,
		the program would exit before the worker even started.
	*/
	done <- true
}

func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}
