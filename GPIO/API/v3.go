package main

import (
	"github.com/gorilla/mux"
	"github.com/warthog618/gpiod"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/{id}={val}", func(w http.ResponseWriter, r *http.Request) {
		chip, err := gpiod.NewChip("gpiochip0", gpiod.WithConsumer("myapp"))
		if err != nil {
			log.Fatal(err)
		}
		defer chip.Close()
		params := mux.Vars(r)

		pin, _ := strconv.Atoi(params["id"])
		val, _ := strconv.Atoi(params["val"])

		l, err := c.RequestLine(pin, gpiod.AsOutput())
		if err != nil {
			log.Fatal(err)
		}

		l.SetValue(val)
	})
	// Create default web handler, and call a starting web page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web_static.html")
		println("Default Web Page")
	})
	// start a listening on port 8081
	log.Fatal(http.ListenAndServe("8081", nil))

}
