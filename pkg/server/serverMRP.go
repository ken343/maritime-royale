package server

import (
	"encoding/json"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

func sendElemMap(conn net.Conn) {
	myMap := gamestate.GetWorld()

	for k, myElem := range myMap {
		bytes, err := json.Marshal(myElem)
		if err != nil {
			log.Fatal(err)
		}

		myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
		conn.Write(myMRP.MRPToByte())
		if k%1000 == 0 {
			myMRP := mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
			conn.Write(myMRP.MRPToByte())
		}
	}

	myMRP := mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
	conn.Write(myMRP.MRPToByte())
}
