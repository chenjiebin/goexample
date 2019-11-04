// 演示golang中随机数
// 作者：陈杰斌
// 参考地址：
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())

	// rand.Int函数随机一个
	fmt.Println(rand.Int())

	// rand.Intn随机[0, n)
	fmt.Println(rand.Intn(100))
}
