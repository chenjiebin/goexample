// 判断文件是否存在
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat("file.log")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalln(err)
	}
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return
	}
	fmt.Println("文件存在")
}
