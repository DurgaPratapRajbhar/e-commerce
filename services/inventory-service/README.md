# Inventory Service

The Inventory Service manages product inventory and tracks inventory transactions in the e-commerce platform.

## Features

- **Inventory Management**: Track product quantities, reserved quantities, and warehouse locations
- **Transaction Tracking**: Record all inventory changes with types (in, out, reserved)
- **Low Stock Alerts**: Identify items with low inventory levels
- **Reservation System**: Reserve inventory for pending orders
- **RESTful API**: Full CRUD operations for inventory management
- **Swagger Documentation**: Comprehensive API documentation

## Database Schema

### inventory
- `id` (Primary Key)
- `product_id` (Foreign Key to products)
- `variant_id` (Foreign Key to product variants)
- `quantity` (Available quantity)
- `reserved_quantity` (Quantity reserved for pending orders)
- `warehouse_location` (Storage location)
- `last_updated` (Timestamp)
- `created_at` (Timestamp)

### inventory_transactions
- `id` (Primary Key)
- `product_id` (Foreign Key to products)
- `variant_id` (Foreign Key to product variants)
- `transaction_type` (in, out, reserved)
- `quantity` (Amount of change)
- `reference_id` (Order ID, etc.)
- `notes` (Additional information)
- `created_at` (Timestamp)

## API Endpoints

### Inventory Management
- `POST /api/v1/inventory` - Create inventory record
- `GET /api/v1/inventory/product/:product_id/variant/:variant_id` - Get inventory by product and variant
- `PUT /api/v1/inventory/update` - Update inventory
- `GET /api/v1/inventory/low-stock` - Get low stock items
- `DELETE /api/v1/inventory/product/:product_id/variant/:variant_id` - Delete inventory

### Inventory Transactions
- `POST /api/v1/inventory/transactions` - Create transaction
- `GET /api/v1/inventory/transactions/product/:product_id` - Get transactions by product
- `GET /api/v1/inventory/transactions/product/:product_id/variant/:variant_id` - Get transactions by product and variant
- `GET /api/v1/inventory/transactions/reference/:reference_id` - Get transactions by reference ID
- `GET /api/v1/inventory/transactions/recent` - Get recent transactions

### Documentation
- `GET /swagger/index.html` - Swagger UI
- `GET /health` - Health check

## Environment Variables

- `INVENTORY_DB_DSN` - Database connection string (default: `root:@tcp(localhost:3306)/inventory_service_db?parseTime=true`)

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run cmd/main.go
```

The service will start on port 8082.