package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type App struct {
	Err      error
	TcpAddr  *net.TCPAddr
	Listener net.Listener
	Conn     net.Conn
}

func (a *App) init(port string) {
	a.TcpAddr, a.Err = net.ResolveTCPAddr("tcp4", ":"+port)
	checkError(a.Err)
	a.Listener, a.Err = net.ListenTCP("tcp", a.TcpAddr)
	checkError(a.Err)
	for {
		a.Conn, a.Err = a.Listener.Accept()
		if a.Err != nil {
			continue
		}
		go a.handleClient()
	}
}

func (a *App) handleClient() {
	a.Conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                            // set maximum request length to 128B to prevent flood based attacks
	defer a.Conn.Close()                                    // close connection before exit
	for {
		read_len, err := a.Conn.Read(request)

		if err != nil {
			log.Println(err)
			break
		}
		if read_len == 0 {
			break
		}
		switch string(request[:read_len]) {
		case "timestamp":
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			a.Conn.Write([]byte(daytime))

			log.Println(string(request[:read_len]))

		case "clone":
			a.Conn.Write(request[:read_len])
		case "shutdown":
			a.Conn.Write([]byte("Bye"))
			os.Exit(1)

		default:
			daytime := time.Now().String()
			a.Conn.Write([]byte(daytime))

			log.Println(read_len)
		}

	}
}
func write(conn net.Conn, data string) {
	conn.Write([]byte(data))
}
func checkError(err error) {
	if err != nil {
		log.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	a := App{}
	a.init("1200")
}
