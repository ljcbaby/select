package router

import (
	"github.com/ljcbaby/select/controller"
	"github.com/ljcbaby/select/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 创建控制器实例
	Controller := &controller.Controller{}

	// 中间件
	r.Use(middleware.CORS())

	// 路由配置
	r.GET("/", Controller.Index)
	r.GET("/pools", Controller.GetPools)
	r.POST("/select/:poolID", Controller.DrewSelect)

	r.POST("/pool", Controller.Pool.CreatePool)
	r.GET("/pool/:poolID", Controller.Pool.GetPool)
	r.PATCH("/pool/:poolID", Controller.Pool.UpdatePool)
	r.DELETE("/pool/:poolID", Controller.Pool.DeletePool)
	r.GET("/pool/:poolID/results", Controller.Pool.GetResults)

	r.POST("/pool/:poolID/selections", Controller.Selection.CreateSelection)
	r.GET("/pool/:poolID/selections", Controller.Selection.GetSelections)
	r.PATCH("/pool/:poolID/selections/:selectionID", Controller.Selection.UpdateSelection)
	r.DELETE("/pool/:poolID/selections/:selectionID", Controller.Selection.DeleteSelection)

	r.POST("/pool/:poolID/groups", Controller.Group.CreateGroup)
	r.GET("/pool/:poolID/groups", Controller.Group.GetGroups)
	r.PATCH("/pool/:poolID/groups/:groupID", Controller.Group.UpdateGroup)
	r.DELETE("/pool/:poolID/groups/:groupID", Controller.Group.DeleteGroup)

	r.POST("/pool/:poolID/roles", Controller.Role.CreateRole)
	r.GET("/pool/:poolID/roles", Controller.Role.GetRoles)
	r.PATCH("/pool/:poolID/roles/:roleID", Controller.Role.UpdateRole)
	r.DELETE("/pool/:poolID/roles/:roleID", Controller.Role.DeleteRole)

	return r
}
