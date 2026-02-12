package gIface

//创建Server的接口

type IServer interface {
	//服务开启
	Start()
	//服务停止
	Stop()
	//服务启动
	Serve()
	//添加路由功能,给当前服务注册一个路由方法，供客户端使用
	AddRouter(router IRouter)
}
