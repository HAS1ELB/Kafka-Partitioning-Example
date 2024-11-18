
# Kafka Partitioning Example Project

This project demonstrates how to implement Kafka producers and consumers in Go with partitioning, using the `sarama` library. The goal is to show how to create topics with multiple partitions, produce messages to these partitions, and consume them.

## Project Structure

```
kafka-partitioning-example/
├── consumer/
│   └── consumer.go         # Code for consuming messages from Kafka partitions
├── producer/
│   └── producer.go         # Code for producing messages to Kafka
├── create_topic.go         # Code to create Kafka topic with multiple partitions
└── README.md               # This file
```

## Requirements

- Kafka server running locally or remotely (configured in the Go code)
- Go (v1.16 or above)
- Sarama library for Go: `github.com/IBM/sarama`

## Setup and Installation

1. **Install Kafka** (if not already installed):

   - You can install and run Kafka locally or use a remote Kafka service.
   - Make sure Kafka is running on the default port (`localhost:9092`).

2. **Install Go and Sarama Library**:

   Ensure you have Go installed and then run the following command to install the Sarama library:
   ```bash
   go get github.com/IBM/sarama
   ```

3. **Create the Topic**:

   Run `create_topic.go` to create a topic with multiple partitions:

   ```bash
   go run create_topic.go
   ```

   This will create a topic called `new-3partitioned-topic` with 3 partitions.

4. **Run the Producer**:

   The producer sends messages to the Kafka topic, distributing them across partitions. To run the producer, use:

   ```bash
   cd producer
   go run producer.go
   ```

5. **Run the Consumer**:

   The consumer reads messages from all partitions of the topic. To run the consumer, use:

   ```bash
   cd consumer
   go run consumer.go
   ```

   The consumer will print the messages from each partition with their offset, key, and value.

## How It Works

### Producer (`producer.go`):
- The producer connects to Kafka and sends messages to a topic with multiple partitions (`new-3partitioned-topic`).
- The messages are distributed across partitions based on a key provided in the producer code.

### Consumer (`consumer.go`):
- The consumer connects to Kafka and consumes messages from all partitions of the topic.
- The consumer prints out the partition, offset, key, and value for each consumed message.

### Topic Creation (`create_topic.go`):
- This code creates a new Kafka topic with 3 partitions using the Sarama library's admin client.

## Kafka Topics

- **Topic**: `new-3partitioned-topic`
- **Partitions**: 3

## Troubleshooting

- Ensure that your Kafka broker is running and accessible on `localhost:9092`.
- If you encounter errors like "no brokers available", double-check the Kafka server settings.
- Verify that the topic exists before running the producer or consumer.

## Contributing

Feel free to contribute by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License.

