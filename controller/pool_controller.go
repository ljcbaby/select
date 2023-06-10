package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
)

type PoolController struct{}

func (c *PoolController) CreatePool(ctx *gin.Context) {
	var req model.PoolBase
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}

	if req.Name == "" || req.Type < 1 || req.Type > 4 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "Name or Type error.",
		})
		return
	}
	if len(req.Name) > 32 || len(req.Description) > 255 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 3,
			"msg":  "Name or Description too long.",
		})
		return
	}

	service := service.PoolService{}
	id, err := service.CreatePool(req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	var data struct {
		Id int64 `json:"id"`
	}
	data.Id = id
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func (c *PoolController) GetPool(ctx *gin.Context) {
	// 处理获取pool的请求
}

func (c *PoolController) UpdatePool(ctx *gin.Context) {
	// 处理更新pool的请求
}

func (c *PoolController) DeletePool(ctx *gin.Context) {
	id := ctx.Param("poolID")
	poolID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Invalid pool ID.",
		})
		return
	}

	service := service.PoolService{}
	err = service.DeletePool(poolID)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "Pool not found.",
			})
			return
		}
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
	})
}

func (c *PoolController) GetResults(ctx *gin.Context) {
	// 处理获取抽签结果的请求
}
