package Algorithm

import (
	"crypto/elliptic"
	"hash"
)

// 0x1800312A IOS 849
// 0x1800312A IOS 849

// Mozilla/5.0 (iPhone; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.49(0x18003127) NetType/WIFI Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.49(0x18003127) NetType/WIFI Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.48(0x18003030) NetType/4G Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.42(0x18002a32) NetType/4G Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 16_7_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.48(0x1800302c) NetType/WIFI Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.49(0x18003129) NetType/4G Language/zh_HK
// Mozilla/5.0 (iPhone; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.48(0x18003030) NetType/4G Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.48(0x18003030) NetType/WIFI Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.48(0x18003030) NetType/WIFI Language/zh_CN
// Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.49(0x1800312a) NetType/WIFI Language/zh_CN
// ————————————————
// 8.0.49.2600(0x28003157) WeChat/arm64

//浏览器版本
//[]byte("Windows-QQBrowser")

// var MmtlsShortHost = "hkshort.weixin.qq.com" // "extshort.weixin.qq.com"	// "szshort.weixin.qq.com"
var MmtlsShortHost = "szshort.weixin.qq.com" // "extshort.weixin.qq.com"	// "szshort.weixin.qq.com"
var MmtlsLongHost = "szlong.weixin.qq.com"
var MmtlsLongPort = 443

// 设备类型
var IPadDeviceType = "pad-android-31"

// var IPadDeviceType = "iPhone iOS16.1.2"
var IPadModel = "pad-android-31"
var IPadOsVersion = "13.3"

// var IPhoneDeviceType = "iPhone iOS13.3"
// var IPhoneModel = "iPhone9,1"
var AndroidDeviceType = "pad-android-31"
var AndroidManufacture = "HUAWEI"
var AndroidModel = "CAM-TL00"
var AndroidRelease = "8"
var AndroidIncremental = "1"
var MacDeviceType = "pad-android-31"

// 版本号
var IPadVersion = 0x28003157
var IPhoneVersion = 0x1800312A
var AndroidVersion = 0x28003157
var MacVersion = 671101235

var RSA182_N = "D153E8A2B314D2110250A0A550DDACDCD77F5801F3D1CC21CB1B477E4F2DE8697D40F10265D066BE8200876BB7135EDC74CDBC7C4428064E0CDCBE1B6B92D93CEAD69EC27126DEBDE564AAE1519ACA836AA70487346C85931273E3AA9D24A721D0B854A7FCB9DED49EE03A44C189124FBEB8B17BB1DBE47A534637777D33EEC88802CD56D0C7683A796027474FEBF237FA5BF85C044ADC63885A70388CD3696D1F2E466EB6666EC8EFE1F91BC9353F8F0EAC67CC7B3281F819A17501E15D03291A2A189F6A35592130DE2FE5ED8E3ED59F65C488391E2D9557748D4065D00CBEA74EB8CA19867C65B3E57237BAA8BF0C0F79EBFC72E78AC29621C8AD61A2B79B"
var RSA182_E = "010001"

type HYBRID_STATUS int32

const (
	HYBRID_ENC HYBRID_STATUS = 0
	HYBRID_DEC HYBRID_STATUS = 1
)

type Client struct {
	PubKey     []byte
	Privkey    []byte
	InitPubKey []byte
	Externkey  []byte

	Version    int
	DeviceType string

	clientHash hash.Hash
	serverHash hash.Hash

	curve elliptic.Curve

	Status HYBRID_STATUS
}

type PacketHeader struct {
	PacketCryptType byte
	Flag            uint16
	RetCode         uint32
	UICrypt         uint32
	Uin             uint32
	Cookies         []byte
	Data            []byte
}

type PackData struct {
	Reqdata          []byte
	Cgi              int
	Uin              uint32
	Cookie           []byte
	ClientVersion    int
	Sessionkey       []byte
	EncryptType      uint8
	Loginecdhkey     []byte
	Clientsessionkey []byte
	Serversessionkey []byte
	UseCompress      bool
	MMtlsClose       bool
}
