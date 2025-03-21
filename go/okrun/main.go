package main

import (
	"errors"
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
			case <-time.After(time.Millisecond):
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
		g.Add(func() error {
			time.Sleep(900 * time.Millisecond)
			fmt.Printf("The second actor is returning immediately\n")
			return errors.New("immediate teardown")
		}, func(err error) {
			// Note that this interrupt function is called, even though the
			// corresponding execute function has already returned.
			fmt.Printf("The second actor was interrupted with: %v\n", err)
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}
