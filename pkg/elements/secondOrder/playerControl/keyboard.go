package playerControl

import (
	"net"

	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/advancePos"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/gamestate"
	"github.com/ken343/maritime-royale/pkg/networking/connection"
)

//KeyboardMover is the component that handles all
//keyboard movement
type KeyboardMover struct {
	container *elements.Element
	posData   elements.Component
	Speed     float64
	Type      string
}

func init() {
	var mover = new(KeyboardMover)
	gamestate.MRPMAP["KeyboardMover"] = mover
}

//NewKeyboardMover creates a KeyboardMover which is
//the component that handles all keyboard movement
func NewKeyboardMover(container *elements.Element, speed float64) *KeyboardMover {
	return &KeyboardMover{
		container: container,
		posData:   container.GetComponent(new(advancePos.AdvancePosition)),
		Speed:     speed,
		Type:      "KeyboardMover",
	}
}

func (mover *KeyboardMover) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewKeyboardMover(finalElem, 1)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (mover *KeyboardMover) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (mover *KeyboardMover) OnUpdate(xOffset float64, yOffset float64) error {
	if mover.container.ID != connection.GetID() || mover.posData == nil {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		mover.posData.(*advancePos.AdvancePosition).VX += -mover.posData.(*advancePos.AdvancePosition).Speed
		mover.container.Same = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		mover.posData.(*advancePos.AdvancePosition).VX += mover.posData.(*advancePos.AdvancePosition).Speed
		mover.container.Same = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		mover.posData.(*advancePos.AdvancePosition).VY += -mover.posData.(*advancePos.AdvancePosition).Speed
		mover.container.Same = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		mover.posData.(*advancePos.AdvancePosition).VY += mover.posData.(*advancePos.AdvancePosition).Speed
		mover.container.Same = false
	}

	return nil
}

func (mover *KeyboardMover) OnCheck(elemC *elements.Element) error {
	return nil
}

func (mover *KeyboardMover) OnUpdateServer() error {
	return nil
}

func (mover *KeyboardMover) OnMerge(compM elements.Component) error {
	return nil
}

func (mover *KeyboardMover) SetContainer(container *elements.Element) error {
	mover.container = container
	mover.posData = container.GetComponent(new(advancePos.AdvancePosition))
	return nil
}

func (mover *KeyboardMover) MakeCopy() elements.Component {
	myComp := *mover
	return &myComp
}
