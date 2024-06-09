package main

import "fmt"

func main() {
	var a string = "abcdefg"

	// len() 字符串长度
	fmt.Println(len(a))
	// a[num] 打印第几个字节
	fmt.Printf("%c\n", a[3])

	// for range 循环，一般用于数组 / 切片
	// 会返回数组/切片的下标及值
	// 接受一个值则为数组下标，接收两个值则为数组下标及下标对应的值
	for i, v := range a {
		// i 是下标
		fmt.Println(i)
		// v 是下标对应的值
		fmt.Printf("%c\n", v)
	}

	// string 字符串是不能修改的
	// a[3] = 'A' 报错
}
