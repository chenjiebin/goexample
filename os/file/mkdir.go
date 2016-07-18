// golang创建目录
package main

import (
	"log"
	"os"
)

func main() {
	// 创建当个目录
	err := os.Mkdir("tmp", 0755)
	if err != nil {
		log.Fatalln(err)
	}

	// 递归创建目录
	err = os.MkdirAll("tmp/tmp1/tmp2", 0755)
	if err != nil {
		log.Fatalln(err)
	}
}
