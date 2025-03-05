package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cpuMaxProcs := 0
	numGoRoutines := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	start := time.Now()

	totalIterations := 10_000_000

	for range totalIterations {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			cpuMaxProcs = cpuMaxProcs + runtime.NumCPU()
			numGoRoutines = numGoRoutines + runtime.NumGoroutine()
			mu.Unlock()
		}()
	}

	wg.Wait() // Wait for all Goroutines to finish
	fmt.Println("cpuMaxProcs ", cpuMaxProcs/totalIterations)
	fmt.Println("numGoRoutines ", numGoRoutines/totalIterations)
	fmt.Println(time.Since(start).Seconds(), " seconds")
}
