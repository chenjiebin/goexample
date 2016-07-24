// 执行系统命令
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	exec.
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	cmd := exec.Command("ls", "-a", "-l")
	// 重定向标准输出、标准错误输出、标准输入
	// 这里都是重定向到终端(terminal)
	// 也可以重定向到文件、管道这类支持IO的类型
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	// 执行命令
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
