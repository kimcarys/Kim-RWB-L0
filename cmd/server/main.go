package main

import (
	"log"
	"order-service/internal/cache"
	"order-service/internal/database"
	"order-service/internal/handler"
	"order-service/internal/nats"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Инициализация кэша
	orderCache := cache.NewCache()
	
	// Восстановление кэша из БД
	err = orderCache.RestoreFromDB(db)
	if err != nil {
		log.Printf("Warning: Failed to restore cache from DB: %v", err)
	}

	// Подключение к NATS Streaming
	err = nats.ConnectAndSubscribe(orderCache, db)
	if err != nil {
		log.Fatal("Failed to connect to NATS:", err)
	}

	log.Println("Order service started successfully!")
	
	// Запуск HTTP сервера
	handler.StartServer(orderCache)
}