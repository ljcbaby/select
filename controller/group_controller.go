package controller

import (
	"github.com/gin-gonic/gin"
)

type GroupController struct{}

func (c *GroupController) CreateGroup(ctx *gin.Context) {
	// 处理创建group的请求
}

func (c *GroupController) GetGroups(ctx *gin.Context) {
	// 处理获取groups的请求
}

func (c *GroupController) UpdateGroup(ctx *gin.Context) {
	// 处理更新group的请求
}

func (c *GroupController) DeleteGroup(ctx *gin.Context) {
	// 处理删除group的请求
}
