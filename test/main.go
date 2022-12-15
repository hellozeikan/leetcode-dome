package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WgTest(n int) {
	// 减计数
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println(n)
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My first lucky number is", rand.Intn(10))
	// 协程数

	var numGoroutine int = 20
	wg.Add(numGoroutine)
	for i := 0; i < numGoroutine; i++ {
		go WgTest(i)
	}
	// 使用wg.Wait()阻塞主协程
	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {
// 	//这里我们假设数据是int类型，缓存格式设为100
// 	dataChan := make(chan int, 100)
// 	go func() {
// 		for {
// 			select {
// 			case data := <-dataChan:
// 				fmt.Println("data:", data)
// 				time.Sleep(1 * time.Second) //这里延迟是模拟处理数据的耗时
// 			}
// 		}
// 	}()

// 	//填充数据
// 	for i := 0; i < 100; i++ {
// 		dataChan <- i
// 	}

// 	//这里循环打印查看协程个数
// 	for {
// 		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
// 		time.Sleep(2 * time.Second)
// 	}
// }
