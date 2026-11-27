/*
变量来源于数学，是计算机语言中能储存计算结果或能表示值抽象概念。

变量可以通过变量名访问。

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用 var 关键字：
var identifier type
可以一次声明多个变量：
var identifier1, identifier2 type
*/

package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	// 没有初始化就为零值
    var b int
    fmt.Println(b)

    // bool 零值为 false
    var c bool
    fmt.Println(c)

	/*
	以下几种类型为 nil：
		var a *int
		var a []int
		var a map[string] int
		var a chan int
		var a func(string) int
		var a error // error 是接口
	*/
	

}