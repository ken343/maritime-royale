package units

type jet struct {
	symbol string
}

func NewJet() jet {
	return jet{symbol: "J"}
}

func (s jet) OnDrawServer() string {
	return s.symbol
}
