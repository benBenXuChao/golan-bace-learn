/*
 * @Author: your name
 * @Date: 2020-12-09 09:42:04
 * @LastEditTime: 2020-12-09 10:51:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/1_slice方法.go
 */
package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name string
	Age  int
}
type st []student

func (a st) Len() int {
	return len(a)
}
func (a st) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a st) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func Initg() {
	sortList := []int{8, 5, 4, 7, 9, 5, 4, 3, 6}
	// 复制
	// copyList := make([]int,10,10)
	copyList := []int{9: 10}
	fmt.Printf("co:%T\n", copyList)
	copy(copyList, sortList)
	fmt.Printf("copy:%v\n", copyList)
	// 排序
	sort.Ints(sortList)
	fmt.Printf("int类型排序:%v\n", sortList)
	// 查询是否有某个值
	people := []student{
		{"tian", 22},
		{"he", 46},
		{"qiang", 18},
		{"jie", 56},
	}
	sort.Sort(st(people))
	fmt.Printf("结构体切片排序：%v\n", people)

}
