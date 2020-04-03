package render

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//Rotator is the component that handles all
//rendering of sprites onto the screen
type Rotator struct {
	container *elements.Element
	XPrev     float64
	YPrev     float64
	Type      string
}

//NewRotator creates a SpriteRenderer which
//is the component that handles all rendering of
//sprites onto the screen
func NewRotator(container *elements.Element) *Rotator {

	return &Rotator{
		container: container,
		XPrev:     0,
		YPrev:     0,
		Type:      "Rotator",
	}
}

//OnDraw Draws the stored texture file onto the screen
func (rot *Rotator) OnDraw(screen *ebiten.Image) error {
	return nil
}

//OnUpdate is used to qualify SpriteRenderer as a component
func (rot *Rotator) OnUpdate() error {
	if rot.container.ID != connection.GetID() {
		return nil
	}
	if rot.container.YPos == rot.YPrev && rot.container.XPos == rot.XPrev {
	} else {
		rot.container.Rotation = math.Atan2((rot.container.YPos - rot.YPrev), (rot.container.XPos - rot.XPrev))
	}
	rot.XPrev = rot.container.XPos
	rot.YPrev = rot.container.YPos
	return nil
}

func (rot *Rotator) OnCheck(elemC *elements.Element) error {
	return nil
}
