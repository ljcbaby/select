package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
)

type RoleController struct{}

func (c *RoleController) returnRoles(ctx *gin.Context, poolID int64) {
	rs := service.RoleService{}
	var roles []model.GroupRole
	err := rs.GetRoles(poolID, &roles)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"roles": roles,
		},
	})
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	pc := PoolController{}
	v, poolID := pc.checkPoolType(ctx)
	if !v {
		return
	}
	var req model.GroupRole
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}
	if len(req.Name) < 1 || len(req.Name) > 32 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "Name error.",
		})
		return
	}
	rs := service.RoleService{}
	err = rs.CreateRole(poolID, req.Name)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	c.returnRoles(ctx, poolID)
}

func (c *RoleController) GetRoles(ctx *gin.Context) {
	pc := PoolController{}
	v, poolId := pc.checkPoolType(ctx)
	if !v {
		return
	}
	c.returnRoles(ctx, poolId)
}

func (c *RoleController) UpdateRole(ctx *gin.Context) {
	pc := PoolController{}
	v, poolID := pc.checkPoolType(ctx)
	if !v {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("roleID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "RoleID error.",
			"data": err.Error(),
		})
		return
	}
	rs := service.RoleService{}
	v, err = rs.VerifyRole(poolID, id)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	if !v {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "RoleID error.",
		})
		return
	}
	var req model.GroupRole
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}
	if len(req.Name) < 1 || len(req.Name) > 32 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 3,
			"msg":  "Name error.",
		})
		return
	}
	err = rs.UpdateRole(id, req.Name)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	c.returnRoles(ctx, poolID)
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	pc := PoolController{}
	v, poolID := pc.checkPoolType(ctx)
	if !v {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("roleID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "RoleID error.",
			"data": err.Error(),
		})
		return
	}
	rs := service.RoleService{}
	v, err = rs.VerifyRole(poolID, id)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	if !v {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "RoleID error or not exist.",
		})
		return
	}
	err = rs.DeleteRole(id)
	if err != nil {
		returnMySQLError(ctx, err)
		return
	}
	c.returnRoles(ctx, poolID)
}
