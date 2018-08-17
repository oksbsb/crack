package util

import (
	"crack/model"
	"net"
	"sync"
	"time"
)

func Portcheck(numberworks int, hostports <-chan string) chan model.ServerResult {

	var wg sync.WaitGroup

	results := make(chan model.ServerResult, 0)
	for w := 0; w < numberworks; w++ {
		wg.Add(1)
		go func() {
			scanport(hostports, results)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

func scanport(jobs <-chan string, results chan<- model.ServerResult) {

	for host := range jobs {
		results <- tcpscan(host)
	}
}

func tcpscan(s string) model.ServerResult {
	res := model.ServerResult{Hostport: s}
	conn, err := net.DialTimeout("tcp", s, 2*time.Second)
	if err != nil {
		Error("%s is close", s)
		res.Open = false
		return res
	}
	defer conn.Close()
	res.Open = true
	Info("%s is open", s)
	return res
	//TODO 端口服务
	//bannerBuffer := make([]byte,256)
	//n,err := conn.Read(bannerBuffer)
	//if err !=nil {
	//
	//}

}
