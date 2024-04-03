package main

import (
	"log"
	"net/http"
	"text/template"
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
	fileServer := http.FileServer(http.Dir("template/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
	}
}
