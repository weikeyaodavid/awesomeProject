package main

import "fmt"

func main() {
	var flag bool = true
	var a string = "aa"

	var b int = 1
	// uint 只包含正数
	var bb uint = 1
	// 等同于 uint8 很小
	var bbb byte = 1

	var c float64 = 3.1415

	fmt.Println(flag, a, b, bb, bbb, c)
	// 浮点类型的打印用 %f，默认打印6位小数
	fmt.Printf("%f\n", c)
	// 用 .x 表明保留几位，
	fmt.Printf("%.3f\n", c)

	// 双引号是字符串
	var str = "a"
	// 单引号是字符，输出是在unicode表中的编号，是个int值
	var str2 = 'A'
	fmt.Printf("%T, %T\n", str, str2)
	fmt.Printf("%s, %d\n", str, str2)

	// 字符串拼接
	fmt.Println("aaa" + "bbb")
	// 转义字符
	fmt.Println("aaa \" bbb")

	// 类型转换都必须显示转换，整形不能转换为布尔类型
	typeA := 5
	typeB := float64(typeA)
	fmt.Println(typeA, typeB)
	fmt.Printf("%T, %T\n", typeA, typeB)
}
