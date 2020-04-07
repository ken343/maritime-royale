package playerControl

import (
	"encoding/json"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/gamestate"
	"github.com/ken343/maritime-royale/pkg/networking/connection"
	"github.com/ken343/maritime-royale/pkg/networking/mrp"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type Replicator struct {
	container *elements.Element
	conn      net.Conn
	Type      string
	count     int
}

func init() {
	var replic = new(Replicator)
	gamestate.MRPMAP["Replicator"] = replic
}

//NewReplicator creates a Replicator which is
//the component that handles all replication
//of an element onto a server.
func NewReplicator(container *elements.Element, conn net.Conn) *Replicator {

	return &Replicator{
		container: container,
		conn:      conn,
		Type:      "Replicator",
	}
}

func (replic *Replicator) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewReplicator(finalElem, conn)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (replic *Replicator) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (replic *Replicator) OnUpdate(xOffset float64, yOffset float64) error {

	if replic.container.ID == connection.GetID() && replic.conn != nil && !replic.container.Same {

		bytes, _ := json.Marshal(replic.container)
		myMRP := mrp.NewMRP([]byte("REPLIC"), []byte(bytes), []byte(replic.container.UniqueName))
		replic.conn.Write(myMRP.MRPToByte())

		replic.container.Same = true
	}

	return nil
}

func (replic *Replicator) OnCheck(elemC *elements.Element) error {
	return nil
}

func (replic *Replicator) OnUpdateServer() error {

	if replic.count == 100 && !replic.container.Same {

		go gamestate.UpdateElemToAll(replic.container)

		replic.count = 0
		replic.container.Same = true

	} else if replic.count == 100 {

		replic.count = 0
	}

	replic.count++

	return nil
}

func (replic *Replicator) OnMerge(compM elements.Component) error {
	return nil
}

func (replic *Replicator) SetContainer(container *elements.Element) error {
	replic.container = container
	return nil
}

func (replic *Replicator) MakeCopy() elements.Component {
	myComp := *replic
	return &myComp
}
