package util

import "github.com/99designs/gqlgen/graphql"

type GqlConfig struct {
	APIPath           string
	PlaygroundPath    string
	PlaygroundAPIPath string
	PlaygroundEnabled bool
	Port              string
	ComplexityLimit   int
	Environment       string
	Cache             graphql.Cache
}

type ClientConfig struct {
	BaseURL        string
	CfClientID     string
	CfClientSecret string
}
