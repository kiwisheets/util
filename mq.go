package util

import (
	"fmt"
	"io/ioutil"

	"github.com/cheshir/go-mq"
	"github.com/joho/godotenv"
	"github.com/maxtroughear/goenv"
	"gopkg.in/yaml.v2"
)

func NewMQ() mq.MQ {
	mqYaml, err := ioutil.ReadFile("mq.yaml")
	if err != nil {
		panic(fmt.Errorf("unable to open mq.yaml"))
	}

	var mqConfig mq.Config

	err = yaml.Unmarshal([]byte(mqYaml), &mqConfig)
	if err != nil {
		panic(fmt.Errorf("failed to read config: %s", err))
	}

	godotenv.Load()
	mqConfig.DSN = goenv.MustGetSecretFromEnv("RABBITMQ_DSN")

	messageQ, err := mq.New(mqConfig)
	if err != nil {
		panic(fmt.Errorf("failed to init RabbitMQ: %s", err))
	}

	return messageQ
}
