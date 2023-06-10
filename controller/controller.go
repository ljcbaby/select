package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
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

func (c *Controller) GetPools(ctx *gin.Context) {
	ps := service.PoolService{}
	var pools []model.PoolBase
	err := ps.GetPools(&pools)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pools,
	})
}

func (c *Controller) DrewSelect(ctx *gin.Context) {
	// 处理抽签的请求
}
