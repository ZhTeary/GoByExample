package concurrent

import (
	"fmt"
	"time"
)

/*
We often want to execute Go code at some point in the future,or repeatedly at some interval.
Go’s built-in timer and ticker features make both of these tasks easy
*/

func Itimer() {
	//Timers represent a single event in the future.

	timer1 := time.NewTimer(2 * time.Second)

	//<-timer1.C blocks on the timer`s channel C untils it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	/*
		If you just wanted to wait, you could have used time.Sleep.
		One reason a timer may be useful is that you can cancel the timer before it fires.
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopeed")
	}

	time.Sleep(2 * time.Second)
}
