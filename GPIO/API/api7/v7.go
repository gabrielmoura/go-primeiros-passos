package main

import (
	"encoding/json"
	"github.com/oleksandr/bonjour"
	"github.com/warthog618/gpiod"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Device struct {
	ID      string  `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	GPIO    []GPIO  `json:"gpio,omitempty"`
	Home    float64 `json:"diskhome,omitempty"`
	Root    float64 `json:"diskroot,omitempty"`
	TempCPU string  `json:"tempcpu,omitempty"`
	Memory  float64 `json:"memory,omitempty"`
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
	Err    interface{}
	mDNS   interface{}
}

func (a *App) Initialize(chipname, consumer string) {
	a.Router = mux.NewRouter()

	a.Chip, a.Err = gpiod.NewChip(chipname, gpiod.WithConsumer(consumer))
	if a.Err != nil {
		log.Fatal(a.Err)
	}
	a.Lines, a.Err = a.Chip.RequestLines([]int{list[0].Pin, list[1].Pin}, gpiod.AsOutput(0, 0))
	if a.Err != nil {
		log.Println(a.Err)
	}
	host, _ := os.Hostname()
	info := []string{"Raspberry PI"}
	// Run registration (blocking call)
	a.mDNS, a.Err = bonjour.Register(host, "_rpi._tcp", "", 8010, info, nil)
	if a.Err != nil {
		log.Println(a.Err)
	}
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
func (a *App) info(w http.ResponseWriter, r *http.Request) {
	home, _ := getHomeDiskUsage()
	root, _ := getRootDiskUsage()
	tempcpu, _ := GetCPUTemp()
	memory, _ := GetMemoryUsage()
	host, _ := getHostname()
	info := Device{Home: home, Root: root, TempCPU: tempcpu, Memory: memory, Name: host, GPIO: list}
	respondWithJSON(w, 200, info)
}
func (a *App) getDevice(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, 200, list)
}
func (a *App) getGPIO(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, list)
}
func (a *App) setGPIO(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	n := 0

	if params["pin0"] != "" {
		pin0, _ := strconv.Atoi(params["pin0"])
		val0, _ := strconv.Atoi(params["val0"])
		for range list {
			if list[n].Pin == pin0 {
				list[n].Status = val0
			}
			n++
		}
	}

	if params["pin1"] != "" {
		pin1, _ := strconv.Atoi(params["pin1"])
		val1, _ := strconv.Atoi(params["val1"])
		for range list {
			if list[n].Pin == pin1 {
				list[n].Status = val1
			}
			n++
		}
	}

	a.Lines.SetValues([]int{list[0].Status, list[1].Status})
	respondWithJSON(w, 200, list)

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
	a.Router.HandleFunc("/gpio/list", a.getGPIO).Methods("GET")
	a.Router.HandleFunc("/gpio/info", a.info).Methods("GET")

	a.Router.PathPrefix("/gpio/set").Queries("pin1", "{pin1:[0-9]+}", "val1", "{val1:[0-9]+}").HandlerFunc(a.setGPIO).Methods("GET")
	a.Router.PathPrefix("/gpio/set").Queries("pin0", "{pin0:[0-9]+}", "val0", "{val0:[0-9]+}").HandlerFunc(a.setGPIO).Methods("GET")
	a.Router.PathPrefix("/gpio/set").Queries("pin0", "{pin0:[0-9]+}", "val0", "{val0:[0-9]+}", "pin1", "{pin1:[0-9]+}", "val1", "{val1:[0-9]+}").HandlerFunc(a.setGPIO).Methods("GET")

}
func main() {
	a := App{}
	a.Initialize("gpiochip0", "myapp")
	a.Run(":8010")

}
