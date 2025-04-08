package main

import (
	"context"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/http"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/repo"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/usecase"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	repo := repo.NewMongoRepo(client)
	uc := usecase.NewProductUseCase(repo)
	r := http.SetupRoutes(uc)
	r.Run(":8081") // Запуск на порту 8081
}
