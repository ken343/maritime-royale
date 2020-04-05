package server

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/objects"

	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/mrp"
)

func sendSessionID(conn net.Conn, ID string) {
	myMRP := mrp.NewMRP([]byte("ID"), []byte(ID), []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func spawnStarterShip(conn net.Conn, ID string) {
	newPlayer := objects.NewPlayer(conn)
	newPlayer.ID = ID
	newPlayer.UniqueName = newPlayer.UniqueName + ID
	gamestate.AddUnitToWorld(newPlayer)
	gamestate.PushChunks()
}
