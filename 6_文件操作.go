/*
 * @Author: your name
 * @Date: 2021-01-06 17:26:44
 * @LastEditTime: 2021-01-07 13:42:17
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/6_文件操作.go
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 变量
const (
	A = 1 << iota // 人
	B = 1 << iota // 系
	C = 1 << iota // LA
)

func countfn(status int) { // A人 B系
	switch status {
	case A + B:
		fmt.Println("都通过")
	case A:
		fmt.Println("人工通过，系统没通过")
	case B:
		fmt.Println("系通过，人工没通过")
	default:
		fmt.Println("都没通过")
	}
}
func heji(a int, b int, c int) {
	switch a & b & c {
	case A:
		fmt.Println("A")
	case B:
		fmt.Println("B")
	case C:
		fmt.Println("C")
	case A + B:
		fmt.Println("A+B")
	case A + C:
		fmt.Println("A+C")
	case B + C:
		fmt.Println("B+C")
	case A + B + C:
		fmt.Println("A+B+C")

	}
}

// 文件读取和关闭
func catFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	var bts = make([]byte, 255)
	num, err := file.Read(bts)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了: %d个字节\n", num)
	fmt.Println(string(bts[:num]))
}

// 循环读取文件（可能会乱码）
func cycleRead(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件读取失败")
		return err
	}
	defer file.Close()
	var reads = []byte{}
	for sl := make([]byte, 128); ; { // 128 当比较小时，会出现乱码的情况，但也不能设置无限大
		n, err := file.Read(sl)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			return err
		}
		reads = append(reads, sl[:n]...)
		fmt.Println(string(reads), ";")
	}
	return nil
}

// 缓存读取文件(中间缓存buff的形式)
func bufferRead(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	var strs = []string{}
	for {
		st, err := buf.ReadString('\n')
		if err == io.EOF {
			fmt.Println("over")
			break
		}
		a := []rune(st)
		a = append(a[:len(a)-1], ';', a[len(a)-1])
		strs = append(strs, string(a))
	}
	fmt.Printf("最后的数组: %v\n", strs)
	return nil
}

// 读取以字符串形式返回
func readForStr(path string) error {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("文件读取到:\n %s\n", str)
	return nil
}

// 创建文件并添加值
func createFile(path string, by []byte) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	_, noOK := file.Write(by)
	// file.WriteString("你好呀")
	if noOK != nil {
		return noOK
	}
	return nil
}

// writeFile
func writeFile(path string, con string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	bfo := bufio.NewWriter(file)
	con = strings.Replace(con, "\n", "$$\n", len(con))
	arr := strings.Split(con, "$$")
	for _, v := range arr {
		bfo.WriteString(v)
	}
	bfo.Flush()
	return nil
}

// ioutil WRITE
func ioWrite(path string, con string) error {
	err := ioutil.WriteFile(path, []byte(con), 0666)
	if err != nil {
		return err
	}
	return nil
}

// 实现文件copy
func copyFile(opath string, cpath string) error {
	ofile, err := os.Open(opath)
	if err != nil {
		return err
	}
	defer ofile.Close()
	cfile, err := os.OpenFile(cpath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	io.Copy(cfile, ofile)
	return nil
}
func initqa() {
	fmt.Printf("1: %b\n", os.O_CREATE)
	fmt.Printf("2: %b\n", os.O_TRUNC)
	fmt.Printf("3: %b\n", os.O_RDWR)
	// fmt.Printf(""os.O_CREATE|os.O_TRUNC|os.O_RDWR)
	fmt.Printf("4: %b\n", os.O_CREATE|os.O_TRUNC|os.O_RDWR)

	catFile("./config.ini")
	cycleRead("./word.txt")
	bufferRead("./word.txt")
	readForStr("./word.txt")
	createFile("create.txt", []byte("hello\nword\nbyby\n明天\n"))
	writeFile("create.txt", "我和你\n在一起\n心连心\n")
	ioWrite("create.txt", "今日有酒\t今朝醉\n明日有酒\t看黄昏\n")
	copyFile("./word.txt", "create.txt")
	countfn(C | B)
	heji(A, A|B, A|B|C)
}
