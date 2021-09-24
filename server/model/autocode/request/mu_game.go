package request

import (
	"gin-vue-admin/model/autocode"
	"gin-vue-admin/model/common/request"
)

type MuGameSearch struct{
    autocode.MuGame
    request.PageInfo
}