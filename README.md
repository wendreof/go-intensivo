# Crash Course - Golang: Performance and Multithreading with RabbitMQ

### Starting RabbitMQ
`docker-compose up -d`

### Creating table orders
`sqlite3.exe orders.db`

### Running producer (Produces new messages to RabbitMQ)
`go run cmd/producer/main.go`

### Running consumer (Consume messages from RabbitMQ and insert into db)
`go run cmd/consumer/main.go`

### Adding the packages
`go mod tidy`

### Running tests
`go test ./...`

### Generating executable
`go build cmd/consumer/main.go`