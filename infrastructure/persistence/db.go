package persistence

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitDB inicializa la conexión a MongoDB
func InitDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://Jutonito:CpgxsqUP44a3yw2a@cluster0.zx1vk.mongodb.net/")
	newClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
		return nil, err
	}

	client = newClient
	log.Println("Conexión a MongoDB exitosa.")
	return client, nil
}

// GetDB retorna la base de datos 'Development'
func GetDB() *mongo.Database {
	if client == nil {
		log.Fatal("La conexión a la base de datos no ha sido inicializada")
	}
	return client.Database("Development")
}
