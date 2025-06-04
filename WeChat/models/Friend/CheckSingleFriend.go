package Friend

import (
	"wechatdll/models"
)

// 检测单线好友
func CheckSingleFriend(Wxid string) models.ResponseResult {
	req := GetContractListparameter{
		Wxid:                      Wxid,
		CurrentWxcontactSeq:       0,
		CurrentChatRoomContactSeq: 0,
	}

	contractList := GetContractList(req)
	//contractList.Data

	return models.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    contractList,
	}

}
