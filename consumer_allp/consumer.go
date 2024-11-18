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

	// Récupérer les partitions disponibles pour le topic
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		log.Fatalf("Erreur de récupération des partitions : %v", err)
	}
	fmt.Printf("Partitions disponibles : %v\n", partitions)

	// Consommer depuis toutes les partitions
	for _, partition := range partitions {
		// Consommer la partition en mode "OffsetOldest" pour récupérer tous les messages depuis le début
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			log.Printf("Erreur de consommation de la partition %d : %v", partition, err)
			continue
		}
		defer pc.Close()

		// Affichage des messages
		go func(partition int32, pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Printf("Partition=%d, Offset=%d, Key=%s, Value=%s\n",
					message.Partition, message.Offset, string(message.Key), string(message.Value))
			}
		}(partition, pc)
	}

	// Attente pour consommer les messages pendant 60 secondes
	// Vous pouvez ajuster ce temps en fonction de votre besoin
	fmt.Println("Consommation des messages en cours... Attendez quelques secondes.")
	select {}
}
