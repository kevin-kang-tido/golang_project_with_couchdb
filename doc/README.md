
# Golang with CouchDB CRUD API

This project implements a simple CRUD (Create, Read, Update, Delete) API for managing products, built using Go (`gin-gonic` framework) and CouchDB. The architecture follows Clean Architecture principles inspired by Uncle Bob's design, ensuring separation of concerns and easy maintainability.

## Project Structure

```bash
golang_with_couchdb2/
├── /cmd
│   └── /app
│       └── main.go                    # Entry point of the application
├── /internal
│   ├── /adapters
│   │   └── /db
│   │       └── couchdb.go             # CouchDB connection and repository implementation
│   ├── /usecases
│   │   ├── /ports
│   │   │   └── product_repository.go  # Repository interface for products
│   │   └── /interactors
│   │       └── product_interactor.go  # Business logic for product operations
│   ├── /domain
│   │   └── /entities
│   │       └── product.go             # Product entity definition
│   ├── /delivery
│   │   └── product_handler.go         # HTTP handlers for product CRUD operations
│   └── /router
│       └── router.go                  # Route definitions using Gin
├── /configs
│   └── config.yaml                    # Configuration file (optional)
├── /scripts
│   └── build.sh                       # Build script (optional)
└── go.mod                             # Go modules file
```

## Features

- **Create a product**: `POST /products`
- **Get a product by ID**: `GET /products/{id}`
- **Update a product by ID**: `PUT /products/{id}`
- **Delete a product by ID**: `DELETE /products/{id}`

### Technologies

- **Go**: The programming language for the project.
- **Gin**: A web framework used to handle routing and HTTP requests.
- **CouchDB**: A NoSQL database used to store product data.
- **Clean Architecture**: This project follows Uncle Bob's clean architecture principles, keeping the code modular and maintainable.

---

## Setup Instructions

### Prerequisites

- **Go** (1.16 or higher): [Install Go](https://golang.org/doc/install)
- **CouchDB**: [Install CouchDB](https://couchdb.apache.org/#download) and ensure it is running on `http://localhost:5984`.
- **Gin**: Already included as a dependency in the `go.mod` file.

### Installing Dependencies

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/yourusername/golang_with_couchdb2.git
cd golang_with_couchdb2
```

Install the required dependencies:

```bash
go mod tidy
```

### Running the Application

1. Ensure CouchDB is running on `localhost:5984`.
   
2. Start the API server:

```bash
go run cmd/app/main.go
```

3. The server will start running on port `8080`. You can test the API at `http://localhost:8080`.

---

## API Endpoints

| Method   | Endpoint              | Description                 |
|----------|-----------------------|-----------------------------|
| `POST`   | `/products`            | Create a new product        |
| `GET`    | `/products`            | Get all products (TBD)      |
| `GET`    | `/products/:_id`       | Get a product by ID         |
| `PUT`    | `/products/:_id`       | Update a product by ID      |
| `DELETE` | `/products/:_id`       | Delete a product by ID      |

### Example Requests

- **Create a product:**

```bash
curl -X POST http://localhost:8080/products \
-d '{"name": "Product A", "description": "A sample product", "price": 99.99}' \
-H "Content-Type: application/json"
```

- **Get a product by ID:**

```bash
curl http://localhost:8080/products/{_id}
```

- **Update a product by ID:**

```bash
curl -X PUT http://localhost:8080/products/{_id} \
-d '{"name": "Product A Updated", "description": "Updated description", "price": 120.00}' \
-H "Content-Type: application/json"
```

- **Delete a product by ID:**

```bash
curl -X DELETE http://localhost:8080/products/{_id}
```

---

## Clean Architecture Overview

The project follows Clean Architecture principles:

- **Entities (`internal/domain/entities`)**: Contains core business objects (e.g., `Product`).
- **Use Cases (`internal/usecases/interactors`)**: Implements business logic through interactors.
- **Adapters (`internal/adapters/db`)**: Handles communication with external systems like CouchDB.
- **Delivery (`internal/delivery`)**: Defines HTTP handlers that interact with clients.
- **Router (`internal/router`)**: Sets up HTTP routes using Gin.

This structure ensures high separation of concerns and allows easy modification, testing, and scaling of the project.

---

## Testing

Currently, testing is not implemented, but unit and integration tests can be added by following this modular structure. Use Go's built-in testing framework or other third-party libraries like `testify`.

---

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue if you find any bugs or have feature suggestions.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

### Acknowledgments

This project was built with inspiration from Uncle Bob's Clean Architecture principles.

