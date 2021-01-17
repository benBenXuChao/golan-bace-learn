/*
 * @Author: your name
 * @Date: 2020-12-10 14:12:29
 * @LastEditTime: 2021-01-05 14:35:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/2_struct.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

type a struct {
	nn string
}

// Student 学生
type Student struct {
	Name string
	Li   string
	Age  int
}

func (s *Student) updateSnt(age int) *Student {
	s.Age = age
	return s
}

// Collection 这个被继承的结构体会有自己的方法
type Collection struct {
	City string `json:"city"`
}

func (c *Collection) update(city string) *Collection {
	c.City = city
	return c
}
func initpp() {

	str := "Hello, 世界"
	num := 0
	for i, val := range str {
		fmt.Printf("i:%v,val:%v\t", i, val)
		num = i
	}
	fmt.Println()
	fmt.Println("range =", num+1)                       // 错误
	fmt.Println("bytes =", len(str))                    // 错误
	fmt.Println("utf8  =", utf8.RuneCountInString(str)) // 正确
	fmt.Println("rune  =", len([]rune(str)))            // 正确

	// 使用匿名结构体
	// var p a
	// fmt.Printf("p:%#v", p)

	arr := [4]int{1, 2, 3, 4}
	slice := arr[:]
	newSlice := arr[:]
	fmt.Printf("newSlice:%#v\n", newSlice)
	fmt.Println(cap(slice[:1]))

	slice = append(slice[:2], slice[3:]...)
	newSlice = append(newSlice[1:], 9)

	fmt.Printf("arr:%#v\n", arr)
	// fmt.Printf("newSlice:%#v\n", newSlice)

	// 在结构体中定义结构体,结构体实现继承
	type Address struct {
		Privince string `json:"privince"`
		City     string `json:"city"`
	}
	// 另一个被继承的struct
	type Home struct {
		Area string `json:"area"`
		City string `json:"city"`
	}
	type Person struct {
		Name   string `json:"name"`
		Gender string `json:"gender"`
		// 使用非匿名
		// Address Address
		Address     `json:"address"`    // 匿名字段
		Home        `json:"home"`       // Home和Address同时匿名时，City字段会有冲突
		*Collection `json:"collection"` // 指针的匿名变量
	}
	type People struct {
		Name string
		Collection
	}
	coll := Collection{
		City: "",
	}
	var xiaom = Person{
		Name:   "小明",
		Gender: "男",
		Address: Address{
			Privince: "102123",
			City:     "213221",
		},
		Collection: &coll,
	}
	fmt.Printf("xiaom:%#v\n", xiaom)
	// 使用了匿名字段了以后就可以实现继承了
	// 获取Address和修改Address
	xiaom.Privince = "88888"
	fmt.Printf("xiaom:%v,City:%v\n", xiaom, xiaom.Privince)

	// 嵌套结构体的字段名冲突
	// City字段名冲突
	// xiaom.City
	xiaom.Home.City = "838271" // 有冲突的字段必须要写全，不能使用继承的方式修改和获取值
	fmt.Printf("home.city:%v\n", xiaom.Home.City)

	// 不仅字段可以继承，方法也可以继承
	xiaom.update("55555") // 这个地方修改的是collection的city
	fmt.Printf("collection.City:%s\n", xiaom.Collection.City)
	// 当匿名属性是Collection的指针时
	coll.update("uuuuuu") // 虽然是调用的coll的update方法，但是一样可以修改personal的值
	fmt.Printf("collection.City:%v\n", xiaom.Collection.City)

	// 当people并没有匿名Collection的指针，而是结构体
	col := Collection{
		City: "00000",
	}
	peo := People{
		Name:       "小刚",
		Collection: col,
	}
	peo.update("123456789")
	fmt.Printf("peo:%v\n", peo)
	col.update("bbbbbbb") // 这里就不能改变people的值
	fmt.Printf("peo:%v\n", peo)

	// 将struct结构体序列号
	json, _ := json.Marshal(xiaom)
	fmt.Printf("json:%v\n", string(json))

}

func ccc(as *[]int) {
	*as = []int{1, 2, 34}
	// as[1] = 888

}
