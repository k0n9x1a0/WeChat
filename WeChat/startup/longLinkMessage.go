package startup

import (
	"fmt"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"runtime"
	"wechatdll/TcpPoll"
)

func StartUpInit() {
	longLinkEnabled, _ := beego.AppConfig.Bool("longlinkenabled")
	sysType := runtime.GOOS
	fmt.Println(sysType)
	fmt.Println(longLinkEnabled)
	if sysType == "linux" && longLinkEnabled {
		// LINUX系统
		tcpManager, err := TcpPoll.GetTcpManager()
		if err != nil {
			log.Errorf("TCP启动失败.")
		}
		go tcpManager.RunEventLoop()
	}
	if sysType == "windows" && longLinkEnabled {
		tcpManager, err := TcpPoll.GetTcpManager()
		if err != nil {
			log.Errorf("TCP启动失败.")
		}
		go tcpManager.RunEventLoop()
		log.Errorf("我是windows系统")
	}
}
