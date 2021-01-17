/*
 * @Author: your name
 * @Date: 2020-12-04 10:12:26
 * @LastEditTime: 2020-12-07 10:40:35
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/class-1.go
 */
package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

// Tools 工具函数
func Tools() {
	fmt.Printf("\n就不说\n")
}

// 遍历字符串
func mapStr(str string, china bool) {
	if china {
		for _, val := range str {
			fmt.Printf("v:%v,c:%c,t:%v\t", val, val, val)
		}
	} else {
		for i := 0; i < len(str)-1; i++ {
			fmt.Printf("s: %v,c: %c\t", str[i], str[i])
		}
	}

}

// 替换其中一个字符
func replace(str string, num int, ace rune) string {
	arr := []rune(str)
	for i := range arr {
		fmt.Print(i)
		if i == num {
			arr[i] = ace
		}
	}
	fmt.Println()
	return string(arr) // 将[]rune转换为string
}
func initL() {
	// 测试int8最大值
	int8Max := math.MaxInt8
	fmt.Printf("int8Max: %v\n", int8Max)
	fmt.Println(int8Max == 1<<7-1)
	fmt.Println(math.MaxInt32 == 1<<31-1)
	fmt.Println(math.MinInt8)

	num := 123
	fmt.Printf("type: %T \n", num)
	fmt.Printf("e: %v\n", 1.23e12)
	fmt.Printf("e: %.5f\n", 1.23*math.E)
	fmt.Printf("我的\r 你的又没做\n 还有啥的\t 斜杠/ 反斜杠\\ \n")

	// 字符串方法
	str := "hello word"
	text := "nice gride"
	addStr := str + text
	fmt.Printf("len %v\n", len(str))
	// 字符串拼接
	var t = fmt.Sprintf("%s%s%s", str, " -- ", text)
	// 转换为buffer 再进行拼接
	var buf bytes.Buffer
	buf.WriteString(str)
	buf.WriteString(" -- ")
	buf.WriteString(text)
	bs := buf.String()
	fmt.Printf("t: %v\t bs:%v\t addStr:%v\n", t, bs, addStr)
	fmt.Println(strings.Contains(bs, "o w"))
	fmt.Println(strings.HasPrefix(t, "hello"))
	fmt.Println(strings.HasSuffix(t, "gride"))
	fmt.Println(strings.Index(t, "word"))
	fmt.Println(strings.LastIndex(t, "word"))
	arr := strings.Split(t, " ")
	fmt.Printf("arr: %v\t t: %T\tlen: %d\n", arr, arr, len(arr))
	fmt.Printf("join: %s\n", strings.Join(arr, " "))
	// 遍历一个str
	/**
	英文和数字是uint8 byte类型
	汉字，日文等是uint32 rune类型
	字符串底层byte数组，所以可以和[]byte类型相互转换
	rune类型用来表示utf8字符，一个rune字符由一个或者多个byte组成
	*/
	mapStr("asdfgg15Kkl", false)
	mapStr("你方の", true)
	fmt.Printf(replace("你好啊的", 1, '犇') + "\n")
	// 强制类型转换
	mun1 := 12344
	mun2 := float32(mun1)
	mun1 = int(mun1)
	fmt.Printf("m1:%T,m2:%T\n", mun1, mun2)
}
