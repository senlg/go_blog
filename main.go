package main

import (
	"go_blog/core"
	myFlag "go_blog/flag"
	"go_blog/global"
	"go_blog/routers"
)

func main() {
	// 读取配置文件
	global.Config = core.InitConf()
	// 拦截api路径加载
	global.InterceptApi = core.InitIntercept()
	// 初始化日志配置
	global.Log = core.InitLogger()
	// global.Log.Warnln("wran")
	// global.Log.Errorln("err")
	// global.Log.Infoln("info")
	// global.Log.Debugln("debug")
	// 连接数据库
	global.DB = core.InitGorm()

	myFlag.Parse()
	router := routers.InitRouter()
	router.Run(global.Config.System.GetAdress())
}
