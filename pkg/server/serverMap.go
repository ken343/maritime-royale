package server

import (
	"net"

	"github.com/JosephZoeller/maritime-royale/pkg/weather"

	"github.com/JosephZoeller/maritime-royale/pkg/units"

	"github.com/JosephZoeller/maritime-royale/pkg/terrain"
)

//Square contains all the data about a specific square
type Square struct {
	xPos, yPos int
	terrain    terrain.TerrainServer
	unit       units.UnitServer
	weather    weather.WeatherServer
}

var mapData = map[int]map[int]Square{}

const MAPX, MAPY = 50, 50

func init() {
	for x := 0; x < MAPX; x++ {
		var temp = map[int]Square{}
		for y := 0; y < MAPY; y++ {
			if ((x*50)+y)%2 == 0 {
				temp[y] =
					Square{
						xPos:    x,
						yPos:    y,
						terrain: terrain.NewWater()}
			} else {
				temp[y] =
					Square{
						xPos:    x,
						yPos:    y,
						terrain: terrain.NewIsland()}
			}
		}
		mapData[x] = temp
	}
}

func sendMap(conn net.Conn) {
	var sentMap = ""
	for _, vx := range mapData {
		var line = ""
		for _, vy := range vx {
			line = line + vy.terrain.OnDrawServer()
		}
		sentMap = sentMap + line + "\n"
	}

	conn.Write([]byte(sentMap))
}
