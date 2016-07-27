// 测试写入文件
package main

import (
	"log"
	"os"
)

func main() {
	// 打开一个文件
	// os.O_CREATE 表示文件不存在就会创建
	// os.O_APPEND 表示以追加内容的形式添加
	// os.O_WRONLY 表示只写模式
	// os.O_RDONLY 表示只读模式
	// os.O_RDWR 表示读写模式
	// os.O_EXCL used with O_CREATE, file must not exist
	// os.O_SYNC I/O同步的方式打开
	// os.O_TRUNC if possible, truncate file when opened.
	file, err := os.OpenFile("file.log", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	// 写入字节流
	n, err := file.Write([]byte("你好"))
	if err != nil {
		log.Fatalln(err)
	}
	// 写入字符串
	m, err := file.WriteString(",世界")
	if err != nil {
		log.Fatalln(err)
	}
	// 在指定的偏移处(offset)写入内容
	_, err = file.WriteAt([]byte("!"), int64(n+m))
	if err != nil {
		log.Fatalln(err)
	}
}
