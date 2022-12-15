package main

import (
	"fmt"
	"net"
)

func main() {
	data := []byte("[这里才是一个完整的数据包]")
	fmt.Println(len(data))
	// conn, err := net.DialTimeout("tcp", "localhost:4044", time.Second*30)
	raddr, err := net.ResolveTCPAddr("tcp", "localhost:4044")
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	conn, err := net.DialTCP("tcp", nil, raddr)
	conn.SetNoDelay(true)
	for i := 0; i < 1000; i++ {
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
