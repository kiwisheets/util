package util

import "crypto/ecdsa"

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
	SSLMode        string
	SSLCAPath      string
	Options        string
}
