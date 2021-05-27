package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/index" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./webgame"))
	http.Handle("/", fileServer)
	http.HandleFunc("/index", hello)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

//
