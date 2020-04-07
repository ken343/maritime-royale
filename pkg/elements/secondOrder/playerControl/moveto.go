package playerControl

import (
	"math"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/advancePos"
	"github.com/ken343/maritime-royale/pkg/gamestate"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type MoveTo struct {
	container *elements.Element
	posData   elements.Component
	Type      string
	DestX     float64
	DestY     float64
}

func init() {
	var mov = new(MoveTo)
	gamestate.MRPMAP["MoveTo"] = mov
}

//NewReplicator creates a Replicator which is
//the component that handles all replication
//of an element onto a server.
func NewMoveTo(container *elements.Element, DestX float64, DestY float64) *MoveTo {

	return &MoveTo{
		container: container,
		Type:      "MoveTo",
		posData:   container.GetComponent(new(advancePos.AdvancePosition)),
		DestX:     DestX,
		DestY:     DestY,
	}

}

func (mov *MoveTo) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewMoveTo(finalElem, 0, 0)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (mov *MoveTo) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (mov *MoveTo) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (mov *MoveTo) OnCheck(elemC *elements.Element) error {
	return nil
}

func (mov *MoveTo) OnMerge(compM elements.Component) error {

	return nil
}

func (mov *MoveTo) OnUpdateServer() error {

	hypot := math.Hypot(mov.DestX-mov.container.XPos, mov.DestY-mov.container.YPos)
	if hypot <= 10 {
		mov.container.RemoveComponentType(mov)
	}

	UnitX := (mov.DestX - mov.container.XPos) / hypot
	UnitY := (mov.DestY - mov.container.YPos) / hypot
	mov.posData.(*advancePos.AdvancePosition).VX += mov.posData.(*advancePos.AdvancePosition).Speed * UnitX
	mov.posData.(*advancePos.AdvancePosition).VY += mov.posData.(*advancePos.AdvancePosition).Speed * UnitY

	return nil
}

func (mov *MoveTo) SetContainer(container *elements.Element) error {
	mov.container = container
	mov.posData = container.GetComponent(new(advancePos.AdvancePosition))
	return nil
}

func (mov *MoveTo) MakeCopy() elements.Component {
	myComp := *mov
	return &myComp
}
