package gamestate

import (
	"encoding/json"
	"log"
	"net"

	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/networking/mrp"
)

var elementList []*elements.Element
var elementListTemp []*elements.Element

var connectionList = make(map[int]net.Conn)

//GetWorld returns the elementlist representing the current
//gamestate of the world
func GetWorld() []*elements.Element {
	return elementList
}

func AddElemToWorld(elem *elements.Element) {
	elementList = append(elementList, elem)
	for _, client := range connectionList {
		SendElem(client, elem)
		ForceUpdate(client)
	}
}

//PushElemMap pushes all qued changes in elementListTemp to
//elementList in a safe way.
func PushElemMap() {
	var found bool = false
	for _, elemTemp := range elementListTemp {
		for _, elem := range elementList {
			if elem.UniqueName == elemTemp.UniqueName {
				*elem = *elemTemp
				found = true
				break
			}
		}
		if found == true {
			found = false
		} else {
			elementList = append(elementList, elemTemp)
		}
	}

	elementListTemp = []*elements.Element{}

}

func SendElemMap(conn net.Conn) {
	myMap := GetWorld()

	for _, myElem := range myMap {
		bytes, err := json.Marshal(myElem)
		if err != nil {
			log.Fatal(err)
		}

		myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
		conn.Write(myMRP.MRPToByte())
	}

	ForceUpdate(conn)
}

func SendElem(conn net.Conn, elem *elements.Element) {
	bytes, _ := json.Marshal(&elem)

	myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func ForceUpdate(conn net.Conn) {
	myMRP := mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func NewConnection(conn net.Conn, ID int) {
	connectionList[ID] = conn
}

func UpdateElemToAll(elem *elements.Element) {
	for _, client := range connectionList {
		SendElem(client, elem)
	}
}
