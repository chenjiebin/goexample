// json解码
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json格式字符串
	var jsonUsers = []byte(`[
		{"id": "1", "name": "Anny"},
		{"id": "2", "name": "Tom"}
	]`)

	// 用结构体User进行接收
	// Id属性后的``内的内容表示转化json数据时的一个设定
	// 例如：`json:"id,string"`，表示对应的json格式字符串键值为id，类型为字符串
	// `json:"name"`，表示对应的json格式字符串键值为name
	type User struct {
		Id   int    `json:"id,string"`
		Name string `json:"name"`
	}

	// 因为json格式字符串是个数组格式，这里也要定义数组
	var users []User
	err := json.Unmarshal(jsonUsers, &users)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", users)
	// output: [{Id:1 Name:Anny} {Id:2 Name:Tom}]
}
