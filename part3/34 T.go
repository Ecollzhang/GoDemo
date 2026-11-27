/*
Go 语言泛型
泛型是 Go 语言在 1.18 版本中引入的重要特性，它让开发者能够编写更加灵活和可重用的代码。
泛型主要通过以下两个核心概念来实现：
	类型参数（Type Parameters）：允许你在函数或类型定义中使用一个或多个类型作为参数。
	类型约束（Type Constraints）：指定类型参数必须满足的条件，确保在函数内部可以安全地操作这些类型。
概念			作用											示例
类型参数	在函数或类型名后声明，表示待定的类型。					[T any]
类型约束	定义类型参数必须满足的条件（如支持的操作符或方法）。	int，float64，comparable，constraints.Ordered，any
any			约束类型参数为任何类型。								[T any]
comparable	约束类型参数为可比较的类型。							[K comparable]

泛型（Generics）允许我们编写不依赖特定数据类型的代码。
在引入泛型之前，如果我们想要处理不同类型的数据，通常需要为每种类型编写重复的函数。
传统方式的局限性：
	// 处理 int 类型的函数
	func MaxInt(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 处理 float64 类型的函数
	func MaxFloat(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
使用泛型的解决方案：
	// 一个函数处理多种类型
	func Max[T comparable](a, b T) T {
		if a > b {
			return a
		}
		return b
	}

泛型语法详解
类型参数声明
泛型函数和类型通过类型参数列表来声明，语法为 [类型参数 约束]。
// 基本语法结构
func 函数名[T 约束](参数 T) 返回值类型 {
    // 函数体
}

type 类型名[T 约束] struct {
    // 结构体字段
}
类型参数命名约定
	通常使用大写字母：T、K、V、E 等
	T：表示 Type（类型）
	K：表示 Key（键）
	V：表示 Value（值）
	E：表示 Element（元素）

约束（Constraints）
约束定义了类型参数必须满足的条件，是泛型的核心概念。

内置约束
1. any 约束
any 是空接口 interface{} 的别名，表示任何类型都可以。
func PrintAny[T any](value T) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// 使用示例
PrintAny(42)        // Value: 42, Type: int
PrintAny("hello")   // Value: hello, Type: string
PrintAny(3.14)      // Value: 3.14, Type: float64

2. comparable 约束
comparable 表示类型支持 == 和 != 操作符。
func FindIndex[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

// 使用示例
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(FindIndex(numbers, 3))  // 输出: 2

names := []string{"Alice", "Bob", "Charlie"}
fmt.Println(FindIndex(names, "Bob"))  // 输出: 1

3. 联合约束（Union Constraints）
使用 | 运算符组合多个类型。
// 数字类型约束
type Number interface {
    int | int8 | int16 | int32 | int64 | 
    uint | uint8 | uint16 | uint32 | uint64 | 
    float32 | float64
}

func Add[T Number](a, b T) T {
    return a + b
}

// 使用示例
fmt.Println(Add(10, 20))        // 输出: 30
fmt.Println(Add(3.14, 2.71))    // 输出: 5.85

自定义约束
1. 方法约束
定义需要特定方法的约束。
// 定义 Stringer 约束
type Stringer interface {
    String() string
}

func PrintString[T Stringer](value T) {
    fmt.Println(value.String())
}

// 实现自定义类型
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// 使用示例
person := Person{Name: "Alice", Age: 25}
PrintString(person)  // 输出: Alice (25 years old)

2. 复杂约束
结合类型和方法要求。
// 要求类型是数字且实现 String() 方法
type NumericStringer interface {
    Number
    String() string
}
*/
package main

import (
	"fmt"
	"sync"
)

// 交换两个值
func Swap[T any](a, b T) (T, T) {
    return b, a
}

// 判断切片是否包含元素
func Contains[T comparable](slice []T, target T) bool {
    for _, item := range slice {
        if item == target {
            return true
        }
    }
    return false
}

// 去重函数
func Unique[T comparable](slice []T) []T {
    seen := make(map[T]bool)
    result := []T{}
    
    for _, item := range slice {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    return result
}


// 泛型栈实现
type Stack[T any] struct {
    elements []T
}

// 入栈
func (s *Stack[T]) Push(value T) {
    s.elements = append(s.elements, value)
}

// 出栈
func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
   
    lastIndex := len(s.elements) - 1
    value := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return value, true
}

// 查看栈顶元素
func (s *Stack[T]) Peek() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    return s.elements[len(s.elements)-1], true
}

// 判断栈是否为空
func (s *Stack[T]) IsEmpty() bool {
    return len(s.elements) == 0
}


// 线程安全的泛型映射
type SafeMap[K comparable, V any] struct {
    data map[K]V
    mutex sync.RWMutex
}

// 创建新的 SafeMap
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
    return &SafeMap[K, V]{
        data: make(map[K]V),
    }
}

// 设置键值对
func (m *SafeMap[K, V]) Set(key K, value V) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.data[key] = value
}

// 获取值
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    value, exists := m.data[key]
    return value, exists
}

// 删除键
func (m *SafeMap[K, V]) Delete(key K) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    delete(m.data, key)
}

// 获取所有键
func (m *SafeMap[K, V]) Keys() []K {
    m.mutex.RLock()
    defer m.mutex.RUnlock() // 不是"立刻执行"，而是"延迟到函数返回前才执行。为了避免忘记解锁
   
    keys := make([]K, 0, len(m.data))
    for key := range m.data {
        keys = append(keys, key)
    }
    return keys
}



// 使用示例
func main() {
    // Swap 示例
    a, b := 10, 20
    a, b = Swap(a, b)
    fmt.Printf("a=%d, b=%d\n", a, b)  // 输出: a=20, b=10
    
    // Contains 示例
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(Contains(numbers, 3))  // 输出: true
    
    // Unique 示例
    duplicates := []int{1, 2, 2, 3, 4, 4, 5}
    unique := Unique(duplicates)
    fmt.Println(unique)  // 输出: [1 2 3 4 5]



	// 整数栈
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)
   
    fmt.Println(intStack.Pop())  // 输出: 3 true
   
    // 字符串栈
    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
   
    fmt.Println(stringStack.Pop())  // 输出: world true


	// 创建字符串到整数的映射
    scores := NewSafeMap[string, int]()
    scores.Set("Alice", 95)
    scores.Set("Bob", 87)
   
    if score, exists := scores.Get("Alice"); exists {
        fmt.Printf("Alice's score: %d\n", score)  // 输出: Alice's score: 95
    }
   
    fmt.Println("Keys:", scores.Keys())  // 输出: Keys: [Alice Bob]
}