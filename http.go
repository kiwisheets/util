package util

type GqlConfig struct {
	APIPath           string
	PlaygroundPath    string
	PlaygroundAPIPath string
	PlaygroundEnabled bool
	Port              string
	ComplexityLimit   int
	Environment       string
	Redis             RedisConfig
}

type ClientConfig struct {
	BaseURL        string
	CfClientID     string
	CfClientSecret string
}
