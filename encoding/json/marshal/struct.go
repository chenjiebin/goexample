// json编码
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type ColorGroup struct {
		ID     int      `json:"id,string"`
		Name   string   `json:"name,omitempty"`
		Colors []string `json:"colors"`
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// 如果没有设置Name属性值，则在编码成json的时候会忽略Name属性
	group = ColorGroup{
		ID:     1,
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// 如果没有设置Colors值，因为没有omitempty属性，会输出nil
	group = ColorGroup{
		ID:   1,
		Name: "Reds",
	}
	b, err = json.Marshal(group)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
