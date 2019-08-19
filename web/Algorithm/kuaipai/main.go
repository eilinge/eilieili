package main

import "fmt"

func fabinicc(n int) int {
	if n < 2 {
		return n
	}
	return fabinicc(n-2) + fabinicc(n-1)
}

func main() {
	n := 4
	fmt.Println(fabinicc(n))
}
