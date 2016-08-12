// redis_server project main.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var kv_data map[string]string

func main() {

	kv_data = make(map[string]string)

	//建立socket，监听端口
	netListen, _ := net.Listen("tcp", ":1215")

	defer netListen.Close()

	fmt.Println("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accept a client")
		go handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {
	for {
		str := parseRESP(conn)
		switch value := str.(type) {
		case string:
			if len(value) == 0 {
				goto end
			}
			conn.Write([]byte(value))
		case []string:
			if value[0] == "SET" {
				key := string(value[1])
				val := string(value[2])
				kv_data[key] = val
				conn.Write([]byte("+OK\r\n"))
			} else if value[0] == "GET" {
				key := string(value[1])
				val := string(kv_data[key])
				val_len := strconv.Itoa(len(val))
				str := "$" + val_len + "\r\n" + val + "\r\n"
				conn.Write([]byte(str))
			} else {
				conn.Write([]byte("+OK\r\n"))
			}
			break
		default:

		}
	}
end:
	conn.Close()
}

func parseRESP(conn net.Conn) interface{} {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}
	cmd_type := string(line[0])
	cmd_txt := strings.Trim(string(line[1:]), "\r\n")

	switch cmd_type {
	case "*":
		count, _ := strconv.Atoi(cmd_txt)
		var data []string
		//"*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"
		for i := 0; i < count; i++ {
			line, _ := r.ReadString('\n')
			cmd_txt := strings.Trim(string(line[1:]), "\r\n")
			c, _ := strconv.Atoi(cmd_txt)
			length := c + 2
			str := ""
			for length > 0 {
				block, _ := r.Peek(length)
				if length != len(block) {

				}
				r.Discard(length)
				str += string(block)
				length -= len(block)
			}

			data = append(data, strings.Trim(str, "\r\n"))
		}
		return data
	default:
		return cmd_txt
	}
}
