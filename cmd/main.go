package main

import (
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"

	"monitor-service/internal/adapters/database"
	"monitor-service/internal/adapters/queue"
	"monitor-service/internal/adapters/rpc"
	"monitor-service/internal/adapters/websocket"
	"monitor-service/internal/api"
	service "monitor-service/internal/service"
)

func main() {
	// 🔹 Configuração inicial
	rpcURL := "ws://localhost:8546"
	mongoURI := "mongodb://admin:password@localhost:27017/"
	mongoDBName := "monitor-service"

	contractAddresses := map[string]string{
		"financial_compliance":   "0xb726288807e3D9D8a1b820C6Cb59b00546Eec6Ec",
		"modular_compliance":     "0x727547e5E17756425A257827DeA9832f845Bb6e9",
		"ident_registry_storage": "0x1E39d4bCda350162C60DDb56a49241eC1Fb0F8b0",
		"registry_md":            "0x64998624b832cF3F95A81F9b383a9cFc7c58D064",
		"identity_registry":      "0x0D8CA8E4c0ECE2a1107d712ec76F6F91814Be243",
		"token_rwa":              "0x64998624b832cF3F95A81F9b383a9cFc7c58D064",
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
	queueService, err := queue.NewRabbitMQ("monitor-events")
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
	port := ":8082"
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
