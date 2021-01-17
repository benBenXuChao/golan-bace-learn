# go语言笔记

#### 创建文件

1. bin（装一些exe文件【windows】）
2. pkg
3. src（写代码的目录）
   * golang.org
   * github.com（自己的GitHub地址）
     * 前端组
     * 后端组（自己的项目名）
     * 基础架构

#### 查看go环境变量 go env

1. 看GOPATH
2. 看GOROOT

#### 编译写好的go

1. go build（要在你的main.go同级目录下运行）
2. 在其他地方go build （go build github.com/benbe）在src下接着写，因为src已经配置了环境变量
3. 打包后换一个名字 go build -o hellow.exe

#### 其他执行命令

1. go run main.go 编译main.go文件
2. go install 分为两步
   1. 先编译得到一个可执行文件
   2. 将可执行文件拷贝到GOPATH/bin目录中

#### 交叉编译

1. 在windows平台编译一个能在Linux平台执行的可执行文件

   ```bash
   SET CGO_ENABLED=0  // 禁用CGO
   SET GOOS=linux  // 目标平台是linux
   SET GOARCH=amd64  // 目标处理器架构是amd64
   ```

   



#### main.go中的错误语法

```go
    // 函数外只能放（变量、常量、函数、类型）的定义
    var count = "你好"
    func init() {
        fmt.Println(count)
    }
    // 不能使用写表达式，运算，调用函数等
    // fmt.Println("在外面不能调用函数")
    func main() {
        fmt.Println("hello word")
    }
```

#### 变量和常量

1. 关键字

   ```go
   func  interface  select  defer  struct  go  chan
   goto  package  fallthrough  range 
   ```

   

2. 保留字

   ```go
   Constants: iota  nil
   Types: int int8 int16 int32 int64
   	   uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error
   Functions: make len new append copy delete
              complex real imag panic recover
   ```

3. 变量的特点

   1. Go语言的变量必须先声明再使用

      ```go
          var number string // 生命变量时要定义类型
          var ( // 批量声明变量
              name string
              age int
              visiable bool 
          )
          func main(){
              name = "你好"
              age = 22
              visiable = true
              fmt.Println()
          }
      ```

      

   2. 声明后的变量一定要使用，不然报错（在全局中可以声明了不使用，非全局变量不使用会报错）

   3. Go语言中推荐使用驼峰式命名（小驼峰）

   4. fmt中的Print家族

      ```go
          fmt.Print(visiable) // 在终端打印变量
          fmt.Printf("name:%s",name) // %s: 占位符，最后的打印会把%s替换掉
          fmt.Println(age) // 打印完指定的内容后会有一个换行符
      ```

   5. 同一个作用域重复声明同名的变量

      ```go
       func main(){
              var str = "test"
              // str := "不能再声明"
          }
      ```
   
      

   6. 声明变量的同时赋值

      ```go
       func main(){
              var str string = "xss"
              var si = "ssh" // 类型推论，可以不定义类型
              fmt.Println(str)
          }
      ```
   
      

   7. 短变量声明（只能在函数中用）

      ```go
       func main(){
              s1 := "可以省略了var，而且可以类型推断" // 只能在函数中使用
              fmt.Printf("s1: %s",s1)
              // 一般只用批量声明和短变量声明
          }
      ```
   
      

   8. 匿名变量

      ```go
       func foo(){
              return 25,"小明"
          }
          func main(){
              _,name := foo() // _ 表示一个占位符
              age,_ := foo() // 省略你不关心的变量，常用于nil
              fmt.Printf("name: %s \n",name)
              fmt.Println(age)
          }
      ```
   
      

4. 常量：常量时恒定不变的量

   ```go
       const PI = 3.1415926 // PI不能再赋值
       // 批量声明常量
       const(
           OK = 200
           NotFound = 404
           BadRequest = 400
       )
       // 批量声明常量时，如果某一行声明没有赋值，默认和上一行一样
       const(
          n1 = 100
          n2
          n3
       )
       // iota 在const关键字出现时将被重置为0.const中每新增一行常量声明将使iota计数一次
       const(
           a1 = iota 
           a2 = iota
           a3
       )
   
   ```

   iota中的面试题

   ```go
       // 匿名
       const(
           b1 = iota // 0
           b2 // 1
           _
           _
           b3 // 2
       )
       // 插队
       const(
           d1 = iota // 0
           d2 = 100  // 100
           d3        // 100
           d4 = iota // 3 每新增一行常量声明计数一次
           d5		  // 4
       )
       // 一行多个声明,和空行
       const(
           e1,e2 = iota+1, iota+2
   
           e3,e4 = iota+3, iota+iota
       )
       // 常用的iota
       const(
           _  = iota
           MB = 1<< 10*iota
           GB = 1<< 10*iota
           TB = 1<< 10*iota
       )
   ```

   

#### go语言的基本数据类型

1. 整型

   1. 带正负号

      ```go
          int // 32位操作系统是int32，64位操作系统就是int64
          int8 // -128   128
          int16 // -32768  32768
          int32  int64 // 2的32和2的64
          // 其中int8就是熟知的byte类型，byte是int8的别名
      ```

      

   2. 不带正负号

      ```go
          uint // 32位操作系统是uint32，64位操作系统就是uint64
          uint8  uint16  uint32  uint64
      ```

   3. 注意事项

      - 获取对象长度的内建len()函数，返回的长度可以根据不同的平台的字节长度进行变化，
      - 在切片或者map的元素数量等都可以用int来表示，
      - 为了规避在不同编译目标平台字节长度影响，不要使用int和uint

   4. 进制数

      ```go
          var i1 = 10
          fmt.Printf("i1: %d",i1) // 10进制
          fmt.Printf("i1: %b",i1) // 2进制
          fmt.Printf("i1: %o",i1) // 8进制
          fmt.Printf("i1: %x",i1) // 16进制
          var i2 int8 = 016 // 0开头 多用于文件
          i3 = 0x34ef // 0x开头 多用于地址
          fmt.Printf("i3: %d，%d",i2,i3) // 转换为10进制
          fmt.Printf("i3: %T",i3) // 查看类型
          // 明确指定类型
          i4 := int16(127) // int16
      ```

      

2. 浮点数和布尔值

   1. 获取最大浮点数

      ```go
      max8 := math.MaxInt8
      maxf32 := math.MaxFloat64
      ```

      

   2. go语言中的小数默认是float64

   3. 不同的float32值和float64不能相互赋值

      ```go
          var f64 := 1.9878 // 默认为float64
          var f32 := float32(f64)
          // f64 = f32 不能将f32赋值给f64
      ```

      

3. 布尔值

   1. 默认值是false

      ```go
          b1 := true
          b2 bool
          fmt.Printf("类型：%T，值：%v",b2,b2) // %v表示不管是什么类型，只需要打印它的值
      ```

#### 数组

#### 函数

1. 函数的参数及返回值

   1. 函数传参时必须要有类型，返回值同样也要有类型
   2. 传参的类型可以简写
   3. 当参数不确定有多少时，可以使用不定参数
   4. 函数的返回值为具名的时，返回结果可以直接使用这个具名变量，但return是不能省的

   ```go
   	// 函数形参类型的简写
       func argument(age, source int, name, cname string) {
           fmt.Printf("age:%T,\t source:%T,\t name:%T,\t cname:%T\n", age, source, name, cname)
       }
   
       // 可变参数
       func sum(x int, y ...int) int {
           total := x
           for _, val := range y {
               total += val
           }
           return total
       }
   
       // 具名返回值
       func reduce(big int, small int) (res int) { // 具名返回值要是用()
           // 具名返回已经定义了变量直接可以使用
           res = big - small
           return // 可以不写 return res 但是一定要写return
       }
   
       // 多个返回值要使用()
       func multiple(ori int, num int) (int, int) {
           return ori / num, ori % num
       }
   
   ```

   

2. 作用域

   1. go函数的作用域与js类似，都是往上一层找
   2. go的块级作用域，if、for、switch语句

   ```go
   	// 函数作用域和js类似
   	// 块级作用域
   	if i := 2; i > 3 {
   		// i只能在if内部，和if这一行使用
   	}
   	// fmt.Printf("i:%d",i) 在外面是无法使用的
   	// for也是块级作用域
   	for i := 0; i < 10; i++ {
   		// 这个i也是只能在for内部，及for这一行使用
   		for j := 0; j < i+1; j++ {
   			// 在这个for里可以对i进行重新定义，类似与函数作用域
   			if j == 9 && i == 9 {
   				fmt.Printf("i父块及作用域：%d\n", i)
   				i := 1000 // 重新定义i
   				fmt.Printf("i重新赋值：%d\n", i)
   			}
   		}
   	}
   ```

   

3. 在函数中定义函数

   1. 函数中是无法再定义一个函数的
   2. 可以使用匿名函数 或者 立即执行函数来解决无法定义函数的问题

   ```go
   func init(){
       // 可以在init中使用自调用（立即执行）函数
       func(a, b int) {
   		fmt.Printf("a:%d,b:%d\n", a, b)
   	}(10, 20)
       // 可以在init中使用匿名函数，再将匿名函数赋值给一个变量
       fnn := func(c, d int) {
   		fmt.Printf("c:%d,d:%d\n", c, d)
   	}
   	fnn(1, 3)
   }
   ```

   

4. 一等公民的理解

   1. 函数的类型就是一个func()，但是这个func类型可能有形参和返回值，在写func类型是都要带上
   2. 函数可以作为另一个函数的参数
   3. 函数可以作为另一个函数的返回值
   4. 当函数返回值是map、slice、chnal类型时，返回nil也是合法的

   ```go
   	// 没有参数，也没有返回值
       func f1() {
   
       }
   
       // 有参数有返回值
       func f2(a int) (b int) {
           b = a * 10
           return
       }
   
       // f3的参数是一个函数，注意函数类型的写法（要把形参和返回值都写上）
       func f3(callback func(int) int) {
           ret := callback(5)
           fmt.Printf("ret:%d\n", ret)
       }
   
       // f4返回值是一个函数
       func f4(a int) func(int) int {
           // 返回的是一个函数
           // 以后每次调用改返回函数时，都不需要传入a，因为a被永久的保存了
           return func(b int) int {
               return a + b
           }
       }
   
   	// 函数的类型是什么
   	fmt.Printf("f1的类型：%T\n", f1)
   	// 虽然f1和f2都是函数可是它两的类型不一样，因为f2是有参有返回值
   	fmt.Printf("f2的类型：%T\n", f2)
   	// 函数也是可以当作参数传入函数内
   	f3(f2)
   	// 直接传入一个匿名的函数
   	f3(func(a int) int {
   		return a
   	})
   	// 函数同样也可以作为一个返回值返回
   	rest := f4(5)
   	fmt.Printf("获取闭包函数的值：%d\n", rest(6))
   
   	// 当函数返回值是map、slice、chnal类型时，返回nil也是合法的
       func f5(str string) []string {
           if str == "" {
               return nil
           }
           return strings.Split(str, " ")
       }
   	slice := f5("")
   	if slice == nil {
   		fmt.Printf("传入的是一个空字符串\n")
   	}
   ```

5. 闭包的理解

   1. 闭包指的是一个函数和与其相关的引用环境组合而成的实体，即闭包=函数+引用环境
   2. 当使用闭包时，引用环境的变量会被保存，即一直有效
   3. 每次函数时都会产生一个新的环境，但返回的func只会记忆它执行时的环境

   ```go
   	// 对记忆的环境变量进行操作（加减）
       func option(n int) (add func(int) int, red func(int) int) {
           add = func(an int) int { // 加
               n += an
               return n
           }
           red = func(rn int) int { // 减
               n -= rn
               return n
           }
           return
       }
   
       // 自动补全格式
       func complex(str string) func(string) string {
           return func(s string) string {
               if !strings.HasSuffix(s, str) {
                   return s + str
               }
               return s
           }
       }
   	// 当使用闭包时，引用环境的变量会被保存，即一直有效
   	add0, red0 := option(12)
   	fmt.Printf("add0后的值：%d\n", add0(2)) // 14
   	fmt.Printf("red0后的值：%d\n", red0(3)) // 11
   	fmt.Printf("add0后的值：%d\n", add0(4)) // 15
   
   	// 每调用一次complex都会产生一个新的环境，但返回的func中使用的环境变量是旧的
   	suftxt := complex(".txt")
   	sufhtml := complex(".hmlt")
   	fmt.Printf("suftxt方法：%s\n", suftxt("remember"))  // remember.text
   	fmt.Printf("sufhtmlt方法：%s\n", sufhtml("canvas")) // canvas.hmlt
   ```

   

#### defer和panic和recover

1. 函数中延迟处理defer

   1. Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。
   2. 在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行
   3. defer后面调用的方法的参数计算并不会延迟处理

   ```go
       // 验证defer是延时处理
       func df1() {
           defer fmt.Println("步骤：1")
           fmt.Println("步骤：2")
           fmt.Println("步骤：3")
           defer fmt.Println("步骤：4")
       }
       // 注意4在1之前被打印，类似与栈先进后出
       df1() // 2,3,4,1的顺序
   
       // 参数并不会延迟，defer的表达式会先将参数计算好
       func fd5() {
           a, b := 3, 1
           defer fmt.Printf("1. %d\t%d\n", a, b)
           a = 0
           defer fmt.Printf("2. %d\t%d\n", a, b)
           b = 2
       }
   	fd5() // 2. 0    1   a和b跟正常逻辑一样，不会延迟执行     
   		  // 1. 3    1
   ```

2. defer的执行时机

![](C:\Users\xuchao\Desktop\go-mark\go语言笔记.assets\defer.png)

1. 验证defer的执行时机

   ```go
   func fd2(x int) (y int) {
   	defer func() {
   		y++ // 在return之前，y会将赋值后的值再加一
   	}()
   	return x + 10
   }
   func fd3(x int) int {
   	defer func() {
   		x++ // 在y对x赋值之后，对x的操作不会影响到y
   	}()
   	return x + 10
   }
   func fd4(x int) *int {
   	defer func() {
   		x++ // 当y等于x的地址时，x再赋值之后依然会影响y
   	}()
   	x += 10
   	return &x
   }
   fmt.Printf("df2执行结果:%d\n", fd2(3))  // 14
   fmt.Printf("df3执行结果:%d\n", fd3(3))  // 13
   fmt.Printf("df4执行结果:%d\n", *fd4(3)) // 14
   ```

   

2. 捕获错误defer中recover一个panic

   1. 程序运行时出错，就会报一个panic
   2. 在没有任何捕获错误的机制下，panic会中断后面程序的运行
   3. 使用defer和recover，在报panic时，会尝试重新恢复程序，这样此panic只会中断当前func的程序，后面的func程序会继续进行
   4. recover必须和defer配合使用
   5. defer函数必须在错误之前声明

```go
	func pAndR(){
        // recover必须和defer配合使用
        // defer函数必须在错误之前声明
        defer func(){
            err := recover()
            fmt.Printf("错误被捕获到：%v",err)
        }() // 自调用
        var arr []int
        arr[0] = 10 // 这个在运行时会报painc错误，越界
    }

	// 也可以手动抛一个panic
func handPnc(){
    defer func(){
        if err := recover();!=nil{
            // 如果程序出现了错误recover可以尝试恢复程序，如果可以恢复，后面的函数就可以执行
            fmt.Printf("已经recover%v",err)
        }
    }()
    painc("手动抛出错误")
}
```



#### map

1. map是一种无序的基于`key-value`的数据结构
2. Go语言中 `map`的定义语法如下：

```go
map[KeyType]ValueType
```

- keyType:表示键的类型
- valueType:表示键对应的值的类型

1. map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
make(map[KeyType]ValueType, [cap])
```



1. map的初始化方式

   1. 初始化并赋值
   2. 使用make初始化

   ```go
   	// 初始化并赋值
   	mp := map[string]int{
   		"小明": 22,
   		"李晨": 31,
   		"小刚": 16,
   	}
   	fmt.Printf("mp:%#v\n", mp)
   	// 使用make进行初始化
   	mp1 := make(map[string]int, 5)
   	mp1["jack"] = 23
   	mp1["rose"] = 22
   	fmt.Printf("mp1:%#v\n", mp1)
   ```

   

2. map的取值和删除

```go
	age := mp["小明"]
	fmt.Printf("age:%d\n", age)
	// 但是一般要先判断该map是否有这个key
	val, ok := mp1["piter"]
	if !ok {
		// 当ok为false时，表示没有这个key，val得到的时默认值
		fmt.Printf("val:%d\n", val)
	}
	// 使用内置方法delete删除键值对
	delete(mp, "李晨")
	delete(mp1, "hi")
	fmt.Printf("mp:%#v\n", mp)
	fmt.Printf("mp1:%#v\n", mp1)
```



1. 顺序遍历map

```go
	rand.Seed(time.Now().UnixNano()) // time.Now().UnixNano()当前纳秒
	mp2 := make(map[string]int, 50)
	// 产生随机的key和val
	for i := 0; i < 50; i++ {
		mp2[fmt.Sprintf("mp%02d", rand.Intn(50))] = rand.Intn(50)
	}
	// 将所有的key放进切片中
	var slice = make([]string, 100)
	for key := range mp2 {
		slice = append(slice, key)
	}
	// 对切片进行排序
	sort.Strings(slice)
	// 遍历排序后的切片，并取对应的值
	for _, val := range slice {
		fmt.Printf("key:%s\tval:%d\n", val, mp2[val])
	}
```



1. 切片类型的map和map类型的切片

   1. 切片类型的map-->map的值是切片类型
   2. map类型的切片-->切片的值是map

   ```go
   	// map类型的切片
   	slmp := make([]map[string]int, 3, 5) // 初始化
   	mp3 := make(map[string]int, 3)       // 内部的map也需要初始化
   	slmp = append(slmp, mp3) // 由于是引用类型，map再后面进行修改
   	mp3["李青"] = 33
   	fmt.Printf("slmp:%#v\n", slmp)
   	// 定义且初始化
   	slmp1 := []map[string]int{
   		map[string]int{
   			"lev": 22,
   		},
   	}
   	fmt.Printf("slmp1:%#v\n", slmp1)
   	// 切片类型的map
   	mpsl := make(map[string][]int, 5) // 初始化
   	sl := []int{1, 2, 3}              // 内部的切片也需要初始化
   	mpsl["gor"] = sl
   	fmt.Printf("mpsl:%#v\n", mpsl)
   	// 定义且初始化
   	mpsl1 := map[string][]int{
   		"brt": []int{3, 2, 1},
   	}
   	fmt.Printf("mpsl1:%#v\n", mpsl1)
   ```

   

#### 切片

1. 创建切片

   ```go
   	// 直接初始化
   	slice := []int{1, 2, 3} // 数组需要在定义时设置长度，或者...
   	// 使用make关键字
   	makeSlice := make([]int, 3, 10) // 类型、长度、容量 [0,0,0]
   	// 用数组去切割得到切片
   	arr := [3]int{1, 2, 3}
   	formArr := arr[1:]
   	// 切片再切割得到切片
   	formSlice := formArr[:1]
   
   	fmt.Printf("slice:%#v\n", slice)         // []int{1, 2, 3}
   	fmt.Printf("makeSlice:%#v\n", makeSlice) // []int{0, 0, 0}
   	fmt.Printf("formArr:%#v\n", formArr)     // []int{2, 3}
   	fmt.Printf("formSlice:%#v\n", formSlice) // []int{2}
   
   	// 切片是引用类型，都指向底层的数组，修改数组的值，切片也会变得
   	arr[1] = 88 // formArr也会变
   	fmt.Printf("formArr<--arr:%#v,arr:%#v\n", formArr, arr)
   	// 由于切片是不保存数据的，所有的数据都交给底层数组来存储，所以修改切片也会影响底层数组
   	formArr = append(formArr, 90)
   	fmt.Printf("formArr-->arr:%#v,指针：%p,arr:%#v\n", formArr, formArr, arr) // [4]int{1, 88, 3, 90} 
   ```

   

2. 切片三要素：指针、长度、容量

   ```go
   	// 切片三要素 指针、长度、容量
   	slice := []int{1, 2, 3}
   	fmt.Printf("slice指针：%p,长度：%d,容量：%d\n", slice, len(slice), cap(slice)) // 指针：0xc0000044a0,长度：3,容量：3
   	// 切割切片 slice[low,high,max]
   	// len如何计算 high-low
   	// cap如何技术 原始的cap-low1-low2....  如果设置了max（max合法情况下） max-low
   	// low <= high <= max <= 原始切片的cap
   	origin := []int{0, 1, 2, 3, 4, 5, 6}
   	or1 := origin[1:2] //[1] len=2-1  cap=7-1
   	or2 := or1[2:6]    //[3,4,5,6] len=6-2 cap=7-1-2 high最大值为cap
   	or3 := or1[3:4:5]  //[4] len=1 cap=5-3 可以通过第三个参数减少容量，且high<=max
   	fmt.Printf("or1值：%#v,指针：%p,长度：%d,容量：%d\n", or1, or1, len(or1), cap(or1))
   	fmt.Printf("or2值：%#v,指针：%p,长度：%d,容量：%d\n", or2, or2, len(or2), cap(or2))
   	fmt.Printf("or3值：%#v,指针：%p,长度：%d,容量：%d\n", or3, or3, len(or3), cap(or3))
   ```

   

3. 切片方法append

   ```go
   	// 使用append对切片进行操作
   	// append操作后需要对切片重新赋值，因为有可能扩容
   	// append追加元素，原来的底层数组放不下的时候，go会申请一个更大长度的底层数组
   	var appSli = []int{1, 2, 3, 4, 5, 6}
   	appSli = append(appSli, 33) // push
   	fmt.Printf("appSli值：%#v,指针：%p,容量：%d\n", appSli, appSli, cap(appSli))
   	appSli = append(appSli[1:]) // shift
   	fmt.Printf("appSli值：%#v,指针：%p,容量：%d\n", appSli, appSli, cap(appSli))
   	appSli = append(appSli[:len(appSli)-1]) // pop
   	fmt.Printf("appSli值：%#v,指针：%p,容量：%d\n", appSli, appSli, cap(appSli))
   	appSli = append(appSli[:3], appSli[4:]...) // remove
   	fmt.Printf("appSli值：%#v,指针：%p,容量：%d\n", appSli, appSli, cap(appSli))
   ```

   

4. append之后切片扩容规律

   - 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
   - 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
   - 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
   - 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

5. 切片的copy

   ```go
   	// 切片是引用类型，直接赋值赋值的是地址，赋值后两者会有关联，如想深copy需要使用copy
   	equalSlice := appSli
   	var copySlice = make([]int, 10, 10) // copySlice必须初始化过
   	copy(copySlice, appSli)
   	appSli[0] = 99
   
   	fmt.Printf("equalSlice值：%v,指针：%p,容量：%d,指针是否相等：%t\n", equalSlice, &equalSlice, cap(equalSlice), &appSli == &equalSlice) // 意料之外的是 指针不相同
   
   	fmt.Printf("copySlice值：%v,指针：%p,容量：%d\n", copySlice, copySlice, cap(copySlice))
   ```

   

6. 切片的排序sort

   ```go
   	// 要使用sort包来进行排序，必须是int类型
   	var floatSlice = []float64{5.1, 3, 4.2, 7, 2}
   	var intSlice = []int{4, 3, 6, 2, 8}
   	sort.Ints(intSlice)
   	sort.Float64s(floatSlice)
   	fmt.Printf("floatSlice:%v", floatSlice)
   	fmt.Printf("intSlice:%v", intSlice)
   ```

   

7. 切片的比较

   - 两个切片是无法进行比较的
   - 切片唯一可以比较的是nil，如果一个切片的值为nil，那么该切片是没有底层的数组的
   - 一个nil值的切片的长度和容量都是0，但是长度和容量都为0的切片一定是nil
   - 判断切片是否为空，使用len，不要使用是否等于nil来判断

8. 切片理解图

   ```go
   a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
   ```

   

   - s1 := a[:5]

   ![切片图](C:\Users\xuchao\Desktop\go-mark\go语言笔记.assets\slice_01.png)

   - s2 := a[3:6]

     ![slice_02](C:\Users\xuchao\Desktop\go-mark\go语言笔记.assets\slice_02.png)

7. 切片的本质
   - 切片就是一个框，框住了一块连续的内存，只能保存相同类型的变量
   - 切片的长度是动态的
   - 切片属于引用类型，真正的数据是保存在底层的数组里

#### make和new关键字（得到指针）

1. new初始化一个值类型的指针

   ```go
       // make和new都是实例化一个变量，并为这个变量分配一个内存地址，实例化得到的值也会分配一个地址
   	// new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，
       // 返回的是新值的地址，也就是 T 类型的指针 *T
   	var a = new(int) // 得到一个int类型的指针
       fmt.Printf("a:%v",a) // 打印的是a指针指向的内存地址，因为无法直接打印指针，只能打印a指向的地址
       *a = 10 // 对a指针指向的内存地址赋值为一个值
       fmt.Printf("*a: %d",*a) // 
   
       // 初始化一个数组指针
       var arr = new([3]int)
       arr[2] = 10 // [0,0,10] // 直接对指针地址的第几个进行赋值
       arr[1] = 10                   // 由于go语言进行了封装所以接赋值
       // *arr[2] = 20 错误写法因为arr[2]可能是一个表达式
       (*arr)[2] = 29 // 必须要用()括起来
       fmt.Printf("arr:%v",*arr) // 获取指针指向的指需要使用* [0,10,29]
   ```

   

2. make初始化一个引用型的指针

   ```go
   	// make 只能用于 slice，map，channel 三种类型    
   	var b = make([]int,0,10) // 得到一个切片类型的指针
       fmt.Printf("b:%v\n",b) // 打印的是指针指向的地址
   
   ```

3. 注意new和mack的区别

   ```go
   	// 切片是一个引用类型，相对于数组不需要再定义长度
   	arr := [3]int{1, 2, 3}  // 数组需要在定义时设置长度，或者...
   	pointArr := new([3]int) // new返回的是指针不是值
   	slice := []int{1, 2, 3}
   	pointSli := make([]int, 5, 10) // make返回的是值不是指针
   	fmt.Printf("arr值：%v,类型：%T\n", arr, arr) // [3]int
   	fmt.Printf("pointArr值：%v,类型：%T\n", pointArr, pointArr) // *[3]int
   	fmt.Printf("slice值：%v,类型：%T\n", slice, slice) // []int
   	fmt.Printf("pointSli值：%v,类型：%T\n\n", pointSli, pointSli) // []int
   
	// go底层帮助我们补*号
   	var str = new(int)
   	var num = new(int)
   	var arr = new([3]int)
   	var stru = new(struct{ Name string })
   	// 虽然都是获得指针，但是打印时并不都是指针指向的地址
   	// 因此可以看出go对struct和array处理过
   	fmt.Println(str)  // 0xc0000a21a0 打印出地址 
   	fmt.Println(num)  // 0xc0000a21a8 打印出地址 
   	fmt.Println(arr)  // &[0 0 0] 打印出值，但他会在值前面加&
   	fmt.Println(stru) // &{} 打印出值，但他会在值前面加&
   ```
   
   

#### struct结构体和方法

1. 创建一个自定义类型

   ```go
       // 使用type关键字创建一个自己的类型，不是内置的
       type myInt int8
       // 注意它与类型别名的区别
       type alinsInt = int 
       // 创建一个struct 是自定义的
       type student struct{ // type +名字+ struct
           Name string
           Age int
           Hobby []string
           Good bool
       }
   ```

   

2. 创建一个结构体和赋值

   ```go
   	// 定义且初始化
   	var jack = student{
   		Name:  "jack", // 需要写 ,
   		Age:   23,
   		Hobby: []string{"开心", "沮丧"},
   		Good:  true,
   	}
   	fmt.Printf("jack结构体：%v\n", jack)
   
   	// 使用new关键字，初始化为默认值
   	var rose = new(student)
   	fmt.Printf("rose初始值：%v\n", rose)
   	// 对rose赋值
   	rose.Name = "rose"
   	rose.Age = 20
   	rose.Hobby = []string{"化妆", "吃"}
   	rose.Good = false
   	fmt.Printf("rose结构体：%v\n", rose)
   
   	// 和new一样的写法
   	var ming = &student{} // 和new(strudent)一样
   ```

   

3. 结构体的初始化

   ```go
   	// 初始化结构体(省略key初始化)
   	var ning = student{
   		"明",
   		28,
   		[]string{"打野", "辅助"},
   		true,
   	}
   	fmt.Printf("ning: %v\n", ning)
   	// 使用key和val初始化时，可以省略某些字段，这些字段为默认值
   	kai := student{
   		Name:  "kai",
   		Hobby: []string{"打野"},
   	}
   	fmt.Printf("kai: %v\n", kai)
   	// 两者不可混合使用
   	/**
   	jie := student{
   		Name: "kai",
   		23,
   		Hobby: []string{"打野"},
   		true,
   	}
   	fmt.Printf("kai:%v\n", jie)
   	*/
   ```

   

4. 结构体中的方法

   - Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。

   - 这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。
   - 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写

   ```go
   	// Student 学生
       type Student struct {
           Name  string
           Age   int
           Hobby []string
           Good  bool
       }
   	/**
   	方法最多三个()号
   	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
       	函数体...
   	}
   	*/ 
   	// go语言中的方法不是定义在结构体内部的
   	func (s *Student) structFunc(name string, age int) (*Student) {
           s.Name = name
           s.Age = age
           return s
   	}
   	// 结构体的方法
   	var newStudent = Student{
   		Name: "小刚",
   		Age:  22,
   	}
   	newStudent = *newStudent.structFunc("hai", 10)
   	fmt.Printf("newStudent的值：%v\n\n", newStudent) // 
   ```

   

5. 结构体的内存布局

   1. 内存是以字节为单位的十六进制数
   2. 一字节 = 8位 = 8bit  （int8占一个字节，int16占两个字节）

6. 结构体是一个值类型

   ```go
   	var jin = student{}
   	var xia = jin
   	xia.Name = "xia"
   	fmt.Printf("xia:%v\n", xia)
   	fmt.Printf("jin:%v\n", jin) // 修改失败
   	var niu = &jin
   	niu.Name = "niu"
   	fmt.Printf("niu:%v\n", *niu)
   	fmt.Printf("jin:%v\n", jin) // 修改成功
   ```

   

7. 值接收者和指针接收者，以及返回值和指针

   1. 如果函数接受的不是一个指针，那么函数的形参会申请另一个地址，修改形参值不会对原有值产生影响
   2. 如果返回的值不是一个指针，那么需要一个新的地址接受这个值

   ```go
   	// 传入结构体值，返回值和指针
       func coypStruct(st Student) (Student, *Student) {
           // 当传入的是值的话，函数内部会复制的是结构体的值，
           // 修改结构体的值，不会对函数外结构体有影响
           fmt.Printf("st指针：%p\n", &st)
           st.Name = "new-name"
           prtSt := &st
           return st, prtSt
       }
   
       // 传入结构体指针，且返回结构体指针
       func copyPrt(st *Student) *Student {
           // 当传入的是指针的话，函数内部会复制的是结构体的指针，
           // 修改指针指向地址内的值，会对函数外结构体有影响
           fmt.Printf("st指针：%p\n", st)
           st.Name = "prt-name"
           return st
       }
   	// 直接向函数中传入一个结构体，在函数内得到的是一个副本
   	oriSt := Student{
   		Name: "jack",
   	}
   	// P源-->表示要修改的结构体  P内-->表示函数内部的指针
   	ori, prtOri := coypStruct(oriSt)
   	fmt.Printf("oriSt值：%#v,指针:%p\n", oriSt, &oriSt)  // 值没有发生变更
   	fmt.Printf("oriSt值：%#v,指针:%p\n", ori, &ori)       // P源 != P内 =!P ori
   	fmt.Printf("prtOri值：%#v,指针:%p\n", prtOri, prtOri) // P源 != P内 ==P prtOri
   	returnPrt := copyPrt(&oriSt)
   	fmt.Printf("oriSt值：%#v,指针:%p\n", oriSt, &oriSt) // 值发生了变更
   	fmt.Printf("returnPrt值：%#v,指针:%p\n", *returnPrt, returnPrt) // P源 == P内 ==P returnPrt
   ```

   

8. 什么时候传入结构体指针，什么时候传结构体

   - 需要修改接收者中的值
   - 接收者是拷贝代价比较大的大对象
   - 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

9. 使用fmt打印指针时

   1. 获取值类型的指针时需要使用&取址符
   2. 如果使用的是new关键字，则直接得到了指针
   3. 获取引用性类型的指针时，直接可以打印%p，（因为一个引用性类型的值包含指针）
   4. 如果一个变量等于一个引用性类型的值a，那么这个变量的地址和a的地址是不同的
   5. 如果一个变量b等于数组或struct，b的地址就是第一个值或者第一个属性的地址，这个地址与后面的地址都是连续的

   ```go
   	// 使用fmt打印指针时，引用类型和值类型不同
   	// 引用类型直接可以用，值类型要取址符
   	fmtInt := 1
   	fmtArr := [2]int{1, 2}
   	fmtStr := "hi"
   	fmtSli := []string{"haha"}
   	fmtMap := map[string]string{"name": "jack", "sor": "tien"}
   	fmt.Printf("fmtInt指针：%p\n", &fmtInt)
   	fmt.Printf("fmtArr指针：%p\n", &fmtArr)
   	fmt.Printf("fmtStr指针：%p\n", &fmtStr)
   	fmt.Printf("fmtSli指针：%p\n", fmtSli)
   	fmt.Printf("fmtMap指针：%p\n", fmtMap)
   
   	// 变量地址和变量保存的值的地址，不是一个概念（一个变量也需要一个地址来保存）
   	fmt.Printf("str值的地址：%p\n", str)   // str值的地址：0xc0000141b8
   	fmt.Printf("str变量的地址：%p\n", &str) // str变量的地址：0xc000006030
   
   	fmt.Printf("stru值的地址：%p\n", str)    // stru值的地址：0xc0000141b8
   	fmt.Printf("stru变量的地址：%p\n", &stru) // stru变量的地址：0xc000006038
   
   	// 上面的过程类似与下面
   	var needntNew = 1                               // needntNew保存的是一个值
   	prtNeed := &needntNew                           // prtNeed保存的是一个指针指向的地址
   	fmt.Printf("needntNew变量分配的地址：%p\n", &needntNew) // needntNew变量的地址：0xc00008c030
   	fmt.Printf("prtNeed变量保存的地址：%p\n", prtNeed)      // prtNeed变量的地址：0xc00008c030
   	fmt.Printf("prtNeed变量分配的地址：%p\n", &prtNeed)     // prtNeed变量的地址：0xc00008a010
   
   	// 如何获取(结构体字段地址，数组每一项)的地址
   	type testStruc struct {
   		a byte
   		b byte
   		c byte
   	}
   	var areaStruc = &testStruc{
   		1,
   		'a',
   		'2',
   	}
   	// 可以发现a,b,c都是连续的，并且变量areaStruc储存的地址就是a的地址
   	fmt.Printf("areaStruc的地址：%p\n", areaStruc) // areaStruc的地址：0xc000014140
   	fmt.Printf("a的地址: %p\n", &(areaStruc.a))   // a的地址: 0xc000014140
   	fmt.Printf("b的地址：%p\n", &(areaStruc.b))    // b的地址：0xc000014141
   	fmt.Printf("c的地址：%p\n", &(areaStruc.c))    // c的地址：0xc000014142
   
   	var areaArr = &[3]byte{'a', 3, 'e'}
   	fmt.Printf("areaArr的地址：%p\n", areaArr)  // areaArr的地址：0xc00009211b
   	fmt.Printf("0的地址: %p\n", &(areaArr[0])) // 0的地址: 0xc00009211b
   	fmt.Printf("1的地址：%p\n", &(areaArr[1]))  // 1的地址：0xc00009211c
   	fmt.Printf("2的地址：%p\n", &(areaArr[2]))  // 2的地址：0xc00009211d
   ```

   

   ![](C:\Users\xuchao\Desktop\go-mark\go语言笔记.assets\ptr.png)

10. 构造函数

    1. 构造函数可以简便的创建结构体对象
    2. 构造函数可以返回新“new”出来结构体的值，也可以返回结构体指针
    3. 一般返回结构体指针，这样会尽量减少开辟新的内存空间

   ```go
       // 构造函数
       func People(name string, age int, hobby []string, good bool) student {
           return student{
               name,
               age,
               hobby,
               good,
           }
       }
   	aila := People("aila", 22, []string{"网剧"}, true)
   	fmt.Printf("aila:%v\n", aila)
   ```

   

11. 在结构体中定义结构体、实现继承、结构体方法的继承、提前准备tag，为json序列化做准备

    1. 结构体中定义结构体有三种方式
       1. 定义具名字段的结构体
       2. 定义匿名字段的结构体
       3. 结构体字段对应的是一个结构体指针
    2. 实现继承必须要使用匿名结构体
    3. 不仅可以继承struct的属性，还可以继承struct的方法

    ```go
    	// 在结构体中定义结构体，实现一个对象的嵌套
    	// 一个将被继承的struct，类似基类
    	type Address struct {
    		Privince string `json:"privince"`
    		City     string `json:"city"` // 
    	}
    	// 另一个将被继承的struct，类似积基类
    	type Home struct {
    		Area string `json:"area"`
    		City string `json:"city"`
    	}
    	// 一个类似派生类的struct
    	type Person struct {
    		Name   string `json:"name"`
    		Gender string `json:"gender"`
    		// 使用非匿名，结构体中定义结构体
    		// Address Address
    		Address     `json:"address"`    // 匿名字段，将来可以实现继承
    		Home        `json:"home"`       // Home和Address同时匿名时，City字段会有冲突
    		*Collection `json:"collection"` // 匿名字段，但值是一个指针类型
    	}
    	// Collection初始化
    	coll := Collection{
    		City: "",
    	}
    	// 定义一个小明结构体
    	var xiaom = Person{
    		Name:   "小明",
    		Gender: "男",
    		Address: Address{
    			Privince: "102123",
    			City:     "213221",
    		},
    		Collection: &coll,
    	}
    	// 使用了匿名字段了以后就可以实现继承了
    	// 获取Address和修改Address
    	xiaom.Privince = "88888" // xiaom直接可以修改privince
    	fmt.Printf("xiaom:%v,City:%v\n", xiaom, xiaom.Privince)
    
    	// 嵌套结构体的字段名冲突（必须都是匿名字段才会出现冲突）
    	// City字段名冲突
    	// xiaom.City已经存在冲突了
    	xiaom.Home.City = "838271" // 有冲突的字段必须要写全，不能使用继承的方式修改和获取值
    	fmt.Printf("home.city:%v\n", xiaom.Home.City)
    
    	// 不仅仅可以继承struc的属性，还可以继承方法，这个代码要写在Personal的前面
        type Collection struct {
            City string `json:"city"`
        }
        // Collection的update方法
        func (c *Collection) update(city string) *Collection {
            c.City = city
            return c
        }
    	// 小明可以直接调用继承来的update方法
    	xiaom.update("55555")
    	// 由于这是Collection的方法所以修改的是Collection中city的属性
    	fmt.Printf("collection.City:%s\n", xiaom.Collection.City)
    
    	// 嵌套结构体中传入指针和传入struct值的区别
    	// 虽然是调用的coll的update方法，但是一样可以修改personal的值（因为Personal中传入的是指针）
    	coll.update("uuuuuu")
    	// 当传入的是struct值时
    	type People struct {
    		Name string
    		Collection
    	}
    	col := Collection{
    		City: "00000",
    	}
    	peo := People{
    		Name:       "小刚",
    		Collection: col,
    	}
    	col.update("bbbbbbb") // 这里就不能改变people的值，只修改的col的值
    	
    ```

    

12. 序列化结构体

    1. 需要提前在定义结构体时打上tag，告诉json包如何处理struct的字段
    2. Marshal会把结果返回

    ```go
    	// 将struct结构体序列号
    	json, _ := json.Marshal(xiaom)
    	fmt.Printf("json:%v\n", string(json))
    ```

    


#### 接口interface

1. interface是什么

   - 接口实际上就是一个类型，它使用type来声明
   - 接口与一般的类型不同，它不是一种动态的存储类型

2. 为什么要使用interface

   - 引入问题（方法中无法传入一个具体的类型）
   - 使用interface解决问题

   ```go
   	// 定义一个有speak方法的接口
   	type speaker interface {
   		speak()
   	}
   	// 结构体猫
   	type cat struct {
   		talk string
   	}
   	// 结构体狗
       type dog struct {
           talk string
       }
   	// 狗有speak方法
       func (d dog) speak() {
           fmt.Println(d.talk)
       }
   	// 猫有speak方法
       func (c cat) speak() {
           fmt.Println(c.talk)
       }
   	// 引入问题，我们需要一个方法，它接受一个参数，这个参数是一个有speak方法的结构体
   	// 我们定义方法的时候无法直接定义参数是什么结构体，因为可能传猫，也可能传狗
   	// 使用接口就可以解决问题（定义参数是一个speaker接口类型）
   	var hi = func(some speaker) {
   		some.speak()
   	}
   	mao := cat{
   		"喵喵喵",
   	}
   	gou := dog{
   		"汪汪汪",
   	}
   	hi(mao) // 可以传猫
   	hi(gou) // 可以传狗
   ```

   

3. 定义接口类型的值

   - 定义一个值的类型为一个interface，那么这个类型就是一个动态类型，只有在赋值那一刻才能确定这个类型是什么
   - 只要实现了接口所有的方法，就可以赋值

   ```go
   	var piter speaker // 定义了一个类型为speaker接口的变量
   	piter = mao // mao有speak方法，所以可以赋值 
   				//（只有在赋值的时候才能确定，存储的类型为cat，值为mao）
   	piter.speak()
   	piter = gou // gou有speak方法，所以也可以赋值
   	piter.speak()
   ```

   

4. 怎样才算实现了这个interface

   - 实现了接口中所有的方法
   - 实现方法必须和接口中的方法类型一样（包括方法名，参数，返回值）

   ```go
       // 定义一个runner接口
   	type runner interface {
           run(speed, unit string) []int // run方法可以省去形参，但是必须写两个string
       }
   	// dog实现了runner中的run方法（方法的类型是一样的）
       func (d dog) run(speed, unit string) []int {
           fmt.Printf("狗跑了%s%s", speed, unit)
           return nil
       }
   	// cat有run方法，但并没有实现runner接口的方法
       func (c cat) run() {
           fmt.Println("猫会跑")
       }
       var hali runner
       hali = gou // gou可以赋值给hali
       hali.run("20", "m")
       // hali = mao // 猫虽然也有run方法但是与interface中的run方法不是一个类型
       // hali.run()
   ```

   

5. 一个机构体可以实现多个接口

6. 接口的继承

7. 接口定义的方法，值接收者和指针接收者

#### 反射

1. 什么是反射

   - 程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分，因此程序在运行程序时，程序无法获取自身的信息，例如（字段名称，存储类型，结构体属性名等）
   - 如果一个语言支持反射，那么该语言在编译期会将变量的反射信息（字段名等）整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并有能力修改他们
   - go语言就支持反射，可以使用reflect包访问程序的反射信息

2. 变量的内存机制

   - 类型信息：预先定义好的元信息
   - 值信息：程序运行过程中可动态变化的

3. 任何接口值都是由 `具体类型` 和`具体类型的值`两部分组成的，reflect包提供了 `reflect.TypeOf()` 和`reflect.ValueOf()`两个函数来获取任意对象的Value和Type

4.  `reflect.TypeOf()` 方法获取接口值的类型

   - 反射中类型分为两种
     - 类型（Type）type关键字构造的很多自定义类型
     - 种类（Kind）指底层的类型

   - 当需要区分指针、结构体等大品种的类型时，就会用到`种类（Kind）`。 

   - Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的`.Name()`都是返回`空`。

   - 在reflect包中定义的Kind类型如下

     ```go
     type Kind uint
     const (
         Bool                 // 布尔型
         Int                  // 有符号整型
         Int8                 // 有符号8位整型
     	...
         Uint                 // 无符号整型
         Uint8                // 无符号8位整型
         Float32              // 单精度浮点数
         Float64              // 双精度浮点数
         Array                // 数组
         Chan                 // 通道
         Func                 // 函数
         Interface            // 接口
         Map                  // 映射
         Ptr                  // 指针
         Slice                // 切片
         String               // 字符串
         Struct               // 结构体
         Complex64            // 64位复数类型
         Complex128           // 128位复数类型
         Uintptr              // 指针
         UnsafePointer        // 底层指针
     )
     ```

     

5. 常用到的反射

   1. json数据的解析
   2. orm数据的解析
   3. 配置文件的解析


