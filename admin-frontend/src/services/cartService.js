import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const cartService = {
  // Cart Management
  getCart: (cartId) => api.get(API_ENDPOINTS.CART.BY_ID(cartId)),
  getCartByUser: (userId) => api.get(API_ENDPOINTS.CART.BY_USER(userId)),
  createCart: (cartData) => api.post(API_ENDPOINTS.CART.BASE, cartData),
  updateCart: (cartId, cartData) => api.put(API_ENDPOINTS.CART.BY_ID(cartId), cartData),
  deleteCart: (cartId) => api.delete(API_ENDPOINTS.CART.BY_ID(cartId)),
  clearUserCart: (userId) => api.delete(API_ENDPOINTS.CART.CLEAR_USER_CART(userId)),
};