package server

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/ken343/maritime-royale/pkg/gamestate"
	"github.com/ken343/maritime-royale/pkg/networking/mrp"
)

//Server starts a server on the selected port and acts
//as the main entrance into the server package.
func Server(port string) {

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	fmt.Println("Server Now Listening on Port:", port)

	newConnSignal := make(chan string)

	connections := 0
	for {
		go session(server, newConnSignal, connections)
		fmt.Println(<-newConnSignal)
		connections++
	}
}

func session(ln net.Listener, newConnSignal chan string, sessionID int) {
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err.Error())
		log.Panic()
	}

	newConnSignal <- "New Connection Made"

	gamestate.NewConnection(conn, sessionID)
	sendSessionID(conn, strconv.Itoa(sessionID))

	go mrp.ReadMRPFromConn(conn, gamestate.HandleMRP)

	gamestate.SendElemMap(conn)
	spawnStarterShip(conn, strconv.Itoa(sessionID))

	closeConnection := make(chan string)
	fmt.Println(<-closeConnection)
}
