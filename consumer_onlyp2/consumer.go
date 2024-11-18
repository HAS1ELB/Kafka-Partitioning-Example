package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Configuration du consommateur Kafka
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Connexion au broker Kafka
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Erreur de connexion au consommateur Kafka : %v", err)
	}
	defer consumer.Close()

	// Le nom du topic
	topic := "new-3partitioned-topic"

	// Consommer depuis la partition 2 uniquement
	partition := int32(2)

	// Consommer la partition en mode "OffsetOldest" pour récupérer tous les messages depuis le début
	pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Erreur de consommation de la partition %d : %v", partition, err)
	}
	defer pc.Close()

	// Affichage des messages
	fmt.Printf("Consommation des messages depuis la partition %d...\n", partition)
	for message := range pc.Messages() {
		fmt.Printf("Partition=%d, Offset=%d, Key=%s, Value=%s\n",
			message.Partition, message.Offset, string(message.Key), string(message.Value))
	}

	// Attente pour consommer les messages
	select {}
}
