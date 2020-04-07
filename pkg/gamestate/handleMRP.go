package gamestate

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/networking/connection"
	"github.com/ken343/maritime-royale/pkg/networking/mrp"
)

var serverConnection net.Conn
var MRPMAP = make(map[string]elements.Component)

//Dial setsup a gamestate to be controlled by the server dialed
//via the address variable.
func Dial(address string) {
	var err error

	serverConnection, err = net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	go mrp.ReadMRPFromConn(serverConnection, HandleMRP)
}

func HandleMRP(mrpItem *mrp.MRP, conn net.Conn) {
	switch mrpItem.GetRequest() {
	case "ELEM":
		bytesMaster := []byte(mrpItem.GetBody())

		var finalElem = new(elements.Element)
		handleELEMCreates(bytesMaster, finalElem)

		if mrpItem.GetFooters()[0] == "NIL" {
			for _, elem := range GetEntireWorld() {
				if elem.UniqueName == finalElem.UniqueName {
					blacklistedNames = append(blacklistedNames, elem.UniqueName)
					RemoveElem(elem)
					break
				}
			}
		} else {
			AddUnitToWorld(finalElem)

			PushChunks()

		}

	case "REPLIC":
		world := GetEntireWorld()
		handleREPLIC(mrpItem, conn, world)

	case "ID":
		connection.SetID(mrpItem.GetBody())

	case "END":
		PushChunks()

	default:
		fmt.Println("Command Not Understood")
	}

}

func handleELEMCreates(bytesMaster []byte, finalElem *elements.Element) {

	var tempElem = *new(map[string]interface{})

	json.Unmarshal(bytesMaster, &tempElem)

	//fmt.Println(tempElem)

	test := tempElem["Components"].([]interface{})
	for _, comp := range test {

		if comp != nil {

			//var myComp elements.Component
			kindOfComp := comp.(map[string]interface{})["Type"].(string)
			myComp := MRPMAP[kindOfComp]
			if myComp != nil {
				myComp.MRP(finalElem, serverConnection)
			}
		}
	}

	json.Unmarshal(bytesMaster, &finalElem)
}

func handleREPLIC(mrpItem *mrp.MRP, conn net.Conn, world []*elements.Element) {
	for _, elem := range world {
		if elem.UniqueName == mrpItem.GetFooters()[0] {
			var elemTemp = new(elements.Element)
			handleELEMCreates([]byte(mrpItem.GetBody()), elemTemp)

			if elem.Check(elemTemp) == nil {
				elemTemp.Merge(elem)
			}

			break
		}
	}

}
