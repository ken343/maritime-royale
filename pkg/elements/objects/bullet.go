package objects

import (
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements/thirdOrder/explode"

	"github.com/jtheiss19/project-undying/pkg/gamestate"

	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/advancePos"
	"github.com/jtheiss19/project-undying/pkg/elements/firstOrder/attack"

	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
)

func init() {
	gamestate.ObjectMap["Bullet"] = NewBullet(nil, 0, 0)
}

func NewBullet(conn net.Conn, DestX, DestY float64) *elements.Element {
	bullet := &elements.Element{}

	bullet.XPos = 0
	bullet.YPos = 0

	bullet.Active = true

	bullet.UniqueName = "MyBullet"

	//--FIRST ORDER--------------------------------------------//

	aPos := advancePos.NewAdvancePosition(bullet, 5)
	bullet.AddComponent(aPos)

	dam := attack.NewDamage(bullet)
	bullet.AddComponent(dam)

	//--SECOND ORDER-------------------------------------------//

	sr := render.NewSpriteRenderer(bullet, "carrier.png")
	bullet.AddComponent(sr)

	rot := render.NewRotator(bullet)
	bullet.AddComponent(rot)

	coli := physics.NewCollider(bullet)
	bullet.AddComponent(coli)

	mov := playerControl.NewMoveTo(bullet, -400, -400)
	bullet.AddComponent(mov)

	//--THIRD ORDER--------------------------------------------//

	explo := explode.NewExplosion(bullet)
	bullet.AddComponent(explo)

	replic := playerControl.NewReplicator(bullet, conn)
	bullet.AddComponent(replic)

	return bullet
}
