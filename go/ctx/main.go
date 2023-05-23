package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	var fn = func() {
		<-ctx.Done()
		wg.Done()
	}

	wg.Add(3)
	go fn()
	go fn()
	go fn()

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
