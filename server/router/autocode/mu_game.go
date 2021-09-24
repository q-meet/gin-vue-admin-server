package autocode

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type MuGameRouter struct {
}

// InitMuGameRouter 初始化 MuGame 路由信息
func (s *MuGameRouter) InitMuGameRouter(Router *gin.RouterGroup) {
	muGameRouter := Router.Group("muGame").Use(middleware.OperationRecord())
	var muGameApi = v1.ApiGroupApp.AutoCodeApiGroup.MuGameApi
	{
		muGameRouter.POST("createMuGame", muGameApi.CreateMuGame)   // 新建MuGame
		muGameRouter.DELETE("deleteMuGame", muGameApi.DeleteMuGame) // 删除MuGame
		muGameRouter.DELETE("deleteMuGameByIds", muGameApi.DeleteMuGameByIds) // 批量删除MuGame
		muGameRouter.PUT("updateMuGame", muGameApi.UpdateMuGame)    // 更新MuGame
		muGameRouter.GET("findMuGame", muGameApi.FindMuGame)        // 根据ID获取MuGame
		muGameRouter.GET("getMuGameList", muGameApi.GetMuGameList)  // 获取MuGame列表
	}
}
