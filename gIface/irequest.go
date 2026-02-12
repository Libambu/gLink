package gIface

type IRequest interface {
	//获取当前的连接对象
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}
