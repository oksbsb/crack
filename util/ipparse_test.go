package util

import (
	"fmt"
	"testing"
	"crack/model"
)

func TestMakeip(t *testing.T) {
	ip, _ := Makeip("192.168.1.1")
	fmt.Println(ip)
	fmt.Print(model.ScanResult{})

}

func TestMakeip2(t *testing.T) {
	fmt.Print("a")
}
