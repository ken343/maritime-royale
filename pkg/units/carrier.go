package units

type carrier struct {
	symbol string
}

func NewCarrier() carrier {
	return carrier{symbol: "C"}
}

func (s carrier) Draw() {

}
