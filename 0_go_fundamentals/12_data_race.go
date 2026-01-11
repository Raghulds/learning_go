package gofundamentals

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int = 0
var countAtomic int64 = 0 // sync/atomic requires exact type

const (
	CONCURRENTS         = 10
	ITERATIONS_FOR_EACH = 10_00
)

func DataRace() {
	// RaceIssue()
	fmt.Println("---------------------------")
	SyncPackage_Mutex()
	fmt.Println("---------------------------")
	SyncPackage_Atomic()
	fmt.Println("---------------------------")
}

func RaceIssue() {
	var wg sync.WaitGroup
	for range CONCURRENTS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ITERATIONS_FOR_EACH {
				count++ // fetch count -> increment count -> store count
				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Data race issue - count now: ", count)
}

/*
SYNC.MUTEX:
Coordinates accesses; Ensures only one goroutine executes the critical section at a time
Lock acquisition
Possible goroutine parking
Context switches
*/
func SyncPackage_Mutex() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range CONCURRENTS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ITERATIONS_FOR_EACH {
				mu.Lock() // Critical section
				count++
				mu.Unlock() // No defer as it's a loop block
			}
		}()
	}
	wg.Wait()
	fmt.Println("Mutex lock used - count now: ", count)
}

/*
SYNC/ATOMIC:
Single CPU instruction
No blocking
No scheduler
*/
func SyncPackage_Atomic() {
	var wg sync.WaitGroup

	for range CONCURRENTS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ITERATIONS_FOR_EACH {
				atomic.AddInt64(&countAtomic, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("sync/atmoic used - count now: ", countAtomic)
}
