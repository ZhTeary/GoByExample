package concurrent

import (
	"fmt"
	"time"
)

func NonblockingChannelOperation() {
	messages := make(chan string)
	signals := make(chan bool)

	//non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message444", msg)
	default:
		fmt.Println("no message received")
	}

	//non-blocking send
	/*
			BUT here msg cannot be sent to the messages channel,
		because the channel has no buffer and there is no receiver.
		Therefore, the default case is selected.
	*/
	msg := "x"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	/*
		    use multiple cases above the default clause to implement a multi-way non-blocking select.
		Here we attempt non-blocking receives on both messages and signals.
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	/*
			如果想体验无缓冲通道传递消息并且在select中接收到，如下demo
		注意：
			1. 无缓冲通道必须读写同时存在
			2. 想要select接收到，必须把读出放在前面，否则会死锁，因为：
				Select关键字将会阻塞，指导出现一个receiver（接收者）来接受消息通道，
				但是如果receiver在select之后，select就看不到接收者，也就不会执行select
	*/
	testChan := make(chan string)
	testMsg := "x"
	go func() {
		fmt.Println("Why out first?", <-testChan)
	}()

	time.Sleep(3 * time.Second)
	select {
	case testChan <- testMsg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

}
