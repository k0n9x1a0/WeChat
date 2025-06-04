package controllers

import (
	"encoding/json"
	"fmt"
	"wechatdll/models"
	"wechatdll/models/Wxapp"
)

// 微信小程序模块
type WxappController struct {
	BaseController
}

// @Summary 授权小程序(返回授权后的code)
// @Param	body			body	Wxapp.DefaultParam	 true		""
// @Failure 200
// @router /JSLogin [post]
func (c *WxappController) JSLogin() {
	var Data Wxapp.DefaultParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
	if err != nil {
		Result := models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}
	WXDATA := Wxapp.JSLogin(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 添加小程序到历史列表
// @Param	body			body	Wxapp.DefaultParam	 true		""
// @Failure 200
// @router /updatewxausagerecord [post]
func (c *WxappController) Updatewxausagerecord() {
	var Data Wxapp.DefaultParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
	if err != nil {
		Result := models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}
	WXDATA := Wxapp.Updatewxausagerecord(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 小程序操作
// @Param	body			body	Wxapp.JSOperateWxParam	 true		" "
// @Failure 200
// @router /Wxapp/JSOperateWxData [post]
func (c *WxappController) JSOperateWxData() {
	var Data Wxapp.JSOperateWxParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &Data)
	if err != nil {
		Result := models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
		c.Data["json"] = &Result
		c.ServeJSON()
		return
	}
	WXDATA := Wxapp.JSOperateWx(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}