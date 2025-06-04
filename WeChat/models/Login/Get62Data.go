package Login

import (
	"wechatdll/comm"
	"wechatdll/lib"
)

func Get62Data(Wxid string) string {
	D, err := comm.GetLoginata(Wxid)
	if err != nil {
		return err.Error()
	}
	return lib.Get62Data(D.Deviceid_str)
}
