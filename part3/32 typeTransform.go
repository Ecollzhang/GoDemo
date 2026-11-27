/*
Go 语言类型转换
类型转换用于将一种数据类型的变量转换为另外一种类型的变量。

Go 语言类型转换基本格式如下：
	type_name(expression)
	type_name 为类型，expression 为表达式。

*/

package main

import (
	"fmt"
	"strconv"
)

// 定义一个接口 Writer
type Writer interface {
    Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
    str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
    sw.str += string(data)
    return len(data), nil
}

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	
	mean = float32(sum)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)

   	var str string = "10"
	var num int
	num, _ = strconv.Atoi(str)
	fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	/*
		以上代码将字符串变量 str 转换为整型变量 num。
		注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
	*/
	str2 := "123"
    num2, err2 := strconv.Atoi(str)
    if err2 != nil {
        fmt.Println("转换错误:", err2)
    } else {
        fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str2, num2)
    }

	// 以下实例将整数转换为字符串
	num3 := 123
    str3 := strconv.Itoa(num3)
    fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num3, str3)

	// 以下实例将字符串转换为浮点数
	str4 := "3.14"
    num4, err4 := strconv.ParseFloat(str4, 64)
    if err4 != nil {
        fmt.Println("转换错误:", err4)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str4, num4)
    }

	// 将浮点数转换为字符串
	num5 := 3.14
    str5 := strconv.FormatFloat(num5, 'f', 2, 64)
    fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num5, str5)

	/*
		接口类型转换
			接口类型转换有两种情况：类型断言和类型转换。

		类型断言
			类型断言用于将接口类型转换为指定类型，其语法为
		
		value.(type) 
			或者 
		value.(T)

		其中 value 是接口类型的变量，type 或 T 是要转换成的类型。
		如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。
	*/
	var i interface{} = "Hello, World"
    str6, ok := i.(string)
    if ok {
        fmt.Printf("'%s' is a string\n", str6)
    } else {
        fmt.Println("conversion failed")
    }
	/*
		以上实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息。
	*/

	/*
	类型转换
		类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
		T(value)
		T 是目标接口类型，value 是要转换的值。
	*/
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
    var w Writer = &StringWriter{}
    // 将 Writer 接口类型转换为 StringWriter 类型
    sw := w.(*StringWriter)
    // 修改 StringWriter 的字段
    sw.str = "Hello, World"
    // 打印 StringWriter 的字段值
    fmt.Println(sw.str)
	/*
	解析：
		定义接口和结构体：
			Writer 接口定义了 Write 方法。
			StringWriter 结构体实现了 Write 方法。
		类型转换：
			将 StringWriter 实例赋值给 Writer 接口变量 w。
			使用 w.(*StringWriter) 将 Writer 接口类型转换为 StringWriter 类型。
		访问字段：
			修改 StringWriter 的字段 str，并打印其值。
	*/
	/* 
	空接口类型
		空接口 interface{} 可以持有任何类型的值。在实际应用中，空接口经常被用来处理多种类型的值。
	*/
	printValue(42)
    printValue("hello")
    printValue(3.14)

}

func printValue(v interface{}) {
    switch v := v.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

