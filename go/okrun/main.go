package main

import (
	"context"
	"fmt"
	"time"

	"github.com/oklog/run"
)

func main() {
	var g run.Group
	{
		cancel := make(chan struct{})
		g.Add(func() error {
			select {
			case <-time.After(1000 * time.Millisecond):
				fmt.Printf("The first actor had its time elapsed\n")
				return fmt.Errorf("ooo")
			case <-cancel:
				fmt.Printf("The first actor was canceled\n")
				return nil
			}
		}, func(err error) {
			fmt.Printf("The first actor was interrupted with: %v\n", err)
			close(cancel)
		})
	}
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			select {
			case <-sl(ctx):
				fmt.Printf("The first actor had its time elapsed\n")
				return fmt.Errorf("ooo")
			case <-ctx.Done():
				fmt.Printf("The first actor was canceled\n")
				return nil
			}
		}, func(err error) {
			// Note that this interrupt function is called, even though the
			// corresponding execute function has already returned.
			fmt.Printf("The second actor was interrupted with: %v\n", err)
			cancel()
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}

func sl(ctx context.Context) <-chan int {
	out := make(chan int)
	fmt.Println("staring sl")
	go func() {
		select {
		case <-time.After(5 * time.Second): // Reduced timeout for testing
			out <- 42
			close(out)
		case <-ctx.Done():
			close(out) // Close the channel if the context is canceled
			return
		}

	}()
	return out
}
