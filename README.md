# Crash Course - Golang: Performance and Multithreading with RabbitMQ, SQLite3, Prometheus and Grafana


[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) 

[![codecov](https://codecov.io/gh/wendreof/go-intensivo/branch/main/graph/badge.svg?token=UVqukYJNTJ)](https://codecov.io/gh/wendreof/go-intensivo) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=bugs)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) 
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo) [![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=wendreof_go-intensivo&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=wendreof_go-intensivo)

### Starting RabbitMQ
```bash
docker-compose up -d
```

### Containerizing the app
```bash
docker build -t wendreof/go-intensivo:latest -f Dockerfile.prod .
```

### Accessing services
`http://localhost:15672/`
`http://localhost:3000/`
`http://localhost:9090/graph`

### Accessing endpoints
`http://localhost:8080/total`

### Grafana dashboard to import
`https://grafana.com/grafana/dashboards/10991-rabbitmq-overview/`

### Creating table orders
```bash
sqlite3.exe orders.db
```

### Running producer (Produces new messages to RabbitMQ)
```bash
go run cmd/producer/main.go
```

### Running consumer (Consume messages from RabbitMQ and insert into db)
```bash
go run cmd/consumer/main.go
```

### Adding the packages
```bash
go mod tidy
```

### Running tests
```bash
go test ./...
```

### Generating executable
```bash
go build cmd/consumer/main.go
```