// 对文件进行md5编码
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	md5File()
}

// md5编码字符串
func md5String() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x\n", h.Sum(nil))
	// output: e2c569be17396eca2a2e3c11578123ed

	// 直接使用md5 ew对象的Write方式也是一样的
	h2 := md5.New()
	h2.Write([]byte("The fog is getting thicker!"))
	h2.Write([]byte("And Leon's getting laaarger!"))
	fmt.Printf("%x\n", h2.Sum(nil))
	// output: e2c569be17396eca2a2e3c11578123ed
}

// md5编码文件
func md5File() {
	file, err := os.Open("md5test.log")
	if err != nil {
		panic(err)
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", h.Sum(nil))
	// output: 43c6359298645ded23f3c2ee44acf564

	// 经过io.Copy操作后，file的偏移量(seek)被指向了最后面
	// 如果还需要使用则需要修改file色偏移量(seek)
	// 该行代码输出为空，因为file的seed已经位于最后了
	io.Copy(os.Stdin, file)
	// output:

	file.Seek(0, 0)

	// 该行输出文件的内容，因为file的偏移量(seek)被设置为0了
	io.Copy(os.Stdin, file)
	// output: md5test.log

}
