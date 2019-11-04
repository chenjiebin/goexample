package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type User struct {
	Id   int64
	Name string
	Age  int
}

func main() {
	// 初始化用户
	user := User{1, "Tom", 18}
	fmt.Println("init user: ", user)

	// 进行序列化
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(user); err != nil {
		log.Fatal("encode error:", err)
	}
	fmt.Println("gob encode result: ", buffer)

	fmt.Println()

	// 进行反序列化
	var user2 User
	fmt.Println("define user2 for decode: ", user2)
	dec := gob.NewDecoder(&buffer)
	if err := dec.Decode(&user2); err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println("gob decode result: ", user2)
}
