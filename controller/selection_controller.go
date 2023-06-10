package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ljcbaby/select/model"
	"github.com/ljcbaby/select/service"
)

type SelectionController struct{}

func (c *SelectionController) returnSelections(ctx *gin.Context, poolID int64) {
	ss := service.SelectionService{}
	var selections []model.Selection
	err := ss.GetSelections(poolID, &selections)
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
			"selections": selections,
		},
	})
}

func (c *SelectionController) CreateSelection(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	var req model.Selection
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}
	if len(req.Name) < 1 || len(req.Name) > 32 || req.Number < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
		})
		return
	}
	ss := service.SelectionService{}
	err = ss.CreateSelection(poolID, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnSelections(ctx, poolID)
}

func (c *SelectionController) GetSelections(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	c.returnSelections(ctx, poolID)
}

func (c *SelectionController) UpdateSelection(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	ID, err := strconv.ParseInt(ctx.Param("selectionID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "selectionID error.",
		})
		return
	}
	ss := service.SelectionService{}
	v, err = ss.VerifySelection(poolID, ID)
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
			"msg":  "selectionID error.",
		})
		return
	}
	var req model.Selection
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
			"data": err.Error(),
		})
		return
	}
	if len(req.Name) < 1 || len(req.Name) > 32 || req.Number < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Body error.",
		})
		return
	}
	err = ss.UpdateSelection(ID, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnSelections(ctx, poolID)
}

func (c *SelectionController) DeleteSelection(ctx *gin.Context) {
	ps := PoolController{}
	v, poolID := ps.checkPoolType(ctx)
	if !v {
		return
	}
	ID, err := strconv.ParseInt(ctx.Param("selectionID"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "selectionID error.",
		})
		return
	}
	ss := service.SelectionService{}
	v, err = ss.VerifySelection(poolID, ID)
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
			"msg":  "selectionID error.",
		})
		return
	}
	err = ss.DeleteSelection(ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "MySQL error.",
			"data": err.Error(),
		})
		return
	}
	c.returnSelections(ctx, poolID)
}
