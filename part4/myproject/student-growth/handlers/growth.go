package handlers

import (
	"net/http"
	"student-growth/models"

	"github.com/gin-gonic/gin"
)

// @Summary 添加学生成长记录
// @Description 录入一条学生的成长轨迹记录
// @Tags 成长记录管理
// @Accept json
// @Produce json
// @Param record body models.GrowthRecord true "成长记录信息"
// @Success 200 {object} models.GrowthRecord
// @Router /growth/add [post]
func AddGrowthRecord(c *gin.Context) {
	var record models.GrowthRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&record)
	c.JSON(http.StatusOK, record)
}

// @Summary 获取所有学生成长记录
// @Description 查询全部的成长轨迹列表（后续可扩展分页和按学生查询）
// @Tags 成长记录管理
// @Produce json
// @Success 200 {array} models.GrowthRecord
// @Router /growth/list [get]
func ListGrowthRecords(c *gin.Context) {
	var records []models.GrowthRecord
	DB.Find(&records)
	c.JSON(http.StatusOK, records)
}