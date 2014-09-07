package main

import (
	"log"
	"net/http"
)

var (
	outputFile string
)

func main() {
	http.HandleFunc("/amstatus", amandaServer)
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: " + err.Error())
	}
}
