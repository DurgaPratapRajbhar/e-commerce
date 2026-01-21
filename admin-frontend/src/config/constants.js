// Application Constants
export const APP_NAME = 'E-commerce Admin Panel';
export const APP_VERSION = '1.0.0';
export const API_STATUS = {
  IDLE: 'idle',
  LOADING: 'loading',
  SUCCEEDED: 'succeeded',
  FAILED: 'failed',
};

export const PERMISSIONS = {
  READ_USERS: 'read:users',
  WRITE_USERS: 'write:users',
  READ_PRODUCTS: 'read:products',
  WRITE_PRODUCTS: 'write:products',
  READ_ORDERS: 'read:orders',
  WRITE_ORDERS: 'write:orders',
  READ_INVENTORY: 'read:inventory',
  WRITE_INVENTORY: 'write:inventory',
};

export const ROLES = {
  ADMIN: 'admin',
  MODERATOR: 'moderator',
  USER: 'user',
};

export const STATUS = {
  ACTIVE: 'active',
  INACTIVE: 'inactive',
  PENDING: 'pending',
  DRAFT: 'draft',
  ARCHIVED: 'archived',
};

export const ROUTES = {
  DASHBOARD: '/dashboard',
  LOGIN: '/login',
  LOGOUT: '/logout',
  PRODUCTS: {
    LIST: '/products',
    ADD: '/products/add',
    EDIT: (id) => `/products/edit/${id}`,
    IMAGES: (productId) => `/products/${productId}/product-images`,
  },
  CATEGORIES: {
    LIST: '/categories',
    ADD: '/categories/add',
    EDIT: (id) => `/categories/edit/${id}`,
  },
  USERS: {
    LIST: '/users',
    ADD: '/users/add',
    EDIT: (id) => `/users/edit/${id}`,
  },
  ORDERS: {
    LIST: '/orders',
    VIEW: (id) => `/orders/${id}`,
  },
};

export const MESSAGES = {
  SUCCESS: {
    LOGIN: 'Login successful',
    LOGOUT: 'Logout successful',
    SAVE: 'Data saved successfully',
    DELETE: 'Item deleted successfully',
  },
  ERROR: {
    LOGIN: 'Login failed',
    NETWORK: 'Network error occurred',
    VALIDATION: 'Please check your input',
    UNAUTHORIZED: 'Unauthorized access',
  },
};