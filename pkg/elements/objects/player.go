package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/health"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
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

	//--FIRST ORDER--------------------------------------------//

	aPos := advancePos.NewAdvancePosition(player, playerSpeed)
	player.AddComponent(aPos)

	hp := health.NewHealth(player, 100)
	player.AddComponent(hp)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(player, "destroyer.png")
	player.AddComponent(sr)

	shot := playerControl.NewShooter(player)
	player.AddComponent(shot)

	mover := playerControl.NewKeyboardMover(player, playerSpeed)
	player.AddComponent(mover)

	coli := physics.NewCollider(player)
	player.AddComponent(coli)

	rot := render.NewRotator(player)
	player.AddComponent(rot)

	//--THIRD ORDER--------------------------------------------//

	//--NETWORKING---------------------------------------------//

	replic := playerControl.NewReplicator(player, conn)
	player.AddComponent(replic)

	return player
}
