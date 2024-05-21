package main

import (
	"fmt"
	"git.cydevcloud.com/crypto-trading/kafka-to-elastic/config"
	"git.cydevcloud.com/crypto-trading/kafka-to-elastic/kafka"
)

var env config.EnvironmentVariables

func init() {
	env = config.GetEnvironmentVariables()
}

func main() {
	kafkaClient := kafka.NewKafkaConsumer(env.KafkaBrokers, env.KafkaTopics, env.KafkaGroupID)
	readers := kafkaClient.Consume()
	for {
		for topic, reader := range readers {
			fmt.Println("topic: ", topic)
			fmt.Println("reader stats : ", reader.Stats())

		}
	}

}
