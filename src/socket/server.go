package main

import (
	"bufio"
	"net"
)

// https://tonybai.com/2015/11/17/tcp-programming-in-golang/

func sconnHandler(conn net.Conn) {
	//var b []byte
	//n, err := conn.Read(b)
	//if err != nil {
	//	println(err)
	//}
	//println("read ", n, string(b))
	//conn.Write(b)

	reader := bufio.NewReader(conn)
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(line), isPrefix)

	//bufio.NewWriter(conn)

}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		println(err)
		return
	}
	for {
		// will block in listener accept until accept one connection
		conn, err := listener.Accept()
		if err != nil {
			println(err)
		}
		go sconnHandler(conn)
	}

}
