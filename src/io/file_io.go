package main

import (
	"bufio"
	"io"
	"os"
	"unsafe"
)

func main() {
	path := "/Users/guanhao/Downloads/TODO"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	println(file.Name())

	fileInfo, err := os.Stat(path)
	println(fileInfo.Size())

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil || io.EOF == err {
			break
		}

		// 字符切片[]byte转换成string 这样会产生一次内存拷贝
		newStr := string(line)
		// 字符切片[]byte转换成string 这样二者会利用同一片内存
		oldStr := (*string) (unsafe.Pointer(&line))

		println(newStr)
		println(*oldStr)
	}

}
