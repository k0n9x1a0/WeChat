package Login

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"time"
	"wechatdll/Algorithm"
	"wechatdll/Cilent/mm"
	"wechatdll/Cilent/wechat"
	"wechatdll/Mmtls"
	"wechatdll/baseinfo"
	cecdh "wechatdll/clientsdk/cecdn"
	clientsdk "wechatdll/clientsdk/hybrid"
	"wechatdll/comm"
	"wechatdll/lib"
	"wechatdll/models"
	"wechatdll/models/Tools"
)

// Secautoauth二次登录111
func GetSecautoauthReq(userInfo *comm.LoginData) ([]byte, *Algorithm.Client, error) {

	userInfo.EcPublicKey, userInfo.EcPrivateKey = cecdh.GenerateEccKey()

	//基础设备信息
	Imei := userInfo.DeviceInfo.Imei
	SoftType := userInfo.DeviceInfo.SoftTypeXML
	tmpTime := int(time.Now().UnixNano() / 1000000000)
	tmpTimeStr := strconv.Itoa(tmpTime)
	ClientSeqId := string(userInfo.DeviceInfo.Imei + "-" + tmpTimeStr)
	WCExtInfoseq := GetExtPBSpamInfoData(userInfo)

	req := &wechat.AutoAuthRequest{
		RsaReqData: &wechat.AutoAuthRsaReqData{
			AesEncryptKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(userInfo.Sessionkey))),
				Buffer: userInfo.Sessionkey,
			},
			PubEcdhKey: &wechat.ECDHKey{
				Nid: proto.Uint32(713),
				Key: &wechat.SKBuiltinString_{
					Len:    proto.Uint32(uint32(len(userInfo.EcPublicKey))),
					Buffer: userInfo.EcPublicKey,
				},
			},
		},
		AesReqData: &wechat.AutoAuthAesReqData{
			BaseRequest: &wechat.BaseRequestPlus{
				ClientVersion: proto.Int32(int32(userInfo.ClientVersion)),
				DeviceId:      userInfo.DeviceInfo.DeviceID,
				OsType:        &userInfo.DeviceInfo.OsType,
				SessionKey:    userInfo.Sessionkey,
				Uin:           &userInfo.Uin,
				Scene:         proto.Uint32(2),
			},
			BaseReqInfo: &wechat.BaseAuthReqInfo{},
			AutoAuthKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(userInfo.Autoauthkey))),
				Buffer: userInfo.Autoauthkey,
			},
			Imei:         &Imei,
			SoftType:     &SoftType,
			BuiltinIpSeq: proto.Uint32(0),
			ClientSeqId:  &ClientSeqId,
			Signature:    proto.String(""),
			DeviceName:   proto.String(userInfo.DeviceInfo.DeviceName),
			DeviceType:   proto.String("iPhone"),
			Language:     proto.String("Zh"),
			TimeZone:     proto.String("8.0"),
			ExtSpamInfo: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(WCExtInfoseq))),
				Buffer: WCExtInfoseq,
			},
		},
	}
	reqdata, err := proto.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	hecData := Tools.Pack(userInfo, reqdata, 763, 1)

	//hec := &Algorithm.Client{}
	//hec.InitPlus("IOS", int(baseinfo.ClientVersion), userInfo.DeviceInfo.OsType)
	//hecData := hec.HybridEcdhPackIosEnPlus(763, userInfo.Uin, userInfo.Cooike, reqdata)
	//fmt.Println(hecData)

	httpclient := Mmtls.GenNewHttpClient(userInfo.MmtlsKey, userInfo.MmtlsHost)
	recvData, err := httpclient.MMtlsPost(userInfo.MmtlsHost, "/cgi-bin/micromsg-bin/secautoauth", hecData, userInfo.Proxy)

	//httpclient := Mmtls.GenNewHttpClient(D.MmtlsKey, D.MmtlsHost)
	//recvData, err := httpclient.MMtlsPost(D.MmtlsHost, "/cgi-bin/micromsg-bin/pushloginurl", hecData, D.Proxy)

	if err != nil {

	}
	if len(recvData) < 32 {

	}

	// ---------------------
	hec := &Algorithm.Client{}
	//hec.InitPlus("IOS", int(baseinfo.ClientVersion), userInfo.DeviceInfo.OsType)
	//hecData := hec.HybridEcdhPackIosEnPlus(763, userInfo.Uin, userInfo.Cooike, reqdata)
	fmt.Println(hecData)
	return hecData, hec, nil
}

// 二次登录-new
func GetSecautouthReq(userInfo *comm.LoginData) ([]byte, *SecLoginKeyMgr, error) {
	userInfo.EcPublicKey, userInfo.EcPrivateKey = cecdh.GenerateEccKey()
	autoAuthKey := &wechat.AutoAuthKey{}
	err := proto.Unmarshal(userInfo.Autoauthkey, autoAuthKey)
	if err != nil {
		return nil, nil, err
	}
	userInfo.Sessionkey = autoAuthKey.EncryptKey.Buffer
	var tmpNid uint32 = 713
	var key wechat.SKBuiltinString_
	key.Buffer = userInfo.EcPublicKey
	var tmpLen = (uint32)(len(userInfo.EcPublicKey))
	key.Len = &tmpLen
	// ClientSeqId
	//tmpTime := int(time.Now().UnixNano() / 1000000000)
	//tmpTimeStr := strconv.Itoa(tmpTime)
	//var strClientSeqID = string(userInfo.DeviceInfo.Imei + "-" + tmpTimeStr)

	//基础设备信息
	Imei := Algorithm.IOSImei(userInfo.Deviceid_str)
	// TODO: 放到初始化上下文中生成
	SoftType := Algorithm.SoftType_iPad(userInfo.Deviceid_str, Algorithm.IPadOsVersion, Algorithm.IPadModel)
	ClientSeqId := lib.GetClientSeqId(userInfo.Deviceid_str)

	// extSpamInfo
	var extSpamInfo wechat.SKBuiltinString_
	extSpamInfo.Buffer = GetExtPBSpamInfoData(userInfo)
	extSpamInfoLen := uint32(len(extSpamInfo.Buffer))
	extSpamInfo.Len = &extSpamInfoLen

	req := &wechat.AutoAuthRequest{
		RsaReqData: &wechat.AutoAuthRsaReqData{
			AesEncryptKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(autoAuthKey.EncryptKey.Buffer))),
				Buffer: autoAuthKey.EncryptKey.Buffer,
			},
			PubEcdhKey: &wechat.ECDHKey{
				Nid: proto.Uint32(tmpNid),
				Key: &key,
			},
		},
		AesReqData: &wechat.AutoAuthAesReqData{
			BaseRequest: &wechat.BaseRequestPlus{
				SessionKey:    []byte{},
				Uin:           proto.Uint32(userInfo.Uin),
				DeviceId:      userInfo.Deviceid_byte,
				ClientVersion: proto.Int32(int32(userInfo.ClientVersion)),
				//DeviceType:    []byte(userInfo.DeviceType),
				Scene: proto.Uint32(0),
			},
			BaseReqInfo: &wechat.BaseAuthReqInfo{},
			AutoAuthKey: &wechat.SKBuiltinString_{
				Len:    proto.Uint32(uint32(len(userInfo.Autoauthkey))),
				Buffer: userInfo.Autoauthkey,
			},
			Imei:         &Imei,
			SoftType:     &SoftType,
			BuiltinIpSeq: proto.Uint32(0),
			ClientSeqId:  &ClientSeqId,
			DeviceName:   proto.String(userInfo.DeviceInfo.DeviceName),
			DeviceType:   proto.String("iPhone"),
			Language:     proto.String("zh_CN"),
			TimeZone:     proto.String("8.0"),
			ExtSpamInfo:  &extSpamInfo,
		},
	}
	reqData, err := proto.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	secKeyMgr := NewSecLoginKeyMgrByVer(146)
	//加密
	encrypt, epKey, token, ecdhpairkey, err := clientsdk.HybridEncrypt(reqData, secKeyMgr.WeChatPubKey)
	if err != nil {
		return nil, nil, err
	}
	/*ecdhPacket := &wechat.EcdhPacket{
		Type: proto.Uint32(1),
		Key: &wechat.BufferT{
			ILen:   proto.Uint32(415),
			Buffer: ecdhpairkey.PubKey,
		},
		Token:        token,
		Url:          proto.String(""),
		ProtobufData: encrypt,
	}*/
	ecdhPacket := &wechat.HybridEcdhRequest{
		Type: proto.Int32(1),
		SecECDHKey: &wechat.BufferT{
			ILen:   proto.Uint32(415),
			Buffer: ecdhpairkey.PubKey,
		},
		Randomkeydata:       token,
		Randomkeyextenddata: epKey,
		Encyptdata:          encrypt,
	}
	secKeyMgr.PubKey = ecdhpairkey.PubKey
	secKeyMgr.PriKey = ecdhpairkey.PriKey
	secKeyMgr.SourceData = reqData
	secKeyMgr.FinalSha256 = append(secKeyMgr.FinalSha256, epKey[24:]...)
	secKeyMgr.FinalSha256 = append(secKeyMgr.FinalSha256, reqData...)
	ecdhDataPacket, err := proto.Marshal(ecdhPacket)
	if err != nil {
		return nil, nil, err
	}

	packHeader := Tools.CreatePackHead(userInfo, baseinfo.MMPackDataTypeUnCompressed, 763, ecdhDataPacket, ecdhDataPacket, uint32(len(ecdhDataPacket)), 12, uint32(0x4e))
	//设置Hybrid 加密密钥版本
	packHeader.HybridKeyVer = secKeyMgr.WeChatPubKeyVersion
	//开始组头
	retData := Tools.PackHeaderSerialize(packHeader, false)
	return retData, secKeyMgr, nil
}

func SecautoauthNew(Wxid string) models.ResponseResult {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("异常：%v", err.Error()),
			Data:    nil,
		}
	}

	if len(D.Autoauthkey) <= 0 {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: "账号异常：Autoauthkey读取失败",
			Data:    nil,
		}
	}
	retData, hec, err := GetSecautoauthReq(D)
	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	httpclient := Mmtls.GenNewHttpClient(D.MmtlsKey, D.MmtlsHost)
	respData, err := httpclient.MMtlsPost(D.MmtlsHost, "/cgi-bin/micromsg-bin/secautoauth", retData, D.Proxy)
	fmt.Println(hec)

	//resp, err := Mmtls.MMHTTPPostData(D.GetMMInfo(), "/cgi-bin/micromsg-bin/secautoauth", hecData)
	//if err != nil {
	//	return nil, err
	//}

	if err != nil {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}
	fmt.Println(444444)
	if len(respData) <= 31 {
		return models.ResponseResult{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("组包异常3, 返回31字节"),
			Data:    nil,
		}
	}

	ph1 := hec.HybridEcdhPackIosUn(respData)
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

	//loginRes := UnifyAuthResponse
	//
	//if loginRes.GetBaseResponse().GetRet() != 0 || loginRes.BaseResponse == nil || loginRes.AuthSectResp == nil {
	//	return models.ResponseResult{
	//		Code:    -8,
	//		Success: false,
	//		Message: "登录失败：您可能已退出微信",
	//		Data:    loginRes,
	//	}
	//}
	//
	//Wx_loginecdhkey := Algorithm.DoECDH713Key(prikey, loginRes.GetAuthSectResp().GetSvrPubEcdhkey().GetKey().GetBuffer())
	//m := md5.New()
	//m.Write(Wx_loginecdhkey)
	//D.Loginecdhkey = Wx_loginecdhkey
	//ecdhdecrptkey := m.Sum(nil)
	//D.Cooike = ph1.Cookies
	//D.Sessionkey = Algorithm.AesDecrypt(loginRes.GetAuthSectResp().GetSessionKey().GetBuffer(), ecdhdecrptkey)
	//D.Autoauthkey = loginRes.GetAuthSectResp().GetAutoAuthKey().GetBuffer()
	//D.Autoauthkeylen = int32(loginRes.GetAuthSectResp().GetAutoAuthKey().GetILen())
	//D.Serversessionkey = loginRes.GetAuthSectResp().GetServerSessionKey().GetBuffer()
	//D.Clientsessionkey = loginRes.GetAuthSectResp().GetClientSessionKey().GetBuffer()
	//
	//err = comm.CreateLoginData(*D, D.Wxid, 0)
	//
	//if err != nil {
	//	return models.ResponseResult{
	//		Code:    -8,
	//		Success: false,
	//		Message: fmt.Sprintf("系统异常：%v", err.Error()),
	//		Data:    nil,
	//	}
	//}

	return models.ResponseResult{
		Code:    1,
		Success: false,
		Message: "登陆成功",
		Data:    "",
	}
}
