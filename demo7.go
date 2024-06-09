package main

import "fmt"

// func 方法名 入参 返回值
func test2(a int, b string) (string, int) {
	return b, a
}

// 可变参数 name... type     可变参数一个函数中最多只有一个，并且必须写在入参列表最后
func test3(a string, b ...int) string {
	for _, i := range b {
		fmt.Println(i)
	}
	return a
}

func main() {
	s, i := test2(1, "aa")
	t, _ := test2(1, "aa")
	fmt.Println(s, i, t)
	fmt.Println(test3("aaaa", 1, 2, 3))

	// 值传递：基本数据类型，arr，struct
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("before", arr)
	// 传递拷贝的值，而不是引用，原始值不变
	update(arr)
	// 原始不变
	fmt.Println("later", arr)

	// 切片：可以扩容的数组
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	// 引用传递：slice,map,chan...
	update2(slice)
	// 原始不变
	fmt.Println("later", slice)

}

func update(arr2 [4]int) {
	arr2[2] = 4
	fmt.Println("update later ", arr2)
}
func update2(arr2 []int) {
	arr2[2] = 4
	fmt.Println("update later ", arr2)
}
