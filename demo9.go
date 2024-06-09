package main

import "fmt"

func main() {
	// 函数本身是一种数据类型 func()， 函数名后不加（）函数就是一个变量
	fmt.Printf("%T\n", main)

	var callAgain func(name string)
	callAgain = callName
	callAgain("John")
	// 两个函数变量的地址一样，代表函数其实也是引用传递，callAgain和callName都是函数指针，指向函数体的内存地址
	fmt.Println(callAgain)
	fmt.Println(callName)

	// 匿名函数，然后f3指向这个匿名函数，再调用这个f3函数
	f3 := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	f3("David")

	// 匿名函数直接调用自身
	func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}("David too")

	// 也可以有返回值
	num := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(num)

	// 回调函数，把函数当参数传入
	result2 := add(8, 3, sub)
	fmt.Println(result2)

	result := add(3, 4, func(a, b int) int {
		return a + b
	})
	fmt.Println(result)
}

func callName(name string) {
	fmt.Println(name)
}

func add(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

func sub(a, b int) int {
	return a - b
}
