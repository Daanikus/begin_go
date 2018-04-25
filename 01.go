package main

import "fmt"

// 函数外定义必须使用关键词 var， 不可用 冒号
// 同时，这定义的并不是全局变量， 而是包内变量， 也就是 package main
var aa = 1
var bb = "hello ll"
//cc := true
// 也可用括号来统一定义变量
var (
	qq = 1
	ww = 2
	ee = true
)

func variableZeroValue()  {
	// 第一种定义变量方式， 定义的变量必须使用
	// 变量名在前， 类型在后
	// 未赋值默认为0
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a, b int  = 3, 8
	var s string = "ooxx"
	fmt.Println(a, b, s)
}

func variableTypeDeduction()  {
	// 简化上面的定义类型写法， 由编译器判断
	var a, b, c, s  = 3, 4, true, "an"
	fmt.Println(a, b, c, s)
}

func variableShort()  {
	// 等价于上面的定义方式
	// 初次定义变量时以冒号定义
	a, b, c, s := 3, 4, true, "def"
	// 再次定义时不可用 var 关键词， 视为重复定义
	a = 1
	fmt.Println(a, b, c, s)
}


func main() {
	fmt.Println("Hello, World!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
	fmt.Println(aa, bb, qq, ww, ee)
}