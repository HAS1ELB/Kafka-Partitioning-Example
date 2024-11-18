package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Configuration Kafka
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0

	// Connecter à Kafka
	brokers := []string{"localhost:9092"}
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		log.Fatalf("Erreur : %v", err)
	}
	defer admin.Close()

	// Définir les détails du topic
	topicName := "new-3partitioned-topic"
	topicDetail := sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 1,
	}

	// Créer le topic
	if err := admin.CreateTopic(topicName, &topicDetail, false); err != nil {
		log.Fatalf("Erreur de création du topic : %v", err)
	}
	log.Printf("Le topic '%s' a été créé avec succès.", topicName)
}
