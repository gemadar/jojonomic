package handlers

import (
	"buyback-storage-service/models"
	"buyback-storage-service/repositories"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
)

type Consumer struct {
	store *repositories.TrxStorage
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		norek := string(msg.Key)
		var trx models.Transaksi

		err := json.Unmarshal(msg.Value, &trx)
		if err != nil {
			log.Printf("failed to unmarshal: %v", err)
			continue
		}
		consumer.store.Add(norek, trx)
		repositories.HandleInputTransaksi(trx)
		repositories.HandleUpdateSaldo(trx.Gram, trx.Norek)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func SetupConsumerGroup(ctx context.Context, store *repositories.TrxStorage) {
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
