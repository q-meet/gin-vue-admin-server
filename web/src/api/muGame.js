import service from '@/utils/request'

// @Tags MuGame
// @Summary 创建MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MuGame true "创建MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /muGame/createMuGame [post]
export const createMuGame = (data) => {
  return service({
    url: '/muGame/createMuGame',
    method: 'post',
    data
  })
}

// @Tags MuGame
// @Summary 删除MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MuGame true "删除MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /muGame/deleteMuGame [delete]
export const deleteMuGame = (data) => {
  return service({
    url: '/muGame/deleteMuGame',
    method: 'delete',
    data
  })
}

// @Tags MuGame
// @Summary 删除MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /muGame/deleteMuGame [delete]
export const deleteMuGameByIds = (data) => {
  return service({
    url: '/muGame/deleteMuGameByIds',
    method: 'delete',
    data
  })
}

// @Tags MuGame
// @Summary 更新MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MuGame true "更新MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /muGame/updateMuGame [put]
export const updateMuGame = (data) => {
  return service({
    url: '/muGame/updateMuGame',
    method: 'put',
    data
  })
}

// @Tags MuGame
// @Summary 用id查询MuGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MuGame true "用id查询MuGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /muGame/findMuGame [get]
export const findMuGame = (params) => {
  return service({
    url: '/muGame/findMuGame',
    method: 'get',
    params
  })
}

// @Tags MuGame
// @Summary 分页获取MuGame列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取MuGame列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /muGame/getMuGameList [get]
export const getMuGameList = (params) => {
  return service({
    url: '/muGame/getMuGameList',
    method: 'get',
    params
  })
}