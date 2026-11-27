package main

import "fmt"
/*
静态类型与类型推断
Go 是静态类型语言，所有变量类型在编译时即确定。但与 C/Java 不同的是，Go 支持类型推断——编译器可以自动推导变量类型，开发者无需反复声明类型名。

这种设计在保持类型安全的同时，让代码写起来像动态语言一样简洁。
*/
func main() {
	// 完整声明：显式指定类型
	var name string = "runoob"

	// 类型推断：根据值自动推断变量类型
	var version = 1.17 // 推导为float64类型
	count := 100 // 简短声明：自动推导类型为int，必须在函数内部使用
	message := "Hello, Go!" // 简短声明：自动推导类型为string，必须在函数内部使用

	// 多变量声明
	x, y, z := 1, 2, 3 //同时声明多个变量：自动推导类型为int，必须在函数内部使用
	fmt.Println(name, version, count, message, x, y, z)
	
}
