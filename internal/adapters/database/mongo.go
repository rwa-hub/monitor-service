package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(uri string, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(30 * time.Second).
		SetConnectTimeout(30 * time.Second).
		SetSocketTimeout(60 * time.Second).
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(60 * time.Second).
		SetRetryWrites(true).
		SetRetryReads(true).
		SetReadPreference(readpref.Primary()).
		SetHeartbeatInterval(10 * time.Second).
		SetAppName("monitor-service")

	var client *mongo.Client
	var err error
	maxRetries := 5
	backoff := time.Second

	for i := 0; i < maxRetries; i++ {
		client, err = mongo.Connect(ctx, clientOptions)
		if err == nil {
			break
		}

		log.Printf("⚠️ Tentativa %d/%d de conectar ao MongoDB falhou: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(backoff)
			backoff *= 2
		}
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao MongoDB após %d tentativas: %v", maxRetries, err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer pingCancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("erro ao verificar conexão com MongoDB: %v", err)
	}

	log.Println("✅ Conexão com MongoDB estabelecida com sucesso!")

	db := client.Database(dbName)
	mongodb := &MongoDB{client, db}

	go mongodb.monitorConnection()

	return mongodb, nil
}

func (m *MongoDB) monitorConnection() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := m.client.Ping(ctx, readpref.Primary()); err != nil {
			log.Printf("⚠️ Problema detectado na conexão com MongoDB: %v", err)
			// Aqui podemos adicionar métricas ou alertas se necessário
		}
		cancel()
	}
}

func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return m.db.Collection(collectionName)
}

func (m *MongoDB) InsertEvent(collection string, event interface{}) error {
	maxRetries := 3
	var lastErr error
	backoff := time.Second

	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		_, err := m.db.Collection(collection).InsertOne(ctx, event)
		cancel()

		if err == nil {
			return nil
		}

		lastErr = err
		log.Printf("❌ Tentativa %d/%d - Erro ao inserir evento na coleção [%s]: %v\n", i+1, maxRetries, collection, err)

		if i < maxRetries-1 {
			time.Sleep(backoff)
			backoff *= 2 // Exponential backoff
		}
	}

	return fmt.Errorf("falha após %d tentativas: %v", maxRetries, lastErr)
}

// ✅ Fecha a conexão com o banco de dados
func (m *MongoDB) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := m.client.Disconnect(ctx); err != nil {
		log.Println("❌ Erro ao fechar conexão com MongoDB:", err)
	} else {
		log.Println("✅ Conexão com MongoDB fechada com sucesso!")
	}
}

// ✅ Busca eventos com filtros e paginação
func (m *MongoDB) FindEvents(collection string, filter interface{}, findOptions *options.FindOptions) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := m.db.Collection(collection).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar eventos: %v", err)
	}
	defer cur.Close(ctx)

	var events []bson.M
	if err := cur.All(ctx, &events); err != nil {
		return nil, fmt.Errorf("erro ao decodificar eventos: %v", err)
	}

	return events, nil
}

func (db *MongoDB) FindEventsInAllCollections(filters bson.M, findOptions *options.FindOptions) ([]bson.M, error) {
	var results []bson.M

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collections, err := db.db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Println("❌ Erro ao listar coleções:", err)
		return nil, err
	}

	for _, collectionName := range collections {
		log.Printf("🔍 Buscando eventos na coleção [%s]...\n", collectionName)

		collection := db.db.Collection(collectionName)
		searchCtx, searchCancel := context.WithTimeout(context.Background(), 30*time.Second)

		cursor, err := collection.Find(searchCtx, filters, findOptions)
		if err != nil {
			log.Printf("⚠️ Erro ao buscar eventos na coleção [%s]: %v\n", collectionName, err)
			searchCancel()
			continue
		}

		var collectionResults []bson.M
		if err := cursor.All(searchCtx, &collectionResults); err != nil {
			log.Printf("⚠️ Erro ao decodificar eventos da coleção [%s]: %v\n", collectionName, err)
			cursor.Close(searchCtx)
			searchCancel()
			continue
		}

		cursor.Close(searchCtx)
		searchCancel()
		results = append(results, collectionResults...)
	}

	log.Printf("✅ Consulta concluída! Total de eventos encontrados: %d\n", len(results))
	return results, nil
}
