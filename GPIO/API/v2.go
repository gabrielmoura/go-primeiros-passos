package main

import (
	"encoding/json"
	"github.com/hashicorp/mdns"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/warthog618/gpiod"
)


type Device struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	GPIO *GPIO  `json:"gpio,omitempty"`
}
type GPIO struct {
	Pin    int `json:"pin,omitempty"`
	Status int `json:"status,omitempty"`
}

var Gpio []GPIO
var list = append(Gpio, GPIO{Pin: 4, Status: 1})

func GetGPIO(w http.ResponseWriter, r *http.Request) {
	c, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("myapp"))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(c.Lines())

}
func GetGPIO2(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(list)
}

func SetGPIO(w http.ResponseWriter, r *http.Request) {
	chip, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("myapp"))
	if err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r)

	pin, _ := strconv.Atoi(params["id"])
	val, _ := strconv.Atoi(params["val"])

	l, _ := chip.RequestLine(pin, gpiod.AsOutput(1))
	defer func() {
		l.Reconfigure(gpiod.AsInput)
		l.Close()
		chip.Close()
	}()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//l.SetValue(val)
	list[0].Pin = pin
	list[0].Status = val

	l.Reconfigure(gpiod.AsOutput(val))
	json.NewEncoder(w).Encode(l.Info)

}

// função principal para executar a api
func main() {
	host, _ := os.Hostname()
	info := []string{"Raspberry PI"}
	service, _ := mdns.NewMDNSService(host, "_rpi._tcp", "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	router := mux.NewRouter()
	//	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	//	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	//	router.HandleFunc("/contato", GetPeople).Methods("GET")
	//	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	//	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	//	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	router.HandleFunc("/gpio", GetGPIO2).Methods("GET")
	router.HandleFunc("/gpio/{id}={val}", SetGPIO).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
