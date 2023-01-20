package main

import (
	"log"
	"net"
	"test1/api"
	"test1/input.go"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()

	log.Printf("server started on :8080")
	api.LoadFile()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		log.Printf("client has connected %s", conn.RemoteAddr().String())
		api.LoadFile()
		go input.ReadInput(conn)
	}

}
