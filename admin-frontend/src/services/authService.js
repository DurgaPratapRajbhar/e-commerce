import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const authService = {
  register: (userData) => api.post(API_ENDPOINTS.AUTH.REGISTER, userData),
  login: (credentials) => api.post(API_ENDPOINTS.AUTH.LOGIN, credentials),
  getMe: () => api.get(API_ENDPOINTS.AUTH.ME),
  refresh: () => api.post(API_ENDPOINTS.AUTH.REFRESH),
  logout: () => api.post(API_ENDPOINTS.AUTH.LOGOUT),
};