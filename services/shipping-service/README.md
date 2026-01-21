# Shipping Service

The Shipping Service manages shipment tracking and delivery status in the e-commerce platform.

## Features

- **Shipment Management**: Create, update, and track shipments
- **Tracking Events**: Record and retrieve shipment tracking events
- **Status Updates**: Update shipment status (pending, in_transit, delivered, returned)
- **Tracking Number Lookup**: Find shipments by tracking number
- **RESTful API**: Full CRUD operations for shipment management
- **Swagger Documentation**: Comprehensive API documentation

## Database Schema

### shipments
- `id` (Primary Key)
- `order_id` (Foreign Key to orders)
- `tracking_number` (Unique tracking identifier)
- `carrier` (Shipping carrier)
- `shipping_method` (Shipping method)
- `status` (pending, in_transit, delivered, returned)
- `estimated_delivery` (Expected delivery date)
- `actual_delivery` (Actual delivery date)
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

### tracking_events
- `id` (Primary Key)
- `shipment_id` (Foreign Key to shipments)
- `event_type` (Type of tracking event)
- `location` (Location of event)
- `description` (Event description)
- `timestamp` (Event timestamp)
- `created_at` (Timestamp)

## API Endpoints

### Shipments
- `POST /api/v1/shipments` - Create shipment
- `GET /api/v1/shipments/:id` - Get shipment by ID
- `GET /api/v1/shipments/order/:order_id` - Get shipment by order ID
- `GET /api/v1/shipments/tracking/:tracking_number` - Get shipment by tracking number
- `PUT /api/v1/shipments/:id` - Update shipment
- `PATCH /api/v1/shipments/:id/status` - Update shipment status
- `DELETE /api/v1/shipments/:id` - Delete shipment
- `GET /api/v1/shipments` - Get all shipments with pagination
- `GET /api/v1/shipments/status/:status` - Get shipments by status

### Tracking
- `POST /api/v1/tracking` - Create tracking event
- `GET /api/v1/tracking/:shipment_id` - Get tracking events for shipment
- `GET /api/v1/tracking/:shipment_id/latest` - Get latest tracking event for shipment

### Documentation
- `GET /swagger/index.html` - Swagger UI
- `GET /health` - Health check

## Environment Variables

- `SHIPPING_DB_DSN` - Database connection string (default: `root:@tcp(localhost:3306)/shipping_service_db?parseTime=true`)

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run cmd/main.go
```

The service will start on port 8083.