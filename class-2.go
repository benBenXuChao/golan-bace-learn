/*
 * @Author: your name
 * @Date: 2020-12-04 17:56:34
 * @LastEditTime: 2020-12-07 13:50:38
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/class-2.go
 */
package main

import "fmt"

func initP() {
	// 控制语句
	var slic = []string{"小米", "华为", "oppole"}
	for _, v := range slic {
		if str := v + "$$"; str == "华为$$" {
			fmt.Println(str)
		}
		// switch
		switch v {
		case "小米":
			fmt.Println(v)
			break
		case "华为":
			fmt.Println(v)
			break
		default:
			fmt.Println(v)
		}
	}
	list := [...]int{1, 2, 3, 4, 5}
	var (
		i   = 0
		num = 0
	)
	for ; i < len(list); i++ {
		num += list[i]
		switch {
		case list[i] == 1:
			fmt.Printf("list[i] == 1: %d\n", list[i])
			fallthrough
		case 1 < list[i] && list[i] < 4:
			fmt.Printf("1 < list[i] && list[i] < 4: %d\n", list[i])
			break
		case list[i] > 4:
			fmt.Printf("list[i] > 4: %d\n", list[i])
			break
		default:
			fmt.Printf("default: %d\n", list[i])
		}
	}
	fmt.Println(num)
	// 排序
	numList := [6]int{5, 4, 7, 6, 5, 7}
	// goto跳出循环，goto语句可以快速跳出循环，避免重复退出
	for i := 0; i < len(numList); i++ {
		for j := i + 1; j < len(numList); j++ {
			fmt.Println(numList[i], numList[j])
			if numList[i]+numList[j] > 11 {
				goto tag
			}
		}
	}
tag:
	fmt.Println("goto跳转了")
breakTag: // break语句可以再语句后面添加标签，表示退出某个标签对应的代码，【标签要求必须定义在对应for、switch、select的代码上，必须相邻】
	for _, val := range numList {
		for _, item := range numList {
			if item == 5 {
				continue // 如果是5,下面的不执行，并进行下次循环
			}
			println(val, item)
			if val+item > 11 {
				break breakTag

			}
		}
	}
	fmt.Println("break跳转了")
	for i := 0; i < len(numList); i++ {
		for j := i + 1; j < len(numList); j++ {
			if numList[i] < numList[j] {
				var t = numList[i]
				numList[i] = numList[j]
				numList[j] = t
			}
		}
	}
	fmt.Println(numList)
	// 打印99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
