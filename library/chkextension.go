// 检测文件后缀名

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ChkExtension("/data/upload/log/2016/04/13/1ea00428e2a1043f69ca1a01d58b0380.zip", "zip"))
}

// 检测文件后缀名
func ChkExtension(path string, e string) bool {
	return strings.LastIndex(path, "."+e)+len("."+e) == len(path)
}
