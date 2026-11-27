/*
Go 语言条件语句
条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

Go 语言提供了以下几种条件判断语句：

语句				描述
if 语句			if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句		你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句		switch 语句用于基于不同条件执行不同动作。
select 语句		select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
*/



package main

import (
	"fmt"
	"time"
)

func main() {
	// ========== 1. if 语句 ==========
	fmt.Println("===== 1. if 语句 =====")
	score := 85
	if score >= 60 {
		fmt.Println("if 语句：成绩及格")
	}

	// ========== 2. if...else 语句 ==========
	fmt.Println("\n===== 2. if...else 语句 =====")
	age := 17
	if age >= 18 {
		fmt.Println("if...else：已成年")
	} else {
		fmt.Println("if...else：未成年")
	}

	// 还可以配合 else if 使用
	if score >= 90 {
		fmt.Println("等级：优秀")
	} else if score >= 75 {
		fmt.Println("等级：良好")
	} else if score >= 60 {
		fmt.Println("等级：及格")
	} else {
		fmt.Println("等级：不及格")
	}

	// ========== 3. if 嵌套语句 ==========
	fmt.Println("\n===== 3. if 嵌套语句 =====")
	username := "admin"
	password := "123456"
	isActive := true

	if username == "admin" {
		if password == "123456" {
			if isActive {
				fmt.Println("if 嵌套：登录成功，账号已激活")
			} else {
				fmt.Println("if 嵌套：账号未激活")
			}
		} else {
			fmt.Println("if 嵌套：密码错误")
		}
	} else {
		fmt.Println("if 嵌套：用户名不存在")
	}

	// ========== 4. switch 语句 ==========
	fmt.Println("\n===== 4. switch 语句 =====")
	day := 3
	switch day {
	case 1:
		fmt.Println("switch：星期一")
	case 2:
		fmt.Println("switch：星期二")
	case 3:
		fmt.Println("switch：星期三")
	case 4, 5:
		fmt.Println("switch：星期四或星期五")
	default:
		fmt.Println("switch：周末或其他")
	}

	// switch 也支持无表达式的写法（类似 if-else 链）
	switch {
	case score >= 90:
		fmt.Println("switch(无表达式)：优秀")
	case score >= 60:
		fmt.Println("switch(无表达式)：及格")
	default:
		fmt.Println("switch(无表达式)：不及格")
	}

	// ========== 5. select 语句 ==========
	fmt.Println("\n===== 5. select 语句 =====")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 ch1 的消息"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自 ch2 的消息"
	}()

	// select 会等待可运行的 case，多个就绪时随机选一个
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("select 收到：", msg1)
		case msg2 := <-ch2:
			fmt.Println("select 收到：", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("select：超时")
		}
	}

	// select 带 default 的写法（非阻塞）
	fmt.Println("\n===== select 带 default =====")
	ch3 := make(chan int)
	select {
	case v := <-ch3:
		fmt.Println("收到值：", v)
	default:
		fmt.Println("select default：通道无数据，不阻塞")
	}

	fmt.Println("\n程序执行完毕")
}