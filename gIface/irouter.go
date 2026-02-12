package gIface

/**
抽象路由接口
路由里数据都是IRequest
*/

type IRouter interface {
	//在处理业务之前执行
	PreHandler(request IRequest)
	//处理conn的业务钩子方法
	Handler(request IRequest)
	//在处理coon业务之后执行的方法
	PostHandler(request IRequest)
}
