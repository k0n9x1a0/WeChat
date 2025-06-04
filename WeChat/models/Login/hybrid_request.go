package Login

import (
	"encoding/hex"
	"github.com/gogo/protobuf/proto"
	"github.com/lunny/log"
	"strconv"
	"strings"
	"time"
	"wechatdll/Cilent/wechat"
	cecdh "wechatdll/clientsdk/cecdn"
	"wechatdll/clientsdk/extinfo"
	clientsdk "wechatdll/clientsdk/hybrid"
	"wechatdll/comm"
	"wechatdll/models/baseutils"
)

type SecLoginKeyMgr struct {
	WeChatPubKeyVersion byte
	WeChatPubKey        string `json:"-"`
	SourceData          []byte `json:"-"`
	PriKey              []byte `json:"-"`
	PubKey              []byte `json:"-"`
	FinalSha256         []byte `json:"-"`
}

func (Sec *SecLoginKeyMgr) Reset() {
	Sec.PubKey = []byte{}
	Sec.PriKey = []byte{}
	Sec.FinalSha256 = []byte{}
	Sec.SourceData = []byte{}
}

func (Sec *SecLoginKeyMgr) SetKey() {
	if 146 == Sec.WeChatPubKeyVersion {
		Sec.WeChatPubKeyVersion = 145
		Sec.WeChatPubKey = clientsdk.WeChatPubKey_145
	} else {
		Sec.WeChatPubKeyVersion = 146
		Sec.WeChatPubKey = clientsdk.WeChatPubKey_146
	}
}

// GetBaseRequest 获取baserequest
func GetBaseRequest(userInfo *comm.LoginData) *wechat.BaseRequest {
	ret := &wechat.BaseRequest{}
	ret.SessionKey = []byte(userInfo.Sessionkey)
	ret.Uin = &userInfo.Uin
	if !strings.HasPrefix(userInfo.LoginDataInfo.LoginData, "A") && userInfo.DeviceInfo != nil {
		ret.DeviceId = userInfo.DeviceInfo.DeviceID
		ret.ClientVersion = proto.Int32(int32(userInfo.ClientVersion))
		ret.DeviceType = []byte(userInfo.DeviceType)
		ret.Scene = proto.Uint32(0)
		//log.Info("ios is base request")
	} else {
		ret.ClientVersion = proto.Int32(int32(userInfo.ClientVersion))
		ret.DeviceType = []byte(userInfo.DeviceType)
		ret.DeviceId = userInfo.Deviceid_byte
		ret.Scene = proto.Uint32(1)
		//log.Info("android is base request")
	}
	return ret
}

func NewSecLoginKeyMgrByVer(ver byte) *SecLoginKeyMgr {
	sec := &SecLoginKeyMgr{}
	switch ver {
	case 146:
		sec.WeChatPubKeyVersion = 146
		sec.WeChatPubKey = clientsdk.WeChatPubKey_146
		break
	case 145:
		sec.WeChatPubKeyVersion = 145
		sec.WeChatPubKey = clientsdk.WeChatPubKey_145
		break
	}
	return sec
}

func NewSecLoginKeyMgr() *SecLoginKeyMgr {
	return &SecLoginKeyMgr{
		WeChatPubKeyVersion: 146,
		WeChatPubKey:        clientsdk.WeChatPubKey_146,
	}
}

func GetExtPBSpamInfoData(userInfo *comm.LoginData, wxId ...string) []byte {
	wxId_ := ""
	if len(wxId) == 0 {
		wxId_ = userInfo.GetUserName()
	} else {
		wxId_ = wxId[0]
	}

	retData, err := extinfo.GetCCDPbLib(
		userInfo.DeviceInfo.OsTypeNumber,
		userInfo.DeviceInfo.OsType,
		userInfo.DeviceInfo.UUIDOne,
		userInfo.DeviceInfo.UUIDTwo,
		userInfo.DeviceInfo.DeviceName,
		userInfo.DeviceInfo.DeviceToken,
		hex.EncodeToString(userInfo.DeviceInfo.DeviceID),
		wxId_,
		userInfo.DeviceInfo.GUID2,
		userInfo,
	)
	if err != nil {
		log.Info(err)
	}
	return retData
}

// GetManualAuthAccountProtobuf 组用户登录基本信息
func GetManualAuthRsaReqDataProtobuf(userInfo *comm.LoginData, wxid string, newpass string) *wechat.ManualAuthRsaReqData {
	var tmpNid uint32 = 713
	userInfo.EcPublicKey, userInfo.EcPrivateKey = cecdh.GenerateEccKey()
	authRequest := &wechat.ManualAuthRsaReqData{}
	// aes_key
	var aesKey wechat.SKBuiltinString_
	var tmpAesKeyLen uint32 = 16
	aesKey.Len = &tmpAesKeyLen
	aesKey.Buffer = []byte(userInfo.Sessionkey)
	authRequest.RandomEncryKey = &aesKey

	// 其它参数
	var ecdhKey wechat.ECDHKey
	var key wechat.SKBuiltinString_
	key.Buffer = userInfo.EcPublicKey
	var tmpLen = (uint32)(len(userInfo.EcPublicKey))
	key.Len = &tmpLen
	ecdhKey.Nid = &tmpNid
	ecdhKey.Key = &key
	authRequest.CliPubEcdhkey = &ecdhKey
	authRequest.UserName = &wxid
	//判断是否为iPad登录的伪密码
	if !strings.HasPrefix(newpass, "extdevnewpwd_") && !strings.HasPrefix(newpass, "strdm@") {
		newpass = baseutils.Md5Value(newpass)
	}
	authRequest.Pwd = &newpass
	return authRequest
}

// GetManualAuthAesReqProtobuf 生成自动登陆aesreq项
func GetManualAuthAesReqDataProtobuf(userInfo *comm.LoginData) *wechat.ManualAuthAesReqData {
	// if userInfo.DeviceInfoA16 != nil {
	// 	return GetManualAuthAesReqDataProtobufA16(userInfo)
	// }
	zeroUint32 := uint32(0)
	zeroInt32 := int32(0)
	emptyString := string("")

	var aesRequest wechat.ManualAuthAesReqData
	// BaseRequest
	baseReq := GetBaseRequest(userInfo)
	var tmpScene uint32 = 1
	baseReq.Scene = &tmpScene
	baseReq.SessionKey = []byte{}
	baseReq.Uin = proto.Uint32(0)
	aesRequest.BaseRequest = baseReq
	inputType := uint32(2)
	aesRequest.InputType = &inputType
	aesRequest.BaseReqInfo = &wechat.BaseAuthReqInfo{}
	if userInfo.Ticket != "" {
		aesRequest.BaseReqInfo = &wechat.BaseAuthReqInfo{
			AuthTicket: proto.String(userInfo.Ticket),
		}
		aesRequest.InputType = proto.Uint32(1)
	}

	// imei
	aesRequest.Imei = &userInfo.DeviceInfo.Imei
	aesRequest.TimeZone = &userInfo.DeviceInfo.TimeZone
	aesRequest.DeviceName = &userInfo.DeviceInfo.DeviceName
	aesRequest.DeviceType = &userInfo.DeviceInfo.DeviceName
	aesRequest.Channel = &zeroInt32
	aesRequest.Language = &userInfo.DeviceInfo.Language
	aesRequest.BuiltinIpseq = &zeroUint32
	aesRequest.Signature = &emptyString
	aesRequest.SoftType = &userInfo.DeviceInfo.SoftTypeXML
	aesRequest.DeviceBrand = &userInfo.DeviceInfo.DeviceBrand
	aesRequest.RealCountry = &userInfo.DeviceInfo.RealCountry
	aesRequest.BundleId = &userInfo.DeviceInfo.BundleID
	aesRequest.AdSource = &userInfo.DeviceInfo.AdSource

	// ClientSeqId
	tmpTime := int(time.Now().UnixNano() / 1000000000)
	tmpTimeStr := strconv.Itoa(tmpTime)
	var strClientSeqID = string(userInfo.DeviceInfo.Imei + "-" + tmpTimeStr)
	aesRequest.ClientSeqId = &strClientSeqID

	// TimeStamp
	tmpTime2 := uint32(time.Now().UnixNano() / 1000000000)
	aesRequest.TimeStamp = &tmpTime2

	// extSpamInfo
	var extSpamInfo wechat.SKBuiltinString_
	extSpamInfo.Buffer = GetExtPBSpamInfoData(userInfo)
	extSpamInfoLen := uint32(len(extSpamInfo.Buffer))
	extSpamInfo.Len = &extSpamInfoLen
	aesRequest.ExtSpamInfo = &extSpamInfo

	return &aesRequest
}

type WCExtInfo struct {
	CcData      wechat.SKBuiltinString_
	DeviceToken wechat.SKBuiltinString_
	BehaviorID  wechat.SKBuiltinString_
}
