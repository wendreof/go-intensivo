# Crash Course - Golang: Performance and Multithreading with RabbitMQ

### Starting RabbitMQ
`docker-compose up -d`

### Creating table orders
`sqlite3.exe orders.db`

### Running consumer
`go run cmd/consumer/main.go`

### Adding the packages
`go mod tidy`

### Running tests
`go test ./...`