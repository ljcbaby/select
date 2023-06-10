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

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Ljcbaby's draw backend.",
	})
}

func (c *Controller) DrewSelect(ctx *gin.Context) {
	// 处理抽签的请求
}

func returnMySQLError(ctx *gin.Context, err error) {
	ctx.JSON(200, gin.H{
		"code": 100,
		"msg":  "MySQL error.",
		"data": err.Error(),
	})
}
