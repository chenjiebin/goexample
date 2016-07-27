// 创建文件
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.log")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(file)
}
