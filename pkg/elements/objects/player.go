package objects

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
)

const (
	playerSpeed = 1
)

var playerTexture *ebiten.Image

func NewPlayer(conn net.Conn) *elements.Element {
	player := &elements.Element{}

	player.XPos = 0
	player.YPos = 0

	player.Active = true

	player.Type = "player"
	player.ID = "0"

	sr := render.NewSpriteRenderer(player, "destroyer.png", playerTexture)
	player.AddComponent(sr)

	playerTexture = sr.Tex

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	return player
}
