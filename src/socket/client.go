package main

import (
	"net"
	"time"
)

func connHandler(conn net.Conn) {
	defer conn.Close()
	b := []byte("hello world")
	conn.Write(b)
	println(string(b))

	//writer := bufio.NewWriter(conn)
	//nn, err := writer.WriteString("hello world\r\n")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(nn)
}

func main() {
	// net.DialTimeout("tcp", ":8888", 2 * time.Second)
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		println(err)
		return
	}
	go connHandler(conn)
	time.Sleep(time.Second * time.Duration(3))
}
