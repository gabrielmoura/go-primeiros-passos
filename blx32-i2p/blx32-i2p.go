package main

import (
	"fmt"
	"github.com/majestrate/i2p-tools/sam3"
)

const yoursam = "127.0.0.1:7656" // sam bridge

func client(server sam3.I2PAddr) {
	sam, _ := sam3.NewSAM(yoursam)
	keys, _ := sam.NewKeys()
	stream, _ := sam.NewStreamSession("clientTun", keys, sam3.Options_Small)
	fmt.Println("Client: Connecting to " + server.Base32())
	conn, _ := stream.DialI2P(server)
	conn.Write([]byte("Hello world!"))
	return
}

func main() {
	sam, _ := sam3.NewSAM(yoursam)
	keys, _ := sam.NewKeys()
	go client(keys.Addr())
	stream, _ := sam.NewStreamSession("serverTun", keys, sam3.Options_Medium)
	listener, _ := stream.Listen()
	conn, _ := listener.Accept()
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	fmt.Println(keys)
	fmt.Println("Key string: " + string(keys.String()))
	fmt.Println("Server received: " + string(buf[:n]))
}
