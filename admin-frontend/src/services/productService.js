import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const productService = {
  // Category Management
  getCategories: () => api.get(API_ENDPOINTS.PRODUCT.CATEGORIES),
  getCategoryById: (categoryId) => api.get(API_ENDPOINTS.PRODUCT.CATEGORY_BY_ID(categoryId)),
  createCategory: (categoryData) => api.post(API_ENDPOINTS.PRODUCT.CATEGORIES, categoryData),
  updateCategory: (categoryId, categoryData) => api.put(API_ENDPOINTS.PRODUCT.CATEGORY_BY_ID(categoryId), categoryData),
  deleteCategory: (categoryId) => api.delete(API_ENDPOINTS.PRODUCT.CATEGORY_BY_ID(categoryId)),

  // Product Management
  getProducts: () => api.get(API_ENDPOINTS.PRODUCT.PRODUCTS),
  getProductById: (productId) => api.get(API_ENDPOINTS.PRODUCT.PRODUCT_BY_ID(productId)),
  createProduct: (productData) => api.post(API_ENDPOINTS.PRODUCT.PRODUCTS, productData),
  updateProduct: (productId, productData) => api.put(API_ENDPOINTS.PRODUCT.PRODUCT_BY_ID(productId), productData),
  deleteProduct: (productId) => api.delete(API_ENDPOINTS.PRODUCT.PRODUCT_BY_ID(productId)),

  // Product Image Management
  getProductImages: () => api.get(API_ENDPOINTS.PRODUCT.PRODUCT_IMAGES),
  getProductImagesByProduct: (productId) => api.get(API_ENDPOINTS.PRODUCT.PRODUCT_IMAGES_BY_PRODUCT(productId)),
  createProductImage: (imageData) => api.post(API_ENDPOINTS.PRODUCT.PRODUCT_IMAGES, imageData),
  updateProductImage: (imageId, imageData) => api.put(API_ENDPOINTS.PRODUCT.PRODUCT_IMAGE_BY_ID(imageId), imageData),
  deleteProductImage: (imageId) => api.delete(API_ENDPOINTS.PRODUCT.PRODUCT_IMAGE_BY_ID(imageId)),

  // Product Unit Management
  getProductUnits: () => api.get(API_ENDPOINTS.PRODUCT.PRODUCT_UNITS),
  getProductUnitById: (unitId) => api.get(API_ENDPOINTS.PRODUCT.PRODUCT_UNIT_BY_ID(unitId)),
  createProductUnit: (unitData) => api.post(API_ENDPOINTS.PRODUCT.PRODUCT_UNITS, unitData),
  updateProductUnit: (unitId, unitData) => api.put(API_ENDPOINTS.PRODUCT.PRODUCT_UNIT_BY_ID(unitId), unitData),
  deleteProductUnit: (unitId) => api.delete(API_ENDPOINTS.PRODUCT.PRODUCT_UNIT_BY_ID(unitId)),
};