package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "ELEM":
			sendElemMap(conn)
		case "REPLIC":
			for _, elem := range gamestate.GetWorld() {
				if elem.ID == mrpItem.GetFooters()[0] {
					var elemTemp *elements.Element
					json.Unmarshal([]byte(mrpItem.GetBody()), &elemTemp)

					if elemTemp.XPos-elem.XPos < 50 && elemTemp.XPos-elem.XPos > -50 {
						json.Unmarshal([]byte(mrpItem.GetBody()), &elem)
					}

					bytes, _ := json.Marshal(&elem)

					myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
					conn.Write(myMRP.MRPToByte())

					myMRP = mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
					conn.Write(myMRP.MRPToByte())
				}
			}
		default:
			fmt.Println("Command Not Understood")
		}
	}
}

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
