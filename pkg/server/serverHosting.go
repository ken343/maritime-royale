package server

import (
	"fmt"
	"log"
	"net"
)

//Server starts a server on the selected port and acts
//as the main entrance into the server package.
func Server(port string) {

	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal := make(chan string)

	for {
		go session(server, newConnSignal)
		fmt.Println(<-newConnSignal)
	}

}

func session(l net.Listener, newConnSignal chan string) {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal <- "New Connection Made"

	closeConnection := make(chan string)

	go readMRP(conn)

	fmt.Println(<-closeConnection)
}
