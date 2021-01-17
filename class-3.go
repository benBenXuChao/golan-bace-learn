/*
 * @Author: your name
 * @Date: 2020-12-07 13:51:48
 * @LastEditTime: 2020-12-07 16:20:39
 * @LastEditors: Please set LastEditors
 * @Description: 数组
 * @FilePath: /demo/class-3.go
 */
package main

import (
	"fmt"
)

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}
func findTwo(arr []int) (a int, b int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				a = i
				b = j
				goto br
			}
		}
	}
br:
	fmt.Printf("(%d, %d)\n", a, b)
	return a, b
}
func inito() {
	// 定义一个数组
	// 1. 定义赋值
	var arr = [3]int{1, 2, 3}
	fmt.Printf("arr: %#v\n", arr)
	// 使用new关键字
	newArr := new([3]int)
	newArr[1] = 4
	fmt.Println(*newArr)
	// 先定义，后赋值
	var typeArr [3]string
	typeArr[1] = "hi"
	fmt.Println(typeArr)
	// 两个不同len的数组不能赋值操作
	arr2 := [4]int{5, 6, 7, 8}
	// arr = arr2 发生错误[3]int和[4]int是两个不同的类型
	fmt.Println(arr2)

	// 自行推断数组长度
	auto := [...]byte{'a', 'b', 'c'}
	fmt.Printf("auto: %s\n", string(auto[:]))
	// 指定索引赋值，用于省略中间的值
	at := [...]int{1: 99, 4: 66, 7: 88}
	fmt.Println(at)
	for i, val := range arr {
		fmt.Printf("%d:--%p\t", i, &val)
	}
	fmt.Printf("\narr的地址：%p\n", &arr)

	// 多维数组
	var falt = [2][3]int{
		0: {2: 10},
		1: {1: 20},
	}
	// 修改多维数组
	falt[0][0] = 12
	for _, val := range falt {
		for _, item := range val {
			fmt.Printf("%d\t", item)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...
	a := [...][]string{
		{"小明", "小李", "小刚"}, // 数组使用{}包裹，不再是一个int或者string，用{}代表数组
		{"nice", "bade", "good"},
	}
	for _, val := range a {
		for _, item := range val {
			fmt.Printf("%s\t", item)
		}
		fmt.Println()
	}
	//不支持多维数组的内层使用...
	/**
	b := [3][...]string{ // invalid use of '...'
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(b)
	*/

	// 数组是值类型，复制的时候是值复制
	arr3 := arr
	fmt.Printf("变更前：%t\n", arr3 == arr) // 数组是支持 == != [他是值类型]
	arr3[1] = 77
	fmt.Println(arr, arr3)
	fmt.Printf("变更后：%t\n", arr3 == arr)
	falt1 := falt
	falt1[0][1] = 999
	fmt.Println(falt1, at)
	// 定义一个指针数组
	var (
		n1 = 1
		n2 = 2
	)
	cpt := [2]*int{
		&n1,
		&n2,
	}

	*cpt[0] = 33
	fmt.Printf("修改了指针数组的值n1：%d\n", n1)
	// 求和方法
	var add = [...]int{1, 3, 5, 7, 8}
	fmt.Printf("求和：%d\n", sum(add[:]))
	var two = [...]int{1, 3, 5, 7, 8}
	fmt.Println(findTwo(two[:]))
}
