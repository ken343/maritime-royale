package objects

import (
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/elements/render"
)

func NewIsland(xpos float64, ypos float64, Name string) *elements.Element {
	Island := &elements.Element{}

	Island.XPos = xpos
	Island.YPos = ypos

	Island.Active = true

	Island.UniqueName = Name

	sr := render.NewSpriteRenderer(Island, "island.png")
	Island.AddComponent(sr)

	return Island
}
