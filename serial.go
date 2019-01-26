package main

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/tarm/serial"
)

func loadArgs() *serial.Config {
	if len(os.Args) != 3 {
		log.Fatal("Invalid arguments: serial <port> <baud>")
	}
	baud, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid baud: %s", err)
	}
	return &serial.Config{Name: os.Args[1], Baud: baud}
}

func main() {
	args := loadArgs()
	log.Print("Opening connection...")
	conn, err := serial.OpenPort(args)
	if err != nil {
		log.Fatalf("Can't open port %s", err)
	}
	log.Print("Connection opened")

	// send stdin to the connections
	go func() {
		for {
			if _, err := io.Copy(conn, os.Stdin); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// write messages from the connection to the stdout
	for {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
		}
	}

}
