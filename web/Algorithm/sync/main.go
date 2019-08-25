package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读多写少: 读写锁(节约大量时间)
	wr.RLock()
	defer wr.RUnlock()

	wr.Lock()
	defer wr.Unlock()
*/
var (
	x  int
	wg sync.WaitGroup
	mt sync.Mutex
	wr sync.RWMutex
)

func read() {
	wr.RLock()
	defer wr.RUnlock()
	time.Sleep(time.Millisecond)
	wg.Done()
}

func write() {
	wr.Lock()
	defer wr.Unlock()
	x = x + 1
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	wg.Wait() // 阻塞, 等待goroute结束
	fmt.Println("x: ", x)
}
