package Tools

import (
	"fmt"
	"wechatdll/comm"
	"wechatdll/models"
)

// 全局进行服务的开启和关闭
type SetEnableServiceParam struct {
	Wxid    string
	Disable bool
}

func EnableAccountService(Data SetEnableServiceParam) models.ResponseResult {
	// 得到登录人信息
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}
	// 设置服务状态
	D.EnableService = !(Data.Disable)
	_ = comm.CreateLoginData(*D, D.Wxid, 0)
	return models.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    "服务启用状态  " + fmt.Sprintf("%t", !Data.Disable),
	}
}
