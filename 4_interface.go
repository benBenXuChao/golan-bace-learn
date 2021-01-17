/*
 * @Author: your name
 * @Date: 2021-01-05 14:01:20
 * @LastEditTime: 2021-01-06 17:56:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/4_interface.go
 */
package main

import "fmt"

type Irun interface {
	run()
}
type Italk interface {
	talk()
}
type people struct {
	Name string
}

// 接口继承
type animal interface {
	Irun
	Italk
}
type any = interface{}
type mnt = int

func (p *people) run() {
	fmt.Printf("p.Name: %s会跑\n", p.Name)
}
func (p people) talk() {
	fmt.Printf("p.Name: %s会叫\n", p.Name)
}

func sun(a any) {

	if b, ok := a.(bool); ok {
		fmt.Printf("b: %v\n", b)
	}
	if b, ok := a.(string); ok {
		fmt.Printf("b: %v\n", b)
	}
	if b, ok := a.(mnt); ok {
		fmt.Printf("b: %v\n", b)
	}
	if b, ok := a.(student); ok {
		fmt.Printf("b: %v\n", b)
	}

}
func initcf() {
	var person people = people{
		"李青",
	}
	var (
		p1 Irun
		p2 Italk
		p3 animal
	)
	// people同时实现了两个接口
	p1 = &person
	p2 = person // 结构体和指针都行
	p3 = &person
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p3: %v\n", p3)

	// any类型的作用
	// 1. 传参数可以是任意类型
	const bl = true
	sun(bl)
	sun("你好")
	sun(12)
	sun(student{Name: "jack"})
	// 值是任意类型的字典

	var mp map[string]any
	mp = make(map[string]any)
	mp["name"] = "jack"
	mp["age"] = 32
	mp["isMoney"] = true
	fmt.Printf("mp: %v\n", mp)
	age, ok := mp["age"]
	if ok {

		fmt.Printf("age: %v\n", age)
	}

}
