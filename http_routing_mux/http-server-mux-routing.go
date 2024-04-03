package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is a GET request"))
	})

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Post request"))
	})

var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		w.Write([]byte("Hi " + name))
	})

func main() {
	router := mux.NewRouter()
	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	router.Handle("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")
	http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
}
