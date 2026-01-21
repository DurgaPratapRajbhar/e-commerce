# Admin Frontend API Mapping

## Service Files Overview

### authService.js
Main authentication service with validation error handling

**Usage Example:**
```javascript
import { authService } from './services/authService';

const result = await authService.login('user@example.com', 'password123');
```

**Available Methods:**
- `authService.login(email, password)` - User login
- `authService.register(userData)` - User registration  
- `authService.getMe()` - Get current user profile
- `authService.refresh(token)` - Refresh user token
- `authService.logout()` - User logout

## API Configuration

### apiConfig.js
Centralized API configuration and endpoints

**Base URL Configuration:**
- Default: `http://localhost:8081`
- Configurable via `REACT_APP_API_BASE_URL` environment variable

**Available Endpoints:**

#### Auth Service
- `AUTH.REGISTER` - 'auth/register'
- `AUTH.LOGIN` - 'auth/login'  
- `AUTH.ME` - 'auth/me'
- `AUTH.REFRESH` - 'auth/refresh'

#### User Service
- `USER.PROFILES` - 'user/profiles'
- `USER.PROFILE_BY_ID(id)` - `user/profiles/${id}`
- `USER.ADDRESSES` - 'user/addresses'
- `USER.ADDRESS_BY_ID(id)` - `user/addresses/${id}`

#### Product Service
- `PRODUCT.CATEGORIES` - 'product/categories'
- `PRODUCT.CATEGORY_BY_ID(id)` - `product/categories/${id}`
- `PRODUCT.PRODUCTS` - 'product/products'
- `PRODUCT.PRODUCT_BY_ID(id)` - `product/products/${id}`
- `PRODUCT.PRODUCT_IMAGES` - 'product/product-images'
- `PRODUCT.PRODUCT_IMAGES_BY_PRODUCT(productId)` - `product/product-images/by-product/${productId}`
- `PRODUCT.PRODUCT_REVIEWS` - 'product/product-reviews'
- `PRODUCT.PRODUCT_UNIT_BY_ID(id)` - `product/product-units/${id}`

#### Cart Service
- `CART.BASE` - 'cart'
- `CART.BY_ID(id)` - `cart/${id}`
- `CART.BY_USER(userId)` - `cart/user/${userId}`
- `CART.CLEAR_USER_CART(userId)` - `cart/user/${userId}`

#### Order Service
- `ORDER.BASE` - 'order'
- `ORDER.BY_ID(id)` - `order/${id}`
- `ORDER.BY_USER(userId)` - `order/user/${userId}`
- `ORDER.UPDATE_STATUS(id)` - `order/${id}/status`
- `ORDER.UPDATE_PAYMENT(id)` - `order/${id}/payment`

#### Shipping Service
- `SHIPMENT.BASE` - 'shipment'
- `SHIPMENT.ALL` - 'shipment/all'
- `SHIPMENT.BY_ID(id)` - `shipment/${id}`
- `SHIPMENT.BY_ORDER(orderId)` - `shipment/order/${orderId}`
- `SHIPMENT.STATUS(id)` - `shipment/status/${id}`
- `SHIPMENT.UPDATE_STATUS(id)` - `shipment/${id}/status`
- `SHIPMENT.TRACKING` - 'shipment/tracking'
- `SHIPMENT.TRACKING_BY_ID(id)` - `shipment/tracking/${id}`
- `SHIPMENT.TRACKING_LATEST(id)` - `shipment/tracking/${id}/latest`

#### Inventory Service
- `INVENTORY.BASE` - 'inventory'
- `INVENTORY.LOW_STOCK` - 'inventory/low-stock'
- `INVENTORY.BY_PRODUCT_VARIANT(productId, variantId)` - `inventory/product/${productId}/variant/${variantId}`
- `INVENTORY.TRANSACTIONS` - 'inventory/transactions'
- `INVENTORY.TRANSACTIONS_BY_PRODUCT(productId)` - `inventory/transactions/product/${productId}`
- `INVENTORY.TRANSACTIONS_BY_VARIANT(productId, variantId)` - `inventory/transactions/product/${productId}/variant/${variantId}`
- `INVENTORY.TRANSACTIONS_RECENT` - 'inventory/transactions/recent'
- `INVENTORY.TRANSACTIONS_BY_REFERENCE(refId)` - `inventory/transactions/reference/${refId}`

#### Payment Service
- `PAYMENT.BASE` - 'payment'
- `PAYMENT.BY_ID(id)` - `payment/${id}`
- `PAYMENT.BY_ORDER(orderId)` - `payment/order/${orderId}`
- `PAYMENT.BY_STATUS(status)` - `payment/status/${status}`
- `PAYMENT.UPDATE_STATUS(id)` - `payment/${id}/status`
- `PAYMENT.REFUNDS` - 'payment/refunds'
- `PAYMENT.REFUND_BY_ID(id)` - `payment/refunds/${id}`
- `PAYMENT.REFUNDS_BY_ORDER(orderId)` - `payment/refunds/order/${orderId}`
- `PAYMENT.REFUNDS_BY_STATUS(status)` - `payment/refunds/status/${status}`

## API Handler

### apiHandler.js
Universal API handler with validation error handling

**Features:**
- Automatic token handling from localStorage
- Centralized error handling
- Validation error parsing from backend
- Network error detection

**HTTP Methods:**
- `api.get(endpoint, options)` - GET requests
- `api.post(endpoint, data, options)` - POST requests
- `api.put(endpoint, data, options)` - PUT requests
- `api.delete(endpoint, options)` - DELETE requests

**Validation Error Handling:**
- Parses backend validation errors in format: `{ field: "field_name", message: "user-friendly message" }`
- Combines multiple field errors into a single error message
- Handles network errors gracefully

## Environment Configuration

**Environment Variables:**
- `REACT_APP_API_BASE_URL` - API base URL (default: http://localhost:8081)
- `REACT_APP_API_VERSION` - API version (default: v1)
- `REACT_APP_API_TIMEOUT` - Request timeout in ms (default: 30000)

## Validation Error Format

The API handler specifically handles validation errors from the backend in this format:
```json
{
  "error": {
    "code": "VAL_2001",
    "message": "Validation failed",
    "details": {
      "fields": [
        {
          "field": "password",
          "message": "password is required"
        }
      ]
    }
  }
}
```

These errors are automatically parsed and converted to user-friendly messages.