package terrain

type water struct {
	Type string
}

func NewWater() water {
	return water{Type: "water"}
}
func (s water) Draw() {

}
