// tar压缩
package main

import (
	"archive/tar"
	"bytes"
	"log"
	"os"
)

func main() {
	// 创建一个缓冲区用来保存压缩文件内容
	var buf bytes.Buffer
	// 创建一个压缩文档
	tw := tar.NewWriter(&buf)

	// 定义一堆文件
	// 将文件写入到压缩文档tw
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// 将压缩文档内容写入文件 file.tar.gz
	f, err := os.OpenFile("file.tar.gz", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}
