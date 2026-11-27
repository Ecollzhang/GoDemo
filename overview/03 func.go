
// Go 函数使用 func 关键字声明，支持多返回值——这是 Go 错误处理机制的基础。
package main
// Go 没有 try-catch 异常机制，而是通过函数返回 error 值来显式处理错误。这种设计让错误处理路径一目了然，避免了隐式控制流。
import (
	"fmt"
	"errors"
)

// add 函数接收两个 int 类型参数，返回它们的和
func add(a int, b int) int {
	return a + b
}

// divide 函数接收两个 float64 类型参数，返回它们的商和一个 error 类型的错误信息
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零") // 返回0和自定义错误信息
	}
	return a / b, nil // 返回商和nil表示没有错误
}

// swap 函数接收两个 string 类型参数，返回它们的交换值
func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	// 调用 add 函数
	sum := add(10, 20)
	fmt.Println("10 + 20 =", sum)

	// 调用 divide 函数
	quotient, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("10 / 2 =", quotient)
	}

	// 调用 swap 函数
	x := "Hello"
	y := "World"
	x, y = swap(x, y)
	fmt.Println("交换后 x =", x, "y =", y)

	// 测试 divide 函数的错误处理
	_, err = divide(10.0, 0.0)
	if err != nil {
		fmt.Println("错误:", err)
	}
}