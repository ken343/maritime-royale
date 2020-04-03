package playerControl

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//KeyboardMover is the component that handles all
//keyboard movement
type KeyboardMover struct {
	container *elements.Element
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
func (mover *KeyboardMover) OnUpdate(world []*elements.Element) error {
	if mover.container.ID != connection.GetID() {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		mover.container.XPos -= mover.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		mover.container.XPos += mover.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		mover.container.YPos -= mover.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		mover.container.YPos += mover.Speed
	}

	return nil
}

func (mover *KeyboardMover) OnCheck(elemC *elements.Element) error {
	return nil
}

func (mover *KeyboardMover) OnUpdateServer(world []*elements.Element) error {
	return nil
}
