package main

import "fmt"

func main() {
	/*
		闭包
		匿名函数是闭包：它们可以引用在外层函数中定义的变量。
		然后，这些变量在外层的函数和匿名函数之间共享，只要它们还可以访问，它们就会继续存在。
	*/

	adder := createAdder()
	fmt.Println("Adder:", adder(1))
	fmt.Println("Adder:", adder(3))

	//新创建了一个函数，不会共享之前的
	adder2 := createAdder()
	fmt.Println("Adder2:", adder2(1))
	fmt.Println("Adder2:", adder2(3))
}

func createAdder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
