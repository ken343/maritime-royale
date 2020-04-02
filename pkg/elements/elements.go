package elements

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

type Component interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image) error
}

type Element struct {
	XPos       float64
	YPos       float64
	Rotation   float64
	Active     bool
	Type       string
	ID         string
	Components []Component
}

func (elem *Element) Draw(screen *ebiten.Image) error {
	for _, comp := range elem.Components {
		err := comp.OnDraw(screen)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) Update() error {
	for _, comp := range elem.Components {
		err := comp.OnUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) AddComponent(new Component) {
	for _, existing := range elem.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	elem.Components = append(elem.Components, new)
}

func (elem *Element) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	return nil
}
