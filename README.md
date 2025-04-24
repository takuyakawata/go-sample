# Product Management System

A Domain-Driven Design (DDD) implementation of a product management system in Go.

## Project Structure

This project follows the principles of Domain-Driven Design (DDD) and Clean Architecture:

```
.
├── cmd/                  # Application entry points
│   └── app/              # Main application
│       └── main.go       # Main function
├── internal/             # Private application code
│   ├── domain/           # Domain layer (entities, value objects, repositories interfaces)
│   │   └── product/      # Product domain
│   ├── usecase/          # Application layer (use cases)
│   │   └── product/      # Product use cases
│   ├── interface/        # Interface layer (controllers, presenters)
│   │   └── api/          # API interfaces
│   │       └── handler/  # HTTP handlers
│   └── infrastructure/   # Infrastructure layer (repository implementations, external services)
│       └── repository/   # Repository implementations
│           └── memory/   # In-memory repository implementation
└── README.md             # This file
```

## Domain Model

The core domain model is the Product entity, which is composed of several value objects:

- **Product**: The main entity representing a product
- **ProductID**: Value object for product identifier
- **ProductName**: Value object for product name
- **ProductDescription**: Value object for product description
- **Price**: Value object for product price (amount and currency)
- **Stock**: Value object for product stock quantity

## Use Cases

The application supports the following use cases:

1. Create a new product
2. Update an existing product
3. Delete a product
4. Get a product by ID
5. Get all products

## API Endpoints

The application exposes the following REST API endpoints:

- `POST /products` - Create a new product
- `PUT /products/{id}` - Update an existing product
- `DELETE /products/{id}` - Delete a product
- `GET /products/{id}` - Get a product by ID
- `GET /products` - Get all products

## Running the Application

To run the application:

```bash
go run cmd/app/main.go
```

The server will start on port 8080.

## Example API Requests

### Create a Product

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "id": "prod-001",
    "name": "Smartphone",
    "description": "Latest model smartphone",
    "price": 999,
    "currency": "USD",
    "stock": 100
  }'
```

### Get a Product

```bash
curl -X GET http://localhost:8080/products/prod-001
```

### Update a Product

```bash
curl -X PUT http://localhost:8080/products/prod-001 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Smartphone Pro",
    "description": "Latest model smartphone with pro features",
    "price": 1299,
    "currency": "USD",
    "stock": 50
  }'
```

### Delete a Product

```bash
curl -X DELETE http://localhost:8080/products/prod-001
```

### Get All Products

```bash
curl -X GET http://localhost:8080/products
```

## Design Decisions

1. **Domain-Driven Design**: The project follows DDD principles to focus on the core domain and domain logic.
2. **Clean Architecture**: The codebase is organized in layers (domain, application, interface, infrastructure) to separate concerns.
3. **Value Objects**: Domain concepts are modeled as value objects to encapsulate validation and business rules.
4. **Repository Pattern**: The domain layer defines repository interfaces, while the infrastructure layer provides implementations.
5. **Use Case Pattern**: Application logic is organized into use cases that orchestrate the flow of data to and from the domain.
6. **Dependency Injection**: Dependencies are injected rather than created by the components that use them.
7. **Immutability**: Value objects are immutable to ensure consistency.

## Future Improvements

1. Add a persistent database implementation (PostgreSQL, MongoDB)
2. Add authentication and authorization
3. Implement pagination for listing products
4. Add more comprehensive validation
5. Add logging and monitoring
6. Implement caching for frequently accessed data
7. Add more comprehensive error handling
