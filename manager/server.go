package manager

import (
	"log"
	"net/http"
	"strconv"
)

func ListenAndServe(port int) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", Handler)
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler,
	}
	log.Printf("HTTP: Serv on :%d", port)
	log.Panicf("manager died with: %v", server.ListenAndServe())
}
