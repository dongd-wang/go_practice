package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login Page")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout Page")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error listening")
		return

	}
}
