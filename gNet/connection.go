package gNet

import (
	"fmt"
	"gLink/gIface"
	"gLink/utils"
	"net"
)

type Connection struct {
	//当前连接的socket套接字
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前的连接状态
	isClosed bool
	//监控当前conn是否关闭的Channel
	ExitChan chan bool
	//该链接的Router
	Router gIface.IRouter
}

func NewConnect(conn *net.TCPConn, connID uint32, router gIface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) GReader() {
	fmt.Printf("connId:%d GReader is starting....\n", c.ConnID)
	defer func() {
		//关闭资源
		c.Stop()
		fmt.Printf("GRead connId: %d closed successfully\n", c.ConnID)
	}()
	for true {
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("connId:", c.ConnID, "receiver err :", err)
			break
		}
		//err = c.handleFunc(c.Conn, buf, cnt)
		//if err != nil {
		//	fmt.Println("connId:", c.ConnID, " handleFunc err :", err)
		//	break
		//}
		//3.0调用路由
		req := Request{
			coon: c,
			data: buf,
		}
		c.Router.PreHandler(&req)
		c.Router.Handler(&req)
		c.Router.PostHandler(&req)
	}
}

func (c *Connection) Start() {
	fmt.Printf("ConnId:%d Connect starting .....\n", c.ConnID)
	//TODO 开启一个读gorouter
	go c.GReader()
	//TODO 开启一个写gorouter

}

func (c *Connection) Stop() {
	fmt.Printf("connId:%d Connect closing .....\n", c.ConnID)
	if c.isClosed == false {
		return
	}
	c.isClosed = true
	//关闭套接字
	c.Conn.Close()
	//TODO 我觉得要先给Exit管道发送退出信号
	//关闭管道
	close(c.ExitChan)
	fmt.Printf("connId:%d Connect close successfully .....\n", c.ConnID)
}

func (c *Connection) GetTCPConnect() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
