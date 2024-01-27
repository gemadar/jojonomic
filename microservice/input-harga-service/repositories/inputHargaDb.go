package repositories

import (
	"encoding/json"
	"fmt"
	"input-harga-service/models"
	"os"

	"github.com/IBM/sarama"
)

func SendKafkaMessage(producer sarama.SyncProducer, h models.Harga) error {

	harga := models.Harga{
		AdminId:      h.AdminId,
		HargaTopup:   h.HargaTopup,
		HargaBuyback: h.HargaBuyback,
	}

	hargaJSON, err := json.Marshal(harga)
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: os.Getenv("KAFKA_TOPIC"),
		Key:   sarama.StringEncoder(h.AdminId),
		Value: sarama.StringEncoder(hargaJSON),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
