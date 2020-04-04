package playerControl

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
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
func (replic *Replicator) OnUpdate(world []*elements.Element) error {
	if replic.conn != nil {

		if replic.container.ID == connection.GetID() {
			bytes, _ := json.Marshal(replic.container)
			myMRP := mrp.NewMRP([]byte("REPLIC"), []byte(bytes), []byte(replic.container.ID))
			replic.conn.Write(myMRP.MRPToByte())
		}

	}
	return nil
}

func (replic *Replicator) OnCheck(elemC *elements.Element) error {
	if math.Abs(replic.container.XPos-elemC.XPos) >= 20 {
		fmt.Print("RubberBand")
		return errors.New("DeSync")
	}
	return nil
}

func (replic *Replicator) OnUpdateServer(world []*elements.Element) error {
	return nil
}
