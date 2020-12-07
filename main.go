package util

import "crypto/ecdsa"

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
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
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
