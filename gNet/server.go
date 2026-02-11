package gNet

import (
	"errors"
	"fmt"
	"gLink/gIface"
	"net"
)

type Server struct {
	//定义服务器的名称
	Name string
	//定义服务器ip版本
	IpVersion string
	//定义服务器监听Ip
	Ip string
	//定义服务器端口号port
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start]glink Server Listening at IP : %s,Port : %d \r\n", s.Ip, s.Port)
	//获取addr
	addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("[err]glink resolver err:", err)
		return
	}
	//开始监听
	listener, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		fmt.Println("[err]glink listenTcp  err:", err)
		return
	}
	fmt.Printf("[success]glink Server %s successful \n", s.Name)
	var conId uint32
	conId = 1
	//建立一个携程开始循环等待
	go func() {
		for {
			con, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[err]glink acceptTcp err:", err)
				continue
			}
			//创建链接模块
			dealConnect := NewConnect(con, conId, defaultCallBack)
			go dealConnect.Start()
			conId++
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

func NewServer(name string) gIface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

// TODO 目前这个Hander方法是写死的，应该由用户创建，进行传递
func defaultCallBack(conn *net.TCPConn, data []byte, cnt int) error {
	msg := string(data[:cnt])
	msg = "[echo]" + msg
	_, err := conn.Write([]byte(msg + "\r\n"))
	if err != nil {
		fmt.Println("server writer err", err)
		return errors.New("default CallBack err")
	}
	return nil
}
