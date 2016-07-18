// 获取完整的url路径
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto)
	// output:HTTP/1.1
	fmt.Println(r.TLS)
	// output: <nil>
	fmt.Println(r.Host)
	// output: localhost:9090
	fmt.Println(r.RequestURI)
	// output: /index?id=1

	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	fmt.Println(strings.Join([]string{scheme, r.Host, r.RequestURI}, ""))
	// output: http://localhost:9090/index?id=1
}

func main() {
	http.HandleFunc("/index", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
