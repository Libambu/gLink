package main

import (
	"fmt"
	"gLink/gIface"
	"gLink/gNet"
)

type PingRouter struct {
	gNet.BaseRouter
}

func (p *PingRouter) PreHandler(request gIface.IRequest) {
	fmt.Println("PingRouter Prehandler....")
	msg := string(request.GetData())
	connect := request.GetConnection().GetTCPConnect()
	connect.Write([]byte("PingRouter Prehandler " + msg))
}

func (p *PingRouter) Handler(request gIface.IRequest) {
	fmt.Println("PingRouter handler....")
	msg := string(request.GetData())
	connect := request.GetConnection().GetTCPConnect()
	connect.Write([]byte("PingRouter handler " + msg))
}

func (p *PingRouter) PostHandler(request gIface.IRequest) {
	fmt.Println("PingRouter Posthandler....")
	msg := string(request.GetData())
	connect := request.GetConnection().GetTCPConnect()
	connect.Write([]byte("PingRouter Posthandler " + msg))
}

func main() {
	server := gNet.NewServer("glink0.3")
	server.AddRouter(&PingRouter{})
	server.Serve()
}
