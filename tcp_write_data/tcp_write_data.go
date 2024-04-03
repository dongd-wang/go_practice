package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal("Error starting TCP server : ", err)
		return
	}

	defer listener.Close()
	log.Println("TCP Server started on: ", CONN_HOST+":"+CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	fmt.Println("Message received from the client:", string(message))
	conn.Write([]byte(message + " received \n"))
	// conn.Close()
}
