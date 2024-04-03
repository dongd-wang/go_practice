package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.RequestURI)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
