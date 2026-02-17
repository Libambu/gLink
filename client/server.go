package main

import (
	"fmt"
	"gLink/gIface"
	"gLink/gNet"
)

type PingRouter struct {
	gNet.BaseRouter
}

func (p *PingRouter) Handler(request gIface.IRequest) {
	fmt.Println("PingRouter handler....")
	msg := string(request.GetData())
	connect := request.GetConnection().GetTCPConnect()
	connect.Write([]byte("PingRouter handler " + msg))
}

func main() {
	server := gNet.NewServer()
	server.AddRouter(&PingRouter{})
	server.Serve()
}
