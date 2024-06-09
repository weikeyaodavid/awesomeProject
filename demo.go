package main

import "fmt"

/*
多行注释
*/
func main() {
	//单行注释
	fmt.Println("Hello World!")

	var name string = "zhangsan"
	name = "lisi"
	fmt.Println(name)

	var age int = 18
	fmt.Println(name, age)

	/*
		string 默认值为空字符串
		int 默认值为 0
		bool 默认值为 false
	*/
	var (
		you     string
		me      string
		times   int
		success bool
	)

	fmt.Println(you, me, times, success)

	var a, b, c bool
	fmt.Println(a, b, c)

	//短变量自动推断
	selfName := "David"
	selfAge := 26
	fmt.Println(selfName, age)

	// 常用占位符：https://blog.csdn.net/HHTNAN/article/details/78607180
	fmt.Printf("%T, %T\n", selfName, selfAge)
	// 取地址符 &
	fmt.Printf("age:%d, 内存地址:%p\n", selfAge, &selfAge)
	selfAge = 27
	fmt.Printf("age:%d, 内存地址：:%p", selfAge, &selfAge)
}
