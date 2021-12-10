package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/educleite/imersao5-gateway/adapter/broker/kafka"
	"github.com/educleite/imersao5-gateway/adapter/factory"
	"github.com/educleite/imersao5-gateway/adapter/presenter/transaction"
	"github.com/educleite/imersao5-gateway/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// fmt.Println("Ola")

	// 1 - db
	// 2 - repository
	// 3 - configMapProducer
	// 4 - producer
	// 5 - configMapConsumer
	// 6 - consumer
	// 7 - topic
	// 8 - usecase

	// 1 - db
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)

	}
	// 2 - repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// 3 - configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}

	// 4 - producer
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	// 5 - configMapConsumer
	var msgChan = make(chan *ckafka.Message) // channel

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	// 7 - topic
	topics := []string{"transactions"}

	// 6 - consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)

	// 8 - usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {

		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}

}
