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

func (c *PoolController) checkPoolType(ctx *gin.Context, r bool) (bool, int64) {
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
		if r {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 11,
				"msg":  "Pool type error.",
			})
		}
		return false, poolId
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
	poolID, err := strconv.ParseInt(ctx.Param("poolID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Invalid pool ID.",
		})
		return
	}

	ps := service.PoolService{}
	var pool model.Pool
	err = ps.GetPool(poolID, &pool)
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
	ss := service.SelectionService{}
	err = ss.GetSelections(poolID, &pool.Selections)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	if pool.Type == 3 {
		gs := service.GroupService{}
		err = gs.GetGroups(poolID, &pool.Groups)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
		rs := service.RoleService{}
		err = rs.GetRoles(poolID, &pool.Roles)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pool,
	})
}

func (c *PoolController) UpdatePool(ctx *gin.Context) {
	poolID, err := strconv.ParseInt(ctx.Param("poolID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Invalid pool ID.",
		})
		return
	}
	var req model.PoolBase
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}
	if req.Name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 3,
			"msg":  "Name error.",
		})
		return
	}
	if len(req.Name) > 32 || len(req.Description) > 255 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 4,
			"msg":  "Name or Description too long.",
		})
		return
	}
	ps := service.PoolService{}
	var before model.Pool
	err = ps.GetPool(poolID, &before)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 5,
				"msg":  "Pool not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}
	if before.Type != req.Type && req.Type != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 6,
			"msg":  "Type cannot be changed.",
		})
		return
	} else {
		req.Type = before.Type
	}
	if before.Status > req.Status {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 7,
			"msg":  "Status cannot be backward.",
		})
		return
	}
	err = ps.UpdatePool(poolID, req)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 5,
				"msg":  "Pool not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}
	if req.Type == 4 && before.Status == 1 && req.Status == 2 {
		ss := service.SelectionService{}
		err = ss.GenerateSelections(poolID)
		if err != nil {
			if err.Error() == "PreStatus" {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 8,
					"msg":  "Pre status error.",
				})
				return
			}
			returnMySQLError(ctx, err)
			return
		}
	}
	var pool model.Pool
	err = ps.GetPool(poolID, &pool)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 5,
				"msg":  "Pool not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}
	ss := service.SelectionService{}
	err = ss.GetSelections(poolID, &pool.Selections)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	if pool.Type == 3 {
		gs := service.GroupService{}
		err = gs.GetGroups(poolID, &pool.Groups)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
		rs := service.RoleService{}
		err = rs.GetRoles(poolID, &pool.Roles)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": pool,
	})
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

func (c *PoolController) DrewSelect(ctx *gin.Context) {
	// 处理抽签的请求
}

func (c *PoolController) GetResults(ctx *gin.Context) {
	poolID, err := strconv.ParseInt(ctx.Param("poolID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Invalid pool ID.",
		})
		return
	}
	ps := service.PoolService{}
	var pool model.Pool
	err = ps.GetPool(poolID, &pool)
	if err != nil {
		if err.Error() == "noPool" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "Pool not found.",
			})
			return
		}
		returnMySQLError(ctx, err)
		return
	}
	var results []model.Results
	ss := service.SelectionService{}
	err = ss.GetSelectionsByOrder(poolID, &results)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	var res []model.Result
	if pool.Type == 3 {
		gs := service.GroupService{}
		err = gs.GetGroups(poolID, &pool.Groups)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
	} else {
		ds := service.DrawService{}
		err = ds.GetDraws(poolID, &res)
		if err != nil {
			returnMySQLError(ctx, err)
			return
		}
	}
	for r := range res {
		for i := range results {
			if res[r].UID == results[i].Id {
				results[i].Result = append(results[i].Result, res[r])
			}
		}
	}
	for i := range results {
		if len(results[i].Result) == 0 {
			results[i].Result = make([]model.Result, 0)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": results,
	})
}
