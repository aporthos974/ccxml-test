package main

import (
	"net/http"
	"log"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	log.Printf("Running server...")
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
