package Wxapp

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mm"
	"wechatdll/comm"
	"wechatdll/models"
)

func JSOperateWx(Data JSOperateWxParam) models.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}
	//sts, _ := base64.StdEncoding.DecodeString("eyJhcGlfbmFtZSI6IndlYmFwaV9nZXR1c2VycHJvZmlsZSIsImRhdGEiOnsiYXBwX3ZlcnNpb24iOjE3MywiZGVzYyI6IueUqOS6juS4quS6uuS4reW/g+aYvuekuiIsImxhbmciOiJlbiIsInZlcnNpb24iOiIzLjMuMSJ9LCJvcGVyYXRlX2RpcmVjdGx5IjpmYWxzZSwic2hvd19jb25maXJtIjp0cnVlLCJ3aXRoX2NyZWRlbnRpYWxzIjp0cnVlfQ==")
	sts := Data.Data
	req := &mm.JSOperateWxDataRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    D.Sessionkey,
			Uin:           proto.Uint32(D.Uin),
			DeviceId:      D.Deviceid_byte,
			ClientVersion: proto.Int32(int32(D.ClientVersion)),
			DeviceType:    []byte(D.DeviceType),
			Scene:         proto.Uint32(0),
		},
		Appid:       proto.String(Data.Appid),
		Data:        []byte(sts),
		GrantScope:  proto.String("scope.userInfo"),
		Opt:         proto.Int(Data.Opt),
		VersionType: proto.Int32(0),
		ExtInfo: &mm.WxaExternalInfo{
			HostAppid: proto.String(""),
			Scene:     proto.Int32(1089),
			SourceEnv: proto.Int32(1),
		},
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
		Cgiurl: "/cgi-bin/mmbiz-bin/js-operatewxdata",
		Proxy:  D.Proxy,
		PackData: Algorithm.PackData{
			Reqdata:          reqdata,
			Cgi:              1133,
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
	Response := mm.JSOperateWxDataResponse{}
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
