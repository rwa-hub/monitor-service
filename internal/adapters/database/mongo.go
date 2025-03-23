package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

// ✅ Conexão com MongoDB
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

// ✅ Retorna uma coleção do banco de dados
func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return m.db.Collection(collectionName)
}

// ✅ Insere um evento no MongoDB
func (m *MongoDB) InsertEvent(collection string, event interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := m.db.Collection(collection).InsertOne(ctx, event)
	if err != nil {
		log.Printf("❌ Erro ao inserir evento na coleção [%s]: %v\n", collection, err)
	}

	return err
}

// ✅ Fecha a conexão com o banco de dados
func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.Background()); err != nil {
		log.Println("❌ Erro ao fechar conexão com MongoDB:", err)
	} else {
		log.Println("✅ Conexão com MongoDB fechada com sucesso!")
	}
}

// ✅ Busca eventos com filtros e paginação
func (m *MongoDB) FindEvents(collection string, filter interface{}, findOptions *options.FindOptions) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := m.db.Collection(collection).Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var events []bson.M
	if err := cur.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// ✅ 🔥 Busca eventos em TODAS as coleções
func (db *MongoDB) FindEventsInAllCollections(filters bson.M, findOptions *options.FindOptions) ([]bson.M, error) {
	var results []bson.M

	// 🔹 Obtendo todas as coleções
	collections, err := db.db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Println("❌ Erro ao listar coleções:", err)
		return nil, err
	}

	for _, collectionName := range collections {
		log.Printf("🔍 Buscando eventos na coleção [%s]...\n", collectionName)

		collection := db.db.Collection(collectionName)

		// ✅ Criando um novo contexto para cada busca
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cursor, err := collection.Find(ctx, filters, findOptions)
		if err != nil {
			log.Printf("⚠️ Erro ao buscar eventos na coleção [%s]: %v\n", collectionName, err)
			continue // Continua para a próxima coleção
		}

		var collectionResults []bson.M
		if err := cursor.All(ctx, &collectionResults); err != nil {
			log.Printf("⚠️ Erro ao decodificar eventos da coleção [%s]: %v\n", collectionName, err)
			continue
		}

		results = append(results, collectionResults...)
	}

	log.Printf("✅ Consulta concluída! Total de eventos encontrados: %d\n", len(results))
	return results, nil
}
