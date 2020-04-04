package physics

import (
	"math"
	"net"

	"github.com/jtheiss19/project-undying/pkg/gamestate"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

//Collider is the component that handles all
//keyboard movement
type Collider struct {
	container   *elements.Element
	Type        string
	Radius      float64
	PrevX       float64
	PrevY       float64
	HasCollided bool
}

func init() {
	var coli = new(Collider)
	gamestate.MRPMAP["Collider"] = coli
}

//NewCollider creates a KeyboardMover which is
//the component that handles all keyboard movement
func NewCollider(container *elements.Element) *Collider {
	return &Collider{
		container:   container,
		Type:        "Collider",
		Radius:      50,
		HasCollided: false,
	}
}

func (coli *Collider) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewCollider(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (coli *Collider) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate scans the state of the keyboard and prefroms
//actions based on said state.
func (coli *Collider) OnUpdate(world []*elements.Element) error {

	for _, elem := range world {
		if elem.GetComponent(coli) != nil && elem.ID != coli.container.ID {
			elemComp := elem.GetComponent(coli)
			if isCollison(elemComp.(*Collider), coli) {
				coli.container.XPos = coli.PrevX
				coli.container.YPos = coli.PrevY
				coli.HasCollided = true
			}
		}
	}

	coli.PrevX = coli.container.XPos
	coli.PrevY = coli.container.YPos

	return nil
}

func (coli *Collider) OnCheck(elemC *elements.Element) error {
	return nil
}

func (coli *Collider) OnUpdateServer(world []*elements.Element) error {
	gamestate.GetWorld()
	for _, elem := range world {
		if elem.GetComponent(coli) != nil && elem.ID != coli.container.ID {
			elemComp := elem.GetComponent(coli)
			if isCollison(elemComp.(*Collider), coli) {
				coli.container.XPos = coli.PrevX
				coli.container.YPos = coli.PrevY
				coli.HasCollided = true
			}
		}
	}

	coli.PrevX = coli.container.XPos
	coli.PrevY = coli.container.YPos

	return nil
}

func isCollison(coli1, coli2 *Collider) bool {

	elem1 := coli1.container
	elem2 := coli2.container

	xDiff := elem1.XPos - elem2.XPos
	yDiff := elem1.YPos - elem2.YPos

	totalDiff := math.Hypot(xDiff, yDiff)

	if totalDiff <= coli1.Radius || totalDiff <= coli2.Radius {
		return true
	}
	return false
}
