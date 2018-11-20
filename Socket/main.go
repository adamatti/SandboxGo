package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")

	if err != nil {
		log.Fatal("Error", err)
	}
	log.Print("All good")
	conn.Close()
}
