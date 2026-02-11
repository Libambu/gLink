package gNet

import (
	"fmt"
	"gLink/gIface"
	"net"
)

type Connection struct {
	//当前连接的socket套接字
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//当前的连接状态
	isClosed bool
	//当前连接所绑定的业务处理方法
	handleFunc gIface.HandleFunc
	//监控当前conn是否关闭的Channel
	ExitChan chan bool
}

func NewConnect(conn *net.TCPConn, connID uint32, callBack gIface.HandleFunc) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		handleFunc: callBack,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
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
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("connId:", c.ConnID, "receiver err :", err)
			break
		}
		err = c.handleFunc(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("connId:", c.ConnID, " handleFunc err :", err)
			break
		}
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
