package terrain

type TerrainServer interface {
	OnDrawServer() string
}

type TerrainClient interface {
	OnDrawClient() string
}
