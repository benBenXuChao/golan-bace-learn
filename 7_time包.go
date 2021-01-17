/*
 * @Author: your name
 * @Date: 2021-01-07 15:51:51
 * @LastEditTime: 2021-01-12 13:58:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/7_time包.go
 */
package main

import (
	"fmt"
	"time"
)

// 获取当前时间并转换为年月日
func getTime(now time.Time) {
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 获取毫秒
func getUnixSecond() {
	now := time.Now()
	unixs := now.Unix()
	milli := time.Now().UnixNano() / 1e6
	mic := time.Now().UnixNano() / 1e3
	nanos := now.UnixNano()
	fmt.Printf("now: %s\n", now)
	fmt.Printf("当前  秒数: %d\n", unixs)
	fmt.Printf("当前毫秒数: %d\n", milli)
	fmt.Printf("当前微秒数: %d\n", mic)
	fmt.Printf("当前纳秒数: %d\n", nanos)
}

// 根据当前时间戳获取时间
func getTimeByUnix(t int64) {
	times := time.Unix(t, 0)

	fmt.Printf("times: %s\n", times)

}

// 间隔年月天
func betweenDate(y, m, d int) {
	now := time.Now()
	tim := now.AddDate(y, m, d)
	getTime(tim)

}

// 间隔时、分、秒
func betweenTime(sw time.Duration) {
	now := time.Now()
	later := now.Add(sw)
	getTime(later)
}

// 计算两个时间
func reduce(t1, t2 time.Time) {
	hr := t1.Sub(t2) / time.Hour
	fmt.Printf("相差: %d小时\n", hr)

}

// sleep 定时器（settimeout）
func sleepDem(t time.Duration, fn func()) {
	time.Sleep(t)
	fn()
}

// tickDem 定时器（setinterval）
func tickDem(t time.Duration) {
	tick := time.Tick(time.Millisecond * 100)
	now := time.Now().Add(t)
	for i := range tick {
		fmt.Printf("i: %s\n", i)
		if i.After(now) {
			break
		}
	}
	fmt.Println("setinterval结束")
}

// 设置一个时间
func setTime(format string, setTime string) {
	tim, err := time.ParseInLocation(format, setTime, time.Local) // format的格式是用来解析setTime的
	if err == nil {
		fmt.Printf("设置一个时间并返回格式化: %s\n", tim.Format("2006/01/02 15-04-05")) // 这个format是用来转换tim的
	} else {
		fmt.Println("解析失败")
	}
}

// 将时间戳转换为固定格式
func unixTime(t int64) {
	tmu := time.Unix(t, 0)
	fmt.Printf("秒时间戳转换为固定格式: %s\n", tmu.Format("2006-01-02 15:04:05"))
}
func initsdc() {
	// getTime(time.Now())
	// getUnixSecond()
	// getTimeByUnix(time.Now().Unix())
	// betweenTime(time.Hour * 1)
	// betweenDate(1, 1, 1)
	// reduce(time.Date(2020, 1, 3, 12, 0, 0, 0, time.Local), time.Date(2020, 1, 2, 19, 0, 0, 0, time.Local))
	// tickDem(time.Second * 1)
	// sleepDem(time.Second*1, func() {
	// 	fmt.Println("后执行")
	// })
	// fmt.Println("先执行")
	// ts := time.Now().Format("2006/01/02 15:04:05")
	// hs := time.Now().Format("2006/01/02 03:04:05")
	// fmt.Printf("格式化: %s\n", ts)
	// fmt.Printf("格式化: %s\n", hs)
	setTime("2006年01月02日15时04分05秒", "2020年01月12日11时26分23秒")
	unixTime(time.Now().Unix())
}
