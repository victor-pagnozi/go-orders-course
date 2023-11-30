package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/victor-pagnozi/go-intensivo-fcl/internal/infra/database"
	"github.com/victor-pagnozi/go-intensivo-fcl/internal/usecase"
	"github.com/victor-pagnozi/go-intensivo-fcl/pkg/rabbitmq"
)

type Car struct {
	Model string
	Color string
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}

	defer db.Close() // Espera tudo rodar, para dps fechar a conexão

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel)

	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}

		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Message processed and save on the database", output)
	}
}
