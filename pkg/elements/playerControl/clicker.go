package playerControl

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

//Clicker is the component that handles all
//keyboard movement
type Clicker struct {
	container *elements.Element

	Type string
}

//NewClicker creates a KeyboardMover which is
//the component that handles all keyboard movement
func NewClicker(container *elements.Element) *Clicker {
	return &Clicker{
		container: container,
		Type:      "Clicker",
	}
}

//OnDraw is used to qualify SpriteRenderer as a component
func (clc *Clicker) OnDraw(screen *ebiten.Image) error {
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (clc *Clicker) OnUpdate() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		trc := NewTracker(clc.container, 1, 600, 600)
		clc.container.AddComponent(trc)
	}
	return nil
}

func (clc *Clicker) OnCheck(elemC *elements.Element) error {
	return nil
}
