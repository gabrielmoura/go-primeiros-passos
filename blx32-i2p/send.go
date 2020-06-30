package main

import (
	"fmt"
	"github.com/majestrate/i2p-tools/sam3"
)

const yoursam = "127.0.0.1:7656" // sam bridge

func main() {
	sam, _ := sam3.NewSAM(yoursam)
	keys, _ := sam.NewKeys()
	go client(keys.Addr())
	stream, _ := sam.NewStreamSession("serverTun", keys, sam3.Options_Medium)
	listener, _ := stream.Listen()
	conn, _ := listener.Accept()
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println(string(keys.String()))
	fmt.Println("Server received: " + string(buf[:n]))
}
