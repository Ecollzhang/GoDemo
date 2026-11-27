# GoDemo

一个用于系统学习 Go 语言的学习仓库，按照由浅入深的顺序组织，从最基础的语法开始，逐步过渡到并发、文件操作、接口等高级特性，最后通过一个完整的后端实战项目（学生成长管理系统）把所学内容串起来。

## 目录结构

```
GoDemo/
├── overview/   # 快速概览：用一个文件快速过一遍 Go 的核心概念
├── part1/      # 第一部分：语言基础（结构、类型、变量、流程控制）
├── part2/      # 第二部分：进阶语法（函数、指针、结构体、切片、Map）
├── part3/      # 第三部分：高级特性（递归、接口、泛型、并发、文件）
└── part4/      # 第四部分：实战项目
    └── myproject/
        └── student-growth/   # 学生成长管理系统（Gin + GORM + SQLite + Swagger）
```

## 学习路径

### overview — 快速概览

用一个文件一个主题的方式，快速体验 Go 的核心能力：

| 文件 | 主题 |
| --- | --- |
| `01 control.go` | 流程控制 |
| `02 error.go` | 错误处理 |
| `03 func.go` | 函数 |
| `04 goroutine.go` | 协程（goroutine） |
| `05 hello.go` | Hello World |
| `06 http.go` | HTTP 服务器（标准库 `net/http`） |
| `07 interface.go` | 接口 |
| `08 staticType.go` | 静态类型 |
| `09 structMethod.go` | 结构体与方法 |
| `010 tools.go` | 常用工具 |
| `011 varConst.go` | 变量与常量 |
| `012 workerpool.go` | 并发 Worker 池 |

### part1 — 语言基础

| 文件 | 主题 |
| --- | --- |
| `11 langStruct.go` | Go 语言结构 |
| `12 basic.go` | 基础语法 |
| `13 types.go` | 数据类型 |
| `14 var.go` | 变量 |
| `15 condition.go` | 条件语句 |
| `16 const.go` | 常量 |
| `17 loop.go` | 循环 |
| `18 operator.go` | 运算符 |

### part2 — 进阶语法

| 文件 | 主题 |
| --- | --- |
| `21 func.go` | 函数 |
| `22 scope.go` | 作用域 |
| `23 array.go` | 数组 |
| `24 pointer.go` | 指针 |
| `25 struct.go` | 结构体 |
| `26 slice.go` | 切片（Slice） |
| `27 range.go` | range 遍历 |
| `28 map.go` | Map |

### part3 — 高级特性

| 文件 | 主题 |
| --- | --- |
| `31 recursive.go` | 递归 |
| `32 typeTransform.go` | 类型转换 |
| `33 interface.go` | 接口 |
| `34 T.go` | 泛型 |
| `35 error.go` | 错误处理 |
| `36 goroutine.go` | 并发（goroutine） |
| `37 file.go` | 文件操作 |
| `38 regexp.go` | 正则表达式 |
| `39 assert.go` | 类型断言 |
| `40 inherit.go` | 组合与「继承」模拟 |

> `part3/` 目录下还包含部分示例运行产生的数据文件（如 `config.json`、`log.txt`、`output.txt` 等），属于练习 `file.go` 等文件操作时的产物。

## 实战项目：student-growth（part4）

一个基于 **Gin + GORM + SQLite + Swagger** 的学生成长管理系统后端 API，包含用户模块和学生成长记录模块。

### 技术栈

- [Gin](https://github.com/gin-gonic/gin) — Web 框架
- [GORM](https://gorm.io/) — ORM 框架
- [SQLite](https://github.com/glebarez/sqlite)（纯 Go 实现，无需 CGO）— 数据库
- [swaggo/swag](https://github.com/swaggo/swag) — 自动生成 Swagger 文档

### 项目结构

```
student-growth/
├── main.go                 # 入口：连接数据库、注册路由、启动服务
├── go.mod / go.sum         # 依赖管理
├── models/
│   └── models.go           # 数据模型：User、GrowthRecord
├── handlers/
│   ├── user.go             # 用户注册
│   └── growth.go           # 成长记录的增、查
├── docs/                   # Swagger 自动生成的文档
└── student_growth.db       # SQLite 数据库文件（运行时自动生成）
```

### API 一览

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/user/register` | 教师/管理员注册 |
| POST | `/growth/add` | 添加一条学生成长记录 |
| GET | `/growth/list` | 查询全部成长记录 |
| GET | `/swagger/*any` | Swagger 文档界面 |

### 运行方式

```bash
# 1. 进入项目目录
cd part4/myproject/student-growth

# 2. 安装依赖
go mod tidy

# 3. 启动服务（默认监听 :8080）
go run main.go

# 4. 访问 Swagger 文档
# http://localhost:8080/swagger/index.html
```

启动后会在当前目录自动创建 `student_growth.db` 数据库文件，并通过 GORM 自动迁移（AutoMigrate）创建 `users` 和 `growth_records` 数据表。

### 示例请求

注册用户：

```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"teacher1","password":"123456","role":"teacher"}'
```

添加成长记录：

```bash
curl -X POST http://localhost:8080/growth/add \
  -H "Content-Type: application/json" \
  -d '{"student_name":"张三","event_type":"学业表现","title":"期中考试数学第一","description":"期中考试数学成绩班级第一","score":98}'
```

## 如何运行示例

`overview/`、`part1/`、`part2/`、`part3/` 中的每个 `.go` 文件都是独立的 `main` 包程序，可以直接运行：

```bash
# 以某个示例为例
go run part3/40 inherit.go
```

> 需要本机已安装 Go（版本 >= 1.27）。部分涉及文件操作、HTTP 服务、并发等示例运行后会在当前目录产生输出文件或需要手动停止服务。
