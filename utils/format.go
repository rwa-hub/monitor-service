package utils

import (
	"encoding/json"
	"time"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/logger"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/websocket"
)

func ProcessEvent(eventType string, event interface{}, wsServer *websocket.WebSocketServer, queueService *queue.RabbitMQ, db *database.MongoDB, collectionName string) {

	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao converter evento para JSON")
		return
	}

	var rawData map[string]interface{}
	if err := json.Unmarshal(eventBytes, &rawData); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao decodificar evento JSON")
		return
	}

	raw, ok := rawData["Raw"].(map[string]interface{})
	if !ok {
		logger.Log.Error().Str("eventType", eventType).Msg("❌ Estrutura de evento inválida: campo 'Raw' não encontrado")
		return
	}

	formattedEvent := map[string]interface{}{
		"eventType":       eventType,
		"timestamp":       time.Now().UTC().Format(time.RFC3339),
		"contractAddress": raw["address"],
		"transactionHash": raw["transactionHash"],
		"blockNumber":     raw["blockNumber"],
	}

	for key, value := range rawData {
		if key != "Raw" {
			formattedEvent[key] = value
		}
	}

	formattedBytes, err := json.Marshal(formattedEvent)
	if err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao formatar evento JSON")
		return
	}

	logger.Log.Info().Str("eventType", eventType).RawJSON("formatted_data", formattedBytes).Msg("📢 Evento capturado")

	wsServer.Broadcast(formattedBytes)

	if err := queueService.Publish(formattedBytes); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao publicar evento no RabbitMQ")
	}

	if err := db.InsertEvent(collectionName, formattedEvent); err != nil {
		logger.Log.Error().Err(err).Str("eventType", eventType).Msg("❌ Erro ao salvar evento no MongoDB")
	}
}
