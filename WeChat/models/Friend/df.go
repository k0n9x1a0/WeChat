package Friend

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mili811"
	"wechatdll/comm"
	"wechatdll/models"
)

func BatchDelFriend(Data BatchDelFriendParam) models.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	req := &mili811.BatchModContactTypeRequest{
		BaseRequest: &mili811.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(D.ClientVersion)),
			DeviceType:    []byte(D.DeviceType),
			Scene:         proto.Uint32(0),
		},
		ModTypeInfoCount: proto.Uint32(uint32(len(Data.ToWxids))),
		ModTypeInfoList:  []*mili811.ModContactTypeInfo{},
		Scene:            proto.Uint32(0),
	}

	for _, userName := range Data.ToWxids {
		modContactTypeInfo := &mili811.ModContactTypeInfo{
			UserName:     proto.String(userName),
			ModType:      proto.Uint32(1),
			ModOperation: proto.Uint32(2),
		}
		req.ModTypeInfoList = append(req.ModTypeInfoList, modContactTypeInfo)
	}

	reqdata, err := proto.Marshal(req)

	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	//发包
	protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
		Ip:     D.Mmtlsip,
		Host:   D.MmtlsHost,
		Cgiurl: "/cgi-bin/micromsg-bin/batchmodcontacttype",
		Proxy:  D.Proxy,
		PackData: Algorithm.PackData{
			Reqdata:          reqdata,
			Cgi:              3990,
			Uin:              D.Uin,
			Cookie:           D.Cooike,
			Sessionkey:       D.Sessionkey,
			EncryptType:      5,
			Loginecdhkey:     D.RsaPublicKey,
			Clientsessionkey: D.Clientsessionkey,
			UseCompress:      false,
		},
	}, D.MmtlsKey)

	if err != nil {
		return models.ResponseResult{
			Code:    errtype,
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}

	//解包
	Response := mili811.BatchModContactTypeResponse{}
	err = proto.Unmarshal(protobufdata, &Response)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	return models.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    Response,
	}
}
