.PHONY: all build test clean deps

# Variáveis
GO=go
BINARY_NAME=monitor-service
MAIN_FILE=cmd/main.go

all: deps build

# Instala todas as dependências do projeto
deps:
	$(GO) mod download
	$(GO) mod tidy

# Constrói o projeto
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_FILE)

# Executa os testes
test:
	$(GO) test ./... -v

# Limpa os arquivos gerados
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

# Executa o projeto
run:
	$(GO) run $(MAIN_FILE)

# Verifica o código
lint:
	$(GO) vet ./...
	golangci-lint run

# Atualiza todas as dependências
update-deps:
	$(GO) get -u all
	$(GO) mod tidy   