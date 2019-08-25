package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
goroutine是由Go的运行时(runtime)调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。

goroutine调度
	GPM是Go语言运行时(runtime)层面的实现, 是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

		G很好理解, 就是个goroutine的, 里面除了存放本goroutine信息外 还有与所在P的绑定等信息。

		P管理着一组goroutine队列, P里面会存储当前goroutine运行的上下文环境(函数指针, 堆栈地址及地址边界),
			P会对自己管理的goroutine队列做一些调度(比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等)
			当自己的队列消费完了就去全局队列里取, 如果全局队列里也消费完了会去其他P的队列里抢任务。

		M(machine)是Go运行时(runtime)对操作系统内核线程的虚拟,  M与内核线程一般是一一映射的关系,  一个groutine最终是要放到M上执行的;

		P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时, runtime会新建一个M, 阻塞G所在的P会把其他的G 挂载在新建的M上。
		当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

P的个数是通过runtime.GOMAXPROCS设定(最大256)

单从线程调度讲, Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的,
goroutine则是由Go运行时(runtime)自己的调度器调度的, 这个调度器使用一个称为m:n调度的技术
(复用/调度m个goroutine到n个OS线程)。 其一大特点是goroutine的调度是在用户态下完成的,
不涉及内核态与用户态之间的频繁切换, 包括内存的分配与释放, 都是在用户态维护着一块大的内存池,
不直接调用系统的malloc函数(除非内存池需要改变), 成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源,
近似的把若干goroutine均分在物理线程上,  再加上本身goroutine的超轻量, 以上种种保证了go调度方面的性能。
*/
/*
goroute 实现同步机制
	1. 无缓存的channel
	2. WaitGroup--管理协程执行(替换channel/time.sleep())
		WaitGroup 有3个方法--用来控制计数器的数量: Add(), Done(), Wait()
		Add()    增加计算器
		Done()   计算器减1
		Wait()   计算器非0 一直阻塞

		1. 一定要通过指针传值, 不然进程会进入死锁状态
		2. 计数器不能为负数
*/
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello lili")
	wg.Done()
}

func hello1() {
	fmt.Println("hello1 lili")
	wg.Done()
}

func main01() {
	wg.Add(2)
	go hello()
	go hello1()
	fmt.Println("hello main")
	wg.Wait()
}

func main02() {
	// wg.Add(1000)
	for i := 0; i < 1000; i++ {
		/*
			假使在闭包中直接调用i, 会导致多个i相同
			原因: 闭包中的i, 每次回到外层函数中去找, 会导致多个goroute使用同一个i
			go func() {
				fmt.Println("hello main", i)
				wg.Done()
			}()
		*/
		go func(i int) {
			wg.Add(1)
			fmt.Println("hello main", i)
			wg.Done()
		}(i)
	}
	fmt.Println("hello main")
	wg.Wait()
}

func a() {
	for i := 0; i < 20; i++ {
		fmt.Println("A: ", i)
	}
	wg.Done()
}
func b() {
	for i := 0; i < 20; i++ {
		fmt.Println("B: ", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(4)
	/*
		当使用多核时, 执行a(), b()会混乱
	*/
	wg.Add(2)
	/*
		Go1.5版本之后, 默认使用全部的CPU逻辑核心数
			B:  15
			B:  16
			B:  17
			B:  18
			B:  19
			A:  0
			A:  1
			A:  2
			A:  3
	*/
	go a()
	go b()
	fmt.Println("hello main")
	wg.Wait()
}
