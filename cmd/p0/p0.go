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

func EchoServer() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("echo server starting on %s\n", PORT)

	defer listener.Close()

	for {
		socket, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		go handleSocket(socket)
	}

}

func handleSocket(socket net.Conn) {
	defer socket.Close()

	buffer := make([]byte, 10000)

	for {
		packet := make([]byte, 1024)
		length, err := socket.Read(packet)

		// EOF
		if length == 0 {
			fmt.Println("recv EOF")
			break
		}

		// Other errors
		if err != nil {
			fmt.Printf("error %s\n", err)
			break
		}

		buffer = append(buffer, packet...)
	}

	// write data to responsete
	socket.Write(buffer)
}
