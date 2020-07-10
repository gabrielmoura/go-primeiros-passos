package main

import (
	"log"
	"net"
	"time"
)

type App struct {
	Conn net.Conn
	Err  error
}

func (a *App) init(ip, port string) {

	a.Conn, a.Err = net.Dial("tcp", ip+":"+port)
	if a.Err != nil {
		log.Panic(a.Err)
	}

}

func (a *App) writer() {
	for {
		data := []byte("love")
		log.Printf("enviando: %s\n", data)
		_, a.Err = a.Conn.Write(data)
		if a.Err != nil {
			panic(a.Err)
		}

		time.Sleep(time.Duration(400) * time.Millisecond)
	}
}

func (a *App) reader() {
	buf := make([]byte, 1024)
	for {
		var n int
		n, a.Err = a.Conn.Read(buf)
		if a.Err != nil {
			panic(a.Err)
		}
		log.Printf("recebido: %s\n", buf[:n])
	}

}
func main() {
	a := App{}
	a.init("127.0.0.1", "1200")
	go a.reader()
	a.writer()

	defer a.Conn.Close()
}
