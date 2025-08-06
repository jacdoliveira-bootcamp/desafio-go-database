default: dev

start:
	@echo "Iniciando aplicação Go..."
	go run ./cmd/main.go

dev:
	air

up-build:
	@echo "Subindo containers (rebuild)..."
	docker-compose up --build

up:
	@echo "Subindo containers..."
	docker-compose up -d

down:
	@echo "Derrubando containers..."
	docker-compose down

down-v:
	@echo "Derrubando containers e removendo volumes..."
	docker-compose down -v

logs:
	@echo "Exibindo logs dos containers..."
	docker-compose logs -f

ps:
	@echo "Listando containers ativos..."
	docker-compose ps