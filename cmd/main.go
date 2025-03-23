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
	rpcURL := "ws://host.docker.internal:8546"
	mongoURI := "mongodb://admin:password@host.docker.internal:27017/"
	mongoDBName := "monitor-service"

	contractAddresses := map[string]string{
		"financial_compliance":   "0xF235Bb30Ad375F279248AaFC89F4A92899A900De",
		"modular_compliance":     "0x850F95B0f32E9dB5AA484d160CB58f8A52103dc2",
		"ident_registry_storage": "0xd9bB4402537D044709BC80b666F522A6F0ce6435",
		"registry_md":            "0xC8c03647d39a96f02f6Ce8999bc22493C290e734",
		"identity_registry":      "0x9b5b3cCe1f7f8359E026f9573f258782Be577f29",
		"token_rwa":              "0x2078FaF714Fb3727a66bc10F7A9690b5A16CD0bb",
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
	port := ":8080"
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
