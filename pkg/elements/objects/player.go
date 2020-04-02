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

	player.Type = "player"
	player.ID = "0"

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	clc := playerControl.NewClicker(player)
	player.AddComponent(clc)

	//trc := playerControl.NewTracker(player, 1, 600, 600)
	//player.AddComponent(trc)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	return player
}
