package main

import "fmt"

/*
:= (短变量声明)：用于声明一个新变量并同时赋予初始值。编译器会根据右侧的值自动推断变量类型。

= (赋值)：用于修改一个已经声明过的变量的值。它不创建新变量，只是改变现有变量的内容。
*/

func main() {
	// if 语句：条件表达式不需要括号，但大括号必须有
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 75 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// for 循环一：标准三要素形式
	for i := 0; i < 5; i++ {
		fmt.Println("for 循环一，i =", i)
	}

	// for 循环二：类似 while 的形式
	j := 0
	for j < 5 {
		fmt.Println("for 循环二，j =", j)
		j++
	}

	// for 循环三：range遍历
	fruits := []string{"苹果", "香蕉", "橘子"}
	for index, fruit := range fruits {
		fmt.Printf("for 循环三，index = %d, fruit = %s\n", index, fruit)
	}

	// switch 语句：条件表达式不需要括号，但大括号必须有，每个case自动break，无需显式break
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	default:
		fmt.Println("其他")
	}
}