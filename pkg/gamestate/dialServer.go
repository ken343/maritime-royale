package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/elements/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

var serverConnection net.Conn

func Dial(address string) {
	var err error

	serverConnection, err = net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	go mrp.ReadMRPFromConn(serverConnection, handleMRP)
}

func handleMRP(newMRPList []*mrp.MRP, conn net.Conn) {
	for _, mrpItem := range newMRPList {
		switch mrpItem.GetRequest() {
		case "ELEM":
			var tempElem map[string]interface{}
			var finalElem *elements.Element
			json.Unmarshal([]byte(mrpItem.GetBody()), &tempElem)

			switch tempElem["Type"].(string) {

			case "water":
				finalElem = objects.NewWater(0, 0, "")
				json.Unmarshal([]byte(mrpItem.GetBody()), &finalElem)

			case "player":
				finalElem = objects.NewPlayer(conn)
				json.Unmarshal([]byte(mrpItem.GetBody()), &finalElem)

			default:
				fmt.Println("No Match Found for Tile Data Type:", tempElem["Type"].(string))

			}

			test := tempElem["Components"].([]interface{})
			for _, comp := range test {

				var myComp elements.Component

				switch comp.(map[string]interface{})["Type"].(string) {

				case "SpriteRenderer":
					var sr *render.SpriteRenderer
					myComp = finalElem.GetComponent(sr)
					bytes, _ := json.Marshal(comp)
					json.Unmarshal(bytes, &myComp)

				case "KeyboardMover":
					var sr *playerControl.KeyboardMover
					myComp = finalElem.GetComponent(sr)
					bytes, _ := json.Marshal(comp)
					json.Unmarshal(bytes, &myComp)

				case "Replicator":

				default:
					fmt.Println("Component not defined")

				}
			}
			elementListTemp = append(elementListTemp, finalElem)

		case "END":
			PushElemMap()

		}
	}
}

func UpdateGamestateFromServer() {
	myMRP := mrp.NewMRP([]byte("ELEM"), []byte("test"), []byte("test"))
	serverConnection.Write(myMRP.MRPToByte())
}

func GetServerConnection() net.Conn {
	return serverConnection
}
