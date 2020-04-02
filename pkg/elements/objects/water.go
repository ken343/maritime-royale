package objects

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
)

var waterTexture *ebiten.Image

func NewWater(xpos float64, ypos float64, ID string) *elements.Element {
	water := &elements.Element{}

	water.XPos = xpos
	water.YPos = ypos

	water.Active = true

	water.Type = "water"
	water.ID = ID

	sr := render.NewSpriteRenderer(water, "water.png", waterTexture)
	water.AddComponent(sr)

	waterTexture = sr.Tex

	return water
}
