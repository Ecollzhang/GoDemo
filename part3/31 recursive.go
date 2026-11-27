/*
Go 语言递归函数
递归是一种函数直接或间接调用自身的编程技术。

递归函数通常包含两个部分：
	基准条件（Base Case）：这是递归的终止条件，防止函数无限调用自身。
	递归条件（Recursive Case）：这是函数调用自身的部分，用于将问题分解为更小的子问题。
在 Go 语言中，递归的使用与其他语言类似，但需要注意 Go 的一些特性。
语法格式如下：
	func recursion() {
		recursion() // 函数调用自身
	}

	func main() {
		recursion()
	}
Go 语言支持递归，但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。


递归的优缺点
优点
	简洁性：递归代码通常比迭代代码更简洁，易于理解。
	问题分解：递归天然适合解决可以分解为相似子问题的问题，如树遍历、分治算法等。
缺点
	性能开销：递归调用会占用栈空间，可能导致栈溢出，尤其是在深度递归时。
	调试困难：递归代码可能较难调试，尤其是在递归深度较大时。
*/

package main

import "fmt"

// 递归函数计算阶乘
func factorial(n int) int {
    // 基准条件
    if n == 0 {
        return 1
    }
    // 递归条件
    return n * factorial(n-1)
}

func fibonacci(n int) int {
  if n < 2 {
   return n
  }
  return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 输出: 120
	
	var i int
    for i = 0; i < 10; i++ {
       fmt.Printf("%d\t", fibonacci(i))
    }
}
