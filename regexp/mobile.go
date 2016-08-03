// 匹配手机号码
package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	s := []string{"18505921256", "13489594009", "12759029321"}
	for _, v := range s {
		fmt.Println(rgx.MatchString(v))
	}
	// output:
	//	true
	//	true
	//	false
}
