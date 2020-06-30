package main
import  (
	"GPIO/mDNS"
	"github.com/warthog618/gpiod"
	"log"

	// "github.com/warthog618/gpiod/device/rpi"
	"fmt"
	"time")
func main() {
	fmt.Println("Iniciando Aplicação...")

	mDNS.SetDNS()
	c, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("myapp"))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	l,err:=c.RequestLine(4,gpiod.AsOutput(0))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)

	fmt.Println("apaga 1s")
	l.SetValue(1)
	time.Sleep(1*time.Second)

	fmt.Println("Aceso 3 segundos")
	l.SetValue(0)
	time.Sleep(3*time.Second)

	inf, _ := l.Info()

	fmt.Printf("name: %s\n", inf.Name)
	fmt.Printf("deve apagar 3 segundos")
	l.SetValue(1)
	time.Sleep(3*time.Second)
	l.SetValue(0)
	time.Sleep(5*time.Second)

	fmt.Println("Fim da Aplicação...")
}
func handler(evt gpiod.LineEvent) {
	// handle change in line state
}
