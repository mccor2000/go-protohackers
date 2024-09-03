package p0

import (
	"log"
	"net"
	"os"
)


// ---------- TCP server ------------

type TCPServer struct {
	Host string
	Port string
	Listener net.Listener
}

type Handler func(net.Conn)

func New(h string, p string) *TCPServer {
	l, err := net.Listen("tcp", h+":"+p)
	if err != nil {
		log.Fatalf("Failed to listen on %s:%s\n %s", h, p, err.Error())
		os.Exit(1)
	}
	log.Printf("TCP server started on %s:%s\n", h, p)


	return &TCPServer{
		Host: h,
		Port: p,
		Listener: l,
	}
}

func(s TCPServer) Close() {
	s.Listener.Close()
}
// ---------- TCP server ------------

// ----------- P0 -------------------
func StartEchoServer() {
	s := New("0.0.0.0", "5656")
	defer s.Close()

	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection: %s\n", err.Error())
			os.Exit(1)
		}

		go echo(conn)
	}
}

func echo(socket net.Conn) {
	defer socket.Close()

	buffer := make([]byte, 10000)

	for {
		packet := make([]byte, 1024)
		length, err := socket.Read(packet)

		// EOF
		if length == 0 {
			break
		}

		// Other errors
		if err != nil {
			log.Printf("error %s\n", err)
			break
		}

		buffer = append(buffer, packet...)
	}

	// write data to responsete
	socket.Write(buffer)
}
