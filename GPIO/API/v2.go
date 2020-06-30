package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/warthog618/gpiod"
	"log"
	"net/http"
	"strconv"
)

// "Person type" (tipo um objeto)
type Device struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	GPIO *GPIO  `json:"gpio,omitempty"`
}
type GPIO struct {
	Pin    string `json:"pin,omitempty"`
	Status string `json:"status,omitempty"`
}

var Gpio []GPIO

func GetGPIO(w http.ResponseWriter, r *http.Request) {
	c, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("myapp"))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(c.Lines())

}
func cchip(pin, stat) chan GPIO {
	Gpio = append(Gpio, GPIO{Pin: pin, Status: stat})
	ch := make(chan GPIO)
	ch <- GPIO{pin, stat}
	return ch
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
	l.SetValue(val)
	json.NewEncoder(w).Encode(l.Info)

}

// função principal para executar a api
func main() {

	router := mux.NewRouter()
	//	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	//	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	//	router.HandleFunc("/contato", GetPeople).Methods("GET")
	//	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	//	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	//	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	router.HandleFunc("/gpio", GetGPIO).Methods("GET")
	router.HandleFunc("/gpio/{id}={val}", SetGPIO).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
