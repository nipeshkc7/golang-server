package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running!")
	fmt.Println("pong")
}

func handleRequests() {
	http.HandleFunc("/", ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Starting Go Server !")
	handleRequests()
}
