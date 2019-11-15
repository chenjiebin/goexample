// tar解压缩演示
// 这边演示一下从源文件进行解压，然后输出文件内容
// 作者：陈杰斌
// 参考地址：http://www.01happy.com/golang-tar/
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 解压需要使用tar.NewReader方法, 这个方法接收一个io.Reader对象
	// 那边怎么从源文件得到io.Reader对象呢？
	// 这边通过os.Open打开文件,会得到一个os.File对象，
	// 因为他实现了io.Reader的Read方法，所有可以直接传递给tar.NewReader
	file, err := os.Open("file.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// tar对象读取文件内容, 遍历输出文件内容
	tr := tar.NewReader(file)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s文件内容:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
