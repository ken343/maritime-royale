package gamemap

import (
	"math/rand"
	"strconv"

	"github.com/ken343/maritime-royale/pkg/elements/objects"
	"github.com/ken343/maritime-royale/pkg/gamestate"
)

//NewWorld inits the elementList with elements representing
//water and a single player element. This is essentially a
//test enviroment.
func NewWorld() {
	gamestate.CreateChunk()
	for x := -10; x < 10; x++ {
		for y := -10; y < 10; y++ {
			myWater := objects.NewWater(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y))
			gamestate.AddTerrainToWorld(myWater)
			if rand.Intn(20) <= 0 {
				myIsland := objects.NewIsland(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y)+" ")
				gamestate.AddTerrainToWorld(myIsland)
			}
		}
	}
	gamestate.PushChunks()
}
