package main

import "fmt"

func main() {
	sample1()
	fmt.Println("=====================")
	sample2()
	fmt.Println("=====================")
	sample3()
	fmt.Println("=====================")
	sample4()
	fmt.Println("=====================")
}

func sample1() {
	var a string = "Runoob"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
}

func sample2() {
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)
}

func sample3() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// sample4

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func sample4() {
	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
