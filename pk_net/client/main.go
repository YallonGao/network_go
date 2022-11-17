package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var exitChan chan bool = make(chan bool, 1)

func main() {
	fmt.Println("Client Start...")
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	HandleError(err, "Dial start")
	proc(conn)
	if <-exitChan {
		fmt.Println("客户端终止连接")
	}
}
func proc(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		HandleError(err, "Dial Write")
		line = strings.Trim(line, "\r\n")
		if line == "bye" {
			exitChan <- true
			break
		}
		conn.Write([]byte(line))
		fmt.Println("发送:", len(line), "字节数据")
	}
}
func HandleError(err error, s string) {
	if err != nil {
		fmt.Println(s, err)
	}
}
