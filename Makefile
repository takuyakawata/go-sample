GEN := go run ./scripts/gen
gen:            ## run gorm/gen
	$(GEN)

DB_URL := postgres://myapp:myapp@localhost:5432/myapp?sslmode=disable
MIGRATE := migrate -path migrations -database $(DB_URL)

.PHONY: gen migrate-up migrate-down tidy vet test cover

build:
	go build -o bin/myapp ./cmd/app/main.go

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

tidy:
	go mod tidy

vet:
	go vet ./...

test: gen vet
	go test ./... -race -coverprofile=coverage.out

cover:
	go tool cover -html=coverage.out
