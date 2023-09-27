package concurrent

import (
	"fmt"
	"time"
)

/*
	Timers are for when you want to do something once in the furure.
	Tickers are for when you want to do something repeatedly at regular intervals.
*/

func Iticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
