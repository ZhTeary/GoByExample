package concurrent

import (
	"fmt"
	"time"
)

func Timeouts() {
	c1 := make(chan string, 1)
	//在两秒后 将"result 1" 写入到channel c1
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//设置超时时间为1s肯定超时
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1 ")
	}

	c2 := make(chan string, 2)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	//设置超时时间为3s，不会超时
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
