import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const userService = {
  // User Profile API
  getUserProfile: (userId) => api.get(`${API_ENDPOINTS.USER.PROFILES}/${userId}`),
  getUserProfileById: (id) => api.get(`${API_ENDPOINTS.USER.PROFILE_BY_ID(id)}`),
  getAllUserProfiles: (page = 1, limit = 10) => api.get(`${API_ENDPOINTS.ADMIN.USERS}?page=${page}&limit=${limit}`),
  createUserProfile: (profileData) => api.post(API_ENDPOINTS.USER.PROFILES, profileData),
  updateUserProfile: (userId, profileData) => api.put(`${API_ENDPOINTS.USER.PROFILES}/${userId}`, profileData),
  deleteUserProfile: (userId) => api.delete(`${API_ENDPOINTS.USER.PROFILES}/${userId}`),

  // User Address API
  getUserAddresses: (userId) => api.get(`${API_ENDPOINTS.USER.ADDRESSES}/user/${userId}`),
  getAllUserAddresses: () => api.get(API_ENDPOINTS.USER.ADDRESSES),
  getUserAddressById: (id) => api.get(`${API_ENDPOINTS.USER.ADDRESS_BY_ID(id)}`),
  createUserAddress: (addressData) => api.post(API_ENDPOINTS.USER.ADDRESSES, addressData),
  updateUserAddress: (id, addressData) => api.put(`${API_ENDPOINTS.USER.ADDRESS_BY_ID(id)}`, addressData),
  deleteUserAddress: (id) => api.delete(`${API_ENDPOINTS.USER.ADDRESS_BY_ID(id)}`),
  setDefaultAddress: (userId, addressId) => api.put(`${API_ENDPOINTS.USER.ADDRESSES}/user/${userId}/default/${addressId}`),
  getDefaultAddress: (userId) => api.get(`${API_ENDPOINTS.USER.ADDRESSES}/user/${userId}/default`),
};