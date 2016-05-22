// 根据不同的字符串实例化不同的结构体
package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Age int
}
type Bar struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//用于保存实例化的结构体对象
var regStruct map[string]interface{}

func main() {
	str := "Bar"
	if regStruct[str] != nil {
		// 通过反射包中的Type方法获取类型
		t := reflect.ValueOf(regStruct[str]).Type()
		// 根据类型进行实例化
		v := reflect.New(t).Elem()

		// 对其中的Name进行赋值
		f := v.FieldByName("Name")
		if f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
			f.SetString("tom")
		}

		// 依次读取对象内的字段
		for i := 0; i < v.NumField(); i++ {
			valueField := v.Field(i)
			typeField := v.Type().Field(i)
			tag := typeField.Tag

			fmt.Printf("Field Name: %s\n", typeField.Name)
			fmt.Printf("Field Value: %v\n", valueField.Interface())
			fmt.Printf("Tag Value: %s\n", tag.Get("json"))
			fmt.Println("")
			// output:
			//	Field Name: Id
			//	Field Value: 0
			//	Tag Value: id
			//
			//	Field Name: Name
			//	Field Value: tom
			//	Tag Value: name
		}
	}
}

func init() {
	regStruct = make(map[string]interface{})
	regStruct["Foo"] = Foo{}
	regStruct["Bar"] = Bar{}
}
