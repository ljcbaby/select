package controller

import (
	"github.com/gin-gonic/gin"
)

type PoolController struct{}

func (c *PoolController) CreatePool(ctx *gin.Context) {
	// 处理创建pool的请求
}

func (c *PoolController) GetPool(ctx *gin.Context) {
	// 处理获取pools的请求
}

func (c *PoolController) UpdatePool(ctx *gin.Context) {
	// 处理更新pool的请求
}

func (c *PoolController) DeletePool(ctx *gin.Context) {
	// 处理删除pool的请求
}
