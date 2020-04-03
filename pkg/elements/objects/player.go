package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
)

const (
	playerSpeed = 1
)

func NewPlayer(conn net.Conn) *elements.Element {
	player := &elements.Element{}

	player.XPos = 0
	player.YPos = 0

	player.Active = true

	player.UniqueName = "player"
	player.ID = "-1"

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	//tkr := playerControl.NewTracker(player, 1, 600, 600)
	//player.AddComponent(tkr)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	rot := render.NewRotator(player)
	player.AddComponent(rot)

	return player
}
