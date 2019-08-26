package data

import (
	"fmt"
	"testing"
)

const url = "https://github.com/EDDYCJY"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	fmt.Println("b.N ", b.N)
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}
