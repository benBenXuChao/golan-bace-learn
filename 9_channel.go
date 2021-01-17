/*
 * @Author: your name
 * @Date: 2021-01-12 10:47:55
 * @LastEditTime: 2021-01-12 17:56:28
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/9_channel.go
 */
package main

import (
	"fmt"
	"sync"
)

// 消费channel
func consuChan(cn chan int) {
	mun, err := <-cn // err 可写可不写
	if err == true {
		fmt.Printf("channel传过来的值: %d\n", mun)
	}

}

// 多次消费channel
func multChan(cn chan int) {
	m1 := <-cn
	fmt.Printf("第一次消费cn: %d\n", m1)
	m2 := <-cn
	// 当主程序执行完，goroutine后面的代码就不会执行（下面的这句有一定概率执行，取决于主程序是不是先于这个代码）
	fmt.Printf("第二次消费cn: %d\n", m2)
}

// 创建一个channel 管道一般和goroutine配合使用
func createChan() {
	// 先定义，然后make初始化
	var cn chan int
	fmt.Printf("当cn没有make时，是空指针cn: %v\n", cn)
	cn = make(chan int) // 如果没有初始化就开始chan就开始发送（接受也一样），会deadlock
	go consuChan(cn)
	cn <- 1

	// 直接赋值
	cnn := make(chan int)
	go consuChan(cnn)
	cnn <- 5

	// 使用完channel后需要手动关闭
	close(cn)
	// cn <- 2 // panic: send on closed channel 通道关闭后就不可以再发送了，报panic

	// 多次消费channel
	go multChan(cnn)
	cnn <- 10
	cnn <- 15
}

// 有缓存管道
func remmberChan() {
	var cn = make(chan int, 2)
	cn <- 3
	rmcn := <-cn // 同步情况下可以取到
	fmt.Printf("从channel缓存中取数据: %d\n", rmcn)

	// 先存先取，后存后去
	cn <- 6
	cn <- 9
	// cn <- 11 // 最多存2个超过两个直接locked
	rmcn1 := <-cn
	rmcn2 := <-cn
	fmt.Printf("先取6: %d\n", rmcn1)
	fmt.Printf("后取9: %d\n", rmcn2)
	cn <- 12
	close(cn)
	rmcn = <-cn
	fmt.Printf("即使channel关闭了也可以从缓存中取数据: %d\n", rmcn)

}

// 向主程序发送值
func sendMainChan(cn chan int) {
	defer await.Done()
	cn <- 7
	fmt.Printf("发送了一个channel: %d\n", cn)
}

// 连个channel相互统信
func acceptOtherChan(cn chan int) {
	defer await.Done()
	mun := <-cn
	fmt.Printf("接收到另一个channel发过来的值: %d\n", mun)

}

// 定义waitGroup
var await sync.WaitGroup
var at sync.WaitGroup

// 主程序
func mainChan() {
	await.Add(3)
	var cn = make(chan int)
	go sendMainChan(cn)
	mun := <-cn
	fmt.Printf("主程序接受goroutine发过来的channel: %d\n", mun)
	go sendMainChan(cn)
	go acceptOtherChan(cn)
	await.Wait()
}

// 单向通道（acceptOtherChan函数是接受chan，sendMainChan函数是发送chan）
// 改良send
func sendChan(cn chan<- int) {
	defer at.Done()
	cn <- 4
	fmt.Println("单向channel发送值")
}

// 改良accept
func receive(cn <-chan int) {
	defer at.Done()
	m := <-cn
	fmt.Printf("单向channel接受值: %d\n", m)

}

// 单向通道main函数
func hasDirectiveChan() {
	var cn = make(chan int)
	at.Add(2)
	go sendChan(cn)
	go receive(cn)
	at.Wait()
}

// range发送
func rangeSend(cn chan<- int) {
	for i := 0; i < 10; i++ {
		cn <- i + 10
	}
	close(cn) // 必须要close通道
}

// range转换
func rangeTransfor(cn <-chan int, tcn chan<- int) {
	for i := range cn {
		tcn <- i
	}
	close(tcn)
}

// range接受
func rangeReceive(tcn <-chan int) {
	for i := range tcn {
		fmt.Printf("i: %d\n", i)
	}

}

// 通过for range来获取channel连续发过来的值（使用了for range必须要close通道）
func rangeChan() {
	var cn = make(chan int)
	var tcn = make(chan int)
	go rangeSend(cn)
	go rangeTransfor(cn, tcn)
	rangeReceive(tcn)
	fmt.Println()

}
func initfsd() {
	createChan()
	// remmberChan()
	// mainChan()
	// rangeChan()
}
