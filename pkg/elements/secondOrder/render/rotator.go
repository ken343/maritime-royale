package render

import (
	"math"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/advancePos"
	"github.com/ken343/maritime-royale/pkg/gamestate"
)

//Rotator is the component that handles all
//rendering of sprites onto the screen
type Rotator struct {
	container *elements.Element
	posData   elements.Component

	Type string
}

func init() {
	var rot = new(Rotator)
	gamestate.MRPMAP["Rotator"] = rot
}

//NewRotator creates a SpriteRenderer which
//is the component that handles all rendering of
//sprites onto the screen
func NewRotator(container *elements.Element) *Rotator {

	return &Rotator{
		container: container,
		posData:   container.GetComponent(new(advancePos.AdvancePosition)),
		Type:      "Rotator",
	}
}

func (rot *Rotator) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewRotator(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw Draws the stored texture file onto the screen
func (rot *Rotator) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate is used to qualify SpriteRenderer as a component
func (rot *Rotator) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (rot *Rotator) OnCheck(elemC *elements.Element) error {
	return nil
}

func (rot *Rotator) OnMerge(compM elements.Component) error {
	return nil
}

func (rot *Rotator) OnUpdateServer() error {

	if rot.container.YPos == rot.posData.(*advancePos.AdvancePosition).PrevY && rot.container.XPos == rot.posData.(*advancePos.AdvancePosition).PrevX {
	} else if rot.posData.(*advancePos.AdvancePosition).PrevY == 0 || rot.posData.(*advancePos.AdvancePosition).PrevX == 0 {
	} else {
		rot.container.Rotation = math.Atan2((rot.container.YPos - rot.posData.(*advancePos.AdvancePosition).PrevY), (rot.container.XPos - rot.posData.(*advancePos.AdvancePosition).PrevX))
	}

	return nil
}

func (rot *Rotator) SetContainer(container *elements.Element) error {
	rot.container = container
	return nil
}

func (rot *Rotator) MakeCopy() elements.Component {
	myComp := *rot
	return &myComp
}
