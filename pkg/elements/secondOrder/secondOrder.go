package secondOrder

import (
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/physics"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/playerControl"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/render"
)

func Init() {
	render.Init()
	playerControl.Init()
	physics.Init()
}
