package repositories

import (
	"buyback-service/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/sarama"
)

func SendKafkaMessage(producer sarama.SyncProducer, t models.Transaksi) error {
	trx := models.Transaksi{
		Norek:        t.Norek,
		Gram:         t.Gram,
		HargaBuyback: t.HargaBuyback,
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
