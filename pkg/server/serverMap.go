package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/JosephZoeller/maritime-royale/pkg/mrp"

	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
)

//Square contains all the data about a specific square

var mapData = map[int]map[int]terrain.Terrain{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		var temp = map[int]terrain.Terrain{}
		for y := 0; y < MAPY; y++ {
			terrainTemp := terrain.NewIslandServer(x, y)
			temp[y] = &terrainTemp
		}
		mapData[x] = temp
	}
}

func sendMap(conn net.Conn) {

	for x := 0; x < len(mapData); x++ {

		for y := 0; y < len(mapData); y++ {
			body, err := json.Marshal(mapData[x][y])
			if err != nil {
				fmt.Println(err.Error())
				log.Panic()
			}

			var sendingMRP = mrp.NewMRP([]byte("MAP"), body, []byte("/"))

			packet := mrp.MRPToByte(sendingMRP)

			conn.Write(packet)
		}
	}
}
