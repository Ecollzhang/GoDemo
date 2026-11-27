/*
Go 语言文件处理
文件处理是 Go 语言中最常见的操作之一——读取配置、写入日志、数据持久化都离不开它。Go 的标准库提供了一套简洁而强大的文件 I/O 接口，覆盖从单次读写到流式处理的各种场景。
与文件处理相关的核心包有 5 个，各有分工：
	io	通用 Reader/Writer 接口
	bufio	缓冲读写，性能优化
	filepath	跨平台路径处理
1. 文件创建
os.Create 创建一个新文件。如果文件已存在，会被截断（清空内容）。返回的文件对象必须关闭以释放系统资源
2. 文件打开与关闭
os 包提供三种打开文件的方式，适用于不同场景：
函数					模式					说明
os.Open(name)			只读				最简单的方式，只读打开
os.Create(name)		读写 + 创建/截断		 创建新文件或清空已有文件
os.OpenFile(name, flag, perm)	自定义		可指定读写、追加、创建等标志

OpenFile 标志位说明
标志				说明
os.O_RDONLY		只读（默认）
os.O_WRONLY			只写
os.O_RDWR			读写
os.O_APPEND		追加模式，写入内容追加到文件末尾
os.O_CREATE		文件不存在时创建
os.O_TRUNC		打开时清空文件内容
os.O_EXCL		与 CREATE 配合，文件已存在时报错

3. 文件读取
Go 提供了三种主要的文件读取方式，根据文件大小和处理需求选择：
三种文件读取方式
一次性读取 os.ReadFile("file.txt") 适合小文件（配置、JSON等） 自动打开/关闭，最简洁
逐行读取 bufio.NewScanner(file) 适合大文件（日志、CSV等） 内存友好，逐行处理
流式读取 io.ReadAll(reader) 适合网络响应、已打开的文件 通用 Reader 接口
4. 文件写入
Go 提供了多种写入方式，从简单的一次性写入到高性能的缓冲写入
5. 文件追加写入
使用 os.O_APPEND 标志在文件末尾追加内容，而不覆盖已有数据
6. 文件删除与重命名
package main
import (
	"log"
	"os"
)
func main() {
	// 重命名 / 移动文件
	err := os.Rename("old.txt", "new.txt")
	if err != nil {
			log.Fatal("重命名失败:", err)
	}
	log.Println("重命名成功")

	// 删除文件
	err = os.Remove("temp.txt")
	if err != nil {
			log.Fatal("删除失败:", err)
	}
	log.Println("删除成功")
}
7. 文件信息与检查
7.1 获取文件信息
package main
import (
	"fmt"
	"log"
	"os"
)

func main() {
	info, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文件名:    ", info.Name())
	fmt.Println("大小:      ", info.Size(), "字节")
	fmt.Println("权限:      ", info.Mode())
	fmt.Println("修改时间:  ", info.ModTime())
	fmt.Println("是目录吗:  ", info.IsDir())
}
7.2 检查文件是否存在
Go 没有专门的"文件存在"函数，通过 os.Stat + os.IsNotExist 判断：
package main
import (
	"fmt"
	"os"
)
// fileExists 检查文件是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
func main() {
	if fileExists("test.txt") {
			fmt.Println("文件存在")
	} else {
			fmt.Println("文件不存在")
	}
}
8. 目录操作
package main
import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建单个目录
	err := os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 递归创建多级目录（父目录不存在也会自动创建）
	err = os.MkdirAll("path/to/deep/dir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 读取目录内容
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		mark := "&#x1f4c4;"
		if entry.IsDir() {
				mark = "&#x1f4c1;"
		}
		fmt.Printf("%s %s\n", mark, entry.Name())
	}

	// 删除空目录
	os.Remove("newdir")

	// 递归删除目录及其所有内容（&#x26a0; 谨慎使用）
	os.RemoveAll("path")
}
9. 高级操作
9.1 文件复制
使用 io.Copy 在两个文件之间高效复制数据
package main
import (
	"io"
	"log"
	"os"
)
// copyFile 复制文件的通用函数
func copyFile(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
			return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
			return 0, err
	}
	defer dstFile.Close()

	// io.Copy 自动管理缓冲区，高效复制
	return io.Copy(dstFile, srcFile)
}

func main() {
	n, err := copyFile("source.txt", "destination.txt")
	if err != nil {
			log.Fatal(err)
	}
	log.Printf("复制完成，共 %d 字节", n)
}
9.2 临时文件与目录
package main
import (
	"fmt"
	"log"
	"os"
)
func main() {
	// 创建临时文件（自动命名，前缀为 "app-"）
	tmpFile, err := os.CreateTemp("", "app-*.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // 用完后清理
	fmt.Println("临时文件:", tmpFile.Name())

	// 创建临时目录
	tmpDir, err := os.MkdirTemp("", "app-*")
	if err != nil {
			log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // 用完后清理
	fmt.Println("临时目录:", tmpDir)
}
9.3 递归遍历目录
filepath.Walk 可以递归遍历目录树中的所有文件和子目录
9.4 跨平台路径拼接
永远不要用字符串拼接路径（如 dir + "/" + file），使用 filepath.Join 确保跨平台兼容
	// 错误：Windows 上路径分隔符是 \
	path := "dir" + "/" + "file.txt"
	// 正确：自动适配当前操作系统
	path := filepath.Join("dir", "file.txt")

*/

package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
	"time"
)

func main() {
	// os.Create 创建文件（已存在则清空）
	file, err := os.Create("test.txt")
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close() // defer 确保函数结束时关闭文件
	// defer file.Close() 是 Go 文件操作的最佳实践。defer 会确保即使后续代码出现 panic，文件也会被正确关闭，避免文件描述符泄漏。

	// 写入内容验证创建成功
	file.WriteString("文件创建成功\n")
	log.Println("文件创建成功")

	// 方式1：只读打开
	file2, err2 := os.Open("example.txt")
	if err2 != nil {
			fmt.Println("打开失败:", err2)
			return
	}
	defer file2.Close()
	fmt.Println("文件打开成功")

	// 方式2：OpenFile 自定义模式
	// os.O_WRONLY  只写
	// os.O_CREATE  不存在则创建
	// os.O_APPEND  追加模式
	// os.O_TRUNC   存在则截断
	f, err3 := os.OpenFile("log.txt",
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err3 != nil {
			fmt.Println("打开失败:", err3)
			return
	}
	defer f.Close()
	f.WriteString("追加一行日志\n")

	/*3.1 一次性读取（小文件推荐）
	os.ReadFile 是最简洁的方式——自动打开、读取、关闭文件，一步到位*/
	// 一次性读取整个文件（Go 1.16+）
	data, err4 := os.ReadFile("config.json")
	if err4 != nil {
			log.Fatal(err4)
	}
	fmt.Println(string(data))

	/*3.2 逐行读取（大文件推荐）
	对于大文件，使用 bufio.Scanner 逐行处理，避免一次性加载到内存*/
	file5, err5 := os.Open("large-file.log")
	if err5 != nil {
			log.Fatal(err5)
	}
	defer file5.Close()

	// 逐行扫描
	scanner := bufio.NewScanner(file5)
	lineNum := 0
	// 注意：bufio.Scanner 默认最大行长度为 64KB。如果文件中有超长行，需要在调用 Scan() 前设置 scanner.Buffer(make([]byte, 0), maxSize) 来增大缓冲区。
	for scanner.Scan() {
			lineNum++
			line := scanner.Text() // 获取当前行内容
			fmt.Printf("第 %d 行: %s\n", lineNum, line)
	}

	// 检查扫描过程中是否出错
	if err5 := scanner.Err(); err5 != nil {
			log.Fatal("读取错误:", err5)
	}

	/*
	3.3 使用 io.ReadAll 读取
	当你已经有一个打开的 io.Reader（如网络响应、已打开的文件），可以使用 io.ReadAll
	*/
	file6, err6 := os.Open("example.txt")
	if err6 != nil {
			log.Fatal(err6)
	}
	defer file6.Close()

	// 从 Reader 读取所有数据
	data6, err6 := io.ReadAll(file6)
	if err6 != nil {
			log.Fatal(err6)
	}

	fmt.Println(string(data6))

	/*
	4.1 一次性写入
		os.WriteFile 将数据一次性写入文件（覆盖原有内容）
	*/
	content := []byte("Hello, Go!\n这是第二行\n")
	// 0644: 所有者读写，其他用户只读
	err7 := os.WriteFile("output.txt", content, 0644)
	if err7 != nil {
			log.Fatal(err7)
	}
	log.Println("写入成功")

	/*
	4.2 使用 File 对象写入
	通过文件对象写入，可以分多次写入内容
	*/
	file8, err8 := os.Create("output.txt")
	if err8 != nil {
			log.Fatal(err8)
	}
	defer file8.Close()
	// 方式1：写入字符串
	file8.WriteString("直接写入字符串\n")
	// 方式2：写入字节切片
	data = []byte("写入字节切片\n")
	file8.Write(data)
	// 方式3：格式化写入
	fmt.Fprintf(file8, "格式化写入: %d + %d = %d\n", 3, 4, 7)

	/*
	4.3 缓冲写入（大量数据推荐）
		bufio.Writer 会先将数据写入内存缓冲区，攒够后再批量写入磁盘，显著减少 I/O 次数。务必在结束前调用 Flush()
	*/
	file9, err9 := os.Create("buffered-output.txt")
	if err9 != nil {
			log.Fatal(err9)
	}
	defer file9.Close()

	// 创建带缓冲的写入器
	writer := bufio.NewWriter(file9)

	// 写入多行（先进入缓冲区，不会立即写磁盘）
	for i := 0; i < 1000; i++ {
			writer.WriteString("这是第 " + itoa(i) + " 行\n")
	}

	// Flush 将缓冲区剩余数据写入文件
	if err9 := writer.Flush(); err9 != nil {
			log.Fatal("刷新缓冲区失败:", err9)
	}
	log.Println("缓冲写入完成")

	// O_APPEND: 追加模式
	// O_CREATE: 不存在则创建
	// O_WRONLY: 只写
	file10, err10 := os.OpenFile("app.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err10 != nil {
			log.Fatal(err10)
	}
	defer file10.Close()

	// 每次运行都会追加，不会覆盖
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	file10.WriteString("[" + timestamp + "] 应用启动\n")

}
// 简单的 int 转 string 辅助函数
func itoa(n int) string {
        return fmt.Sprintf("%d", n)
}
