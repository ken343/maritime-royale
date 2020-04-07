package explode

import (
	"net"

	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/playerControl"

	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/attack"
	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/health"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/physics"
	"github.com/ken343/maritime-royale/pkg/gamestate"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
)

type Explosion struct {
	container *elements.Element
	Type      string
	coliData  elements.Component
	damData   elements.Component
}

func init() {
	var comp = new(Explosion)
	gamestate.MRPMAP["Explosion"] = comp
}

func (explo *Explosion) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewExplosion(finalElem)
	finalElem.AddComponent(myComp)
}

func NewExplosion(container *elements.Element) *Explosion {
	return &Explosion{
		container: container,
		Type:      "Explosion",
		coliData:  container.GetComponent(new(physics.Collider)),
		damData:   container.GetComponent(new(attack.Damage)),
	}
}

func (explo *Explosion) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

func (explo *Explosion) OnMerge(compM elements.Component) error {
	return nil
}

func (explo *Explosion) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (explo *Explosion) OnCheck(elemC *elements.Element) error {
	return nil
}

func (explo *Explosion) OnUpdateServer() error {
	if explo.coliData.(*physics.Collider).HasCollided {
		for _, elem := range explo.coliData.(*physics.Collider).GetObjectsHit() {
			elemHP := elem.GetComponent(new(health.Health))
			if elemHP != nil && explo.damData != nil {
				elemHP.(*health.Health).TakeDamage(explo.damData.(*attack.Damage).Attack)
			}
		}

		gamestate.RemoveElem(explo.container)
	}

	if explo.container.GetComponent(new(playerControl.MoveTo)) == nil {
		gamestate.RemoveElem(explo.container)
	}
	return nil
}

func (explo *Explosion) SetContainer(container *elements.Element) error {
	explo.container = container
	explo.coliData = container.GetComponent(new(physics.Collider))
	explo.damData = container.GetComponent(new(attack.Damage))
	return nil
}

func (explo *Explosion) MakeCopy() elements.Component {
	myComp := *explo
	return &myComp
}
