package main

import "fmt"

// 全局变量
var name string = "king"

func main() {
	var a int = 1
	var b int = 2
	// 变量交换简便写法
	a, b = b, a
	fmt.Println(a, b)

	// 函数可以返回多个值
	c, d := test()
	fmt.Println(c, d)

	// 匿名变量用 _ 表示，没用，相当于占位符
	e, _ := test()
	fmt.Println(e)

	fmt.Println(name)

	// 常量的定义，不能被修改
	const URL string = "https://www.google.com"
	const COUNT int = 3
	const BOOLEAN bool = true
	fmt.Println(URL, COUNT, BOOLEAN)

	// 几种简略写法，可以让系统自动推导
	const name, address, sex = "aa", true, 1
	var name2, address2, sex2 = "bb", false, 2
	name3 := "cc"

	fmt.Println(name, address, sex)
	fmt.Println(name2, address2, sex2)
	fmt.Println(name3)

	//iota的行为类似于一个计数器，每次在常量声明中出现时自增一次, 每个新的const内都会重新开始计数
	const (
		aa = iota
		bb
		cc = "111"
		dd
		ee = true
		ff
		gg = iota
	)
	const (
		hh = iota
	)
	fmt.Println(aa, gg, bb, cc, dd, ee, ff, gg, hh)
}

func test() (int, int) {
	return 3, 4
}
