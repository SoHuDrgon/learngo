package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("当前系统核数：", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置当前程序执行使用的并发数
	/**
	waitGroup即任务组，它的最要作用就是用来添加需要工作的任务，没完成一次任务就标记一次Done，这样任务组的待完成量会随之减1
	那么主线程就是来判断任务组内是否还有未完成任务，当没有未完成当任务之后主线程就可以结束运行，从而实现了与阻塞队列类似的同步功能
	这里创建了一个空的waitGroup(任务组)
	*/
	wg := sync.WaitGroup{}
	wg.Add(10) //添加10个任务到任务组中
	//这里启动10个线程运行
	for i := 0; i < 10; i++ {
		go goRun(&wg, i)
	}
	wg.Wait()
}

/**
这里需要传入引用类型不能传入值拷贝，因为在子线程中是需要执行Done操作，类似与我们修改结构体中的int变量主词递减，如果是只拷贝的话是不会影响原类型内的数据
这样就会发生死循环导致死锁程序奔溃，报错异常为：fatal error: all goroutines are asleep - deadlock!
*/
func goRun(wg *sync.WaitGroup, index int) {
	a := 1
	//循环叠加1千万次并返回最终结果
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println("线程序号:", index, a)

	wg.Done()
}
