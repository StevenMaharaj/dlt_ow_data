package main

import "sync"

// "os"
// "os/signal"
// "time"

func main() {
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		ob()
	}()

	go func() {
		defer wg.Done()
		trades()
	}()

	wg.Wait()

}
