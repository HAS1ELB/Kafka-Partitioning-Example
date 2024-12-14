# Kafka Partitioning Example

Ce projet démontre l'utilisation de Kafka avec des partitions pour un traitement efficace des messages. Il comprend des exemples de producteurs, de consommateurs et d'administration Kafka écrits en Go.

## Structure du projet

- **admin/** : Contient le script pour créer un topic Kafka avec plusieurs partitions.
  - `create_topic.go`: Script pour créer un topic nommé "new-3partitioned-topic" avec 3 partitions.
  - `go.mod` et `go.sum`: Fichiers de gestion des dépendances.

- **producer/** : Contient le producteur Kafka qui envoie des messages aux partitions.
  - `producer.go`: Script d'envoi de messages avec des clés spécifiques pour répartir les messages entre les partitions.
  - `go.mod` et `go.sum`: Fichiers de gestion des dépendances.

- **consumer_allp/** : Contient un consommateur Kafka qui lit les messages de toutes les partitions.
  - `consumer.go`: Script pour consommer tous les messages depuis toutes les partitions.
  - `go.mod` et `go.sum`: Fichiers de gestion des dépendances.

- **consumer_onlyp2/** : Contient un consommateur Kafka qui lit uniquement les messages d'une partition spécifique.
  - `consumer.go`: Script pour consommer les messages de la partition 2.
  - `go.mod` et `go.sum`: Fichiers de gestion des dépendances.

- **kafka_partitions_example.pdf** : Documentation du projet.

## Prérequis

1. [Kafka](https://kafka.apache.org/) doit être installé et en cours d'exécution.
   - Un broker Kafka doit être accessible sur `localhost:9092`.
2. [Go](https://golang.org/dl/) doit être installé (version 1.23 ou supérieure).

## Installation

1. Clonez le dépôt :
   ```bash
   git clone <url-du-repo>
   cd kafka-partitioning-example
   ```
2. Installez les dépendances pour chaque sous-dossier (par exemple, `admin`, `producer`, etc.) :
   ```bash
   cd admin
   go mod tidy
   ```

## Utilisation

### 1. Création du topic Kafka

Exécutez le script dans le dossier `admin` :
```bash
cd admin
go run create_topic.go
```
Cela crée un topic nommé `new-3partitioned-topic` avec 3 partitions.

### 2. Production des messages

Exécutez le producteur dans le dossier `producer` :
```bash
cd producer
go run producer.go
```
Cela envoie des messages aux partitions en fonction des clés définies.

### 3. Consommation des messages

#### Consommation de toutes les partitions

Exécutez le consommateur dans le dossier `consumer_allp` :
```bash
cd consumer_allp
go run consumer.go
```
Cela lit tous les messages de toutes les partitions.

#### Consommation de la partition 2 uniquement

Exécutez le consommateur dans le dossier `consumer_onlyp2` :
```bash
cd consumer_onlyp2
go run consumer.go
```
Cela lit les messages uniquement de la partition 2.

## Contributions

Les contributions sont les bienvenues ! Veuillez soumettre une PR ou signaler des problèmes dans le suivi des issues.

## Licence

Ce projet est sous licence MIT. Consultez le fichier `LICENSE` pour plus d'informations.
