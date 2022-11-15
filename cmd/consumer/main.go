package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wendreof/gointensivo/internal/order/infra/database"
	"github.com/wendreof/gointensivo/internal/order/usecase"
	"github.com/wendreof/gointensivo/pkg/rabbitmq"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()
	out := make(chan amqp.Delivery) // channel
	go rabbitmq.Consume(ch, out)    // T2

	//forever := make(chan bool)
	numberOfWorkers := 5

	for i := 1; i <= numberOfWorkers; i++ {
		go worker(out, &uc, i)
	}
	
	//<- forever
	
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUseCase := usecase.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUseCase.Execute()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(total)

	})
	
	http.ListenAndServe(":8080", nil)
}

func worker(deliveryMessage <- chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int){
		for msg := range deliveryMessage {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)

		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)

		if err != nil {
			panic(err)
		}
		
		msg.Ack(false)
		fmt.Printf("Worker %d has processed order %s", workerID, outputDTO.ID)
		fmt.Println(outputDTO)
		time.Sleep(1 * time.Second)
	}
}