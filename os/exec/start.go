// Start和Run的的区别
// 作者：陈杰斌
// 参考地址：http://www.01happy.com/golang-os-exec/
package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "5")
	// 如果用Run，执行到该步则会阻塞等待5秒
	// err := cmd.Run()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	// Start，上面的内容会先输出，然后这里会阻塞等待5秒
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
