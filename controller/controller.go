package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Pool      *PoolController
	Selection *SelectionController
	Group     *GroupController
	Role      *RoleController
}

func NewController() *Controller {
	return &Controller{
		// 初始化控制器依赖或成员变量
	}
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "0",
		"msg":  "Ljcbaby's draw backend.",
	})
}

func (c *Controller) GetPools(ctx *gin.Context) {
	// 处理获取pools的请求
}

func (c *Controller) DrewSelect(ctx *gin.Context) {
	// 处理抽签的请求
}

func (c *Controller) GetResult(ctx *gin.Context) {
	// 处理获取result的请求
}
