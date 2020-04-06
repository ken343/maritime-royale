package playerControl

import (
	"math"
	"net"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type Shooter struct {
	container *elements.Element
	Type      string
	HasFired  bool
	DestX     float64
	DestY     float64
	cooldown  int
}

func init() {
	var shoot = new(Shooter)
	gamestate.MRPMAP["Shooter"] = shoot
}

//NewReplicator creates a Replicator which is
//the component that handles all replication
//of an element onto a server.
func NewShooter(container *elements.Element) *Shooter {
	return &Shooter{
		container: container,
		Type:      "Shooter",
		cooldown:  0,
		HasFired:  false,
	}
}

func (shoot *Shooter) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewShooter(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (shoot *Shooter) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (shoot *Shooter) OnUpdate(xOffset float64, yOffset float64) error {
	if shoot.HasFired == true || shoot.container.ID != connection.GetID() {
		return nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shoot.HasFired = true
		w, h := ebiten.CursorPosition()
		shoot.DestX = float64(w) - xOffset
		shoot.DestY = float64(h) - yOffset
	}
	return nil
}

func (shoot *Shooter) OnCheck(elemC *elements.Element) error {
	return nil
}

func (shoot *Shooter) OnMerge(compM elements.Component) error {
	compM.(*Shooter).HasFired = shoot.HasFired
	compM.(*Shooter).DestX = shoot.DestX
	compM.(*Shooter).DestY = shoot.DestY
	return nil
}

var count int

func (shoot *Shooter) OnUpdateServer() error {

	if shoot.HasFired {
		if shoot.cooldown == 0 {
			count++
			shoot.cooldown = 0

			myBullet := gamestate.GetObject("Bullet")
			myBullet.UniqueName = "BULLET" + strconv.Itoa(count)

			mov := NewMoveTo(myBullet, shoot.DestX, shoot.DestY)
			myBullet.AddComponentPostInit(mov)

			rot := math.Atan2(shoot.DestY-shoot.container.YPos, shoot.DestX-shoot.container.XPos)
			uY, uX := math.Sincos(rot)

			myBullet.Rotation = rot
			myBullet.XPos = shoot.container.XPos + uX*70
			myBullet.YPos = shoot.container.YPos + uY*70

			gamestate.AddUnitToWorld(myBullet)
			gamestate.PushChunks()

			shoot.container.Same = false

		}
		shoot.HasFired = false
	}

	if shoot.cooldown > 0 {
		shoot.cooldown -= 1
	}

	return nil
}

func (shoot *Shooter) SetContainer(container *elements.Element) error {
	shoot.container = container
	return nil
}

func (shoot *Shooter) MakeCopy() elements.Component {
	myComp := *shoot
	return &myComp
}
