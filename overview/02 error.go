/*
	Go 编译器不会强制你处理 error，但忽略错误是一种不良实践。至少应在开发阶段记录或打印错误信息。
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    // 不推荐：忽略错误
    data, _ := os.ReadFile("config.txt")
    fmt.Println(string(data))

    // 推荐：显式处理错误
    data, err := os.ReadFile("config.txt")
    if err != nil {
        fmt.Printf("读取文件失败: %v\n", err)
        return
    }
    fmt.Println(string(data))
}