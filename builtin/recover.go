// recover是golang内建函数，用来捕获程序中的异常，避免程序退出
// 作者：陈杰斌
// 参考地址：http://www.01happy.com/golang-builtin-recover/
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10000
	for j := 0; j < 3; j++ {
		// 使用多协程处理，其中可以预见的是除数为0会抛出异常
		go divide(i, j)
	}

	// 为了保证前面线程运行完，这里休眠一下
	for {
		time.Sleep(1 * time.Second)
	}
}

func divide(i, j int) {
	// 定义recover方法，在后面程序出现异常的时候就会捕获
	defer func() {
		if r := recover(); r != nil {
			// 这里可以对异常进行一些处理和捕获
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println(i / j)
}
