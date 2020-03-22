package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/JosephZoeller/maritime-royale/pkg/grid"
	
	"github.com/JosephZoeller/maritime-royale/pkg/mrp"

	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
)

//Square contains all the data about a specific square

var mapData = map[int]map[int]grid.Square{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		var temp = map[int]grid.Square{}
		for y := 0; y < MAPY; y++ {
			if x%2 == 0 {
				temp[y] =
					grid.Square{
						Coords: grid.Coordinate{
							XPos: x,
							YPos: y},
						Terrain: terrain.NewIslandServer()}
			} else {
				temp[y] =
					grid.Square{
						Coords: grid.Coordinate{
							XPos: x,
							YPos: y},
						Terrain: terrain.NewEmpty()}
			}
		}
		mapData[x] = temp
	}

}

func sendMap(conn net.Conn) {

	count := 0

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
	fmt.Println(count)
}
