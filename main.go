package main

import(
	"log"
	"net"
)

func main() {
	//Initialize server
	s := newServer()
	go s.run()

	//Start TCP server
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to connect to server! %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on port :8888")

	//Accept new clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection.  %s", err.Error())
			continue
		}
		//Initialize client
		go s.newClient(conn)
	}
}