package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/alitto/pond/v2"
)

func main() {
	SubmitErrorWithPool()
	// SubmitErrorWithTaskGroup()
}

func SubmitErrorWithTaskGroup() {
	// Create a pool with limited concurrency
	pool := pond.NewPool(2, pond.WithQueueSize(1))

	// Create a task group
	group := pool.NewGroup()

	// Submit a group of tasks
	for i := 0; i < 20; i++ {
		group.SubmitErr(func() error {
			if i == 10 {
				return errors.New("an error occurred")
			}
			fmt.Printf("Running group task #%d\n", i)
			return nil
		})
	}

	// Wait for all tasks in the group to complete or the first error to occur
	err := group.Wait()
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

func SubmitErrorWithPool() {
	queueSize := 10000
	workerCount := 500 // concurrency limit
	totalTasks := 10_000_000

	now := time.Now()
	runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	// Create a pool with limited concurrency
	pool := pond.NewPool(workerCount, pond.WithQueueSize(queueSize))
	queueTotal := 0
	numGoRoutines := 0
	count := 0

	var mu sync.Mutex

	// Submit 1000 tasks
	for i := 0; i < totalTasks; i++ {
		pool.Submit(func() {
			mu.Lock()
			count = count + 1
			queueTotal = queueTotal + pool.QueueSize()
			numGoRoutines = numGoRoutines + runtime.NumGoroutine()

			mu.Unlock()
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	fmt.Println("pool queueSize: ", queueSize)
	fmt.Println("pool workerCount: ", workerCount)
	pool.StopAndWait()
	fmt.Println("\ncount: ", count)
	fmt.Println("totalTasksCompleted: ", pool.CompletedTasks())
	fmt.Println("average goroutines: ", numGoRoutines/totalTasks)
	fmt.Println("average queue size: ", queueTotal/totalTasks)
	fmt.Println(time.Since(now))
}

// interesting results

// Max number of goroutines are limited by the worker workerCount. This is when the go routine has a 1000us sleep
//
// GOMAXPROCS: 8
// pool queueSize:  10000
// pool workerCount:  500
//
// count:  500000
// totalTasksCompleted:  500000
// average goroutines:  500
// average queue size:  10000
// 3.803129959s

// Without the sleep, looks like we get way fewer avg goroutines, probabbly because each one executes so quickly

// GOMAXPROCS: 8
// pool queueSize:  10000
// pool workerCount:  500
//
// count:  500000
// totalTasksCompleted:  500000
// average goroutines:  9
// average queue size:  10000
// 436.276834ms

// When there's fewer cores, there's more go routines

// GOMAXPROCS: 2
// pool queueSize:  10000
// pool workerCount:  500
//
// count:  500000
// totalTasksCompleted:  500000
// average goroutines:  18
// average queue size:  10000
// 427.193417ms

// more tasks, more goroutines

// GOMAXPROCS: 2
// pool queueSize:  10000
// pool workerCount:  500
//
// count:  10000000
// totalTasksCompleted:  10000000
// average goroutines:  35
// average queue size:  10000
// 8.449733333s

// more cores, fewer avg goroutines

// GOMAXPROCS: 8
// pool queueSize:  10000
// pool workerCount:  500
//
// count:  10000000
// totalTasksCompleted:  10000000
// average goroutines:  11
// average queue size:  10000
// 8.969439458s
