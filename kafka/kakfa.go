package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type KafkaConsumer struct {
	Readers map[string]*kafka.Reader
	Context context.Context
	Cancel  context.CancelFunc
	Wg      sync.WaitGroup
}

func NewKafkaConsumer(brokers []string, topics []string, groupID string) *KafkaConsumer {
	readers := make(map[string]*kafka.Reader)
	ctx, cancel := context.WithCancel(context.Background())

	for _, topic := range topics {
		readers[topic] = kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		})
	}

	return &KafkaConsumer{
		Readers: readers,
		Context: ctx,
		Cancel:  cancel,
	}
}

func (kc *KafkaConsumer) Consume() map[string]*kafka.Reader {
	return kc.Readers
}

func (kc *KafkaConsumer) Stop() {
	kc.Cancel()
	kc.Wg.Wait()
	for _, reader := range kc.Readers {
		reader.Close()
	}
}
