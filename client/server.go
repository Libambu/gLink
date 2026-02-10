package main

import "gLink/gNet"

func main() {
	server := gNet.NewServer("glink0.1")
	server.Serve()
}
