package main

import "fmt"

func main() {
	var a int = 101
	if a < 100 {
		fmt.Println("a is less than")
	} else if a > 100 {
		fmt.Println("a is greater than")
	} else {
		fmt.Println("a is equals")
	}

	// switch 默认每个case中都有break，匹配到后不会继续执行
	// 如果需要执行多个则用 fallthrough 关键字，用了之后不会进行判断下一个case是否满足条件而是会直接执行（只能贯穿一个）
	// 若想中止穿透用 break 关键字
	switch a {
	case 100:
		fmt.Println("a is equals 100")
		fallthrough
	case 101:
		fmt.Println("a is less than 100")
		fallthrough
	case 102, 103, 104:
		if a == 100 {
			break
		}
		fmt.Println("a is greater than 100")
	default:
		fmt.Println("a is unknown")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	b := 10
	// 三个条件任一都能省略，如果都省略就同while了
	for {
		b++
		if b == 12 {
			continue
		}
		fmt.Println(b)
		if b > 13 {
			break
		}
	}

	// break 只会跳出当前循环，多重嵌套循环中只会跳出一层
	for c := 0; c < 3; c++ {
		for d := 0; d < 3; d++ {
			fmt.Println(c, d)
			if d == 1 {
				break
			}
		}
		fmt.Println(c)
	}
}
