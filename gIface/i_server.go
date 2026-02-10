package gIface

//创建Server的接口

type IServer interface {
	//服务开启
	Start()
	//服务停止
	Stop()
	//服务启动
	Serve()
}
