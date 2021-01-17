/*
 * @Author: your name
 * @Date: 2021-01-12 17:06:18
 * @LastEditTime: 2021-01-16 21:43:12
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/10_channel高级.go
 */
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count    int
	rwcount  int
	awt      sync.WaitGroup
	lock     sync.Mutex
	rwlock   sync.RWMutex
	iconMap  map[string]image.Image
	mun      int
	syncOnce sync.Once
)

func caseOneChan(cn chan<- int) {
	time.Sleep(time.Second) // 通过time.Sleep证实，不管等多久，select总会等到先回来的那个
	cn <- 3
}
func caseTwoChan(cn chan<- int) {
	time.Sleep(time.Second)
	cn <- 6
}

// 有一个值可以从多个通道中获取，只要其中一个返回了结果就可以了，其他的就不再等
func multChanGetOneData() {
	var cn1 = make(chan int)
	var cn2 = make(chan int)
	var data int
	go caseOneChan(cn1)
	go caseTwoChan(cn2)
	select {
	case data = <-cn1:
		fmt.Printf("cn1先回来: %d\n", data)
	case data = <-cn2:
		fmt.Printf("cn2先回来: %d\n", data)
	}
	fmt.Println(data)
}

// 并发安全，多个goroutine操作同一个对象，这样会发生意想不到的结果
func nosafyGoroutine() {
	defer awt.Done()
	for i := 0; i < 100000; i++ {
		count++
	}
}

// 使用互斥锁来解决操作同一个对象的问题
func useLock() {
	defer awt.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
}

// 数字累加
func countNum() {
	awt.Add(2)
	go nosafyGoroutine()
	go nosafyGoroutine()
	awt.Wait()
	// 结果并不是200000，因为同时操作了count这个变量
	fmt.Printf("两次goroutine累加count的数字: %d\n", count)

	count = 0
	// 使用lock之后
	awt.Add(2)
	go useLock()
	go useLock()
	awt.Wait()
	fmt.Printf("lock之后的计算的count: %d\n", count)
}

// 读锁和写锁
// 如果有多个goroutine同时对一个对象、资源读取，那么就没必要非要等一个goroutine读取之后另一个再读
// 如果有多个goroutine同时对一个对象、资源读写，那么：
// 1. 如果一个goroutine获取的是读锁，其他goroutine获取的不是读锁就可以继续进行，写锁原地等待
// 2. 如果一个goroutine获取的是写锁，其他goroutine无论是什么锁都要等待
// 使用读写锁（当读锁更多时，效果越明显）
func usewlock() {
	defer awt.Done()
	for i := 0; i < 100; i++ {
		rwlock.Lock() // 加写锁
		rwcount++
		time.Sleep(time.Millisecond) // 假设读操作耗时10毫秒
		rwlock.Unlock()
	}
}
func userlock() {
	defer awt.Done()
	for i := 0; i < 1000; i++ {
		rwlock.RLock() // 加读锁
		// rwlock.Lock()  // 依旧使用写锁
		time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
		fmt.Printf("%d\t", rwcount)
		// rwlock.Unlock()
		rwlock.RUnlock()
	}
}

// 即有读又有写
func rwlockMain() {
	start := time.Now()
	awt.Add(3)
	go userlock()
	go usewlock()
	go userlock()
	awt.Wait()
	end := time.Now()
	between := end.Sub(start).Milliseconds()
	fmt.Printf("相差毫秒数: %d\n", between) // 使用Rlock和RUnlock 1179秒，没有使用读锁2253秒

}

// 多个goroutine操作多次操作配置文件问题
func iconConf() {
	iconMap = map[string]image.Image{
		"left":   loadIcon("./image/left.jpg"),
		"right":  loadIcon("./image/right.jpg"),
		"top":    loadIcon("./image/top.jpg"),
		"bottom": loadIcon("./image/bottom.jpg"),
	}
	fmt.Println("iconConf执行了")
	// 这个地方cpu可能解释为
	/*
		iconMap = make(map[string]image.Image)
		iconMap["left"] = loadIcon("left.jpg")
		iconMap["top"] = loadIcon("top.jpg")
		iconMap["right"] = loadIcon("right.jpg")
		iconMap["bottom"] = loadIcon("bottom.jpg")
	**/
	// 这样会导致可能iconMap已经做了make操作，但变量并未初始化
}

// 加载图片
func loadIcon(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}

// 获取单个图片
func singleImg(img string, cn chan<- image.Image) {
	defer awt.Done()
	if iconMap == nil {
		iconConf() // 配置文件只需要init一次就够了，但是这个例子中init了4次，而且是操作同一个文件，但是也不能吧这句代码放在外面，因为
	}
	lock.Lock()
	mun++
	lock.Unlock()
	cn <- iconMap[img]
	if mun >= 4 {
		close(cn)
	}
}

// 改良版获取单个图片
func useSyncOnce(img string, cn chan<- image.Image) {
	defer awt.Done()
	syncOnce.Do(iconConf) // 只初始化一次，当多次goroutine执行这句，只会等第一个执行完了，不再执行（多个goroutine只需一个执行，其他人等待）
	cn <- iconMap[img]
	lock.Lock()
	mun++
	lock.Unlock()
	if mun >= 4 {
		close(cn)
	}
}
func getAllImg() {
	awt.Add(4)
	var cn = make(chan image.Image)
	// go singleImg("left", cn)
	// go singleImg("right", cn)
	// go singleImg("top", cn)
	// go singleImg("bottom", cn)
	go useSyncOnce("left", cn)
	go useSyncOnce("right", cn)
	go useSyncOnce("top", cn)
	go useSyncOnce("bottom", cn)
	for img := range cn {
		fmt.Printf("获取到的图片: %T\n", img)
	}
	awt.Wait()
}

var myMap map[string]int
var mapWait sync.WaitGroup
var mapLoc sync.Mutex
var syncMap = sync.Map{}

// 对字典进行赋值
func setMapVal(val int) {
	defer mapWait.Done()
	syncOnce.Do(func() {
		myMap = make(map[string]int) // 多个goroutine只需要初始化一次
	})
	key := strconv.Itoa(val)
	mapLoc.Lock() // 如果不加锁 fatal error: concurrent map writes
	myMap[key] = val
	mapLoc.Unlock()
}

// 设定一个字典
func initMyMap() {
	var len = 20
	for i := 0; i < len; i++ {
		mapWait.Add(1) // waitGroup是可以累加的
		// go setMapVal(i)
		go useSyncMap(i)
	}
	mapWait.Wait()
	syncMap.Range(func(key, val interface{}) bool {
		fmt.Printf("val: %v\n", val)
		return true
	})
	oneVal, _ := syncMap.Load("4")
	twoVal, ok := syncMap.LoadOrStore("5", 6)
	fmt.Printf("val: %d\n", oneVal)
	if ok {
		fmt.Printf("获取值twoVal: %d\n", twoVal)
	} else {
		fmt.Printf("更新值twoVal: %d\n", twoVal)
	}

}

// 使用go内置的sync.Map来解决对map的保护
func useSyncMap(val int) {
	defer mapWait.Done()
	key := strconv.Itoa(val)
	syncMap.Store(key, val)
}

var proWait sync.WaitGroup

// 如果是值类型，则可以使用原子来替代互斥锁和读写锁
type valueTypeOption interface {
	read() int64
	write(int64) int64
}

// lock的struct
type valueLock struct {
	number int64
	lock   sync.RWMutex
}

// read方法
func (v *valueLock) read() int64 {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.number

}

// write方法
func (v *valueLock) write(n int64) int64 {
	defer v.lock.RUnlock()
	v.lock.RLock()
	v.number = v.number + n
	return v.number
}

// atomic的struct
type valueAtomic struct {
	number int64
}

// atomic操作的读取
func (v *valueAtomic) read() int64 {
	return atomic.LoadInt64(&v.number)
}

// atomic操作的设置
func (v *valueAtomic) write(n int64) int64 {
	atomic.AddInt64(&v.number, n)
	return v.number
}

// gorountine
func rwValue(st valueTypeOption, n int64) {
	defer proWait.Done()
	fmt.Printf("设置前num: %d\n", st.read())
	fmt.Printf("设置后: %d\n", st.write(n))
}

// 对比lock锁和aotmic锁区别
func protectValue() {
	var lov = valueLock{
		number: 10,
	}
	var atv = valueAtomic{
		number: 100,
	}
	for i := 1; i < 5; i++ {
		proWait.Add(1)
		go rwValue(&lov, int64(i))
	}
	for i := 1; i < 5; i++ {
		proWait.Add(1)
		go rwValue(&atv, int64(i))
	}
	proWait.Wait()
}
func initgrt() {
	// multChanGetOneData()
	// countNum()
	// rwlockMain()
	// getAllImg()
	// initMyMap()
	protectValue()
}
