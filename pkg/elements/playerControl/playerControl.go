package playerControl

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

type KeyboardMover struct {
	container *elements.Element
	Speed     float64
	Type      string
}

func NewKeyboardMover(container *elements.Element, speed float64) *KeyboardMover {
	return &KeyboardMover{
		container: container,
		Speed:     speed,
		Type:      "KeyboardMover",
	}
}

func (mover *KeyboardMover) OnDraw(screen *ebiten.Image) error {
	return nil
}

func (mover *KeyboardMover) OnUpdate() error {
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
