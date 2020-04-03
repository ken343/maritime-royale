package secondOrder

import (
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/physics"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/playerControl"
	"github.com/jtheiss19/project-undying/pkg/elements/secondOrder/render"
)

func Init() {
	render.Init()
	playerControl.Init()
	physics.Init()
}
