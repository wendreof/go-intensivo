# Crash Course - Golang: Performance and Multithreading with RabbitMQ

### Starting RabbitMQ
`docker-compose up -d`

### Accessing services
`http://localhost:15672/`
`http://localhost:3000/`
`http://localhost:9090/graph`

### Grafana dashboard to import
`https://grafana.com/grafana/dashboards/10991-rabbitmq-overview/`

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