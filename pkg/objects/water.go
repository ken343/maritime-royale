package objects

type water struct {
	symbol string
}

func NewWater() water {
	return water{symbol: "w"}
}

func (s water) OnDraw() string {
	return s.symbol
}
