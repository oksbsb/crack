package util

import (
	"fmt"
	"testing"
)

func TestPortcheck(t *testing.T) {

	ipli := make(chan string, 0)
	go func() {

		ip, _ := Makeip("103.238.226.173")
		for _, i := range ip {
			ipli <- fmt.Sprintf("%s:2222", i)
		}
		close(ipli)
	}()

	re := Portcheck(100, ipli)
	for i := range re {
		fmt.Println(i)
	}
}
