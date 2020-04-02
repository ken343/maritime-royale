package render

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/jtheiss19/project-undying/pkg/elements"
)

type SpriteRenderer struct {
	container *elements.Element
	Tex       *ebiten.Image

	Type          string
	Width, Height float64
}

func NewSpriteRenderer(container *elements.Element, filename string, masterTex *ebiten.Image) *SpriteRenderer {
	var tex *ebiten.Image
	if masterTex == nil {
		tex = textureFromPNG(filename)
	} else {
		tex = masterTex
	}
	width, height := tex.Size()

	return &SpriteRenderer{
		container: container,
		Tex:       tex,
		Width:     float64(width),
		Height:    float64(height),
		Type:      "SpriteRenderer",
	}
}

func (sr *SpriteRenderer) OnDraw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(sr.Width)/2, -float64(sr.Height)/2)
	op.GeoM.Rotate(1 * math.Pi * sr.container.Rotation / 360)
	op.GeoM.Translate(float64(sr.Width)/2, float64(sr.Height)/2)
	op.GeoM.Translate(sr.container.XPos, sr.container.YPos)

	//op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(sr.Tex, op)
	return nil
}

func (sr *SpriteRenderer) OnUpdate() error {
	return nil
}

func textureFromPNG(filename string) *ebiten.Image {
	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/"+filename, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	masterTexture, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	masterTexture.DrawImage(origEbitenImage, op)
	return masterTexture
}
