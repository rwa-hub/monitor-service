package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(uri string, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return &MongoDB{client, db}, nil
}

func (m *MongoDB) InsertEvent(collection string, event interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.db.Collection(collection).InsertOne(ctx, event)
	if err != nil {
		log.Println("❌ Erro ao inserir evento no MongoDB:", err)
	}

	return err
}

func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.Background()); err != nil {
		log.Println("❌ Erro ao fechar conexão com MongoDB:", err)
	} else {
		log.Println("✅ Conexão com MongoDB fechada com sucesso!")
	}
}
