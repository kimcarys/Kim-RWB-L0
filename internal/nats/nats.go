package nats

import (
	"encoding/json"
	"log"
	"order-service/internal/cache"
	"order-service/internal/database"
	"order-service/internal/models"

	"github.com/nats-io/stan.go"
)

func ConnectAndSubscribe(cache *cache.Cache, db *database.DB) error {
	sc, err := stan.Connect("test-cluster", "order-service-client", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}

	_, err = sc.Subscribe("orders", func(msg *stan.Msg) {
		var order models.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Error unmarshaling order: %v", err)
			return
		}

		// Валидация данных
		if err := order.Validate(); err != nil {
			log.Printf("Invalid order data: %v", err)
			return
		}

		log.Printf("Received order: %s", order.OrderUID)

		// Сохраняем в БД
		if err := db.SaveOrder(&order); err != nil {
			log.Printf("Error saving order to DB: %v", err)
			return
		}

		// Сохраняем в кэш
		cache.Set(&order)

		log.Printf("Order %s saved successfully", order.OrderUID)
	}, stan.DurableName("order-service"))

	if err != nil {
		return err
	}

	log.Println("Subscribed to NATS channel 'orders'")
	return nil
}