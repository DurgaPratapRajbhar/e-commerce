// Permission Configuration
export const PERMISSIONS = {
  // User permissions
  USER: {
    READ: 'user:read',
    WRITE: 'user:write',
    DELETE: 'user:delete',
    MANAGE: 'user:manage',
  },
  
  // Product permissions
  PRODUCT: {
    READ: 'product:read',
    WRITE: 'product:write',
    DELETE: 'product:delete',
    MANAGE: 'product:manage',
  },
  
  // Category permissions
  CATEGORY: {
    READ: 'category:read',
    WRITE: 'category:write',
    DELETE: 'category:delete',
    MANAGE: 'category:manage',
  },
  
  // Order permissions
  ORDER: {
    READ: 'order:read',
    WRITE: 'order:write',
    DELETE: 'order:delete',
    MANAGE: 'order:manage',
  },
  
  // Inventory permissions
  INVENTORY: {
    READ: 'inventory:read',
    WRITE: 'inventory:write',
    DELETE: 'inventory:delete',
    MANAGE: 'inventory:manage',
  },
  
  // Payment permissions
  PAYMENT: {
    READ: 'payment:read',
    WRITE: 'payment:write',
    DELETE: 'payment:delete',
    MANAGE: 'payment:manage',
  },
  
  // System permissions
  SYSTEM: {
    READ: 'system:read',
    WRITE: 'system:write',
    DELETE: 'system:delete',
    MANAGE: 'system:manage',
  },
};

// Role-based permission mapping
export const ROLE_PERMISSIONS = {
  admin: [
    // All permissions for admin
    ...Object.values(PERMISSIONS).flatMap(role => Object.values(role)),
  ],
  
  moderator: [
    PERMISSIONS.USER.READ,
    PERMISSIONS.USER.WRITE,
    PERMISSIONS.PRODUCT.READ,
    PERMISSIONS.PRODUCT.WRITE,
    PERMISSIONS.CATEGORY.READ,
    PERMISSIONS.CATEGORY.WRITE,
    PERMISSIONS.ORDER.READ,
    PERMISSIONS.ORDER.WRITE,
    PERMISSIONS.INVENTORY.READ,
  ],
  
  user: [
    PERMISSIONS.USER.READ,
    PERMISSIONS.PRODUCT.READ,
    PERMISSIONS.CATEGORY.READ,
    PERMISSIONS.ORDER.READ,
  ],
  
  guest: [
    // Only public permissions
  ],
};

// Permission utilities
export const hasPermission = (userPermissions, requiredPermission) => {
  if (!userPermissions || !requiredPermission) return false;
  
  // Check if user has the exact permission or admin permission
  return userPermissions.includes(requiredPermission) || 
         userPermissions.includes(PERMISSIONS.SYSTEM.MANAGE);
};

export const hasAnyPermission = (userPermissions, requiredPermissions) => {
  if (!userPermissions || !requiredPermissions) return false;
  
  return requiredPermissions.some(permission => 
    hasPermission(userPermissions, permission)
  );
};

export const hasAllPermissions = (userPermissions, requiredPermissions) => {
  if (!userPermissions || !requiredPermissions) return false;
  
  return requiredPermissions.every(permission => 
    hasPermission(userPermissions, permission)
  );
};

// Permission-based route access
export const ROUTE_PERMISSIONS = {
  '/users': [PERMISSIONS.USER.READ],
  '/users/add': [PERMISSIONS.USER.WRITE],
  '/users/edit/:id': [PERMISSIONS.USER.WRITE],
  
  '/products': [PERMISSIONS.PRODUCT.READ],
  '/products/add': [PERMISSIONS.PRODUCT.WRITE],
  '/products/edit/:id': [PERMISSIONS.PRODUCT.WRITE],
  
  '/categories': [PERMISSIONS.CATEGORY.READ],
  '/categories/add': [PERMISSIONS.CATEGORY.WRITE],
  '/categories/edit/:id': [PERMISSIONS.CATEGORY.WRITE],
  
  '/orders': [PERMISSIONS.ORDER.READ],
  '/orders/:id': [PERMISSIONS.ORDER.READ],
  '/orders/edit/:id': [PERMISSIONS.ORDER.WRITE],
  
  '/inventory': [PERMISSIONS.INVENTORY.READ],
  '/inventory/low-stock': [PERMISSIONS.INVENTORY.READ],
  
  '/admin': [PERMISSIONS.SYSTEM.MANAGE],
  '/admin/users': [PERMISSIONS.SYSTEM.MANAGE],
  '/admin/roles': [PERMISSIONS.SYSTEM.MANAGE],
  '/admin/permissions': [PERMISSIONS.SYSTEM.MANAGE],
};