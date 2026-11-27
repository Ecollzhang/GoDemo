package main

import (
	"student-growth/handlers"
	"student-growth/models"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	_ "student-growth/docs" // 必须导入生成的文档包
)

// @title 学生成长管理系统 API
// @version 1.0
// @description 这是一个包含用户模块和学生成长记录模块的后端 API 服务
// @host localhost:8080
// @BasePath /
func main() {
	// 1. 连接 SQLite 数据库（会在当前目录生成 student_growth.db 文件）
	db, err := gorm.Open(sqlite.Open("student_growth.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 2. 自动迁移数据表
	db.AutoMigrate(&models.User{}, &models.GrowthRecord{})
	handlers.SetDB(db)

	// 3. 初始化 Gin 引擎
	r := gin.Default()

	// 4. 配置路由分组
	// 用户管理
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handlers.Register)
	}

	// 成长记录管理
	growthGroup := r.Group("/growth")
	{
		growthGroup.POST("/add", handlers.AddGrowthRecord)
		growthGroup.GET("/list", handlers.ListGrowthRecords)
	}

	// 5. 挂载 Swagger 路由
	r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	// 6. 启动服务
	r.Run(":8080")
}