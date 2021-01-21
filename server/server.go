/*
 * @Author: your name
 * @Date: 2021-01-21 17:01:36
 * @LastEditTime: 2021-01-21 17:57:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /service/demo/server/server.go
 */
package server

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) error {
	defer conn.Close()
	var b [256]byte
	ctv := ""
	for {
		n, err := conn.Read(b[:])
		ctv += string(b[:n])
		if err == io.EOF {
			str := fmt.Sprintf("服务端返回数据: %s\n", ctv)
			conn.Write([]byte(str))
			return nil
		}
		if err != nil {
			fmt.Println("客户端读取内容失败")
		}
	}
}
func createServer() error {
	listen, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Printf("服务启动失败: %\n", err)
		return err
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("访问失败!")
			return err
		}
		go process(conn)
	}
}
func init() {
	createServer()
}
