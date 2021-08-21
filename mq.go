package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

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

	mqConfig.DSN = goenv.CanGet("RABBITMQ_DSN", buildDSN())

	messageQ, err := mq.New(mqConfig)
	if err != nil {
		log.Fatalf("failed to init RabbitMQ: %s", err)
	}

	return messageQ
}

func buildDSN() string {
	protocol := goenv.CanGet("RABBITMQ_PROTOCOL", "amqp")
	host := goenv.MustGet("RABBITMQ_HOST")
	port := goenv.CanGetInt("RABBITMQ_PORT", 5672)
	username := goenv.MustGet("RABBITMQ_USERNAME")
	password := goenv.MustGet("RABBITMQ_PASSWORD")
	vhost := goenv.CanGet("RABBITMQ_VHOST", "/")

	return fmt.Sprintf("%s://%s:%s@%s:%s%s", protocol, host, strconv.Itoa(port), username, password, vhost)
}
