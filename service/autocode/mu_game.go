package autocode

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/autocode"
	"gin-vue-admin/model/common/request"
    autoCodeReq "gin-vue-admin/model/autocode/request"
)

type MuGameService struct {
}

// CreateMuGame 创建MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService) CreateMuGame(muGame autocode.MuGame) (err error) {
	err = global.GVA_DB.Create(&muGame).Error
	return err
}

// DeleteMuGame 删除MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService)DeleteMuGame(muGame autocode.MuGame) (err error) {
	err = global.GVA_DB.Delete(&muGame).Error
	return err
}

// DeleteMuGameByIds 批量删除MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService)DeleteMuGameByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.MuGame{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMuGame 更新MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService)UpdateMuGame(muGame autocode.MuGame) (err error) {
	err = global.GVA_DB.Save(&muGame).Error
	return err
}

// GetMuGame 根据id获取MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService)GetMuGame(id uint) (err error, muGame autocode.MuGame) {
	err = global.GVA_DB.Where("id = ?", id).First(&muGame).Error
	return
}

// GetMuGameInfoList 分页获取MuGame记录
// Author [piexlmax](https://github.com/piexlmax)
func (muGameService *MuGameService)GetMuGameInfoList(info autoCodeReq.MuGameSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.MuGame{})
    var muGames []autocode.MuGame
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&muGames).Error
	return err, muGames, total
}
