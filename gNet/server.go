package gNet

import (
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
	//当前Server的绑定的Router对象
	Router gIface.IRouter
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
			dealConnect := NewConnect(con, conId, s.Router)
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

func (s *Server) AddRouter(router gIface.IRouter) {
	s.Router = router
}

func NewServer(name string) gIface.IServer {
	s := &Server{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
