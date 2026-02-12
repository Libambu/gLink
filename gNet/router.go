package gNet

import "gLink/gIface"

/**
为什么要用类而不直接使用接口
因为如要重写接口需要把方法都实现
而使用BaseRouter就可以只实现一部分了
*/

// 实现router，先嵌入先嵌入BaseRouter基类基类，然后根据需要对基类方法进行重写
type BaseRouter struct {
}

func (b BaseRouter) PreHandler(request gIface.IRequest) {

}

func (b BaseRouter) Handler(request gIface.IRequest) {

}

func (b BaseRouter) PostHandler(request gIface.IRequest) {

}
