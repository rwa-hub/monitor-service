package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/internal/api"
	service "monitor-service/internal/service"
)

// checkRequiredEnvVars verifica se todas as variáveis de ambiente necessárias estão definidas
func checkRequiredEnvVars() error {
	required := []struct {
		name        string
		description string
	}{
		{"CONTRACT_COMPLIANCE", "Endereço do contrato Financial Compliance"},
		{"CONTRACT_MODULAR_COMPLIANCE", "Endereço do contrato Modular Compliance"},
		{"CONTRACT_IDENT_REGISTRY_STORAGE", "Endereço do contrato Identity Registry Storage"},
		{"CONTRACT_REGISTRY_MD", "Endereço do contrato Registry MD"},
		{"CONTRACT_IDENTITY_REGISTRY", "Endereço do contrato Identity Registry"},
		{"CONTRACT_TOKEN_RWA", "Endereço do contrato Token RWA"},
	}

	var missingVars []string
	for _, env := range required {
		if os.Getenv(env.name) == "" {
			missingVars = append(missingVars, fmt.Sprintf("%s (%s)", env.name, env.description))
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("as seguintes variáveis de ambiente são obrigatórias mas não foram definidas:\n%s",
			formatMissingVars(missingVars))
	}

	return nil
}

// formatMissingVars formata a lista de variáveis faltantes para exibição
func formatMissingVars(vars []string) string {
	var result string
	for _, v := range vars {
		result += fmt.Sprintf("- %s\n", v)
	}
	return result
}

func main() {
	// 🔹 Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ Arquivo .env não encontrado: %v", err)
	}

	// Verifica variáveis de ambiente obrigatórias
	if err := checkRequiredEnvVars(); err != nil {
		log.Fatalf("❌ Erro de configuração:\n%v", err)
	}

	// 🔹 Configuração inicial
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = "ws://host.docker.internal:8546"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:password@mongodb:27017/"
	}

	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoDBName == "" {
		mongoDBName = "monitor-service"
	}

	// 🔹 Configuração dos endereços dos contratos
	contractAddresses := map[string]string{
		"financial_compliance":   os.Getenv("CONTRACT_COMPLIANCE"),
		"modular_compliance":     os.Getenv("CONTRACT_MODULAR_COMPLIANCE"),
		"ident_registry_storage": os.Getenv("CONTRACT_IDENT_REGISTRY_STORAGE"),
		"registry_md":            os.Getenv("CONTRACT_REGISTRY_MD"),
		"identity_registry":      os.Getenv("CONTRACT_IDENTITY_REGISTRY"),
		"token_rwa":              os.Getenv("CONTRACT_TOKEN_RWA"),
	}

	wsServer := websocket.NewWebSocketServer()
	log.Println("📡 Iniciando WebSocket Server...")

	// 🔹 Inicializa conexão com o RPC
	rpcClient, err := rpc.NewRPCClient(rpcURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao RPC: %v", err)
	}

	client, err := rpcClient.Client()
	if err == nil {
		defer client.Close()
	}

	// 🔹 Conexão com MongoDB
	db, err := database.NewMongoDB(mongoURI, mongoDBName)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao MongoDB: %v", err)
	}

	// 🔹 Conexão com RabbitMQ
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	if queueName == "" {
		queueName = "monitor-events" // valor padrão
	}

	queueService, err := queue.NewRabbitMQ(queueName)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao RabbitMQ: %v", err)
	}
	defer queueService.Close()

	// 🔥 Configuração do Socket.IO
	serverSocket := socketio.NewServer(nil)

	// 🔥 Evento de conexão do Socket.IO
	serverSocket.OnConnect("/", func(s socketio.Conn) error {
		log.Println("🔌 Novo cliente conectado:", s.ID())
		s.Emit("message", "Bem-vindo ao WebSocket!")
		return nil
	})

	// 🔥 Evento de recebimento de mensagem
	serverSocket.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Println("📩 Mensagem recebida:", msg)
		s.Emit("response", "Mensagem recebida com sucesso!")
	})

	// 🔥 Evento de desconexão
	serverSocket.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("❌ Cliente desconectado:", s.ID(), "Motivo:", reason)
	})

	go serverSocket.Serve()
	defer serverSocket.Close()

	// 🔹 Configuração das rotas HTTP
	router := api.SetupRoutes(db)

	// 🔹 Habilita WebSocket `/ws`
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		wsServer.HandleConnections(w, r)
	})

	// 🔹 Rota para Socket.IO `/socket.io/`
	router.Handle("/socket.io/", serverSocket)

	// 🔹 Inicia serviço de eventos
	go service.EventService(rpcClient, wsServer, queueService, db, contractAddresses)

	// 🔹 Inicia servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // valor padrão
	}
	if port[0] != ':' {
		port = ":" + port
	}

	log.Printf("✅ Servidor iniciado na porta %s", port)

	srv := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("❌ Erro ao iniciar o servidor HTTP: %v", err)
	}
}
