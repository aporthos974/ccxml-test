package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Printf("Running server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
