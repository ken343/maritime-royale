package attack

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

type Damage struct {
	container *elements.Element
	Type      string
	Attack    float64
}

func init() {
	var comp = new(Damage)
	gamestate.MRPMAP["Damage"] = comp
}

func (dam *Damage) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewDamage(finalElem)
	finalElem.AddComponent(myComp)
}

func NewDamage(container *elements.Element) *Damage {
	return &Damage{
		container: container,
		Type:      "Damage",
		Attack:    5,
	}
}

func (dam *Damage) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

func (dam *Damage) OnMerge(compM elements.Component) error {
	return nil
}

func (dam *Damage) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (dam *Damage) OnCheck(elemC *elements.Element) error {
	return nil
}

func (dam *Damage) OnUpdateServer() error {
	return nil
}

func (dam *Damage) SetContainer(container *elements.Element) error {
	dam.container = container
	return nil
}

func (dam *Damage) MakeCopy() elements.Component {
	myComp := *dam
	return &myComp
}
