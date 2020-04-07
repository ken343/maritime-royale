package health

import (
	"fmt"
	"net"

	"github.com/ken343/maritime-royale/pkg/gamestate"
	"github.com/ken343/maritime-royale/pkg/networking/connection"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/ken343/maritime-royale/pkg/elements"
)

// AdvancePosition is the component that handles all
// keyboard movement
type Health struct {
	container   *elements.Element
	Type        string
	HealthTotal float64
}

func init() {
	var HP = new(Health)
	gamestate.MRPMAP["Health"] = HP
}

func NewHealth(container *elements.Element, HealthTotal float64) *Health {
	return &Health{
		container:   container,
		Type:        "Health",
		HealthTotal: HealthTotal,
	}
}

func (hp *Health) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewHealth(finalElem, 100)
	finalElem.AddComponent(myComp)
}

func (hp *Health) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	if hp.container.ID != connection.GetID() {
		return nil
	}

	msg := fmt.Sprintf("\n\n\n\n HP: %0.2f", hp.HealthTotal)
	ebitenutil.DebugPrint(screen, msg)

	return nil
}

func (hp *Health) OnMerge(compM elements.Component) error {
	return nil
}

func (hp *Health) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (hp *Health) OnCheck(elemC *elements.Element) error {
	return nil
}

func (hp *Health) OnUpdateServer() error {
	return nil
}

func (hp *Health) SetContainer(container *elements.Element) error {
	hp.container = container
	return nil
}

func (hp *Health) MakeCopy() elements.Component {
	myComp := *hp
	return &myComp
}

func (hp *Health) TakeDamage(damage float64) {
	hp.HealthTotal = hp.HealthTotal - damage
	hp.container.Same = false
	return
}
