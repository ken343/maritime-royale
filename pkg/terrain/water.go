package terrain

type water struct {
	symbol string
}

func NewWater() water {
	return water{symbol: "w"}
}
func (s water) Draw() {

}
