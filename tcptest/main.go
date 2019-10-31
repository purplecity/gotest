package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn,err := net.DialTimeout("tcp","112.74.163.161:50000",2*time.Second)
	if err != nil {
		fmt.Println("dial failed")
		return
	}
	_,err = conn.Write([]byte("AG1408.SHF"))
	if err != nil {
		fmt.Println("write failed")
		return
	}

	for {
		var buf = make([]byte, 128)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed")
			return
		}
		fmt.Println(string(buf[:n]))
	}
}