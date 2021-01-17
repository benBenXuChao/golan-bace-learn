/*
 * @Author: your name
 * @Date: 2021-01-05 10:27:50
 * @LastEditTime: 2021-01-05 14:33:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/3_包.go
 */
package main

import (
	"fmt"

	vi "pkg.deepin.com/web/service/demo/view"
)

func initv() {
	var t = vi.Status + 1
	fmt.Printf("t: %d\n", t)
	vi.GoLan()
	personal := vi.Person{
		Name: "小明",
	}
	var lili vi.Lange
	lili = &personal
	lili.Speak("你好")
	fmt.Printf("vi.Mode: %v\n", vi.Mode)

}
