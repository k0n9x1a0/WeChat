package Algorithm

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
)

func (h *Client) Init(Model string) {

	h.curve = elliptic.P256()
	h.clientHash = sha256.New()
	h.serverHash = sha256.New()

	if Model == "IOS" {
		h.Privkey, h.PubKey = GetECDH415Key()
		h.Version = IPadVersion
		h.DeviceType = IPadDeviceType
		h.InitPubKey, _ = hex.DecodeString("047ebe7604acf072b0ab0177ea551a7b72588f9b5d3801dfd7bb1bca8e33d1c3b8fa6e4e4026eb38d5bb365088a3d3167c83bdd0bbb46255f88a16ede6f7ab43b5")
	}

	if Model == "Android" {
		h.Status = HYBRID_ENC
		h.Version = AndroidVersion
		h.DeviceType = AndroidDeviceType
		h.InitPubKey, _ = hex.DecodeString("0495BC6E5C1331AD172D0F35B1792C3CE63F91572ABD2DD6DF6DAC2D70195C3F6627CCA60307305D8495A8C38B4416C75021E823B6C97DFFE79C14CB7C3AF8A586")
	}

}

func (h *Client) InitPlus(Model string, version int, deviceType string) {
	h.curve = elliptic.P256()
	h.clientHash = sha256.New()
	h.serverHash = sha256.New()
	if Model == "IOS" {
		h.Privkey, h.PubKey = GetECDH415Key()
		h.Version = version
		h.DeviceType = deviceType
		h.InitPubKey, _ = hex.DecodeString("047ebe7604acf072b0ab0177ea551a7b72588f9b5d3801dfd7bb1bca8e33d1c3b8fa6e4e4026eb38d5bb365088a3d3167c83bdd0bbb46255f88a16ede6f7ab43b5")
	}
	if Model == "Android" {
		h.Status = HYBRID_ENC
		h.Version = version
		h.DeviceType = deviceType
		h.InitPubKey, _ = hex.DecodeString("0495BC6E5C1331AD172D0F35B1792C3CE63F91572ABD2DD6DF6DAC2D70195C3F6627CCA60307305D8495A8C38B4416C75021E823B6C97DFFE79C14CB7C3AF8A586")
	}

}
