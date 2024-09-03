package p1

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"

	"github.com/mccor2000/go-protohackers/pkg/p0"
)

func IsPrimeServer() {
	s := p0.New("0.0.0.0", "5656")
	defer s.Close()

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %s\n", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

type RequestPayload struct {
	Method string      `json:"method"`
	Number interface{} `json:"number"`
}

func handleRequest(c net.Conn) {
	defer c.Close()

	for {
		packet := make([]byte, 10000)

		packet_len, err := c.Read(packet)
		if err != nil {
			log.Printf("Error while reading request: %s\n", err.Error())
			return
		}

		var payload RequestPayload

		if err := json.Unmarshal(packet[:packet_len], &payload); err != nil {
			c.Write([]byte("malformed"))
			return
		}
		fmt.Println(payload)

		if payload.Method != "isPrime" {
			c.Write([]byte("malformed"))
			return
		}

		fmt.Println(reflect.TypeOf(payload.Number))

		switch payload.Number.(type) {
		case int:
		case float64:
			c.Write([]byte(fmt.Fprintf("{\"method\":\"isPrime\",\"prime\":}\n", isPrime(i))))
		default:
			c.Write([]byte("malformed"))
		}
	}
}

func isPrime(n int) bool {
  return true
}
