package Login

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"strconv"
	"time"
	"wechatdll/Algorithm"
	"wechatdll/baseinfo"
	"wechatdll/clientsdk/ccdata"
	"wechatdll/lib"
	"wechatdll/models/baseutils"

	"wechatdll/Cilent/mm"
	"wechatdll/comm"
	"wechatdll/models"
)

type GetQRReq struct {
	Proxy      models.ProxyInfo
	DeviceID   string
	DeviceName string
}

type GetQRRes struct {
	baseResponse GetQRResErr
	QrBase64     string
	Uuid         string
	QrUrl        string
	ExpiredTime  string
}

type GetQRResErr struct {
	Ret   int32
	Error string
}

func CreateSoftInfoXML(deviceInfo *baseinfo.DeviceInfo) string {
	// 生成DeviceInfoXML
	var retString string
	retString = retString + "<softtype>"
	retString = retString + "<k3>" + deviceInfo.OsTypeNumber + "</k3>"
	retString = retString + "<k9>" + deviceInfo.DeviceName + "</k9>"
	retString = retString + "<k10>" + strconv.Itoa(int(deviceInfo.CoreCount)) + "</k10>"
	retString = retString + "<k19>" + deviceInfo.UUIDOne + "</k19>"
	retString = retString + "<k20>" + deviceInfo.UUIDTwo + "</k20>"
	retString = retString + "<k22>" + deviceInfo.CarrierName + "</k22>"
	retString = retString + "<k24>" + baseutils.BuildRandomMac() + "/k24"
	retString = retString + "<k33>微信</k33>"
	// <k47>: 网络类型 1-wifi
	retString = retString + "<k47>1</k47>"
	// <k50>: 是否越狱 0-非越狱 1-越狱
	retString = retString + "<k50>0</k50>"
	retString = retString + "<k51>" + deviceInfo.BundleID + "</k51>"
	retString = retString + "<k54>" + deviceInfo.IphoneVer + "</k54>"
	// <k61>: 设备UUID是新的设备，还是老的设备
	retString = retString + "<k61>" + strconv.Itoa(1) + "</k61>"
	retString = retString + "</softtype>"

	return retString
}

func CreateRandomMacAddress() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X", random.Intn(256), random.Intn(256), random.Intn(256), random.Intn(256), random.Intn(256), random.Intn(256))
}

// CreateDeviceInfo 生成新的设备信息 ipad
func createDeviceInfo(deviceId string) *baseinfo.DeviceInfo {
	deviceInfo := &baseinfo.DeviceInfo{}
	if deviceId == "" && len(deviceId) < 2 {
		deviceInfo.Imei = baseutils.RandomSmallHexString(32)
		tmpDeviceID := baseutils.HexStringToBytes(deviceInfo.Imei)
		tmpDeviceID[0] = 0x49
		deviceInfo.DeviceID = tmpDeviceID
	} else {
		tmpImei := deviceId[2:]
		deviceInfo.Imei = baseutils.RandomSmallHexString(2) + tmpImei
		tmpDeviceID := baseutils.HexStringToBytes(deviceId)
		deviceInfo.DeviceID = tmpDeviceID
	}
	deviceInfo.DeviceName = "Redmi Pad" //iPhone
	deviceInfo.DeviceMac = CreateRandomMacAddress()
	deviceInfo.TimeZone = "8.00"
	deviceInfo.Language = "zh_CN" //
	deviceInfo.DeviceBrand = "pad-android-31"
	deviceInfo.RealCountry = "CN"
	deviceInfo.IphoneVer = "Redmi Pro" //iPhone4,7
	deviceInfo.BundleID = "com.tencent.xin"
	deviceInfo.OsTypeNumber = "13.5"     //12.4.6
	deviceInfo.OsType = "pad-android-31" //+ deviceInfo.OsTypeNumber //iPhone
	deviceInfo.CoreCount = 4             // 4核
	deviceInfo.AdSource = baseutils.RandomUUID()
	deviceInfo.UUIDOne = baseutils.RandomUUID()
	deviceInfo.UUIDTwo = baseutils.RandomUUID()
	// 运营商名
	deviceInfo.CarrierName = "(null)"
	deviceInfo.SoftTypeXML = CreateSoftInfoXML(deviceInfo)
	// ClientCheckDataXML
	deviceInfo.ClientCheckDataXML = ccdata.CreateClientCheckDataXML(deviceInfo)
	return deviceInfo
}

func GetQRCODE(Data GetQRReq) models.ResponseResult2 {
	//初始化Mmtls
	httpclient, MmtlsClient, err := comm.MmtlsInitialize(Data.Proxy, Algorithm.MmtlsShortHost)
	if err != nil {
		return models.ResponseResult2{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("MMTLS初始化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	aeskey := []byte(lib.RandSeq(16)) //获取随机密钥
	deviceid := Data.DeviceID
	devicelIdByte, _ := hex.DecodeString(deviceid)

	DeviceToken, err := IPadGetDeviceToken(deviceid, Algorithm.IPadModel, Data.DeviceName, Algorithm.IPadDeviceType, int32(Algorithm.IPadVersion), *httpclient, Data.Proxy, Algorithm.MmtlsShortHost)
	if err != nil {
		DeviceToken = mm.TrustResponse{}
	}

	req := &mm.GetLoginQRCodeRequest{
		BaseRequest: &mm.BaseRequest{
			SessionKey:    []byte{},
			Uin:           proto.Uint32(0),
			DeviceId:      devicelIdByte,
			ClientVersion: proto.Int32(int32(Algorithm.IPadVersion)),
			DeviceType:    []byte(Algorithm.IPadDeviceType),
			Scene:         proto.Uint32(0),
		},
		RandomEncryKey: &mm.SKBuiltinBufferT{
			ILen:   proto.Uint32(uint32(len(aeskey))),
			Buffer: aeskey,
		},
		Opcode:           proto.Uint32(0),
		MsgContextPubKey: nil,
	}

	reqdata, err := proto.Marshal(req)

	if err != nil {
		return models.ResponseResult2{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}

	hec := &Algorithm.Client{}
	hec.Init("IOS")
	hypack := hec.HybridEcdhPackIosEn(502, 0, nil, reqdata)
	recvData, err := httpclient.MMtlsPost(Algorithm.MmtlsShortHost, "/cgi-bin/micromsg-bin/getloginqrcode", hypack, Data.Proxy)
	if err != nil {
		return models.ResponseResult2{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("系统异常：%v", err.Error()),
			Data:    nil,
		}
	}
	ph1 := hec.HybridEcdhPackIosUn(recvData)
	getloginQRRes := mm.GetLoginQRCodeResponse{}

	err = proto.Unmarshal(ph1.Data, &getloginQRRes)

	if err != nil {
		return models.ResponseResult2{
			Code:    -8,
			Success: false,
			Message: fmt.Sprintf("反序列化失败：%v", err.Error()),
			Data:    nil,
		}
	}

	if getloginQRRes.GetBaseResponse().GetRet() == 0 {
		if getloginQRRes.Uuid == nil || *getloginQRRes.Uuid == "" {
			return models.ResponseResult2{
				Code:    -9,
				Success: false,
				Message: "取码过于频繁",
				Data:    getloginQRRes.GetBaseResponse(),
			}
		}

		uuidKey := getloginQRRes.GetUuid()
		//保存redis
		err := comm.CreateLoginData(comm.LoginData{
			Uuid:          uuidKey,
			Aeskey:        aeskey,
			NotifyKey:     getloginQRRes.GetNotifyKey().GetBuffer(),
			Deviceid_str:  deviceid,
			Deviceid_byte: devicelIdByte,
			DeviceName:    Data.DeviceName,
			ClientVersion: Algorithm.IPadVersion,
			Cooike:        ph1.Cookies,
			Proxy:         Data.Proxy,
			MmtlsKey:      MmtlsClient,
			DeviceToken:   DeviceToken,
			DeviceInfo:    createDeviceInfo(""),
			LoginDate:     time.Now().Unix(),
		}, "", 300)

		if err == nil {
			return models.ResponseResult2{
				Code:    1,
				Success: true,
				Message: "成功",
				Data: GetQRRes{
					baseResponse: GetQRResErr{
						Ret:   getloginQRRes.GetBaseResponse().GetRet(),
						Error: getloginQRRes.GetBaseResponse().GetErrMsg().GetString_(),
					},
					QrBase64:    fmt.Sprintf("data:image/jpg;base64,%v", base64.StdEncoding.EncodeToString(getloginQRRes.GetQrcode().GetBuffer())),
					Uuid:        getloginQRRes.GetUuid(),
					QrUrl:       "https://api.qrserver.com/v1/create-qr-code/?data=http://weixin.qq.com/x/" + getloginQRRes.GetUuid(),
					ExpiredTime: time.Unix(int64(getloginQRRes.GetExpiredTime()), 0).Format("2006-01-02 15:04:05"),
				},
				Data62:   lib.Get62Data(deviceid),
				DeviceId: deviceid,
			}
		}
	}

	return models.ResponseResult2{
		Code:    -0,
		Success: false,
		Message: "未知的错误",
		Data:    getloginQRRes,
	}
}
