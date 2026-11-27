// Go 的接口是隐式实现的——只要一个类型实现了接口中所有方法，它就自动满足该接口，无需显式声明 implements。这种设计极大地降低了代码耦合度。
// Go 的接口设计鼓励「小接口」——通常只包含 1-3 个方法。标准库中的 io.Reader（仅一个 Read 方法）和 io.Writer（仅一个 Write 方法）是最经典的例子。
package main
import ("fmt" "math")

// Shape 接口定义了一个计算面积的方法 Area
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 结构体实现了 Shape 接口
type Circle struct {
	Radius float64
}

// Circle 的 Area 方法实现了 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 的 Perimeter 方法实现了 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 结构体实现了 Shape 接口
type Rectangle struct {
	Width, Height float64
}

// Rectangle 的 Area 方法实现了 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 的 Perimeter 方法实现了 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// PrintShapeInfo 函数接收一个 Shape 接口类型的参数，并打印其面积和周长
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// 创建 Circle 和 Rectangle 实例
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// 调用 PrintShapeInfo 函数
	fmt.Println("Circle Info:")
	PrintShapeInfo(circle)
	fmt.Println("Rectangle Info:")
	PrintShapeInfo(rectangle)
}
