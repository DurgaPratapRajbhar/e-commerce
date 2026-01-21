export const API_CONFIG = {
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  version: import.meta.env.VITE_API_VERSION || 'v1',
  timeout: parseInt(import.meta.env.VITE_API_TIMEOUT) || 30000,
};

// Build complete API URL
export const getApiUrl = (endpoint) => {
  const cleanEndpoint = endpoint.startsWith('/') ? endpoint.slice(1) : endpoint;
  return `${API_CONFIG.baseURL}/api/${API_CONFIG.version}/${cleanEndpoint}`;
};

// API Endpoints
export const API_ENDPOINTS = {
  // Auth Service
  AUTH: {
    REGISTER: 'auth/register',
    LOGIN: 'auth/login',
    LOGOUT: 'auth/logout',
    ME: 'auth/me',
    REFRESH: 'auth/refresh',
  },

  // User Service
  USER: {
    PROFILES: 'user/profiles',
    PROFILE_BY_ID: (id) => `user/profiles/${id}`,
    ADDRESSES: 'user/addresses',
    ADDRESS_BY_ID: (id) => `user/addresses/${id}`,
  },

  // Product Service
  PRODUCT: {
    CATEGORIES: 'product/categories',
    CATEGORY_BY_ID: (id) => `product/categories/${id}`,
    PRODUCTS: 'product/products',
    PRODUCT_BY_ID: (id) => `product/products/${id}`,
    PRODUCT_IMAGES: 'product/product-images',
    PRODUCT_IMAGE_BY_ID: (id) => `product/product-images/${id}`,
    PRODUCT_IMAGES_BY_PRODUCT: (productId) => `product/product-images/by-product/${productId}`,
    PRODUCT_REVIEWS: 'product/product-reviews',
    PRODUCT_REVIEW_BY_ID: (id) => `product/product-reviews/${id}`,
    PRODUCT_UNITS: 'product/product-units',
    PRODUCT_UNIT_BY_ID: (id) => `product/product-units/${id}`,
  },

  // Cart Service
  CART: {
    BASE: 'cart',
    BY_ID: (id) => `cart/${id}`,
    BY_USER: (userId) => `cart/user/${userId}`,
    CLEAR_USER_CART: (userId) => `cart/user/${userId}`,
  },

  // Order Service
  ORDER: {
    BASE: 'order',
    BY_ID: (id) => `order/${id}`,
    BY_USER: (userId) => `order/user/${userId}`,
    UPDATE_STATUS: (id) => `order/${id}/status`,
    UPDATE_PAYMENT: (id) => `order/${id}/payment`,
  },

  // Shipping Service
  SHIPMENT: {
    BASE: 'shipment',
    ALL: 'shipment/all',
    BY_ID: (id) => `shipment/${id}`,
    BY_ORDER: (orderId) => `shipment/order/${orderId}`,
    STATUS: (id) => `shipment/status/${id}`,
    UPDATE_STATUS: (id) => `shipment/${id}/status`,
    TRACKING: 'shipment/tracking',
    TRACKING_BY_ID: (id) => `shipment/tracking/${id}`,
    TRACKING_LATEST: (id) => `shipment/tracking/${id}/latest`,
  },

  // Inventory Service
  INVENTORY: {
    BASE: 'inventory',
    LOW_STOCK: 'inventory/low-stock',
    BY_PRODUCT_VARIANT: (productId, variantId) => `inventory/product/${productId}/variant/${variantId}`,
    TRANSACTIONS: 'inventory/transactions',
    TRANSACTIONS_BY_PRODUCT: (productId) => `inventory/transactions/product/${productId}`,
    TRANSACTIONS_BY_VARIANT: (productId, variantId) => `inventory/transactions/product/${productId}/variant/${variantId}`,
    TRANSACTIONS_RECENT: 'inventory/transactions/recent',
    TRANSACTIONS_BY_REFERENCE: (refId) => `inventory/transactions/reference/${refId}`,
  },

  // Payment Service
  PAYMENT: {
    BASE: 'payment',
    BY_ID: (id) => `payment/${id}`,
    BY_ORDER: (orderId) => `payment/order/${orderId}`,
    BY_STATUS: (status) => `payment/status/${status}`,
    UPDATE_STATUS: (id) => `payment/${id}/status`,
    REFUNDS: 'payment/refunds',
    REFUND_BY_ID: (id) => `payment/refunds/${id}`,
    REFUNDS_BY_ORDER: (orderId) => `payment/refunds/order/${orderId}`,
    REFUNDS_BY_STATUS: (status) => `payment/refunds/status/${status}`,
  },

  // Admin Service
  ADMIN: {
    USERS: 'admin/users',
  },
};