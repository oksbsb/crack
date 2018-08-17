package util

import (
	"crack/model"
	"sync"
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
		results <- i.Server(i)
	}
}
