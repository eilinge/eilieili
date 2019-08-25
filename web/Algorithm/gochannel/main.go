package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func print(chan1 <-chan int) {
	// defer wg.Done()
	x := <-chan1
	fmt.Println("x: ", x)
	wg.Done()
}
func main01() {
	var chan1 chan int // 引用类型(channel/slice/map), 需要初始化(make)才能使用
	// defer close(chan1)
	chan1 = make(chan int) // 无缓存区通道, 同步通道: 假使往channel中插入值(主协程), 无其他goroute往channel中取值(子协程), 则会造成堵塞
	wg.Add(1)
	// chan1 = make(chan int, 100) // 有缓存区通道, 异步通道: 假使往channel中插入值(主协程), 主协程中可以继续对channel操作
	/*
		chan1 <- 1
		<-chan1
	*/
	go print(chan1) // 子协程
	chan1 <- 1      // 主协程
	wg.Wait()
	// time.Sleep(time.Second * 3)
}

func f1(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer close(ch2)
	// 从channel中取值(1)
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
}

func main02() {
	ch1 := make(chan int, 1024)
	ch2 := make(chan int, 1024)

	go f1(ch1)
	go f2(ch1, ch2)

	// 从channel中取值(2)
	for v := range ch2 {
		fmt.Println("the value: ", v)
	}
}

func woker(i int, jobs chan int, ch2 chan<- int) {
	defer close(jobs)
	defer close(ch2)
	for {
		tmp, ok := <-jobs
		if !ok {
			break
		}
		fmt.Printf("the goroute: %d, tmp: %d \n", i, tmp)
		ch2 <- tmp * 3
	}
}

func main() {
	ch2 := make(chan int, 5)
	jobs := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		// 开启3个goroute
		go woker(i, jobs, ch2)
	}
	// 5个任务
	for i := 1; i <= 5; i++ {
		go func(i int) {
			jobs <- i
		}(i)
	}
	// ch2无数据时, 需要进行关闭, 否则会造成死锁
	// for v := range ch2 {
	// 	fmt.Println("the value: ", v)
	// }
	for i := 1; i <= 5; i++ {
		tmp := <-ch2
		fmt.Println("the value: ", tmp)
	}
}
