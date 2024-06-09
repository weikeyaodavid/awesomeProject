package main

import "fmt"

func main() {
	// 位运算 - 建立在二进制上，竖着计算。 一般用于加密解密
	// 二进制：只有0和1    逢二进一    1：1   2：10   3：11   4：100
	// & 两者都为 1 才为 1，否则为 0
	// | 两者有一个 1 则为 1，否则为 0
	// ^ 两者相同为 1，否则为 0
	// << 左移运算符   0011 1100 左移两位为 1111 0000
	// >> 右移运算符

	// 60：0011 1100
	// 13：0000 1101
	// 60 & 13 结果为：0000 1100

	var a int = 60
	var b int = 13
	fmt.Println(a & b)
	// %b 打印二进制结果
	fmt.Printf("%b\n", a&b)

	// 0011 1100 左移两位 1111 0000
	fmt.Printf("%b\n", a<<2)

	// 0011 1100 右移两位 0000 1111
	fmt.Printf("%b\n", a>>2)

	var x int
	var y float64
	// 输入   Scanf  格式化输入  Scanln 输入
	// & 代表取变量地址
	fmt.Scanln(&x, &y)
	// 输出   Printf 格式化打印  Println 打印
	fmt.Printf("x: %d, y: %f\n", x, y)

}
