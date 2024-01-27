package repositories

import (
	"encoding/json"
	"fmt"
	"os"
	"topup-service/models"

	"github.com/IBM/sarama"
)

func SendKafkaMessage(producer sarama.SyncProducer, t models.Topup) error {

	trx := models.Topup{
		Norek: t.Norek,
		Gram:  t.Gram,
		Harga: t.Harga,
	}

	trxJSON, err := json.Marshal(trx)
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Key:   sarama.StringEncoder(t.Norek),
		Value: sarama.StringEncoder(trxJSON),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
