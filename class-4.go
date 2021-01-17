/*
 * @Author: your name
 * @Date: 2020-12-07 16:20:29
 * @LastEditTime: 2020-12-09 09:39:51
 * @LastEditors: Please set LastEditors
 * @Description: 切片
 * @FilePath: /demo/class-4.go
 */
package main

func kk(a *int) func() {
	*a = 22
	return func() {

	}
}
func init() {
	defer func() {

	}()
	// 定义且赋值
	// slic := []int{1, 2, 3, 4}
	// fmt.Println(slic)
	// // 仅是定义类型
	// var s []int
	// // s[1] = 55 长度为0时不能修改值
	// s = []int{4, 5} // 后赋值
	// s[1] = 55       // 分配内存后就可以修改值

	// fmt.Println(s)
	// // 使用make
	// sl := make([]string, 0, 10)
	// // sl[0] = "88" 长度为0不能修改值
	// fmt.Println(sl)
	// sl = append(sl, "小明", "小张", "小李")
	// fmt.Println(sl)

	// // 用数组切
	// arr := [...]int{11, 22, 33, 44}
	// sli := arr[1:3] // 顾头不顾尾
	// fmt.Printf("用数组来切：%#v--len:%d--cap:%d\n", sli, cap(sli), len(sli))

	// // 用切片再切
	// slice := sli[1:2]
	// fmt.Printf("用切片切：%#v--len:%d--cap:%d\n", slice, cap(slice), len(slice))

	// // 切片操作简写
	// a := []int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("默认为0:%#v\n", a[:4])          // 1,2,3,4
	// fmt.Printf("默认为len(slice):%#v\n", a[1:]) // 2,3,4,5,6
	// fmt.Printf("都不写：%#v\n", a[:])            // 1,2,3,4,5,6
	// // 切片的完整表达
	// // a[low:high:max]
	// t := a[2:3:4] // max的取值 在灰色部分（最大值为cap）
	// fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	// // 关于切片的len和cap，以及low，high
	// origin := []int{1, 2, 3, 4, 5, 6}
	// s1 := origin[1:3]
	// fmt.Printf("s1: %#v--len：%d--cap:%d\n", s1, len(s1), cap(s1))  // 值 顾头不顾尾2,3    切片长度为 3-1=2   cap 切头留位(切掉一个就剩下5个)
	// s2 := s1[3:5]                                                  // high的最大值为cap
	// fmt.Printf("s2: %#v--len: %d--cap:%d\n", s2, len(s2), cap(s2)) // 切片长度2 值 5,6 cap 2  (从第4个开始切-->2,3,4就是5,长度2个就是5,6)(之前切掉一个，又切掉3个，还剩2个)

	// // 判断切片是否为空使用len
	// // 切片不能使用 == 来进行判断，切片是引用数据类型，只能比较地址
	// // 切片唯一合法的比较是和nil进行比较，当等于nil时，len为0,cap为0
	// // len和cap为0时不能断定切片等于nil
	// var nilSlice []string
	// fmt.Printf("只声明没有赋值的切片等于nil: %t\n", nilSlice == nil) // true
	// emptySlice := []int{}
	// fmt.Printf("做了初始化的切片不等于nil，但是cap和len都为0：%t,%d,%d\n", emptySlice == nil, cap(emptySlice), len(emptySlice)) //false,0,0
	// makeSlice := make([]int, 0)
	// fmt.Printf("使用make创建的切片，分配了内存不等于nil，cap和len也为0：%t,%d,%d\n", makeSlice == nil, cap(makeSlice), len(makeSlice)) // false,0,0

	// // 切片的增删改
	// var appendSli = make([]string, 0, 100)
	// appendSli = append(appendSli, "我", "她") // 做append操作后必须重新赋值，因为有可能会扩容，扩容后地址就变了，所以要赋值
	// fmt.Printf("push操作:%v,%p\n", appendSli, appendSli)
	// appendSli = append([]string{"你"}, appendSli...)
	// fmt.Printf("unshift操作:%v,%p\n", appendSli, appendSli)
	// // appendSli = append(appendSli[:1], 1, appendSli[0:]...)
	// appendSli = append(appendSli[:1], append([]string{"他"}, appendSli[1:]...)...)
	// fmt.Printf("insert操作:%v,%p\n", appendSli, appendSli)
	// appendSli[1] = "咱"
	// fmt.Printf("update操作:%v,%p\n", appendSli, appendSli)
	// appendSli = append(appendSli[:2], appendSli[3:]...)
	// fmt.Printf("remove操作:%v,%p\n", appendSli, appendSli)
	// appendSli = append(appendSli[:0], appendSli[1:]...)
	// fmt.Printf("shift操作:%v,%p\n\n", appendSli, appendSli)
	// //append()添加元素和切片扩容
	// // 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
	// // 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	// // 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	// // 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	// var numSlice []int
	// for i := 0; i < 10; i++ {
	// 	numSlice = append(numSlice, i)
	// 	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	// }
	// numSlice = append(numSlice, 11, 22, 33, 44, 55, 66, 77, 88, 99, 90, 98, 76, 43, 21, 54, 55, 66, 77, 88, 99, 90, 98, 76, 43, 21)
	// fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	// // copy一个slice
	// var copySlice = make([]string, 5, 5)
	// copy(copySlice, appendSli)
	// fmt.Printf("copy:%#v\n", copySlice)
	// copySlice[0] = "哈哈"
	// fmt.Printf("copySlice:%#v，appendSli：%#v\n", copySlice, appendSli)
	// //
	// var strSlice = make([]string, 5, 10)
	// for i := 0; i < 10; i++ {
	// 	strSlice = append(strSlice, fmt.Sprintf("%v", i)) // append是追加
	// }
	// fmt.Printf("strSlice:%#v", strSlice)

	// s := []int{5, 2, 6, 3, 1, 4} // unsorted
	// sort.Ints(s)
	// fmt.Println(sort.IntsAreSorted(s)) // 测试是否是按升序排序的
	// fmt.Println(s)

	// a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	// x := 6
	// i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	// if a[i] == x {
	// 	fmt.Printf("index: %d", i)
	// }

}
