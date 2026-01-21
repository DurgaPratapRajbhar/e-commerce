# Cart Service

This is the cart service for the e-commerce platform. It provides RESTful APIs to manage shopping cart functionality.

## Features

- Add items to cart
- Retrieve cart items by ID or user ID
- Update cart items
- Remove items from cart
- Clear entire cart for a user

## API Endpoints

| Method | Endpoint           | Description                 |
|--------|--------------------|-----------------------------|
| POST   | `/api/carts/`      | Add item to cart            |
| GET    | `/api/carts/:id`   | Get cart item by ID         |
| GET    | `/api/carts/user/:userId` | Get all cart items for a user |
| PUT    | `/api/carts/:id`   | Update cart item            |
| DELETE | `/api/carts/:id`   | Remove item from cart       |
| DELETE | `/api/carts/user/:userId` | Clear user's cart     |

## Data Model

### Cart

| Field      | Type     | Description              |
|------------|----------|--------------------------|
| ID         | uint64   | Primary key              |
| UserID     | uint64   | User identifier          |
| ProductID  | uint64   | Product identifier       |
| VariantID  | *uint64  | Product variant (optional) |
| Quantity   | int      | Item quantity            |
| CreatedAt  | time.Time| Creation timestamp       |
| UpdatedAt  | time.Time| Last update timestamp    |

## Setup

1. Copy `.env.example` to `.env` and configure the values
2. Run the service with `go run cmd/main.go`

## Docker

Build the Docker image:
```bash
docker build -t cart-service .
```

Run the container:
```bash
docker run -p 8080:8080 cart-service
```