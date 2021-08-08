package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type Event struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	PhoneNumber string `json:"phoneNumber"`
}

var events = []Event{
	Event{Id: "111", Name: "birthday", Desc: "jimmy's birthday", PhoneNumber: "+61416918059"},
	Event{Id: "112", Name: "anniversary", Desc: "jimmy's anniversary", PhoneNumber: "+61416918059"},
	Event{Id: "113", Name: "xmas", Desc: "Xmas", PhoneNumber: "+61416918059"},
	Event{Id: "114", Name: "custom", Desc: "custom event", PhoneNumber: "+61416918059"},
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running!")
	fmt.Println("pong")
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("retrieving all events")
	json.NewEncoder(w).Encode(events)
}

func event(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Events")

	switch r.Method {
	case "GET":
		fmt.Println("GET: Events")
	case "POST":
		fmt.Println("POST: Event")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handleRequests() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/getAllEvents", getAllEvents)
	http.HandleFunc("/events", event)

	fmt.Println("Starting app on PORT:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func main() {
	fmt.Println("Starting Go Server !")
	handleRequests()
}
