package Tools

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mm"
	"wechatdll/Cilent/wechat"
	"wechatdll/comm"
	"wechatdll/models"
)

type DownloadVoiceData struct {
	Base64      []byte
	VoiceLength uint32
}

func DownloadVoicePlus(Data DownloadVoiceParam) models.ResponseResult {
	D, err := comm.GetLoginata(Data.Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	offset := 0
	datalen := 0
	total_length := Data.Length
	Bufid := Data.Bufid
	Databuff := make([]byte, 0)
	var VoiceLength uint32
	MasterBufId, _ := strconv.ParseInt(Bufid, 10, 64)
	if total_length > 65535 {
		datalen = 65535
	} else {
		datalen = total_length
	}
	for {
		if offset >= total_length {
			break
		}
		count := 0
		if total_length-offset >= datalen {
			count = datalen
		} else {
			count = total_length - offset
		}

		req := &wechat.DownloadVoiceRequest{
			MsgId:  proto.Uint64(Data.MsgId),
			Offset: proto.Uint32(uint32(offset)),
			Length: proto.Uint32(uint32(count)),
			BaseRequest: &wechat.BaseRequest{
				SessionKey:    []byte{},
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(D.ClientVersion)),
				DeviceType:    []byte(D.DeviceType),
				Scene:         proto.Uint32(0),
			},
			ClientMsgId:  proto.String(""),
			NewMsgId:     proto.Uint64(0),
			ChatRoomName: proto.String(Data.FromUserName),
			MasterBufId:  proto.Int64(MasterBufId),
		}

		//序列化
		reqdata, _ := proto.Marshal(req)

		//发包
		protobufdata, _, errtype, err := comm.SendRequest(comm.SendPostData{
			Ip:     D.Mmtlsip,
			Host:   D.MmtlsHost,
			Cgiurl: "/cgi-bin/micromsg-bin/downloadvoice",
			Proxy:  D.Proxy,
			PackData: Algorithm.PackData{
				Reqdata:          reqdata,
				Cgi:              128,
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
		Response := mm.DownloadVoiceResponse{}
		err = proto.Unmarshal(protobufdata, &Response)
		if err != nil {
			return models.ResponseResult{
				Code:    -8,
				Success: false,
				Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
				Data:    nil,
			}
		}

		var dLen = Response.GetData().GetILen()
		count = int(dLen)
		if count != 0 {
			offset += count
			DataStream := bytes.NewBuffer(Response.GetData().GetBuffer())
			Databuff = append(
				Databuff,
				DataStream.Bytes()...,
			)
			time.Sleep(50)

		}

	}
	rep := DownloadVoiceData{
		Base64:      Databuff,
		VoiceLength: VoiceLength,
	}

	return models.ResponseResult{
		Code:    0,
		Success: true,
		Message: "成功",
		Data:    rep,
	}
}
