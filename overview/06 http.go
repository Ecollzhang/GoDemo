/*
构建 HTTP 服务器
使用标准库 net/http 即可搭建一个功能完整的 Web 服务，无需任何第三方框架。
*/
/*
对于生产级 Web 服务，建议结合实际需求选择合适的框架。小型项目用标准库即可，需要路由分组、中间件等功能时可选用 Gin 或 Echo。
*/
package main
import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

// Response 定义JSON响应结构体
type Response struct {
	Code	int    `json:"code"` // 结构体指定 JSON 标签，确保字段名在 JSON 中为小写
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // omitempty 表示如果 Data 为空，则在 JSON 中省略该字段
}

// helloHandler 处理 /hello 路径的请求，返回 JSON 响应
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 仅允许 GET 请求
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从查询参数获取 name 参数，默认值为 "World"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// 创建响应数据
	response := Response{
		Code:    200,
		Message: "Success",
		Data:    fmt.Sprintf("Hello, %s!", name),
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	// 返回 JSON 响应
	json.NewEncoder(w).Encode(response)
}

func main() {
	// 注册 /hello 路径的处理函数
	http.HandleFunc("/hello", helloHandler)

	addr := ":8080" // 监听端口
	fmt.Printf("Starting server at http://localhost%s/hello\n", addr)
	fmt.Println("You can test it by visiting: http://localhost:8080/hello?name=YourName")

	// 启动 HHTP服务（阻塞，直到服务停止）
	log.Fatal(http.ListenAndServe(":8080", nil)) // 监听端口8080，nil表示使用默认的多路复用器
}
