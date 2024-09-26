docker-compose up --build -d
goose -dir migrations postgres "postgresql://localhost:5432/goose?sslmode=disable" up