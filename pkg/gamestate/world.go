package gamestate

import (
	"strconv"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/objects"
)

var elementList []*elements.Element
var elementListTemp []*elements.Element

func NewWorld() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			myWater := objects.NewWater(float64(x*64), float64(y*64), strconv.Itoa(x)+","+strconv.Itoa(y))
			elementList = append(elementList, myWater)
		}
	}
	myPlayer := objects.NewPlayer(nil)
	elementList = append(elementList, myPlayer)
}

func GetWorld() []*elements.Element {
	return elementList
}

func PushElemMap() {
	var found bool = false
	for _, elemTemp := range elementListTemp {
		for _, elem := range elementList {
			if elem.ID == elemTemp.ID {
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
