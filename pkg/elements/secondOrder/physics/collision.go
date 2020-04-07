package physics

import (
	"math"
	"net"

	"github.com/ken343/maritime-royale/pkg/elements/firstOrder/advancePos"
	"github.com/ken343/maritime-royale/pkg/gamestate"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
)

//Collider is the component that handles all
//keyboard movement
type Collider struct {
	container   *elements.Element
	posData     elements.Component
	Type        string
	Radius      float64
	HasCollided bool
	objectsHit  []*elements.Element
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
		posData:     container.GetComponent(new(advancePos.AdvancePosition)),
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
func (coli *Collider) OnUpdate(xOffset float64, yOffset float64) error {
	return nil
}

func (coli *Collider) OnCheck(elemC *elements.Element) error {
	return nil
}

func (coli *Collider) OnUpdateServer() error {
	if coli.posData == nil {
		return nil
	}

	coli.HasCollided = false
	coli.objectsHit = []*elements.Element{}

	for _, elem := range gamestate.GetEntireWorld() {
		if elem.GetComponent(coli) != nil && elem.UniqueName != coli.container.UniqueName {
			elemComp := elem.GetComponent(coli)
			if isCollison(elemComp.(*Collider), coli) {
				coli.container.XPos = coli.posData.(*advancePos.AdvancePosition).PrevX
				coli.container.YPos = coli.posData.(*advancePos.AdvancePosition).PrevY
				coli.HasCollided = true
				coli.objectsHit = append(coli.objectsHit, elem)
			}
		}
	}

	return nil
}

func (coli *Collider) OnMerge(compM elements.Component) error {
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

func (coli *Collider) SetContainer(container *elements.Element) error {
	coli.container = container
	coli.posData = container.GetComponent(new(advancePos.AdvancePosition))
	return nil
}

func (coli *Collider) MakeCopy() elements.Component {
	myComp := *coli
	return &myComp
}

func (coli *Collider) GetObjectsHit() []*elements.Element {
	return coli.objectsHit
}
