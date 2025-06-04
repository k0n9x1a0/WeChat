package Login

import (
	"crypto/md5"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mm"
	"wechatdll/Mmtls"
	"wechatdll/comm"
	"wechatdll/lib"
	"wechatdll/models"
)

func Secautoauth1(Wxid string) models.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	httpclient := Mmtls.GenNewHttpClient(D.MmtlsKey, D.MmtlsHost)

	if len(D.Autoauthkey) <= 0 {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "账号异常：Autoauthkey读取失败",
			Data:    nil,
		}
	}

	Autoauthkey := &mm.AutoAuthKey{}
	_ = proto.Unmarshal(D.Autoauthkey, Autoauthkey)

	prikey, pubkey := Algorithm.GetEcdh713Key()

	//基础设备信息
	Imei := Algorithm.IOSImei(D.Deviceid_str)
	// TODO: 放到初始化上下文中生成
	SoftType := Algorithm.SoftType_iPad(D.Deviceid_str, Algorithm.IPadOsVersion, Algorithm.IPadModel)
	ClientSeqId := lib.GetClientSeqId(D.Deviceid_str)

	//24算法
	ccData := &mm.CryptoData{
		Version:     []byte("00000006"),
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

	req := &mm.AutoAuthRequest{
		RsaReqData: &mm.AutoAuthRsaReqData{
			AesEncryptKey: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(Autoauthkey.EncryptKey.Buffer))),
				Buffer: Autoauthkey.EncryptKey.Buffer,
			},
			CliPubEcdhkey: &mm.ECDHKey{
				Nid: proto.Int32(713),
				Key: &mm.SKBuiltinBufferT{
					ILen:   proto.Uint32(uint32(len(pubkey))),
					Buffer: pubkey,
				},
			},
		},
		AesReqData: &mm.AutoAuthAesReqData{
			BaseRequest: &mm.BaseRequest{
				SessionKey:    D.Sessionkey,
				Uin:           proto.Uint32(D.Uin),
				DeviceId:      D.Deviceid_byte,
				ClientVersion: proto.Int32(int32(D.ClientVersion)),
				DeviceType:    []byte(D.DeviceType),
				Scene:         proto.Uint32(0),
			},
			BaseReqInfo: &mm.BaseAuthReqInfo{},
			AutoAuthKey: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(D.Autoauthkey))),
				Buffer: D.Autoauthkey,
			},
			Imei:         &Imei,
			SoftType:     &SoftType,
			BuiltinIpseq: proto.Uint32(0),
			ClientSeqId:  &ClientSeqId,
			DeviceName:   proto.String(D.DeviceName),
			DeviceType:   proto.String("pad-android-31"),
			Language:     proto.String("zh_CN"),
			TimeZone:     proto.String("8.0"),
			ExtSpamInfo: &mm.SKBuiltinBufferT{
				ILen:   proto.Uint32(uint32(len(WCExtInfoseq))),
				Buffer: WCExtInfoseq,
			},
		},
	}

	reqdata, err := proto.Marshal(req)

	hec := &Algorithm.Client{}
	hec.Init("IOS")
	hecData := hec.HybridEcdhPackIosEn(763, D.Uin, D.Cooike, reqdata)
	recvData, err := httpclient.MMtlsPost(D.MmtlsHost, "/cgi-bin/micromsg-bin/secautoauth", hecData, D.Proxy)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
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
