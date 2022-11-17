package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("ServerStart...")
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	HandleError(err, "listen start")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		HandleError(err, "listen accept")
		proc(conn)
	}
}
func proc(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端终止连接，地址", conn.LocalAddr().String(), conn.RemoteAddr().String())
			return
		}
		HandleError(err, "Listen Read")
		fmt.Println("接收:", n, "字节数据")
		fmt.Println("内容:", string(buf[:n]))
	}
}
func HandleError(err error, s string) {
	if err != nil {
		fmt.Println(s, err)
	}
}
