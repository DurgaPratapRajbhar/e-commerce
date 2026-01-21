# Auth Service

Authentication microservice for the e-commerce platform.

## Features

- User registration
- User login
- JWT token generation and validation
- Token refresh
- User logout
- Role-based permissions
- Permission checking middleware

## Endpoints

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login with email and password
- `POST /api/auth/refresh` - Refresh access token
- `GET /api/auth/me` - Get current user information
- `POST /api/auth/logout` - Logout user
- `GET /api/permissions/user?role={role}` - Get permissions for a role
- `GET /api/permissions/check?role={role}&permission={permission}` - Check if a role has a specific permission
- `GET /health` - Health check endpoint
- `GET /swagger/index.html` - Swagger API documentation

## Tech Stack

- Go
- Gin Framework
- GORM
- MySQL
- JWT

## Roles and Permissions

### Roles
- `admin` - Full access to all resources
- `merchant` - Access to products, orders, and categories
- `user` - Read access to products and categories, write access to orders

### Permissions
- `read:users`, `write:users`, `delete:users`
- `read:products`, `write:products`, `delete:products`
- `read:orders`, `write:orders`, `delete:orders`
- `read:categories`, `write:categories`, `delete:categories`

## Swagger Documentation

The API is documented using Swagger. You can access the documentation at:

- `http://localhost:8080/swagger/index.html`

The Swagger UI provides interactive documentation for all API endpoints, including request/response schemas and example requests.

## Installation

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up environment variables
4. Run the service: `go run cmd/main.go`

## Environment Variables

- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET_KEY` - Secret key for JWT signing

## Docker

Build and run with Docker:

```bash
docker-compose up -d
```