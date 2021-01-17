/*
 * @Author: your name
 * @Date: 2021-01-16 21:42:50
 * @LastEditTime: 2021-01-17 15:10:15
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /demo/11_goroutine及channel题.go
 */
/**
1. 使用goroutine随机生成int64数据
2. 计算数据的各个位数的和
3. 打印这些数据在主函数
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ranOnce sync.Once
var ranWait sync.WaitGroup

type ranst struct {
	random int64
	count  int64
}

func createRandom(cnRdm chan<- int64) {
	// defer ranWait.Done()
	rand.Seed(time.Nanosecond.Nanoseconds())
	for i := 0; i < 100; i++ {
		cnRdm <- rand.Int63n(10000)
	}
	close(cnRdm)
}
func consumerAndCount(cn <-chan int64, cnCont chan<- ranst) {
	// defer ranWait.Done()
	var (
		random int64
		save   int64
	)
	for i := 0; i < 100; i++ {
		random = <-cn
		save = random
		count := int64(0)
		for random > 0 {
			count += random % 10
			random = random / 10
		}
		cnCont <- ranst{
			random: save,
			count:  count,
		}
	}
	close(cnCont)
}
func ranCountMain() {
	var (
		cnRdm  = make(chan int64, 100)
		cnCont = make(chan ranst, 100)
	)
	ranWait.Add(2)
	go createRandom(cnRdm)
	go consumerAndCount(cnRdm, cnCont)
	for val := range cnCont {
		fmt.Printf("%d计算的位数是：%d\n", val.random, val.count)
	}
	// ranWait.Wait()
}
func initsds() {
	ranCountMain()
}
