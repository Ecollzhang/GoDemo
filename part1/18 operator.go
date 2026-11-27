/*
Go 语言运算符
运算符用于在程序运行时执行数学或逻辑运算。
Go 语言内置的运算符有：
	算术运算符
	关系运算符
	逻辑运算符
	位运算符
	赋值运算符
	其他运算符

&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。	*a; 是一个指针变量

运算符优先级
有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

优先级	运算符
5	* / % << >> & &^
4	+ - | ^
3	== != < <= > >=
2	&&
1	||

*/
package main

import "fmt"

func main() {

   var a int = 21
   var b int = 10
   var c int

   c = a + b
   fmt.Printf("第一行 - c 的值为 %d\n", c )
   c = a - b
   fmt.Printf("第二行 - c 的值为 %d\n", c )
   c = a * b
   fmt.Printf("第三行 - c 的值为 %d\n", c )
   c = a / b
   fmt.Printf("第四行 - c 的值为 %d\n", c )
   c = a % b
   fmt.Printf("第五行 - c 的值为 %d\n", c )
   a++
   fmt.Printf("第六行 - a 的值为 %d\n", a )
   a=21   // 为了方便测试，a 这里重新赋值为 21
   a--
   fmt.Printf("第七行 - a 的值为 %d\n", a )

   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   }
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* Lets change value of a and b */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }

   var a2 bool = true
   var b2 bool = false
   if ( a2 && b2 ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a2 || b2 ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a2 = false
   b2 = true
   if ( a2 && b2 ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a2 && b2) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }

   var a3 int = 21
   var c3 int

   c3 =  a3
   fmt.Printf("第 1 行 - =  运算符实例，c3 值为 = %d\n", c3 )

   c3 +=  a3
   fmt.Printf("第 2 行 - += 运算符实例，c3 值为 = %d\n", c3 )

   c3 -=  a3
   fmt.Printf("第 3 行 - -= 运算符实例，c3 值为 = %d\n", c3 )

   c3 *=  a3
   fmt.Printf("第 4 行 - *= 运算符实例，c3 值为 = %d\n", c3 )

   c3 /=  a3
   fmt.Printf("第 5 行 - /= 运算符实例，c3 值为 = %d\n", c3 )

   c3  = 200;

   c3 <<=  2
   fmt.Printf("第 6行  - <<= 运算符实例，c3 值为 = %d\n", c3 )

   c3 >>=  2
   fmt.Printf("第 7 行 - >>= 运算符实例，c3 值为 = %d\n", c3 )

   c3 &=  2
   fmt.Printf("第 8 行 - &= 运算符实例，c3 值为 = %d\n", c3 )

   c3 ^=  2
   fmt.Printf("第 9 行 - ^= 运算符实例，c3 值为 = %d\n", c3 )

   c3 |=  2
   fmt.Printf("第 10 行 - |= 运算符实例，c3 值为 = %d\n", c3 )

   var a4 int = 4
   var b4 int32
   var c4 float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a4 变量类型为 = %T\n", a4 );
   fmt.Printf("第 2 行 - b4 变量类型为 = %T\n", b4 );
   fmt.Printf("第 3 行 - c4 变量类型为 = %T\n", c4 );

   /*  & 和 * 运算符实例 */
   ptr = &a4     /* 'ptr' 包含了 'a4' 变量的地址 */
   fmt.Printf("a4 的值为  %d\n", a4);
   fmt.Printf("*ptr 为 %d\n", *ptr);
   /*
   	第 1 行 - a4 变量类型为 = int
	第 2 行 - b4 变量类型为 = int32
	第 3 行 - c4 变量类型为 = float32
	a4 的值为  4
	*ptr 为 4
	*/
}

/*
&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。	(A & B) 结果为 12, 二进制为 0000 1100
|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。	A << 2 结果为 240 ，二进制为 1111 0000
>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。	A >> 2 结果为 15 ，二进制为 0000 1111
*/

// package main

// import "fmt"

// func main() {

//    var a uint = 60      /* 60 = 0011 1100 */  
//    var b uint = 13      /* 13 = 0000 1101 */
//    var c uint = 0          

//    c = a & b       /* 12 = 0000 1100 */ 
//    fmt.Printf("第一行 - c 的值为 %d\n", c )

//    c = a | b       /* 61 = 0011 1101 */
//    fmt.Printf("第二行 - c 的值为 %d\n", c )

//    c = a ^ b       /* 49 = 0011 0001 */
//    fmt.Printf("第三行 - c 的值为 %d\n", c )

//    c = a << 2     /* 240 = 1111 0000 */
//    fmt.Printf("第四行 - c 的值为 %d\n", c )

//    c = a >> 2     /* 15 = 0000 1111 */
//    fmt.Printf("第五行 - c 的值为 %d\n", c )
// }

/*
以上实例运行结果：

第一行 - c 的值为 12
第二行 - c 的值为 61
第三行 - c 的值为 49
第四行 - c 的值为 240
第五行 - c 的值为 15
*/