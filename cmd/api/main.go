package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victor-pagnozi/go-intensivo-fcl/internal/entity"
)

func main() {
	fmt.Println("Running on port 8888")
	// vanilla
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", r)

	// framework echo
	e := echo.New()
	e.GET("/order", OrderHandler)
	e.Logger.Fatal(e.Start(":8888"))
}

// framework
func OrderHandler(c echo.Context) error {
	order, _ := entity.NewOrder("1", 10, 1)
	err := order.CalculateFinalPrice()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

// vanilla
// func OrderHandler(c http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// 	order, _ := entity.NewOrder("1", 10, 1)
// 	err := order.CalculateFinalPrice()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}

// 	result := json.NewEncoder(w).Encode(order)

// 	if result != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }
