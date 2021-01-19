package util

import (
	"io/ioutil"
	"log"

	"github.com/cheshir/go-mq"
	"github.com/joho/godotenv"
	"github.com/maxtroughear/goenv"
	"gopkg.in/yaml.v2"
)

func NewMQ() mq.MQ {
	mqYaml, err := ioutil.ReadFile("mq.yaml")
	if err != nil {
		log.Fatalf("unable to open mq.yaml")
	}

	var mqConfig mq.Config

	err = yaml.Unmarshal([]byte(mqYaml), &mqConfig)
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	godotenv.Load()
	mqConfig.DSN = goenv.MustGetSecretFromEnv("RABBITMQ_DSN")

	messageQ, err := mq.New(mqConfig)
	if err != nil {
		log.Fatalf("failed to init RabbitMQ: %s", err)
	}

	return messageQ
}
