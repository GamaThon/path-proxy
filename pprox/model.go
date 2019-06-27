package pprox

type Config struct {
	Listen string  `json:"listen"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Prefix string `json:"prefix"`
	Target string `json:"target"`
}
