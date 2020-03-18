package weather

type WeatherServer interface {
	OnDrawServer() string
}

type WeatherClient interface {
	OnDrawClient() string
}
