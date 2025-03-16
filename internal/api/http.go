package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"monitor-service/internal/adapters/database"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(db *database.MongoDB) *mux.Router {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		handleListEvents(db, w, r)
	}).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/events/{address}", func(w http.ResponseWriter, r *http.Request) {
		handleListEventsByAddress(db, w, r)
	}).Methods("GET", "OPTIONS")

	return router
}

func handleListEvents(db *database.MongoDB, w http.ResponseWriter, r *http.Request) {
	collectionName := r.URL.Query().Get("collection")
	if collectionName == "" {
		http.Error(w, `{"error": "collection is required"}`, http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	filters := bson.M{}

	if eventType := r.URL.Query().Get("eventType"); eventType != "" {
		filters["eventType"] = bson.M{"$regex": "^" + eventType + "$", "$options": "i"}
	}

	if contract := r.URL.Query().Get("contractAddress"); contract != "" {
		filters["contractAddress"] = bson.M{"$regex": "^" + contract + "$", "$options": "i"}
	}

	if userAddress := r.URL.Query().Get("userAddress"); userAddress != "" {
		filters["$or"] = bson.A{
			bson.M{"Agent": userAddress},
			bson.M{"from": userAddress},
			bson.M{"to": userAddress},
		}
	}

	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	if startDateStr != "" && endDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			log.Println("❌ Erro ao analisar startDate:", err)
			http.Error(w, `{"error": "Invalid startDate format"}`, http.StatusBadRequest)
			return
		}

		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			log.Println("❌ Erro ao analisar endDate:", err)
			http.Error(w, `{"error": "Invalid endDate format"}`, http.StatusBadRequest)
			return
		}

		endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

		filters["timestamp"] = bson.M{
			"$gte": startDate.Format(time.RFC3339),
			"$lt":  endDate.Format(time.RFC3339),
		}
	}

	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	log.Println("📡 Filtros aplicados na consulta:", filters)

	events, err := db.FindEvents(collectionName, filters, findOptions)
	if err != nil {
		log.Println("❌ Erro ao buscar eventos:", err)
		http.Error(w, `{"error": "Erro ao buscar eventos"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"events": events,
		"page":   page,
		"limit":  limit,
	})
}

func handleListEventsByAddress(db *database.MongoDB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		http.Error(w, `{"error": "Address is required"}`, http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	filters := bson.M{
		"$or": bson.A{
			bson.M{"From": address},
			bson.M{"To": address},
			bson.M{"Agent": address},
			bson.M{"Owner": address},
			bson.M{"owner": address},
		},
	}

	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	if startDateStr != "" && endDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			log.Println("❌ Erro ao analisar startDate:", err)
			http.Error(w, `{"error": "Invalid startDate format"}`, http.StatusBadRequest)
			return
		}

		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			log.Println("❌ Erro ao analisar endDate:", err)
			http.Error(w, `{"error": "Invalid endDate format"}`, http.StatusBadRequest)
			return
		}

		// 🔥 **Ajustar o final do dia para incluir o período completo**
		endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

		// ✅ **Adiciona o filtro de timestamp nos eventos relacionados ao endereço**
		filters["timestamp"] = bson.M{
			"$gte": startDate.Format(time.RFC3339),
			"$lt":  endDate.Format(time.RFC3339),
		}
	}

	// ✅ **Opções de busca e paginação**
	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

	log.Println("📡 Buscando eventos para o endereço:", address, "com filtros:", filters)

	// ✅ **Executando a busca**
	events, err := db.FindEventsInAllCollections(filters, findOptions)
	if err != nil {
		log.Println("❌ Erro ao buscar eventos por endereço:", err)
		http.Error(w, `{"error": "Erro ao buscar eventos"}`, http.StatusInternalServerError)
		return
	}

	// ✅ **Retornando JSON**
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"events": events,
		"page":   page,
		"limit":  limit,
	})
}
