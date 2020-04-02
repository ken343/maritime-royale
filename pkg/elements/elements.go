package elements

import (
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

//Component is an interface that describes what counts
//as a component. If something can be drawn by having an
//OnDraw() function and can be updated with an OnUpdate()
//function then it counts as a component. These functions
//may be empty.
type Component interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image) error
	OnCheck(*Element) error
}

//Element is the basic atomic structure for all objects.
//Functionality and Excess data is provided to by Components.
//Components extend an Elements functionality.
type Element struct {
	XPos       float64
	YPos       float64
	Rotation   float64
	Active     bool
	Type       string
	ID         string
	Components []Component
}

//Draw loops through all components within the element
//and runs the OnDraw() function for each one.
//Error is returned through the first error from a
//components OnDraw() function.
func (elem *Element) Draw(screen *ebiten.Image) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnDraw(screen)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//Update loops through all components within the element
//and runs the OnUpdate() function for each one.
//Error is returned through the first error from a
//components OnUpdate() function.
func (elem *Element) Update() error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnUpdate()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (elem *Element) Check(elemC *Element) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnCheck(elemC)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//AddComponent adds a component to the component
//slice stored within the element. Panics if the
//component already exists within the slice.
func (elem *Element) AddComponent(new Component) {
	for _, existing := range elem.Components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			//panic(fmt.Sprintf(
			//"attempt to add new component with existing type %v",
			//reflect.TypeOf(new)))
			return
		}
	}
	elem.Components = append(elem.Components, new)
}

//GetComponent gets a component in the component
//slice stored within the element by using the
//component type of the withType component. Returns
//nil if the component does not exist
func (elem *Element) GetComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.Components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	return nil
}
