package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Configuration du producteur Kafka
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner // Utilisation du partitionneur round-robin
	config.Producer.Return.Successes = true

	// Connexion au broker Kafka
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Erreur de connexion au producteur Kafka : %v", err)
	}
	defer producer.Close()

	// Le nom du topic
	topic := "new-3partitioned-topic"

	// Produire des messages et les répartir sur les partitions
	for i := 0; i < 30; i++ {
		// Créer un message à envoyer
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}

		// Envoi du message à la partition (round-robin)
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Erreur d'envoi du message : %v", err)
		} else {
			log.Printf("Message envoyé : Partition=%d, Offset=%d", partition, offset)
		}

		// Attendre un peu entre chaque message
		time.Sleep(1 * time.Second)
	}
}
