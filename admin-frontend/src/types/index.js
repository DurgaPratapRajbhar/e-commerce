// Type definitions for JavaScript using JSDoc

/**
 * @typedef {Object} User
 * @property {number} id - User ID
 * @property {string} email - User email
 * @property {string} username - User username
 * @property {string} role - User role
 * @property {string} firstName - User first name
 * @property {string} lastName - User last name
 * @property {string} createdAt - Creation timestamp
 * @property {string} updatedAt - Update timestamp
 */

/**
 * @typedef {Object} Product
 * @property {number} id - Product ID
 * @property {string} name - Product name
 * @property {string} description - Product description
 * @property {number} price - Product price
 * @property {number} discount - Discount percentage
 * @property {number} stock - Available stock
 * @property {string} sku - Stock keeping unit
 * @property {string} slug - URL-friendly slug
 * @property {string} status - Product status
 * @property {number} categoryId - Category ID
 * @property {string} brand - Product brand
 * @property {string} createdAt - Creation timestamp
 * @property {string} updatedAt - Update timestamp
 * @property {Object} category - Category object
 * @property {Array<Object>} images - Product images
 * @property {Array<Object>} attributes - Product attributes
 * @property {Array<Object>} variants - Product variants
 */

/**
 * @typedef {Object} Category
 * @property {number} id - Category ID
 * @property {string} name - Category name
 * @property {string} description - Category description
 * @property {string} slug - URL-friendly slug
 * @property {number} parent_id - Parent category ID
 * @property {string} image_url - Category image URL
 * @property {string} createdAt - Creation timestamp
 * @property {string} updatedAt - Update timestamp
 */

/**
 * @typedef {Object} Order
 * @property {number} id - Order ID
 * @property {number} user_id - User ID
 * @property {string} order_number - Order number
 * @property {string} status - Order status
 * @property {string} payment_status - Payment status
 * @property {number} total_amount - Total order amount
 * @property {string} currency - Currency code
 * @property {Array<Object>} items - Order items
 * @property {Object} shipping_address - Shipping address
 * @property {Object} billing_address - Billing address
 * @property {string} created_at - Creation timestamp
 * @property {string} updated_at - Update timestamp
 */

/**
 * @typedef {Object} ApiResponse
 * @property {boolean} success - Whether the request was successful
 * @property {Object|Array} data - Response data
 * @property {string} message - Response message
 * @property {Object} meta - Additional metadata
 */

/**
 * @typedef {Object} ApiError
 * @property {boolean} success - Whether the request was successful
 * @property {string} message - Error message
 * @property {Object} error - Error details
 * @property {Array} fields - Field-specific errors
 */

/**
 * @typedef {Object} AuthResponse
 * @property {boolean} success - Whether the authentication was successful
 * @property {User} user - User object
 * @property {string} token - Authentication token
 * @property {string} message - Response message
 */

/**
 * @typedef {Object} PaginationParams
 * @property {number} page - Current page
 * @property {number} limit - Items per page
 * @property {string} sort - Sort field
 * @property {string} direction - Sort direction (asc/desc)
 */

/**
 * @typedef {Object} FilterParams
 * @property {string} search - Search query
 * @property {number} category - Category ID
 * @property {string} status - Status filter
 * @property {number} minPrice - Minimum price
 * @property {number} maxPrice - Maximum price
 */

/**
 * @typedef {Object} LoginCredentials
 * @property {string} email - User email
 * @property {string} password - User password
 */

/**
 * @typedef {Object} RegisterData
 * @property {string} email - User email
 * @property {string} password - User password
 * @property {string} username - User username
 * @property {string} firstName - User first name
 * @property {string} lastName - User last name
 */