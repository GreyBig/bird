package controller

import (
	"bird/common/helper"
	"bird/model"
	"github.com/valyala/fasthttp"
)

type LogController struct{}

// 日志列表
func (that LogController) List(ctx *fasthttp.RequestCtx) {

	s, err := model.LogList()
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, s)
}
