package playerControl

import (
	"encoding/json"
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/mrp"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type Replicator struct {
	container *elements.Element
	conn      net.Conn
	Type      string
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

//OnDraw is used to qualify SpriteRenderer as a component
func (replic *Replicator) OnDraw(screen *ebiten.Image) error {
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (replic *Replicator) OnUpdate() error {
	if replic.conn != nil {
		bytes, _ := json.Marshal(replic.container)
		myMRP := mrp.NewMRP([]byte("REPLIC"), []byte(bytes), []byte(replic.container.ID))
		replic.conn.Write(myMRP.MRPToByte())
	}
	return nil
}
