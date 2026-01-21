# Payment Service

The Payment Service manages payment processing and refund operations in the e-commerce platform.

## Features

- **Payment Processing**: Create and manage payments with multiple payment methods
- **Refund Management**: Process refunds for successful payments
- **Payment Status Tracking**: Track payment status (pending, success, failed, refunded)
- **Transaction Management**: Handle transaction IDs and gateway responses
- **RESTful API**: Full CRUD operations for payment and refund management
- **Swagger Documentation**: Comprehensive API documentation

## Database Schema

### payments
- `id` (Primary Key)
- `order_id` (Foreign Key to orders)
- `user_id` (Foreign Key to users)
- `amount` (Payment amount)
- `payment_method` (card, upi, wallet, cod)
- `payment_status` (pending, success, failed, refunded)
- `transaction_id` (Transaction identifier)
- `gateway_response` (JSON response from payment gateway)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### refunds
- `id` (Primary Key)
- `payment_id` (Foreign Key to payments)
- `order_id` (Foreign Key to orders)
- `amount` (Refund amount)
- `reason` (Refund reason)
- `status` (pending, processed, failed)
- `created_at` (Timestamp)

## API Endpoints

### Payments
- `POST /api/v1/payments` - Create payment
- `GET /api/v1/payments/:id` - Get payment by ID
- `GET /api/v1/payments/order/:order_id` - Get payment by order ID
- `GET /api/v1/payments/transaction/:transaction_id` - Get payment by transaction ID
- `GET /api/v1/payments/user/:user_id` - Get payments by user ID
- `PUT /api/v1/payments/:id` - Update payment
- `PATCH /api/v1/payments/:id/status` - Update payment status
- `DELETE /api/v1/payments/:id` - Delete payment
- `GET /api/v1/payments` - Get all payments with pagination
- `GET /api/v1/payments/status/:status` - Get payments by status

### Refunds
- `POST /api/v1/refunds` - Create refund
- `GET /api/v1/refunds/:id` - Get refund by ID
- `GET /api/v1/refunds/payment/:payment_id` - Get refunds by payment ID
- `GET /api/v1/refunds/order/:order_id` - Get refunds by order ID
- `GET /api/v1/refunds` - Get all refunds with pagination
- `GET /api/v1/refunds/status/:status` - Get refunds by status

### Documentation
- `GET /swagger/index.html` - Swagger UI
- `GET /health` - Health check

## Environment Variables

- `PAYMENT_DB_DSN` - Database connection string (default: `root:@tcp(localhost:3306)/payment_service_db?parseTime=true`)

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run cmd/main.go
```

The service will start on port 8084.