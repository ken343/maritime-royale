package objects

import (
	"net"

	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/physics"

	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/playerControl"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/render"
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

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	coli := physics.NewCollider(player)
	player.AddComponent(coli)

	rot := render.NewRotator(player)
	player.AddComponent(rot)

	return player
}
