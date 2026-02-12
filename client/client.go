package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 8999))
	if err != nil {
		fmt.Println("[err]connect server err : ", err)
		return
	}
	for {
		_, err := con.Write([]byte("hello glink V0.3"))
		if err != nil {
			fmt.Println("[err]write to server err : ", err)
			return
		}
		for i := 0; i < 3; i++ {
			buf := make([]byte, 512)
			cnt, err := con.Read(buf)
			if err != nil {
				fmt.Println("[err]read from server err : ", err)
				return
			}
			msg := string(buf[:cnt])
			fmt.Printf("[receive from server] %s", msg)

		}
		println("------------------------")
		time.Sleep(time.Second * 10)
	}
}
