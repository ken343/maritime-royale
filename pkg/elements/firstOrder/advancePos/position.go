package advancePos

import (
	"errors"
	"fmt"
	"math"
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

//AdvancePosition is the component that handles all
//keyboard movement
type AdvancePosition struct {
	container    *elements.Element
	Type         string
	PrevX, PrevY float64
	Speed        float64
	VX, VY       float64
}

func init() {
	var aPos = new(AdvancePosition)
	gamestate.MRPMAP["AdvancePosition"] = aPos
}

//NewAdvancePosition creates a KeyboardMover which is
//the component that handles all keyboard movement
func NewAdvancePosition(container *elements.Element, Speed float64) *AdvancePosition {
	return &AdvancePosition{
		container: container,
		Type:      "AdvancePosition",
		Speed:     Speed,
		PrevX:     0,
		PrevY:     0,
	}
}

func (aPos *AdvancePosition) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewAdvancePosition(finalElem, 0)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (aPos *AdvancePosition) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

func (aPos *AdvancePosition) OnMerge(compM elements.Component) error {
	compM.(*AdvancePosition).VX = aPos.VX
	compM.(*AdvancePosition).VY = aPos.VY
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (aPos *AdvancePosition) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (aPos *AdvancePosition) OnCheck(elemC *elements.Element) error {
	if math.Abs(aPos.container.XPos-elemC.XPos) >= 200 {
		fmt.Println("RubberBand")
		return errors.New("DeSync")
	}
	return nil
}

func (aPos *AdvancePosition) OnUpdateServer() error {
	aPos.PrevX = aPos.container.XPos
	aPos.PrevY = aPos.container.YPos

	aPos.container.XPos += aPos.VX
	aPos.container.YPos += aPos.VY

	if aPos.VX != 0 || aPos.VY != 0 {
		aPos.container.Same = false
	}

	aPos.VX = 0
	aPos.VY = 0

	if aPos.container.ID != connection.GetID() {
		return nil
	}
	if aPos.container.YPos == aPos.PrevY && aPos.container.XPos == aPos.PrevX {
	} else {
		aPos.container.Rotation = math.Atan2((aPos.container.YPos - aPos.PrevY), (aPos.container.XPos - aPos.PrevX))
	}

	return nil
}

func (aPos *AdvancePosition) SetContainer(container *elements.Element) error {
	aPos.container = container
	return nil
}

func (aPos *AdvancePosition) MakeCopy() elements.Component {
	myComp := *aPos
	return &myComp
}
