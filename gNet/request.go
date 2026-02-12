package gNet

import "gLink/gIface"

type Request struct {
	//已经和客户端建立好连接的Coon
	coon gIface.IConnection
	//客户端请求的数据
	data []byte
}

func (r Request) GetConnection() gIface.IConnection {
	return r.coon
}

func (r Request) GetData() []byte {
	return r.data
}
