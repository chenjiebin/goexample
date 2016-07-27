// 一行行读取文件内容
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("file.log", os.O_RDONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	buff := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, err := buff.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		fmt.Printf("%d line: %s", i, string(line))
		// 文件已经到达结尾
		if err == io.EOF {
			break
		}
	}
	// output: 1 line: 你好,世界!

	fmt.Println()
}
