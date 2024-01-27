package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"topup-storage-service/models"
	"topup-storage-service/repositories"

	"github.com/IBM/sarama"
)

type Consumer struct {
	store *repositories.TopStorage
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		norek := string(msg.Key)
		var top models.Topup
		err := json.Unmarshal(msg.Value, &top)
		if err != nil {
			log.Printf("failed to unmarshal: %v", err)
			continue
		}
		consumer.store.Add(norek, top)
		repositories.HandleInputTransaksi(top)
		repositories.HandleUpdateSaldo(top.Gram, top.Norek)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func SetupConsumerGroup(ctx context.Context, store *repositories.TopStorage) {
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
