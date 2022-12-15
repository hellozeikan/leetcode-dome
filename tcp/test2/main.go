package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client demo
// tcp_demo/client/main.go

func main() {
	// 服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2222")
	if err != nil {
		fmt.Printf("dial failed,err:%v\n", err)
		return
	}
	// 2. 利用该连接进行数据的接受和发送
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 给服务端发消息
		_, err := conn.Write([]byte(s))
		fmt.Println([]byte(s))
		if err != nil {
			fmt.Printf("send failed,err:%v\n", err)
			return
		}
		// 从服务端接收回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		fmt.Println(buf[:n])
		if err != nil {
			fmt.Printf("read failed,err:%v\n", err)
			return
		}
		fmt.Println("收到服务端回复：", string(buf[:n]))
	}
}
