package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"input-harga-storage-service/models"
	"input-harga-storage-service/repositories"
	"log"
	"os"

	"github.com/IBM/sarama"
)

type Consumer struct {
	store *repositories.HargaStorage
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		adminId := string(msg.Key)
		var harga models.Harga
		err := json.Unmarshal(msg.Value, &harga)
		if err != nil {
			log.Printf("failed to unmarshal: %v", err)
			continue
		}
		consumer.store.Add(adminId, harga)
		repositories.HandleInputHarga(harga)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func SetupConsumerGroup(ctx context.Context, store *repositories.HargaStorage) {
	consumerGroup, err := InitializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{os.Getenv("CONSUMER_TOPIC")}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func InitializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{os.Getenv("KAFKA_BROKERS")}, os.Getenv("CONSUMER_GROUP"), config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}
