package util

import (
	"context"
	"crack/model"
	"sync"
	"time"
)

func Passattack(numWorks int, s <-chan model.ScanResult) chan model.ScanResult {
	var wg sync.WaitGroup
	res := make(chan model.ScanResult, 1024)
	for w := 0; w < numWorks; w++ {
		wg.Add(1)
		go func() {
			attackworker(s, res)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
func attackworker(s <-chan model.ScanResult, results chan model.ScanResult) {
	for i := range s {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))

		ch := make(chan struct{})

		go func(ch chan struct{}) {

			results <- i.Server(i)

			ch <- struct{}{}
		}(ch)
		select {
		case <-ch:
		case <-ctx.Done():

			Error("%s \t time out", i.Hostport)
		}
		cancel()
	}
}
