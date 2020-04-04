package objects

import (
	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/elements/secondOrder/render"
)

func NewWater(xpos float64, ypos float64, Name string) *elements.Element {
	water := &elements.Element{}

	water.XPos = xpos
	water.YPos = ypos

	water.Active = true

	water.UniqueName = Name

	sr := render.NewSpriteRenderer(water, "water.png")
	water.AddComponent(sr)

	return water
}
