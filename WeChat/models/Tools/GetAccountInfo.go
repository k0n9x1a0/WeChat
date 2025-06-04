package Tools

import (
	"fmt"
	"time"
	"wechatdll/comm"
	"wechatdll/models"
)

func GetAccountInfo(Wxid string) models.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	ResultObj := make(map[string]interface{})
	// 登录时间
	ResultObj["survive"] = time.Now().Unix() - D.LoginDate

	hbResult := comm.GetTodayMoney(Wxid, 1)
	ResultObj["redPocket"] = hbResult

	zzResult := comm.GetTodayMoney(Wxid, 2)
	ResultObj["transfer"] = zzResult

	return models.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    ResultObj,
	}
}
