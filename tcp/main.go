package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

// tcp 转发
func main() {
	//源端口，目的端口
	var fromport int = 4000
	var toport int = 8000
	fromaddr := fmt.Sprintf("127.0.0.1:%d", fromport)
	toaddr := fmt.Sprintf("127.0.0.1:%d", toport)

	fromlistener, err := net.Listen("tcp", fromaddr)

	if err != nil {
		log.Fatalf("Unable to listen on: %s, error: %s\n", fromaddr, err.Error())
	}
	defer fromlistener.Close()

	for {
		fromcon, err := fromlistener.Accept()
		if err != nil {
			fmt.Printf("Unable to accept a request, error: %s\n", err.Error())
		} else {
			fmt.Println("new connect:" + fromcon.RemoteAddr().String())
		}
		//这边最好也做个协程，防止阻塞
		toCon, err := net.Dial("tcp", toaddr)
		if err != nil {
			fmt.Printf("can not connect to %s\n", toaddr)
			continue
		}

		go handleConnection(fromcon, toCon)
		go handleConnection(toCon, fromcon)
	}

}

func handleConnection(r, w net.Conn) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 1<<20)
	for {
		n, err := r.Read(buffer)
		fmt.Println(bytes.NewBuffer(buffer[:n]).String())
		if err != nil {
			break
		}

		_, err = w.Write(buffer[:n])
		if err != nil {
			break
		}
	}

}
