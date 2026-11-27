package main

import "fmt"
func main() {
	// 方式一：标准声明（显式类型）
	var age int = 30

	// 方式二：声明后赋值（显式类型）
	var height float64
	height = 1.75

	// 方式三：类型推断（隐式类型）
	var weight = 70.5 // 编译器根据赋值自动推断类型为 float64

	// 方式四：简短声明（隐式类型，必须在函数内部使用）
	country := "China" // 编译器根据赋值自动推断类型为 string

	// 方式五：多变量声明（隐式类型，必须在函数内部使用）
	x, y, z := 10, 20.5, "Hello" // 编译器根据赋值自动推断类型为 int、float64 和 string

	// 常量声明
	const pi = 3.14159 // 常量声明，类型为 float64

	// 输出变量和常量的值
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Weight:", weight)
	fmt.Println("Country:", country)
	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Println("Pi:", pi)
}