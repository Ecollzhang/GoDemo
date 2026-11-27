package models

import "gorm.io/gorm"

// 用户表（系统使用者：老师或管理员）
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"` // 反引号（`）里的内容叫 Struct Tag（结构体标签）。它本质上是一段附带在字段上的元数据字符串，Go 语言本身不解析它，而是由反射（Reflection）机制在运行时读取，由各个框架（GORM、JSON、Gin 等）自己解释。
	Password string `json:"password" gorm:"not null"`
	Role string `json:"role" gorm:"default:teacher"`
}

// 学生成长记录表
type GrowthRecord struct {
	gorm.Model
	StudentName string `json:"student_name" gorm:"not null"`
	EventType   string  `json:"event_type"` // 类型：学业表现、品德修养、实践活动等
	Title       string  `json:"title"`      // 记录标题，如："期中考试数学第一"
	Description string  `json:"description"`// 详细描述
	Score       float64 `json:"score"`      // 可选：如果属于考试评测，可以打分
}
