package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"
	"github.com/JosephZoeller/maritime-royale/pkg/units"

	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
)

//Square contains all the data about a specific square

var terrainData = []terrain.Terrain{}
var unitData = []units.Unit{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		for y := 0; y < MAPY; y++ {
			terrainSquare := terrain.NewWater(x, y)
			terrainData = append(terrainData, &terrainSquare)
		}
	}
	unitSquare := units.NewDestroyer(5, 8)
	unitData = append(unitData, &unitSquare)
}

func sendMap(conn net.Conn) {

	for _, v := range terrainData {
		body, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err.Error())
			log.Panic()
		}

		var sendingMRP = mrp.NewMRP([]byte("MAP"), body, []byte("/"))

		packet := mrp.MRPToByte(sendingMRP)

		conn.Write(packet)
	}

	for _, v := range unitData {
		body, err := json.Marshal(v)
		if err != nil {
			fmt.Println(err.Error())
			log.Panic()
		}

		var sendingMRP = mrp.NewMRP([]byte("UNIT"), body, []byte("/"))

		packet := mrp.MRPToByte(sendingMRP)

		conn.Write(packet)
	}
}
