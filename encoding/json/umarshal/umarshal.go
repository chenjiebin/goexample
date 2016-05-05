// json解码
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonUsers = []byte(`[
		{"id": "1", "name": "Anny"},
		{"id": "2", "name": "Tom"}
	]`)
	type User struct {
		Id   int    `json:"id,string"`
		Name string `json:"name"`
	}
	var users []User
	err := json.Unmarshal(jsonUsers, &users)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)
}
