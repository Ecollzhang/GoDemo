/*
Go 语言切片(Slice)
Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

定义切片
你可以声明一个未指定大小的数组来定义切片：
	var identifier []type

切片不需要说明长度。
或使用 make() 函数来创建切片:
	var slice1 []type = make([]type, len)
也可以简写为
	slice1 := make([]type, len)

也可以指定容量，其中 capacity 为可选参数。
	make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 } 
直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。

s := arr[:] 
初始化切片 s，是数组 arr 的引用。

s := arr[startIndex:endIndex] 
将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。

s := arr[startIndex:] 
默认 endIndex 时将表示一直到arr的最后一个元素。

s := arr[:endIndex] 
默认 startIndex 时将表示从 arr 的第一个元素开始。

s1 := s[startIndex:endIndex] 
通过切片 s 初始化切片 s1。

s :=make([]int,len,cap) 
通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

cap 和 len 并不总是一样——它们只在某些创建方式下恰好相等。cap 的核心作用是告诉 append 还有多少空闲空间可用，从而决定是否需要重新分配底层数组。
	s1 := make([]int, 3, 10)   // len=3, cap=10  ← 不一样
	s2 := []int{1, 2, 3}       // len=3, cap=3   ← 一样
	s3 := arr[1:3]             // len=2, cap=len(arr)-1  ← 通常不一样
	s4 := make([]int, 5)       // len=5, cap=5   ← 一样（省略 cap 时默认相等）

append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

*/
package main

import "fmt"

func main() {
	var numbers = make([]int,3,5)
	printSlice(numbers)

	var numbers2 []int
	printSlice(numbers2) // len=0 cap=0 slice=[]
	if(numbers2 == nil){
		fmt.Printf("切片是空的")
	}

	/* 创建切片 */
   nums := []int{0,1,2,3,4,5,6,7,8}  
   printSlice(nums)

   /* 打印原始切片 */
   fmt.Println("nums ==", nums)

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("nums[1:4] ==", nums[1:4])

   /* 默认下限为 0*/
   fmt.Println("nums[:3] ==", nums[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("nums[4:] ==", nums[4:])

   nums1 := make([]int,0,5)
   printSlice(nums1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   nums2 := nums[:2]
   printSlice(nums2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   nums3 := nums[2:5]
   printSlice(nums3)

   /*
   len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
	nums == [0 1 2 3 4 5 6 7 8]
	nums[1:4] == [1 2 3]
	nums[:3] == [0 1 2]
	nums[4:] == [4 5 6 7 8]
	len=0 cap=5 slice=[]
	len=2 cap=9 slice=[0 1]
	len=3 cap=7 slice=[2 3 4]
	*/

	var numberss []int
	printSlice(numberss)

	/* 允许追加空切片 */
	numberss = append(numberss, 0)
	printSlice(numberss)

	/* 向切片添加一个元素 */
	numberss = append(numberss, 1)
	printSlice(numberss)

	/* 同时添加多个元素 */
	numberss = append(numberss, 2,3,4)
	printSlice(numberss)

	/* 创建切片 numberss1 是之前切片的两倍容量*/
	numberss1 := make([]int, len(numberss), (cap(numberss))*2)

	/* 拷贝 numberss 的内容到 numberss1 */
	copy(numberss1,numberss)
	printSlice(numberss1)  
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
   // len=3 cap=5 slice=[0 0 0]


}
