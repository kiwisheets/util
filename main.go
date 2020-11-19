package util

type ServerConfig struct {
	Version     string
	Environment string
	GraphQL     GqlConfig
	JWT         JWTConfig
	Hash        HashConfig
	Database    DatabaseConfig
	Redis       RedisConfig
}

type JWTConfig struct {
	Secret string
}

type HashConfig struct {
	Salt      string
	MinLength int
}

type DatabaseConfig struct {
	Host           string
	User           string
	Password       string
	Database       string
	Port           string
	MaxConnections int
}

type RedisConfig struct {
	Address string
}
