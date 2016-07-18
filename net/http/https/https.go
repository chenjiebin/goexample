// 演示https服务的搭建
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Proto)
	fmt.Println(r.TLS)
	fmt.Println(r.Host)
	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)

	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
