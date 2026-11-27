/*
Go 语言循环语句
在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。

Go 语言提供了以下几种类型循环处理语句：
循环类型	描述
for 循环	重复执行语句块
循环嵌套	在 for 循环中嵌套一个或多个 for 循环

循环控制语句
循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：
控制语句			描述
break 语句		经常用于中断当前 for 循环或跳出 switch 语句
continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句		将控制转移到被标记的语句。
*/

package main

import "fmt"

func main() {
	// ========== 1. for 循环：三种常见形式 ==========
	fmt.Println("===== 1. for 循环的三种形式 =====")

	// 形式一：标准三段式（初始化; 条件; 后置语句）
	fmt.Print("标准 for: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 形式二：只有条件（类似 while）
	fmt.Print("while 式: ")
	n := 0
	for n < 5 {
		fmt.Print(n, " ")
		n++
	}
	fmt.Println()

	// 形式三：无限循环（配合 break 使用）
	fmt.Print("无限循环: ")
	count := 0
	for {
		if count >= 5 {
			break
		}
		fmt.Print(count, " ")
		count++
	}
	fmt.Println()

	// ========== 2. for range：遍历数组、切片、map、字符串 ==========
	fmt.Println("\n===== 2. for range 遍历 =====")

	// 遍历切片
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Printf("切片 index=%d value=%d\n", index, value)
	}

	// 遍历 map（顺序不固定）
	scores := map[string]int{"Alice": 90, "Bob": 85}
	for name, score := range scores {
		fmt.Printf("map key=%s value=%d\n", name, score)
	}

	// 遍历字符串（按 rune 遍历，支持中文）
	for i, ch := range "Go语言" {
		fmt.Printf("字符串 index=%d rune=%c\n", i, ch)
	}

	// ========== 3. 循环嵌套 ==========
	fmt.Println("\n===== 3. 循环嵌套 =====")
	// 打印九九乘法表的前 3 行
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d  ", j, i, i*j)
		}
		fmt.Println()
	}

	// ========== 4. break 语句 ==========
	fmt.Println("\n===== 4. break 语句 =====")

	// 在 for 中 break：找到第一个能被 7 整除的数就停止
	fmt.Print("找第一个能被7整除的数: ")
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			fmt.Println(i)
			break
		}
	}

	// 在 switch 中 break（Go 中 switch 默认自动 break，这里演示显式写法）
	fmt.Print("switch break: ")
	switch 2 {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
		break // 显式 break，效果与默认一致
	case 3:
		fmt.Println("三")
	}

	// break 跳出多层循环：配合标签
	fmt.Print("跳出外层循环: ")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer // 直接跳出 outer 标记的外层循环
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// ========== 5. continue 语句 ==========
	fmt.Println("\n===== 5. continue 语句 =====")

	// 跳过偶数，只打印奇数
	fmt.Print("打印1-10的奇数: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 跳过本次循环剩余部分
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// continue 配合标签：跳到外层循环的下一轮
	fmt.Println("continue 配合标签（跳过 i==1 的整行）:")
outer2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 {
				continue outer2 // 直接进入外层循环下一轮
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// ========== 6. goto 语句 ==========
	fmt.Println("\n===== 6. goto 语句 =====")

	// 使用 goto 跳出多层嵌套循环
	fmt.Print("goto 跳出多层循环: ")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 6 {
				fmt.Printf("找到 i=%d, j=%d\n", i, j)
				goto endLoop // 直接跳转到 endLoop 标签
			}
		}
	}
endLoop:
	fmt.Println("goto 已跳出循环")

	// goto 实现类似循环的效果（不推荐，仅演示）
	fmt.Print("goto 模拟循环: ")
	k := 0
loop:
	if k < 5 {
		fmt.Print(k, " ")
		k++
		goto loop
	}
	fmt.Println()

	fmt.Println("\n程序执行完毕")
}
