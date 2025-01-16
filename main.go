package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mp02/fravega-tech/interfaces"
	"github.com/mp02/fravega-tech/repository"
	"github.com/mp02/fravega-tech/routes"
	"github.com/mp02/fravega-tech/usecases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// @title API ABM Catalog
	// @version 1.0
	// @description Challange Fravega Tech

	// @contact.name Martin Pruyas
	// @contact.url https://www.linkedin.com/in/martin-pruyas/
	// @contact.email o.gema.pg@gmail.com

	// @host localhost:8080
	// @BasePath /v1

	// Configuración de MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}
	fmt.Println(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		fmt.Println("MongoDB connection failed")
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collection := client.Database("productsdb").Collection("products")

	// Inicialización de capas
	repo := repository.NewMongoProductRepository(collection)
	useCase := usecases.NewProductUseCase(repo)
	handler := interfaces.NewProductHandler(useCase)
	// Configuración de rutas con Gin
	router := routes.SetupRoutes(handler)

	// Ejecuta el servidor
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
