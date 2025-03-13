package utils

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/websocket"
)

// ProcessEvent formata e envia eventos para WebSocket, RabbitMQ e MongoDB.
func ProcessEvent(eventType string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, collectionName string) {
	// Converte o evento para JSON
	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao converter evento para JSON")
		return
	}

	// Decodifica o JSON para enriquecer os dados
	var rawData map[string]interface{}
	if err := json.Unmarshal(eventBytes, &rawData); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao decodificar evento JSON")
		return
	}

	// Organiza os dados do evento
	formattedEvent := map[string]interface{}{
		"eventType":       eventType,
		"timestamp":       time.Now().UTC().Format(time.RFC3339),
		"contractAddress": rawData["Raw"].(map[string]interface{})["address"],
		"transactionHash": rawData["Raw"].(map[string]interface{})["transactionHash"],
		"blockNumber":     rawData["Raw"].(map[string]interface{})["blockNumber"],
		"data":            rawData,
	}

	// Converte para JSON formatado
	formattedBytes, err := json.Marshal(formattedEvent)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao formatar evento JSON")
		return
	}

	// 🔥 Log informativo
	logger.Log.Info().Str("eventType", eventType).RawJSON("formatted_data", formattedBytes).Msg("📢 Evento capturado")

	// 🔵 Envia via WebSocket
	wsServer.Broadcast(formattedBytes)

	// 📩 Publica no RabbitMQ
	if err := queueService.Publish(formattedBytes); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao publicar evento no RabbitMQ")
	}

	// 🗄️ Armazena no MongoDB
	if err := db.InsertEvent(collectionName, formattedEvent); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao salvar evento no MongoDB")
	}
}
