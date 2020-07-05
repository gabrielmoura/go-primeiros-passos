package main

import (
	"github.com/warthog618/gpiod"
	"log"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Device struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	GPIO *GPIO  `json:"gpio,omitempty"`
}

type GPIO struct {
	Pin    int `json:"pin,omitempty"`
	Status int `json:"status"`
}

var Gpio []GPIO
var list = append(Gpio, GPIO{Pin: 4, Status: 1}, GPIO{Pin: 17, Status: 1})

type App struct {
	Router *mux.Router
	Chip   *gpiod.Chip
	Lines  *gpiod.Lines
}

func (a *App) Initialize(chipname, consumer string) {
	a.Router = mux.NewRouter()
	a.Chip, _ = gpiod.NewChip(chipname, gpiod.WithConsumer(consumer))
	a.Lines, _ = a.Chip.RequestLines([]int{list[0].Pin, list[1].Pin}, gpiod.AsOutput(0, 0))
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
func (a *App) getDevice(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, list)
}
func (a *App) getGPIO(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, list)
}
func (a *App) setGPIO(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	pin, _ := strconv.Atoi(params["pin"])
	val, _ := strconv.Atoi(params["val"])

	n := 0
	for range list {
		if list[n].Pin == pin {
			list[n].Status = val
		}
		n++
	}
	a.Lines.SetValues([]int{list[0].Status, list[1].Status})
	json.NewEncoder(w).Encode(list)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/gpio", a.getGPIO).Methods("GET")
	a.Router.HandleFunc("/gpio/{pin:[0-9]+}={val:[0-9]+}", a.setGPIO).Methods("GET")

}
func main() {
	a := App{}
	a.Initialize("gpiochip0", "myapp")
	a.Run(":8010")
}
