package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "Mary", Id: "12345"}
	t, _ := template.ParseFiles("template/templates/first-template.html")
	err := t.Execute(w, p)
	if err != nil {
		log.Fatal("Error occured while executing the template: ", err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("template/static"))))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
