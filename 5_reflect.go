/*
 * @Author: your name
 * @Date: 2021-01-05 16:59:23
 * @LastEditTime: 2021-01-17 19:05:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /demo/5_reflect.go
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type man struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Hobby  string  `json:"hobby"`
	Salary float32 `json:"salary"`
}

func (m man) Speak() {
	fmt.Printf("name: %s\n", m.Name)
}

type speaker interface {
	Speak()
}
type aliInt = int32
type mInt int32

// type any interface{}

// GetType 获取类型
func GetType(arg interface{}) (tp, kd string) {
	t := reflect.TypeOf(arg)
	fmt.Printf("type: %s, kind:%s\n", t.Name(), t.Kind())
	return fmt.Sprintf("%s", t.Name()), fmt.Sprintf("%s", t.Kind())
}

// GetValue 获取原始值
func GetValue(arg any) (val any) {
	v := reflect.ValueOf(arg)
	k := v.Kind()

	switch k {
	case reflect.Float32:
		val = float32(v.Float())
	case reflect.Float64:
		val = float64(v.Float())
	case reflect.String:
		val = string(v.String())
	case reflect.Int:
		val = int(v.Int())
	case reflect.Int32:
		val = int32(v.Int())
	case reflect.Bool:
		val = bool(v.Bool())
	}
	fmt.Printf("val: %v, kind: %v\n", val, k)

	return
}

// 修改值
func updateVal(arg any) {
	// old := *arg.(prt)
	v := reflect.ValueOf(arg)
	k := v.Elem().Kind()
	fmt.Println(k)

	switch k {
	case reflect.Float32:
	case reflect.Float64:
		v.Elem().SetFloat(100.21)
		tp := v.Elem().Type().Name()
		fmt.Printf("v的类型: %s\n", tp)

	case reflect.String:
		v.Elem().SetString("liliya")
	case reflect.Int:
	case reflect.Int32:
		v.Elem().SetInt(121)
	case reflect.Bool:
		v.Elem().SetBool(false)
	}
}

// 判断一个接口值是否为空
func isnil(arg any) bool {
	return reflect.ValueOf(arg).IsNil()
}

// 判断接口是否有字段或者方法
func hasProtype(arg any, field string) bool {
	v := reflect.ValueOf(arg)
	_, kind := GetType(arg)
	switch kind {
	case "struct":
		if ok := v.FieldByName(field).IsValid(); ok {
			return true
		}
		if ok := v.MethodByName(field).IsValid(); ok {
			return true
		}
	case "map":
		if ok := v.MapIndex(reflect.ValueOf(field)).IsValid(); ok {
			return true
		}
	}
	return false
}

// 获取struct结构体的key集合
func getAllKyes(sct any, sl *[]any) {
	t := reflect.TypeOf(sct)
	num := t.NumField()
	if num > 0 {
		for i := 0; i < num; i++ {
			std := t.Field(i)
			*sl = append(*sl, std.Name)
			// *sl = append(*sl, [...]any{std.Name, std.Type, std.Index, std.Tag.Get("json")})
		}
	}
}

// 获取struct结构体的值集合
func getAllValues(sct any, sl *[]any) {
	v := reflect.ValueOf(sct)
	num := v.NumField()
	if num > 0 {
		for i := 0; i < num; i++ {
			std := v.FieldByIndex([]int{i})
			var val any
			switch std.Kind() {
			case reflect.Float32:
				val = float32(std.Float())
			case reflect.Float64:
				val = float64(std.Float())
			case reflect.String:
				val = string(std.String())
			case reflect.Int:
				val = int(std.Int())
			case reflect.Int32:
				val = int32(std.Int())
			case reflect.Bool:
				val = bool(std.Bool())
			}
			*sl = append(*sl, val)
		}
	}
}

// 获取所有方法，并调用
func getAllFunc(arg any) {
	v := reflect.ValueOf(arg)
	num := v.NumMethod()
	if num > 0 {
		for i := 0; i < num; i++ {
			v.Method(i).Call([]reflect.Value{})
		}
	}
}

// conf 发的
type conf struct {
	Env  string `json:"env" ini:"ENV"`
	Path string `json:"path" ini:"PATH"`
	Num  string `json:"num" ini:"NUM"`
}

// ReadConfig 实现读取配置文件
func ReadConfig(path string, cner interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	var reft = func(slice []string) {
		t := reflect.TypeOf(cner)
		v := reflect.ValueOf(cner)
		if t.Elem().Kind() == reflect.Struct {
			number := t.Elem().NumField()
			if number > 0 {
				for i := 0; i < number; i++ {
					field := t.Elem().Field(i)
					tag := field.Tag.Get("ini")
					value := v.Elem().Field(i)
					kind := value.Kind()
					if tag == slice[0] {
						switch kind {
						case reflect.String:
							value.SetString(slice[1])
						}

					}
				}
			}
		}
	}
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				slice := strings.Split(line, " = ")
				reft(slice)
			}
			break
		}
		slice := strings.Split(line, " = ")
		reft(slice)

	}
	return nil
}
func getNames(n interface{}) {
	v := reflect.ValueOf(n)
	fmt.Printf("name-elem: %s\n", v.Elem().Type().Name())
}
func initloi() {
	num1 := aliInt(1)
	num2 := mInt(2)
	bl := true
	s := "abc"
	fl := 1.01
	GetType(s)
	GetValue(s)
	updateVal(&s)
	GetType(num1)
	GetValue(num1)
	updateVal(&num1)
	GetType(num2)
	GetValue(num2)
	GetValue(fl)
	updateVal(&fl)
	GetType(bl)
	GetValue(bl)
	updateVal(&bl)
	var st student
	st = student{
		Name: "小刚",
	}
	GetType(st)
	fmt.Printf("hasProtype-sct: %v\n", hasProtype(st, "Name"))

	var ptrin *int
	var in int
	fmt.Printf("ptrin-new之前是否是nil: %v\n", isnil(ptrin))
	fmt.Printf("in是否是nil: %v\n", isnil(&in))
	ptrin = new(int)
	fmt.Printf("ptrin-new之后是否是nil: %v\n", isnil(ptrin))
	var mp map[string]any
	fmt.Printf("slice是否是nil: %v\n", isnil(mp))
	mp = map[string]any{"name": "jack"}
	fmt.Printf("slice是否是nil: %v\n", isnil(mp))
	fmt.Printf("hasProtype-map: %v\n", hasProtype(mp, "aame"))
	per := man{
		Name:   "重耳",
		Age:    22,
		Hobby:  "篮球",
		Salary: 24157.36,
	}
	slicKey := []any{}
	getAllKyes(per, &slicKey)
	slicVal := []any{}
	getAllValues(per, &slicVal)
	fmt.Printf("slicKey: %v\n", slicKey)
	fmt.Printf("slicVal: %v\n", slicVal)
	getAllFunc(per)

	// sp := speaker(per)
	var sp speaker = per
	GetType(sp)
	type sarr [3]int
	arr := sarr{1, 2, 3}
	GetType(arr)

	var (
		str = "abc"
	)
	rt := []byte(str)
	GetType(rt)
	var config = conf{}
	ReadConfig("config.ini", &config)
	fmt.Printf("config: %v\n", config)
}
