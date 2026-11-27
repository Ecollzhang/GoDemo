package handlers

import (
	"net/http"
	"student-growth/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

// @Summary 教师/管理员注册
// @Description 注册一个新账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} map[string]interface{}
// @Router /user/register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		/* Gin 的 c.ShouldBindJSON(&user) 做的事情是：
			1.读取请求体（body）里的 JSON 字符串。
			2.解析这个 JSON。
			3.把解析出来的值，按字段名/tag 映射，写入 &user 指向的内存。
		*/
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败，用户名可能已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功", "username": user.Username})
}
