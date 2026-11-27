/*
Go 错误处理
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
Go 语言的错误处理采用显式返回错误的方式，而非传统的异常处理机制。这种设计使代码逻辑更清晰，便于开发者在编译时或运行时明确处理错误。
Go 的错误处理主要围绕以下机制展开：
	error 接口：标准的错误表示。
	显式返回值：通过函数的返回值返回错误。
	自定义错误：可以通过标准库或自定义的方式创建错误。
	panic 和 recover：处理不可恢复的严重错误。

error 接口
Go 标准库定义了一个 error 接口，表示一个错误的抽象。
error 类型是一个接口类型，这是它的定义：
type error interface {
    Error() string
}
	实现 error 接口：任何实现了 Error() 方法的类型都可以作为错误。
	Error() 方法返回一个描述错误的字符串。

使用 errors 包创建错误
我们可以在编码中通过实现 error 接口类型来生成错误信息。
创建一个简单错误：
	package main
	import (
		"errors"
		"fmt"
	)
	func main() {
		err := errors.New("this is an error")
		fmt.Println(err) // 输出：this is an error
	}
函数通常在最后的返回值中返回错误信息，使用 errors.New 可返回一个错误信息：
	func Sqrt(f float64) (float64, error) {
		if f < 0 {
			return 0, errors.New("math: square root of negative number")
		}
		// 实现
	}
在下面的例子中，我们在调用 Sqrt 的时候传递的一个负数，然后就得到了 non-nil 的 error 对象，将此对象与 nil 比较，结果为 true，所以 fmt.Println(fmt 包在处理 error 时会调用 Error 方法)被调用，以输出错误，请看下面调用的示例代码：
	result, err:= Sqrt(-1)
	if err != nil {
	fmt.Println(err)
	}

显式返回错误
Go 中，错误通常作为函数的返回值返回，开发者需要显式检查并处理。
显式返回错误：
	package main
	import (
			"errors"
			"fmt"
	)
	func divide(a, b int) (int, error) {
			if b == 0 {
					return 0, errors.New("division by zero")
			}
			return a / b, nil
	}
	func main() {
			result, err := divide(10, 0)
			if err != nil {
					fmt.Println("Error:", err)
			} else {
					fmt.Println("Result:", result)
			}
	}

自定义错误
通过定义自定义类型，可以扩展 error 接口。
自定义错误类型：
*/

package main

import (
	"fmt"
)

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func Divide2(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			Dividend: varDividee,
			Divisor:  varDivider,
		}
		errorMsg = dData.Error()
		return
	}
	return varDividee / varDivider, ""
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	if result, errorMsg := Divide2(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	if _, errorMsg := Divide2(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}