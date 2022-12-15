package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) // 用来等待协程结束
	fmt.Println(time.Now())
	go func(ch <-chan time.Time) {
		fmt.Printf("Now is %s\n", <-ch) // 等待信道接收数据，接收到之后打印当前时间
		done <- struct{}{}              // 通知主线程协程退出
	}(time.After(time.Second * 3)) // 调用After，将返回的只读信道传递给协程函数

	<-done
	afterFun()
	userTime()
	close(done)
}

func afterFun() {
	done := make(chan struct{}) // 用来等待协程结束

	go time.AfterFunc(time.Second*3, func() {
		fmt.Println("Print after 3 seconds")
		done <- struct{}{} // 通知主协程结束
	})

	fmt.Println("Print in main afterfunc")

	<-done
	close(done)
}

func userTime() {
	done := make(chan struct{}) // 用来等待协程结束

	timer := time.NewTimer(time.Second * 3)

	go func() {
		fmt.Printf("Now is %s\n", <-timer.C)
		done <- struct{}{}
	}()

	fmt.Println("Print in main usertimer")

	<-done
	close(done)
}
