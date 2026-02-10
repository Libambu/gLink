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
	//建立一个携程开始循环等待
	go func() {
		for {
			con, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("[err]glink acceptTcp err:", err)
				continue
			}
			go func(con *net.TCPConn) {
				defer con.Close()
				for true {
					//V0.1数据echo
					buf := make([]byte, 512)
					cnt, err := con.Read(buf)
					remoteAddr := con.RemoteAddr()
					fmt.Printf("receive msg from %s : %s\n", remoteAddr.String(), string(buf[:cnt]))
					msg := "[echo]" + string(buf[:cnt])
					if err != nil {
						fmt.Println("[err]glink read err:", err)
						continue
					}
					con.Write([]byte(msg + "\r\n"))
				}
			}(con)
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
