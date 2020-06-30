package main

import (


	"fmt"
	"github.com/go-resty/resty"

	"log"

	"strconv"
	"time"
)
const (
	user = "blx32"
	pass = "Ardamaxx"
)
type Conf struct {
	Token string `json:"token"`
}
func  GetToken() *resty.Response {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
	//	SetHeader("Accept", "application/json").
		Get("https://httpbin.org/basic-auth/"+user+"/"+pass)
	if(err!=nil){
		log.Fatal(err)
	}
	return resp
}
func  GetAll() *resty.Response {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page_no": "1",
			"limit": "20",
			"sort":"name",
			"order": "asc",
			"random":strconv.FormatInt(time.Now().Unix(), 10),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("https://httpbin.org/get")
	if(err!=nil){
		log.Fatal(err)
	}
	return resp
}


func main() {
	fmt.Println("Starting the application...")
	fmt.Println(GetAll())
	fmt.Println("Terminating the application...")
}