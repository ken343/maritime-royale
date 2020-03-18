package units

type UnitServer interface {
	OnDrawServer() string
}

type UnitClient interface {
	OnDrawClient() string
}
