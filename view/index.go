/*
 * @Author: your name
 * @Date: 2021-01-05 10:22:45
 * @LastEditTime: 2021-01-05 14:32:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/view/index.go
 */
package view

import "fmt"

// 状态
const (
	Status = 1
	Park   = "jiang"
)

// 变量
var (
	Mode = "vie"
)

// GoLan 外部
func GoLan() {
	fmt.Println("开始了！")
}

// Person 人结构体
type Person struct {
	Name string `json:"name"`
}

// Speak 说
func (p *Person) Speak(s string) {
	fmt.Printf("p.name: %s说%s！\n", p.Name, s)
}

// Lange 语言
type Lange interface {
	Speak(string)
}
