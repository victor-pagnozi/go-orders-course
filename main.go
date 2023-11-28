package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/victor-pagnozi/go-intensivo-fcl/internal/infra/database"
	"github.com/victor-pagnozi/go-intensivo-fcl/internal/usecase"
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

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	// println(output.FinalPrice())
	fmt.Println(output)
}
