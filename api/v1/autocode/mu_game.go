package autocode

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model/autocode"
    "gin-vue-admin/model/common/request"
    autocodeReq "gin-vue-admin/model/autocode/request"
    "gin-vue-admin/model/common/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MuGameApi struct {
}

var muGameService = service.ServiceGroupApp.AutoCodeServiceGroup.MuGameService


// CreateMuGame 创建MuGame
// @Tags MuGame
// @Summary 创建MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MuGame true "创建MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /muGame/createMuGame [post]
func (muGameApi *MuGameApi) CreateMuGame(c *gin.Context) {
	var muGame autocode.MuGame
	_ = c.ShouldBindJSON(&muGame)
	if err := muGameService.CreateMuGame(muGame); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMuGame 删除MuGame
// @Tags MuGame
// @Summary 删除MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MuGame true "删除MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /muGame/deleteMuGame [delete]
func (muGameApi *MuGameApi) DeleteMuGame(c *gin.Context) {
	var muGame autocode.MuGame
	_ = c.ShouldBindJSON(&muGame)
	if err := muGameService.DeleteMuGame(muGame); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMuGameByIds 批量删除MuGame
// @Tags MuGame
// @Summary 批量删除MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /muGame/deleteMuGameByIds [delete]
func (muGameApi *MuGameApi) DeleteMuGameByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := muGameService.DeleteMuGameByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMuGame 更新MuGame
// @Tags MuGame
// @Summary 更新MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MuGame true "更新MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /muGame/updateMuGame [put]
func (muGameApi *MuGameApi) UpdateMuGame(c *gin.Context) {
	var muGame autocode.MuGame
	_ = c.ShouldBindJSON(&muGame)
	if err := muGameService.UpdateMuGame(muGame); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMuGame 用id查询MuGame
// @Tags MuGame
// @Summary 用id查询MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.MuGame true "用id查询MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /muGame/findMuGame [get]
func (muGameApi *MuGameApi) FindMuGame(c *gin.Context) {
	var muGame autocode.MuGame
	_ = c.ShouldBindQuery(&muGame)
	if err, remuGame := muGameService.GetMuGame(muGame.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remuGame": remuGame}, c)
	}
}

// GetMuGameList 分页获取MuGame列表
// @Tags MuGame
// @Summary 分页获取MuGame列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocodeReq.MuGameSearch true "分页获取MuGame列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /muGame/getMuGameList [get]
func (muGameApi *MuGameApi) GetMuGameList(c *gin.Context) {
	var pageInfo autocodeReq.MuGameSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := muGameService.GetMuGameInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
