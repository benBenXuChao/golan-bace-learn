/*
 * @Author: your name
 * @Date: 2020-12-09 10:48:25
 * @LastEditTime: 2020-12-10 14:04:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/书籍管理系统.go
 */
package main

import (
	"fmt"
	"os"
	"time"
)

type book struct {
	title      string
	author     string
	publish    bool
	updateTime time.Time
}

var (
	index  *int
	manage = []book{}
)

type op int

const (
	find op = iota
	update
	remove
)

func getScan(bk *book) int {
	fmt.Print("请输入书名：")
	fmt.Scanln(&bk.title)
	num := findBook(*bk)
	if *index == 3 || *index == 4 {
		if num < 0 {
			fmt.Printf("没有找到名为《%s》的书籍：\n", bk.title)
			return num
		}
		if *index == 3 {
			fmt.Print("请输入新的书名：")
			fmt.Scan(&bk.title)
		}
	} else if *index == 1 && num > -1 {
		fmt.Printf("库里已存在名为《%s》的书籍：\n", bk.title)
		return num
	}
	fmt.Print("请输入作者：")
	fmt.Scan(&bk.author)
	fmt.Print("请输入是否上架：")
	fmt.Scan(&bk.publish)
	bk.updateTime = time.Now()
	return num
}
func addBook(b book) {
	manage = append(manage, b)
}
func lookBooks() {
	fmt.Println("图书馆的所有书籍：")
	for _, val := range manage {
		fmt.Printf("书名：%s，作者：%s，是否上架：%t,更新时间：%s\n", val.title, val.author, val.publish, val.updateTime)
	}
}
func findBook(bk book) int {
	for i, val := range manage {
		if val.title == bk.title {
			return i
		}
	}
	return -1
}

func initc() {
	for {
		fmt.Printf(`
欢迎来到图书管理系统
1. 添加书籍
2. 查看所有书籍
3. 修改书籍
4. 删除书籍
5. 退出
`)
		var check int
		index = &check
		bk := book{}
		fmt.Print("请选择：")
		fmt.Scanln(&check)
		find := -1
		switch check {
		case 1:
			find = getScan(&bk)
			if find == -1 {
				addBook(bk)
			}
			break
		case 2:
			lookBooks()
			break
		case 3:
			find = getScan(&bk)
			if find > -1 {
				manage[find] = bk
			}
			break
		case 4:
			find = getScan(&bk)
			if find > -1 {
				manage = append(manage[:find], manage[find+1:]...)
			}
			break
		case 5:
			os.Exit(0)
			break
		default:
			fmt.Println("输入错误")
		}

	}
}
