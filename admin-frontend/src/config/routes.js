// Route Configuration
export const ROUTES = {
  // Public routes
  PUBLIC: {
    LOGIN: '/login',
    REGISTER: '/register',
    FORGOT_PASSWORD: '/forgot-password',
    RESET_PASSWORD: '/reset-password',
  },
  
  // Protected routes
  PROTECTED: {
    DASHBOARD: '/dashboard',
    PROFILE: '/profile',
    
    // Product routes
    PRODUCTS: {
      LIST: '/products',
      ADD: '/products/add',
      EDIT: (id) => `/products/edit/${id}`,
      IMAGES: (productId) => `/products/${productId}/product-images`,
    },
    
    // Category routes
    CATEGORIES: {
      LIST: '/categories',
      ADD: '/categories/add',
      EDIT: (id) => `/categories/edit/${id}`,
    },
    
    // User routes
    USERS: {
      LIST: '/users',
      ADD: '/users/add',
      EDIT: (id) => `/users/edit/${id}`,
    },
    
    // Order routes
    ORDERS: {
      LIST: '/orders',
      VIEW: (id) => `/orders/${id}`,
      EDIT: (id) => `/orders/edit/${id}`,
    },
    
    // Inventory routes
    INVENTORY: {
      LIST: '/inventory',
      LOW_STOCK: '/inventory/low-stock',
    },
    
    // Settings routes
    SETTINGS: {
      GENERAL: '/settings',
      PROFILE: '/settings/profile',
      SECURITY: '/settings/security',
    },
  },
  
  // Admin routes
  ADMIN: {
    USERS: '/admin/users',
    ROLES: '/admin/roles',
    PERMISSIONS: '/admin/permissions',
  },
  
  // Default routes
  DEFAULT: {
    NOT_FOUND: '/404',
    UNAUTHORIZED: '/403',
  },
};

// Route access levels
export const ACCESS_LEVELS = {
  PUBLIC: 'public',
  AUTHENTICATED: 'authenticated',
  ADMIN: 'admin',
  MODERATOR: 'moderator',
};

// Route metadata
export const ROUTE_METADATA = {
  [ROUTES.PUBLIC.LOGIN]: {
    title: 'Login',
    requiresAuth: false,
  },
  [ROUTES.PUBLIC.REGISTER]: {
    title: 'Register',
    requiresAuth: false,
  },
  [ROUTES.PROTECTED.DASHBOARD]: {
    title: 'Dashboard',
    requiresAuth: true,
  },
  [ROUTES.PROTECTED.PRODUCTS.LIST]: {
    title: 'Products',
    requiresAuth: true,
  },
  [ROUTES.PROTECTED.CATEGORIES.LIST]: {
    title: 'Categories',
    requiresAuth: true,
  },
  [ROUTES.PROTECTED.USERS.LIST]: {
    title: 'Users',
    requiresAuth: true,
  },
  [ROUTES.PROTECTED.ORDERS.LIST]: {
    title: 'Orders',
    requiresAuth: true,
  },
};