package gofundamentals

import (
	"fmt"
	"time"
)

func Concurrency() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := range 3 {
		// Prior to Go 1.22 this was a bug
		go func() {
			fmt.Println("grroutine: ", i)
		}()
	}

	// Leads to deadlock as the no receiver found when send was started
	// ch := make(chan int)
	// ch <- 5   // send
	// v := <-ch // receive
	// fmt.Println("Received: ", v)

	ch := make(chan int)
	go func() {
		ch <- 5 // send
	}()
	v := <-ch // receive
	fmt.Println("Received: ", v)

	fmt.Println(sleepSort([]int{20, 30, 10}))

	time.Sleep(10 * time.Millisecond)

	deadlockDueToNoSendClose()
}

/*
Channel semantics
- send/receive to/from a channel will block untill opposite operation
- Guaranteed delivery

- Receiving from a closed channel will return a zero value without blocking

- Sending to a closed channel will panic
- Closing a closed/nil channel will panic

- send/receive to a nil channel will block forever
*/

// This would print all and then, throw deadlock error
func deadlockDueToNoSendClose() {
	ch := make(chan int)
	go func() {
		for i := range 4 {
			ch <- i + 1
		}

		// Close the channel so the for loop will end. Else it throws deadlock
		close(ch)
	}()

	for v := range ch {
		fmt.Println("Received - ", v)
	}

	v := <-ch // Receiving from a closed channel receives zero value for the type
	fmt.Println("Closed channel: ", v)

	// Using comma ok expression
	v1, ok := <-ch // Receiving from a closed channel receives zero value for the type
	fmt.Println("Closed channel: ", v1, "ok: ", ok)
	/*
		for {
		  v, ok := <-ch
		  if !ok {
		    break
		  }
		  fmt.Println("Received: ", v)
		}
	*/

	// fmt.Println("Sending to a closed channel will panic")
	// ch <- 12

	fmt.Println("Ending..")
}

/*
- For every value "n" in values, spin a goroutine that
  - sleeps for "n" ms
  - sends "n" over a channel

- collect all values from the channel to a slice &
*/
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, i := range values {
		go func() {
			time.Sleep(time.Duration(i) * time.Millisecond)
			ch <- i
		}()
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}
