package weather

type storm struct {
	symbol string
}

func NewStorm() storm {
	return storm{symbol: "#"}
}

func (s storm) OnDrawServer() string {
	return s.symbol
}
