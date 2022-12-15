package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go message(ch1, ch2)
	go WriteTxt(ch1, ch2)
	time.Sleep(1000 * time.Second)
}

func message(ch1 chan int, ch2 chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		ch1 <- r.Intn(10000)
		ch2 <- r.Intn(10000)
	}
}

func WriteTxt(ch1 chan int, ch2 chan int) {
	file, err := os.OpenFile("J:\\king\\file\\a.docx", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件
	defer file.Close()
	for {
		select {
		case v := <-ch1:
			_, err := file.WriteString(strconv.Itoa(v) + "x" + "\r\n")
			if err != nil {
				log.Fatal(err)
			}
		case v := <-ch2:
			_, err := file.WriteString(strconv.Itoa(v) + "q" + "\r\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
