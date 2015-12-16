package main

import (
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "192.168.0.101:8899")
	
	if err != nil {
		panic(err)
	}
	
	conn, err := net.DialUDP("udp", nil, addr)
	
	conn.Write([]byte{ 0x42 })
	

	
	i := uint8(0); 
	for {
		time.Sleep(time.Millisecond * 50)
		conn.Write([]byte{ 0x4E, 0xFF, 0x55 })
		conn.Write([]byte{ 0x40, byte(i), 0x55 })
		i++
	}
	
}