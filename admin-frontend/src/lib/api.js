// import axios from 'axios';
// import { API_CONFIG } from '../config/apiConfig';
// import { ApiError } from '../utils/apiErrorHandler';

// export const api = axios.create({
//   baseURL: `${API_CONFIG.baseURL}/api/${API_CONFIG.version}`,
//   timeout: API_CONFIG.timeout,
//   withCredentials: true, // 🔥 HttpOnly cookie ke liye
//   headers: { 'Content-Type': 'application/json' },
// });


import axios from 'axios';
import { API_CONFIG } from '../config/apiConfig';
import { ApiError } from '../utils/apiErrorHandler';

export const api = axios.create({
  baseURL: `${API_CONFIG.baseURL}/api/${API_CONFIG.version}`,
  timeout: API_CONFIG.timeout,
  withCredentials: true, // ✅ HttpOnly cookies
  headers: {
    'Content-Type': 'application/json',
  },
});

// 🔥 Response Interceptor
api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    const apiError = new ApiError(error?.response?.data || {
      message: error.message,
      status: error.response?.status,
    });

    return Promise.reject(apiError);
  }
);
