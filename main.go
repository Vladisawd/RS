package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		servises    = []string{"1 servis", "2 servis", "3 servis", "4 servis"}
		ready       string
		result      = make(chan string)
		wg          sync.WaitGroup
		ctx, cancel = context.WithCancel(context.Background())
	)

	defer cancel()

	for i := range servises {
		servisNumber := servises[i]

		wg.Add(1)
		go func() {
			finish(ctx, servisNumber, result)
			wg.Done()
		}()
	}

	go func() {
		ready = <-result
		cancel()
	}()

	wg.Wait()
	log.Printf("Ready %q servis", ready)

}

func finish(ctx context.Context, servisNum string, result chan string) {
	time.Sleep(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Stop wait in %q (%v)", servisNum, ctx.Err())
			return
		default:
			if rand.Float64() > 0.75 {
				result <- servisNum
				return
			}

			continue
		}
	}
}
