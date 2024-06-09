package main

import "fmt"

// 全局变量（不能使用自动推导）
var nameCount int = 9

func main() {
	// 函数体内的局部变量
	temp := 100

	//语句内的局部变量
	if temp := 5; temp < 8 {
		temp++
		//就近原则
		fmt.Println(temp)
	}
	// 语句内的局部变量在语句外失效
	fmt.Println(temp + nameCount)

	// defer 函数，延迟到最后执行。当函数执行到最后时，剩下的defer函数会开始倒序执行
	call(1)
	defer call(2)
	a := 3
	call(a)
	a++
	// 注意虽然被延迟执行，但是值已经被传入了，不会被后续函数执行修改
	defer call(a)
	a++
	call(a)
}

func call(a int) {
	fmt.Println(a)
}
