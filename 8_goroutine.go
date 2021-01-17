/*
 * @Author: your name
 * @Date: 2021-01-11 13:16:38
 * @LastEditTime: 2021-01-11 13:53:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/8_goroutine.go
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个goroutine
func goroutine() {
	fmt.Println("同时进行--goroutine")
}

// CreateGoroutine 创建一个goroutine
func CreateGoroutine() {
	fmt.Println("主线开始")
	go goroutine() // 开启了一个goroutine，可以同步执行任务
	for i := 0; i < 30; i++ {
		fmt.Println("-------")
	}
	fmt.Println("主线结束")     // 不注释for循环的话有事goroutine会在主线任务结束之前打印
	time.Sleep(time.Second) // 需要确保goroutine执行结束，所以要等goroutine
}

var aw sync.WaitGroup

func goroutineSyc() {
	defer aw.Done()
	fmt.Println("使用了async")
}

// 使用time.Sleep比较傻，但又必须得等，可以使用sync，当goroutine执行完立即结束，这样不用傻等
func useSync() {
	fmt.Println("useAsync开始")
	aw.Add(1)
	go goroutineSyc()
	for i := 0; i < 50; i++ {
		fmt.Println("++++++++")
	}
	fmt.Println("useAsync结束")
	aw.Wait() // 不需要在那一直等
}
func goroutineMore() {
	defer aw.Done()
	fmt.Println("又一个goroutine")
}

// multipleSync 启动多个goroutine
func multipleSync() {
	fmt.Println("multipleSync开始")
	aw.Add(2)
	go goroutineSyc()
	go goroutineMore()
	fmt.Println("multipleSync结束")
	aw.Wait()
}
func initfdsfdsf() {
	// CreateGoroutine()
	// useSync()
	multipleSync()
}
