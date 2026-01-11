package gofundamentals

import (
	"context"
	"fmt"
	"time"
)

func MultipleChannels() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 105*time.Millisecond)
	defer cancel()

	// goroutine leaking - Timeout happened and goroutines are in the fly and gets stuck in memory
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- 2
	}()
	// Approach for goroutine leak - Using select & a timeout case within, buffered channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		ctx3, cancel3 := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel3()
		select {
		case ch3 <- 3:
			fmt.Println("Channel 3 - out")
		case <-ctx3.Done():
			fmt.Println("Channel 3 - Timeout")
		}
	}()

	select {
	case <-ch1:
		fmt.Println("Channel 1 out")
	case <-ch2:
		fmt.Println("Channel 2 out")
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
}
