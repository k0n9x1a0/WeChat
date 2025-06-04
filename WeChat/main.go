package main

import (
	"fmt"
	"wechatdll/comm"
	"wechatdll/middleware"
	_ "wechatdll/routers"

	"github.com/astaxie/beego"
)

func main() {

	comm.RedisInitialize()
	_, err := comm.RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("【Redis】连接失败，ERROR：%v", err.Error()))
	}
	fmt.Println("仅供内部学习研究交流使用 联系方式：408449830")
	// 项目初始化
	//startup.StartUpInit()

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/"] = "swagger"

	beego.SetLogFuncCall(false)

	// 最后一个参数必须设置为false 不然无法打印数据
	beego.InsertFilter("/*", beego.BeforeRouter, middleware.BaseAuthLog, false)

	//自定义错误页面
	beego.Run()
}
