package controllers

import (
	"encoding/json"
	"fmt"
	"wechatdll/models"
	"wechatdll/models/TenPay"
)

type TenPayController struct {
	BaseController
}

// @Summary 查看红包
// @Param	body	body	TenPay.Openwxhb true
// @Success 200
// @router /Qrydetailwxhb [post]
func (c *TenPayController) Qrydetailwxhb() {
	var Data TenPay.QrydetailwxhbParam
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
	WXDATA := TenPay.Qrydetailwxhb(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 查看红包领取列表入口
func (c *TenPayController) GetRedPacketListApi() {
	var ParamData TenPay.HongBaoDetail
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
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
	WXDATA := TenPay.GetRedPacketListApi(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 打开红包
// @Param	body	body	TenPay.ReceivewxhbParam true
// @Success 200
// @router /Receivewxhb [post]
func (c *TenPayController) Receivewxhb() {
	var Data TenPay.ReceivewxhbParam
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
	WXDATA := TenPay.Receivewxhb(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 拆开红包
// @Param	body	body	TenPay.OpenwxhbParam true
// @Success 200
// @router /Openwxhb [post]
func (c *TenPayController) Openwxhb() {
	var Data TenPay.OpenwxhbParam
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
	WXDATA := TenPay.Openwxhb(Data)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}

// @Summary 抢红包
// @Param	body		body 	TenPay.HongBaoParam    true	"注意参数"
// @Success 200
// @router /OpenHongBao [post]
func (c *TenPayController) AutoHongBao() {
	var ParamData TenPay.HongBaoParam
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &ParamData)
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
	WXDATA := TenPay.AutoHongBao(ParamData)
	c.Data["json"] = &WXDATA
	c.ServeJSON()
}
