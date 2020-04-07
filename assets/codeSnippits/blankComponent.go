package sample

import (
	"net"

	"github.com/ken343/maritime-royale/pkg/gamestate"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
)

type MyComp struct {
	container *elements.Element
	Type      string
}

func init() {
	var comp = new(MyComp)
	gamestate.MRPMAP["MyComp"] = comp
}

func (comp *MyComp) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewMyComp(finalElem)
	finalElem.AddComponent(myComp)
}

func NewMyComp(container *elements.Element) *MyComp {
	return &MyComp{
		container: container,
		Type:      "MyComp",
	}
}

func (comp *MyComp) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

func (comp *MyComp) OnMerge(compM elements.Component) error {
	return nil
}

func (comp *MyComp) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (comp *MyComp) OnCheck(elemC *elements.Element) error {
	return nil
}

func (comp *MyComp) OnUpdateServer() error {
	return nil
}

func (comp *MyComp) SetContainer(container *elements.Element) error {
	comp.container = container
	return nil
}

func (comp *MyComp) MakeCopy() elements.Component {
	myComp := *comp
	return &myComp
}
