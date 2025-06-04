package Login

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
	"time"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mm"
	"wechatdll/Cilent/wechat"
	"wechatdll/Mmtls"
	"wechatdll/comm"
	"wechatdll/models"
)

func Secautoauth(Wxid string) models.ResponseResult {

	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	//初始化Mmtls
	var httpclient *Mmtls.HttpClientModel
	if D.MmtlsKey == nil {
		httpclient, D.MmtlsKey, err = comm.MmtlsInitialize(D.Proxy, D.MmtlsHost)
		if err != nil {
			return models.ResponseResult{
				Code:    -8,
				Success: false,
				Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
				Data:    nil,
			}
		}
	} else {
		httpclient = Mmtls.GenNewHttpClient(D.MmtlsKey, D.MmtlsHost)
	}

	// 请求设备device_token
	if D.DeviceToken.TrustResponseData == nil || D.DeviceToken.TrustResponseData.DeviceToken == nil || *D.DeviceToken.TrustResponseData.DeviceToken == "" {
		D.DeviceToken, err = IPadGetDeviceToken(D.Deviceid_str, D.RomModel, D.DeviceName, D.DeviceType, int32(D.ClientVersion), *httpclient, D.Proxy, D.MmtlsHost)
		if err != nil {
			// 请求失败则放空结构
			D.DeviceToken = mm.TrustResponse{}
		}
	}

	prikey, pubkey := Algorithm.GetEcdh713Key()

	ClientSeqId := fmt.Sprintf("%v_%v", D.Deviceid_str, time.Now().Unix())

	ccData := &mm.CryptoData{
		Version:     []byte("00000003"),
		Type:        proto.Uint32(1),
		EncryptData: Algorithm.GetiPhoneNewSpamData(D.Deviceid_str, D.DeviceName, D.DeviceToken),
		Timestamp:   proto.Uint32(uint32(time.Now().Unix())),
		Unknown5:    proto.Uint32(5),
		Unknown6:    proto.Uint32(0),
	}
	ccDataseq, _ := proto.Marshal(ccData)

	Wcstf := Algorithm.IphoneWcstf(Wxid)
	Wcste := Algorithm.IphoneWcste(0, 0)

	DeviceTokenCCD := &mm.DeviceToken{
		Version:   proto.String(""),
		Encrypted: proto.Uint32(1),
		Data: &mm.SKBuiltinStringT{
			String_: proto.String(D.DeviceToken.GetTrustResponseData().GetDeviceToken()),
		},
		TimeStamp: proto.Uint32(uint32(time.Now().Unix())),
		Optype:    proto.Uint32(2),
		Uin:       proto.Uint32(0),
	}
	DeviceTokenCCDPB, _ := proto.Marshal(DeviceTokenCCD)

	WCExtInfo := &mm.WCExtInfo{
		Wcstf: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(Wcstf))),
			Buffer: Wcstf,
		},
		Wcste: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(Wcste))),
			Buffer: Wcste,
		},
		CcData: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(ccDataseq))),
			Buffer: ccDataseq,
		},
		DeviceToken: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(DeviceTokenCCDPB))),
			Buffer: DeviceTokenCCDPB,
		},
	}

	WCExtInfoseq, _ := proto.Marshal(WCExtInfo)

	req := &wechat.AutoAuthRequest{
		RsaReqData: &wechat.AutoAuthRsaReqData{
			AesEncryptKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(D.Sessionkey))),
				Buffer: D.Sessionkey,
			},
			PubEcdhKey: &wechat.ECDHKey{
				Nid: proto.Uint32(713),
				Key: &wechat.SKBuiltinString_{
					Len:    proto.Uint32(uint32(len(pubkey))),
					Buffer: pubkey,
				},
			},
		},
		AesReqData: &wechat.AutoAuthAesReqData{
			BaseRequest: &wechat.BaseRequestPlus{
				ClientVersion: proto.Int32(int32(D.ClientVersion)),
				DeviceId:      D.Deviceid_byte,
				OsType:        &D.DeviceInfo.OsType,
				SessionKey:    D.Sessionkey,
				Uin:           &D.Uin,
				Scene:         proto.Uint32(2),
			},
			BaseReqInfo: &wechat.BaseAuthReqInfo{},
			AutoAuthKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(D.Autoauthkey))),
				Buffer: D.Autoauthkey,
			},
			Imei:         &D.Imei,
			SoftType:     &D.SoftType,
			BuiltinIpSeq: proto.Uint32(0),
			ClientSeqId:  &ClientSeqId,
			Signature:    proto.String(""),
			DeviceName:   proto.String(D.DeviceName),
			DeviceType:   proto.String(D.DeviceType),
			Language:     proto.String("Zh"),
			TimeZone:     proto.String("8.0"),
			ExtSpamInfo: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(WCExtInfoseq))),
				Buffer: WCExtInfoseq,
			},
		},
	}

	reqdata, err := proto.Marshal(req)

	hec := &Algorithm.Client{}
	hec.Init("IOS")
	hecData := hec.HybridEcdhPackIosEn(763, D.Uin, D.Cooike, reqdata)

	// 遇到mmtls失败, 则重新握手
	retrys := 3
	doRetry := true
	var recvData []byte
	for retrys > 0 && doRetry == true {
		doRetry = false
		retrys--
		recvData, err = httpclient.MMtlsPost(D.MmtlsHost, "/cgi-bin/micromsg-bin/secautoauth", hecData, D.Proxy)
		if err != nil && strings.Contains(err.Error(), "MMTLS") {
			// mmtls异常, 重新握手
			httpclient, D.MmtlsKey, err = comm.MmtlsInitialize(D.Proxy, D.MmtlsHost)
			if err != nil {
				return models.ResponseResult{
					Code:    -8,
					Success: false,
					Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
					Data:    nil,
				}
			}
			// 重新提交
			doRetry = true
		}
	}


	if len(recvData) <= 31 {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("组包异常, 返回31字节"),
			Data:    nil,
		}
	}

	ph1 := hec.HybridEcdhPackIosUn(recvData)
	//解包
	UnifyAuthResponse := mm.UnifyAuthResponse{}
	err = proto.Unmarshal(ph1.Data, &UnifyAuthResponse)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	loginRes := UnifyAuthResponse

	if loginRes.GetBaseResponse().GetRet() != 0 || loginRes.BaseResponse == nil || loginRes.AuthSectResp == nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "登录失败：您可能已退出微信",
			Data:    loginRes,
		}
	}

	Wx_loginecdhkey := Algorithm.DoECDH713Key(prikey, loginRes.GetAuthSectResp().GetSvrPubEcdhkey().GetKey().GetBuffer())
	m := md5.New()
	m.Write(Wx_loginecdhkey)
	D.Loginecdhkey = Wx_loginecdhkey
	ecdhdecrptkey := m.Sum(nil)
	D.Cooike = ph1.Cookies
	D.Sessionkey = Algorithm.AesDecrypt(loginRes.GetAuthSectResp().GetSessionKey().GetBuffer(), ecdhdecrptkey)
	D.Autoauthkey = loginRes.GetAuthSectResp().GetAutoAuthKey().GetBuffer()
	D.Autoauthkeylen = int32(loginRes.GetAuthSectResp().GetAutoAuthKey().GetILen())
	D.Serversessionkey = loginRes.GetAuthSectResp().GetServerSessionKey().GetBuffer()
	D.Clientsessionkey = loginRes.GetAuthSectResp().GetClientSessionKey().GetBuffer()

	err = comm.CreateLoginData(*D, D.Wxid, 0)

	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	return models.ResponseResult{
		Code:    1,
		Success: false,
		Message: "登陆成功",
		Data:    loginRes,
	}
}
