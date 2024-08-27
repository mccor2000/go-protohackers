package p0

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
	TYPE = "tcp"
)

func StartTCPServer() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	fmt.Println("received a connection")
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		if _, err := conn.Read(buffer); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("received msg %s\n", string(buffer))

		// write data to response
		conn.Write(buffer)
	}
}
