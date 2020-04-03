package gamemap

import (
	"strconv"

	"github.com/jtheiss19/project-undying/pkg/elements/objects"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
)

//NewWorld inits the elementList with elements representing
//water and a single player element. This is essentially a
//test enviroment.
func NewWorld() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			myWater := objects.NewWater(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y))
			gamestate.AddElemToWorld(myWater)
			if x%2 == 1 {
				myIsland := objects.NewIsland(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y)+" ")
				gamestate.AddElemToWorld(myIsland)
			}
		}
	}
}
