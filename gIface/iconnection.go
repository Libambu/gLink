package gIface

import "net"

type IConnection interface {
	//启动连接，让当前连接准备开始工作
	Start()
	//停止连接 结束当前连接的工作
	Stop()
	//获取当前连接绑定的socket conn
	GetTCPConnect() *net.TCPConn
	//获取当前连接模块的链接ID
	GetConnID() uint32
	//获取远程客户端的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	//发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

/*
	在 Go 语言中，函数也是一种类型，就像 int、string 或 struct 一样。
	这行代码的意思是：定义一个名为 HandleFunc 的函数类型。
	任何函数，只要它的参数和返回值满足以下条件，它就属于 HandleFunc 类型：
	第一个参数是 *net.TCPConn
	第二个参数是 []byte（接收到的数据）
	第三个参数是 int（数据的长度）
	返回值是一个 error
	定义的业务函数MyBusiness 就可以被当做 HandleFunc 类型来传递
*/

/*
	那为什么不写在接口里面呢
	如果写在接口里：
	假设在 IConnection 里定义一个方法 Handle()。那么每一个实现 IConnection 的结构体（比如 Connection）都必须死死地写掉这个处理逻辑。
	如果你想换一种处理逻辑（比如从“回显”改成“聊天”），你得重新改写底层的 Connection 类。
	写在外面的好处：
	你可以把 HandleFunc 当做一个参数传给 Connection。
	// 伪代码：在创建连接时，把业务逻辑注入进去
	func NewConnection(conn *net.TCPConn, callback gIface.HandleFunc) {
		// ...
	}
	这样，底层的 Connection 就不需要关心业务逻辑了。它只管：“我收到数据了，然后调用一下传给我的那个 HandleFunc 就行了。”
*/

// 定义一个处理来连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
