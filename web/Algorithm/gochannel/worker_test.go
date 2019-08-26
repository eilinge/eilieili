package main

import (
	"testing"
	// "eilieili/web/Algorithm/gochannel"
)

func TestWorker(t *testing.T) {
	n, m, y := ConstPrint()
	if n != 1 || m != 1024 || y != 1048576 {
		t.Error("the n is error")
	}
}
