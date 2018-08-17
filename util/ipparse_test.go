package util

import (
	"fmt"
	"testing"
)

func TestMakeip(t *testing.T) {
	ip, _ := Makeip("192.168.1.1.1.1")
	fmt.Println(ip)
}
