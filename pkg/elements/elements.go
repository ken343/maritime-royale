package elements

import (
	"net"
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

//Component is an interface that describes what counts
//as a component. If something can be drawn by having an
//OnDraw() function and can be updated with an OnUpdate()
//function then it counts as a component. These functions
//may be empty.
type Component interface {
	OnUpdate(xOffset float64, yOffset float64) error
	OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error
	OnCheck(*Element) error
	OnMerge(Component) error
	OnUpdateServer() error
	MRP(finalElem *Element, conn net.Conn)
	SetContainer(*Element) error
	MakeCopy() Component
}

//Element is the basic atomic structure for all objects.
//Functionality and Excess data is provided to by Components.
//Components extend an Elements functionality.
type Element struct {
	XPos       float64
	YPos       float64
	Rotation   float64
	Active     bool
	UniqueName string
	ID         string
	Components []Component
	Same       bool
}

//Draw loops through all components within the element
//and runs the OnDraw() function for each one.
//Error is returned through the first error from a
//components OnDraw() function.
func (elem *Element) Draw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnDraw(screen, xOffset, yOffset)
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
func (elem *Element) Update(xOffset float64, yOffset float64) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnUpdate(xOffset, yOffset)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (elem *Element) UpdateServer() error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnUpdateServer()
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

func (elem *Element) Merge(elemM *Element) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.OnMerge(elemM.GetComponent(comp))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (elem *Element) SetContainer(container *Element) error {
	for _, comp := range elem.Components {
		if comp != nil {
			err := comp.SetContainer(container)
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

func (elem *Element) AddComponentPostInit(new Component) {
	elem.RemoveComponentType(new)
	potentialReplic := elem.Components[len(elem.Components)-1]
	temp := elem.Components[:len(elem.Components)-1]
	elem.Components = append(temp, new)
	elem.Components = append(elem.Components, potentialReplic)
}

func (elem *Element) RemoveComponentType(badComp Component) {
	for k, existing := range elem.Components {
		if reflect.TypeOf(badComp) == reflect.TypeOf(existing) {
			copy(elem.Components[k:], elem.Components[k+1:])
			elem.Components[len(elem.Components)-1] = nil
			elem.Components = elem.Components[:len(elem.Components)-1]
		}
	}
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

func (elem *Element) MakeCopy() *Element {

	myBlankElem := *elem
	myBlankElem.Components = *new([]Component)

	for _, myComp := range elem.Components {
		if myComp != nil {
			newCopy := myComp.MakeCopy()
			myBlankElem.AddComponent(newCopy)
		}
	}

	myBlankElem.SetContainer(&myBlankElem)

	return &myBlankElem
}
