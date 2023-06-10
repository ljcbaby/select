package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
)

type PoolController struct{}

func (c *PoolController) GetPools(ctx *gin.Context) {
	ps := service.PoolService{}
	var pools []model.PoolBase
	err := ps.GetPools(&pools)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pools,
	})
}

func (c *PoolController) checkPoolType(ctx *gin.Context) (bool, int64) {
	poolId, err := strconv.ParseInt(ctx.Param("poolID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10,
			"msg":  "PoolID error.",
			"data": err.Error(),
		})
		return false, -10
	}
	ps := service.PoolService{}
	poolType, err := ps.GetPoolType(poolId)
	if err != nil {
		returnMySQLError(ctx, err)
		return false, -100
	}
	if poolType != 3 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 11,
			"msg":  "Pool type error.",
		})
		return false, -11
	}
	return true, poolId
}

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

	ps := service.PoolService{}
	id, err := ps.CreatePool(req)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"id": id,
		},
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

	ps := service.PoolService{}
	err = ps.DeletePool(poolID)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2,
				"msg":  "Pool not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (c *Controller) DrewSelect(ctx *gin.Context) {
	// 处理抽签的请求
}

func (c *PoolController) GetResults(ctx *gin.Context) {
	// 处理获取抽签结果的请求
}
