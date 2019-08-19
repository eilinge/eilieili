package main

import (
	"fmt"
)

// 递归的本质: 每次调用该递归函数时, 开辟一块新的内存空间, 将此时n的值进行推进至栈底, 进行出栈时, 再进行读取n的值
// 向外部进行逐一提取, 参与函数内部逻辑
func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("the n value: ", n)
	}

}

func main() {
	n := 4
	test(n)
}
