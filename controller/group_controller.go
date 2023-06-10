package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
)

type GroupController struct{}

func (c *GroupController) returnGroups(ctx *gin.Context, poolID int64) {
	gs := service.GroupService{}
	var groups []model.GroupRole
	err := gs.GetGroups(poolID, &groups)
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
		"data": gin.H{
			"groups": groups,
		},
	})
}

func (c *GroupController) CreateGroup(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
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
	gs := service.GroupService{}
	err = gs.CreateGroup(poolID, req.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnGroups(ctx, poolID)
}

func (c *GroupController) GetGroups(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	c.returnGroups(ctx, poolID)
}

func (c *GroupController) UpdateGroup(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "GroupID error.",
			"data": err.Error(),
		})
		return
	}
	gs := service.GroupService{}
	v, err = gs.VerifyGroup(poolID, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	if !v {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "GroupID error.",
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
	err = gs.UpdateGroup(id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnGroups(ctx, poolID)
}

func (c *GroupController) DeleteGroup(ctx *gin.Context) {
	pc := PoolController{}
	v, poolID := pc.checkPoolType(ctx)
	if !v {
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "GroupID error.",
			"data": err.Error(),
		})
		return
	}
	gs := service.GroupService{}
	v, err = gs.VerifyGroup(poolID, id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	if !v {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "GroupID error or not exist.",
		})
		return
	}
	err = gs.DeleteGroup(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnGroups(ctx, poolID)
}
