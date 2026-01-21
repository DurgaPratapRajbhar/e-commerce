# API Documentation

## Authentication

### Login
- **Endpoint**: `POST /api/v1/auth/login`
- **Request**: `{ "email": "string", "password": "string" }`
- **Response**: `{ "success": true, "user": {...} }`
- **Cookie**: `auth_token` (HttpOnly, Secure)

### Register
- **Endpoint**: `POST /api/v1/auth/register`
- **Request**: `{ "email": "string", "password": "string", "username": "string" }`
- **Response**: `{ "success": true, "user": {...} }`

### Logout
- **Endpoint**: `POST /api/v1/auth/logout`
- **Response**: `{ "success": true }`
- **Cookie**: Clears `auth_token`

### Get Me
- **Endpoint**: `GET /api/v1/auth/me`
- **Headers**: Authorization Bearer token or Cookie
- **Response**: `{ "user": {...} }`

## Products

### Get Products
- **Endpoint**: `GET /api/v1/product/products`
- **Query Parameters**: `page`, `limit`, `search`, `category`, `status`
- **Response**: `{ "data": [...], "total": number }`

### Create Product
- **Endpoint**: `POST /api/v1/product/products`
- **Request**: Product data object
- **Response**: `{ "product": {...} }`

## Categories

### Get Categories
- **Endpoint**: `GET /api/v1/product/categories`
- **Response**: `{ "data": [...] }`

### Create Category
- **Endpoint**: `POST /api/v1/product/categories`
- **Request**: `{ "name": "string", "slug": "string", "parent_id": number }`
- **Response**: `{ "category": {...} }`

## Users

### Get Users
- **Endpoint**: `GET /api/v1/user/profiles`
- **Response**: `{ "data": [...] }`

### Get User by ID
- **Endpoint**: `GET /api/v1/user/profiles/{id}`
- **Response**: `{ "user": {...} }`

## Orders

### Get Orders
- **Endpoint**: `GET /api/v1/order`
- **Response**: `{ "data": [...] }`

### Get Order by ID
- **Endpoint**: `GET /api/v1/order/{id}`
- **Response**: `{ "order": {...} }`

## Cart

### Get Cart by User
- **Endpoint**: `GET /api/v1/cart/user/{userId}`
- **Response**: `{ "cart": {...} }`

## Inventory

### Get Low Stock Items
- **Endpoint**: `GET /api/v1/inventory/low-stock`
- **Response**: `{ "data": [...] }`

## Payments

### Get Payments by Order
- **Endpoint**: `GET /api/v1/payment/order/{orderId}`
- **Response**: `{ "data": [...] }`

## Shipments

### Get Shipment by Order
- **Endpoint**: `GET /api/v1/shipment/order/{orderId}`
- **Response**: `{ "shipment": {...} }`