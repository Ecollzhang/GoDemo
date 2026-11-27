/*
	Go 没有类的概念，使用 struct（结构体）组织数据，通过为类型定义方法来实现面向对象风格的行为。方法是带有接收者（receiver）的函数。
*/
package main

import "fmt"

// User 定义用户结构体
type User struct {
	Name string // 首字母大写表示导出，外部包可以访问
	Email string
	age  int // 首字母小写表示未导出，外部包无法访问
}

// Greet 方法为 User 类型定义了一个方法，接收者是 User 的值类型
func (u User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and my email is %s. I am %d years old.", u.Name, u.Email, u.age)
}

// Birthday 方法为 User 类型定义了一个方法，接收者是 User 的指针类型
func (u *User) Birthday() {
	u.age++ // 修改接收者的 age 字段，必须使用指针接收者才能修改原始对象
	fmt.Printf("Happy Birthday! I am now %d years old.\n", u.age)
}

func main() {
	// 创建 User 实例
	user := User{Name: "Alice", Email: "alice@example.com", age: 25}
	fmt.Println(user.Greet())
	user.Birthday()
}