// 自动生成模板MuGame
package autocode

import (
	"gin-vue-admin/global"
	"time"
)

// MuGame 结构体
// 如果含有time.Time 请自行import time包
type MuGame struct {
	global.GVA_MODEL
	Title        string     `json:"title" form:"title" gorm:"column:title;comment:游戏标题;type:varchar(255);"`
	Images       string     `json:"images" form:"images" gorm:"column:images;comment:游戏logo;type:varchar(255);"`
	Content      string     `json:"content" form:"content" gorm:"column:content;comment:介绍内容;type:varchar(255);"`
	Url          string     `json:"url" form:"url" gorm:"column:url;comment:游戏连接;type:varchar(255);"`
	Type         *bool      `json:"type" form:"type" gorm:"column:type;comment:游戏分类:对应mu_nav的id;type:tinyint"`
	CreateTime   *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;type:datetime"`
	Gameopentype *int       `json:"gameopentype" form:"gameopentype" gorm:"column:gameopentype;comment:游戏连接打开方式：1.app打开；2.浏览器打开;type:int"`
}

// TableName MuGame 表名
func (MuGame) TableName() string {
	return "mu_game"
}
