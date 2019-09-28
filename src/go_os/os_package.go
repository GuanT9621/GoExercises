package main

import (
	"os"
	"runtime"
)

/**
os.Chdir()     //chdir将当前工作目录更改为dir目录．
os.Getwd()     //获取当前目录，类似linux中的pwd
os.Chmod()     //更改文件的权限（读写执行，分为三类：all-group-owner）
os.Chown()     //更改文件拥有者owner
os.Chtimes()   //更改文件的访问时间和修改时间，atime表示访问时间，mtime表示更改时间
os.Clearenv()  //清除所有环境变量（慎用）
os.Environ()   //返回所有环境变量
os.Exit()      //系统退出，并返回code，其中０表示执行成功并退出，非０表示错误并退出，其中执行Exit后程序会直接退出，defer函数不会执行．
os.Expand()    //Expand用mapping 函数指定的规则替换字符串中的${var}或者$var（注：变量之前必须有$符号）。比如，os.ExpandEnv(s)等效于os.Expand(s, os.Getenv)。
*/
func main() {
	properties := os.Environ()
	for index, value := range properties {
		println(index, value)
	}
	// 默认情况下，调度器仅使用单线程，要想发挥多核处理器的并行处理能力，必须调用
	// runtime.GOMAXPROCS(n)来设置可并发的线程数，也可以通过设置环境变量 GOMAXPROCS 达到相同的目的。
	runtime.GOMAXPROCS(5)
}
