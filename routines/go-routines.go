package routines

import (
	"fmt"
	"time"

	// "runtime"
)

func LearnRoutines() {

	// use go to create a go routine
	// create with explicitly declared function
	// if the main routine ends, all sub routines will immediately ends
	
	/* 	
	go NewTask()

	i := 0
	for {
		i++
		fmt.Println("main go routine: i = ", i)
		time.Sleep(1 * time.Second)
	} 
	*/

	// create go routine using anonymous method
	/* 	
	go func(a, b int) {
		defer fmt.Println("A.defer")

		fmt.Printf("a = %d, b = %d\n", a, b)

		func () {
			defer fmt.Println("B.defer")

			// use runtime.Goexit() to exit a go routine
			// runtime.Goexit()

			fmt.Println("B")
		}()

		fmt.Println("A")
	}(10, 100) // use () to invoke a anonymous method and pass params

	for {
		time.Sleep(1 * time.Second)
	}
	*/

	// to use the go routine's result, we need to use channel to pass data through defferent routines

	// create channel
	// channel has two types: non-buffered channel and buffered channel
	// channel's declaration should point which kind of data type passing through it
	// expression: make(chan Type, [capacity])
	nonBufferedChan := make(chan int) // equals make(chan int, 0)
	bufferedChan := make(chan int, 10)
	bufferedChanToClose := make(chan int, 10)

	// pass data through channel
	// use chan <- value to pass a value to chan
	// use value := <-chan to receive a value from chan, 
	// also: <-chan(don't use value), value, notClosed := <-chan(get whether chan is closed)
	
	// channel with buffer would not block any of the sending or receiving go routine unless the buffer is full 
	// or the buffer is empty while receiving channel is waiting
	go func ()  {
		for i := 0; i < 4; i++ {
			bufferedChan <- i
			bufferedChanToClose <- i
		}
		fmt.Println("buffered-chan send value ends")

		close(bufferedChanToClose)
	}()

	for i := 0; i < 4; i++ {
		value := <-bufferedChan
		fmt.Println("value of buffered-chan: ", value)
	}

	// channel with no buffer will block the sending and receiving go routine
	// if data is sent but not received, the sending routine will block till the data is received
	// if a routine is waiting for channel's data but doesn't get, the receiving routine will block till the data is sent
	go func ()  {
		for i := 0; i < 4; i++ {
			nonBufferedChan <- i
		}

		fmt.Println("non-buffered-chan send value ends")
	}()
	
	for i := 0; i < 4; i++ {
		value := <-nonBufferedChan
		fmt.Println("value of non-buffered-chan: ", value)
	}

	// close channel
	// use close(chan) to close a channel
	// if a channel is closed, it will panic when sending data to it. But we still can receive data from it
	// this will receiving data from channel unless the channel is closed
	// if the channel is not closed, it will cause a dead lock
	/* 
	for {
		if data, notClosed := <-bufferedChanToClose; notClosed {
			fmt.Println("data: ", data)
		} else {
			return
		}
	}
	*/

	// use range to optimize the upper option
	// it will automatically stop when the channel is closed
	for v := range bufferedChanToClose {
		fmt.Println("value: ", v)
	}
	
	// use select to monitor multi-channel's status
	// in a single process we could just monitor one channel's status at a time
	dataChan := make(chan int)
	quitChan := make(chan int)

	go func ()  {
		for i := 0; i < 6; i++ {
			fmt.Println(<-dataChan)
		}

		quitChan <- 0
	}()

	fibonacci(dataChan, quitChan)
}

func fibonacci(dataChan, quitChan chan int)  {
	x, y := 0, 1

	for {
		select {
		case dataChan <- x:
			// if dataChan is writtable, this case will be executed
			temp := y
			y = x
			x += temp
		case <-quitChan:
			// if quitChan is readable, this case will be executed
			fmt.Println("quit")
			return
		}
	}
}

func NewTask() {
	i := 0

	for {
		i++
		fmt.Println("new go routine: i = ", i)

		time.Sleep(1 * time.Second)
	}
}