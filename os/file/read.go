// 按字节读取文件示例
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("file.log", os.O_RDONLY, 0600)
	if err != nil {
		log.Println(err)
	}
	// 按字节读取
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// output: read 14 bytes: "你好,世界!"
}
