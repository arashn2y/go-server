# Go Server - Product Management API

A RESTful API server built with Go for managing products. This project demonstrates modern Go development practices including clean architecture, database migrations, and type-safe SQL queries with sqlc.

## Features

- ✅ RESTful API for Product CRUD operations
- ✅ PostgreSQL database integration with pgx/v5
- ✅ Type-safe SQL queries using sqlc
- ✅ Database migrations with Goose
- ✅ Environment-based configuration
- ✅ Structured logging with Logrus
- ✅ HTTP routing with Chi router
- ✅ API testing with Bruno

## Tech Stack

- **Language**: Go 1.25.1
- **Web Framework**: Chi v5 (lightweight HTTP router)
- **Database**: PostgreSQL
- **Database Driver**: pgx/v5
- **SQL Generator**: sqlc
- **Migrations**: Goose v3
- **Configuration**: godotenv
- **Logging**: Logrus
- **API Client**: Bruno

## Prerequisites

Before running this project, ensure you have the following installed:

- Go 1.25.1 or higher
- PostgreSQL (running instance)
- sqlc (for generating type-safe SQL code)
- Bruno (optional, for API testing)

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/arashn2y/go-server.git
   cd go-server
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   Create a `.env` file in the root directory following the `.env.example` template:

   ```env
   ADDR=8080
   DSN=postgres://username:password@localhost:5432/database_name?sslmode=disable
   ```

4. **Set up the database**

   Create a PostgreSQL database and run migrations:

   ```bash
   # first add env variables to goose with this command
   goose -env .env
   # Then run all migrations manually using goose
   goose up
   ```

5. **Generate SQL code with sqlc**
   ```bash
   sqlc generate
   ```

## Running the Server

Start the server:

```bash
go run cmd/*.go
```

The server will start on the port specified in your `.env` file (default: 8080).

You can verify the server is running by accessing:

```
http://localhost:8080/health
```

## API Endpoints

### Health Check

```
GET /health
```

Returns server status.

### Products

| Method | Endpoint         | Description          |
| ------ | ---------------- | -------------------- |
| GET    | `/products`      | Get all products     |
| GET    | `/products/{id}` | Get a product by ID  |
| POST   | `/products`      | Create a new product |
| PUT    | `/products/{id}` | Update a product     |
| DELETE | `/products/{id}` | Delete a product     |

### Product Schema

```json
{
  "id": "uuid",
  "name": "string",
  "description": "string",
  "priceInCents": "integer",
  "createdAt": "timestamp"
}
```

### Example Request

**Create Product:**

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sample Product",
    "description": "A sample product description",
    "priceInCents": 1999
  }'
```

**Get All Products:**

```bash
curl http://localhost:8080/products
```

## Project Structure

```
go-server/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   └── api.go              # Application setup and routing
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── db/
│   │   ├── migrations/         # Database migrations
│   │   ├── queries/            # SQL query files for sqlc
│   │   └── schema/             # Database schema definitions
│   ├── form/
│   │   └── product.go          # Form validation
│   ├── json/
│   │   └── json.go             # JSON utilities
│   ├── models/
│   │   └── product.go          # Domain models
│   ├── products/
│   │   ├── handlers.go         # HTTP handlers
│   │   └── service.go          # Business logic
│   ├── repository/
│   │   ├── db.go               # Database connection
│   │   ├── models.go           # Generated models
│   │   └── product.sql.go      # Generated queries (sqlc)
│   └── utils/
│       └── helper.go           # Utility functions
├── Go-server/                  # Bruno API collection
│   ├── bruno.json
│   ├── environments/
│   └── Product/                # Product API tests
├── go.mod
├── sqlc.yaml                   # sqlc configuration
└── README.md
```

## Development

### Using sqlc

This project uses [sqlc](https://sqlc.dev/) to generate type-safe Go code from SQL queries.

1. Write your SQL queries in `internal/db/queries/`
2. Define your schema in `internal/db/schema/`
3. Run `sqlc generate` to generate Go code

The generated code will be placed in `internal/repository/`.

### Database Migrations

Database migrations are located in `internal/db/migrations/`. The project uses Goose for managing migrations.

### API Testing with Bruno

The project includes a Bruno collection for API testing located in the `Go-server/` directory. Open the collection with Bruno to test all API endpoints.

## Configuration

The application uses environment variables for configuration:

| Variable | Description                  | Example                                      |
| -------- | ---------------------------- | -------------------------------------------- |
| `ADDR`   | Server port                  | `8080`                                       |
| `DSN`    | PostgreSQL connection string | `postgres://user:pass@localhost:5432/dbname` |

## Middleware

The server includes the following middleware:

- Request ID tracking
- Real IP detection
- Request logging
- Panic recovery
- Request timeout (60 seconds)

## License

This project is licensed under the terms specified in the LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

[arashn2y](https://github.com/arashn2y)
