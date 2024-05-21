package config

import (
	"os"
	"strings"
)

type EnvironmentVariables struct {
	KafkaBrokers []string
	KafkaTopics  []string
	KafkaGroupID string
	ElasticHost  string
	ElasticPort  string
}

func GetEnvironmentVariables() EnvironmentVariables {
	brokers := os.Getenv("KAFKA_BROKERS")
	brokersSlice := strings.Split(brokers, ",")
	return EnvironmentVariables{
		KafkaBrokers: brokersSlice,
		KafkaTopics:  []string{"telegram-channel"},
		KafkaGroupID: "telegram-consumer-group",
		ElasticHost:  os.Getenv("ELASTIC_HOST"),
		ElasticPort:  os.Getenv("ELASTIC_PORT"),
	}
}
