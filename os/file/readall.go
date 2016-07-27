// 一次性读取文件的所有内容
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("file.log")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	// output: 你好,世界!
}
