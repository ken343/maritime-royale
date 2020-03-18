package units

type submarine struct {
	symbol string
}

func NewSubmarine() submarine {
	return submarine{symbol: "S"}
}

func (s submarine) OnDrawServer() string {
	return s.symbol
}
