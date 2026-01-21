export const loginAPI = async (credentials) => {
  const API_BASE_URL = import.meta.env.VITE_AUTH_SERVICE || 'http://localhost:8001';
  const response = await fetch(`${API_BASE_URL}/api/v1/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(credentials),
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    
    // Handle validation errors from the backend
    if (response.status === 400 && errorData.error && errorData.error.details && errorData.error.details.fields) {
      const validationErrors = errorData.error.details.fields;
      const errorMessages = validationErrors.map(field => `${field.field}: ${field.message}`).join(', ');
      throw new Error(errorMessages);
    }
    
    // Handle other error cases
    if (errorData.message) {
      throw new Error(errorData.message);
    }
    
    throw new Error('Invalid credentials');
  }

  return response.json();
};

export const registerAPI = async (userData) => {
  const API_BASE_URL = import.meta.env.VITE_AUTH_SERVICE || 'http://localhost:8001';
  const response = await fetch(`${API_BASE_URL}/api/v1/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(userData),
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    
    // Handle validation errors from the backend
    if (response.status === 400 && errorData.error && errorData.error.details && errorData.error.details.fields) {
      const validationErrors = errorData.error.details.fields;
      const errorMessages = validationErrors.map(field => `${field.field}: ${field.message}`).join(', ');
      throw new Error(errorMessages);
    }
    
    // Handle other error cases
    if (errorData.message) {
      throw new Error(errorData.message);
    }
    
    throw new Error('Registration failed');
  }

  return response.json();
};
