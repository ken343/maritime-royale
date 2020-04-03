package playerControl

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//KeyboardMover is the component that handles all
//keyboard movement
type KeyboardMover struct {
	container *elements.Element
	Speed     float64
	Type      string
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

//OnDraw is used to qualify SpriteRenderer as a component
func (mover *KeyboardMover) OnDraw(screen *ebiten.Image) error {
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (mover *KeyboardMover) OnUpdate() error {
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
