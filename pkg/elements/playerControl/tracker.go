package playerControl

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

//Tracker is the component that handles all
//keyboard movement
type Tracker struct {
	container *elements.Element
	Speed     float64
	Type      string
	DestX     float64
	DestY     float64
}

//NewTracker creates a KeyboardMover which is
//the component that handles all keyboard movement
func NewTracker(container *elements.Element, speed float64, destX float64, destY float64) *Tracker {
	return &Tracker{
		container: container,
		Speed:     speed,
		Type:      "Tracker",
		DestX:     destX,
		DestY:     destY,
	}
}

//OnDraw is used to qualify SpriteRenderer as a component
func (trc *Tracker) OnDraw(screen *ebiten.Image) error {
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (trc *Tracker) OnUpdate() error {
	if (trc.container.XPos - trc.DestX) > 0 {
		trc.container.XPos -= trc.Speed
	} else {
		trc.container.XPos += trc.Speed
	}

	if (trc.container.YPos - trc.DestY) > 0 {
		trc.container.YPos -= trc.Speed

	} else {
		trc.container.YPos += trc.Speed

	}

	return nil
}

func (trc *Tracker) OnCheck(elemC *elements.Element) error {
	return nil
}

func (trc *Tracker) GetContainer() *elements.Element {
	return trc.container
}

func (trc *Tracker) SetContainer(elem *elements.Element) {
	trc.container = elem
}
