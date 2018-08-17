package util

import (
	"fmt"
	"testing"
)

func TestMakefilelist(t *testing.T) {
	li, _ := Makelist("/etc/passwd")
	fmt.Println(len(li))
}
