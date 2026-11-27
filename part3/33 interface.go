/*
Go 语言接口
	接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。
	Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
	Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具。
接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。

接口的特点
隐式实现：
	Go 中没有关键字显式声明某个类型实现了某个接口。
	只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。

接口类型变量：
	接口变量可以存储实现该接口的任意值。
	接口变量实际上包含了两个部分：
		动态类型：存储实际的值类型。
		动态值：存储具体的值。

零值接口：
	接口的零值是 nil。
	一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。

空接口：
	定义为 interface{}，可以表示任何类型。

接口的常见用法
	多态：不同类型实现同一接口，实现多态行为。
	解耦：通过接口定义依赖关系，降低模块之间的耦合。
	泛化：使用空接口 interface{} 表示任意类型。

接口定义和实现
接口定义使用关键字 interface，其中包含方法声明。
// 定义接口 
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

// 定义结构体
type struct_name struct {
   // variables 
}

// 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
   // 方法实现
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   // 方法实现
}


*/
package main

import (
        "fmt"
        "math"
)

/*
空接口
空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集。
	任意类型都实现了空接口。
	常用于需要存储任意类型数据的场景，如泛型容器、通用参数等。
*/
func printValue(val interface{}) {
        fmt.Printf("Value: %v, Type: %T\n", val, val)
}

// 定义接口
type Shape interface {
        Area() float64
        Perimeter() float64
}

// 定义一个结构体
type Circle struct {
        Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
        return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
        return 2 * math.Pi * c.Radius
}


/*
类型选择（type switch）
type switch 是 Go 中的语法结构，用于根据接口变量的具体类型执行不同的逻辑。
*/
func printType(val interface{}) {
        switch v := val.(type) {
        case int:
                fmt.Println("Integer:", v)
        case string:
                fmt.Println("String:", v)
        case float64:
                fmt.Println("Float:", v)
        default:
                fmt.Println("Unknown type")
        }
}

/*
接口组合
	接口可以通过嵌套组合，实现更复杂的行为描述。
*/
type Reader interface {
        Read() string
}

type Writer interface {
        Write(data string)
}
// 接口体里还可以直接嵌入另一个接口名，这叫接口嵌入。ReadWriter 的方法集 = Reader 的方法集 ∪ Writer 的方法集
type ReadWriter interface {
        Reader
        Writer
}
type File struct{}
func (f File) Read() string {
        return "Reading data"
}
func (f File) Write(data string) {
        fmt.Println("Writing data:", data)
}


func main() {
        c := Circle{Radius: 5}
        var s Shape = c // 接口变量可以存储实现了接口的类型
        fmt.Println("Area:", s.Area())
        fmt.Println("Perimeter:", s.Perimeter())

		printValue(42)         // int
        printValue("hello")    // string
        printValue(3.14)       // float64
        printValue([]int{1, 2}) // slice

		/*
		类型断言
		类型断言用于从接口类型中提取其底层值。

		基本语法:
		value := iface.(Type)
		iface 是接口变量。
		Type 是要断言的具体类型。
		如果类型不匹配，会触发 panic。
		*/
		var i interface{} = "hello"
        str := i.(string) // 类型断言
        fmt.Println(str)  // 输出：hello

		/*
		带检查的类型断言
		为了避免 panic，可以使用带检查的类型断言：
		value, ok := iface.(Type)
		ok 是一个布尔值，表示断言是否成功。
		如果断言失败，value 为零值，ok 为 false。
		*/
		var i2 interface{} = 42
        if str2, ok := i2.(string); ok {
                fmt.Println("String:", str2)
        } else {
                fmt.Println("Not a string")
        }

		printType(42)
        printType("hello")
        printType(3.14)
        printType([]int{1, 2, 3})

		var rw ReadWriter = File{} // 合法：File 满足 ReadWriter 的全部方法。File多一个方法没影响，少一个方法就赋值失败（编译错误）
        fmt.Println(rw.Read())
        rw.Write("Hello, Go!")

		/*
		动态值和动态类型
		接口变量实际上包含了两部分：
			动态类型：接口变量存储的具体类型。
			动态值：具体类型的值。
		*/
		var i3 interface{} = 42
        fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i3, i3)

		/*
		接口的零值
		接口的零值是 nil。
		当接口变量的动态类型和动态值都为 nil 时，接口变量为 nil。
		*/
		var i4 interface{}
        fmt.Println(i4 == nil) // 输出：true

}