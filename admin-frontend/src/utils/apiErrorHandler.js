//D:\PROJECT\e-commerce\admin-frontend\src\utils\apiErrorHandler.js

import React from 'react';

export class ApiError extends Error {
  constructor(errorData) {
    super(errorData?.error?.message || 'An error occurred');
    this.name = 'ApiError';
    this.code = errorData?.error?.code;
    this.details = errorData?.error?.details;
    this.timestamp = errorData?.timestamp;
    this.requestId = errorData?.request_id;
    this.fullResponse = errorData;
  }
}

/**
 * Extract user-friendly error messages from API response
 */
export const getErrorMessages = (errorResponse) => {
  const fields = errorResponse?.error?.details?.fields;
  
  if (!fields || !Array.isArray(fields)) {
    return [errorResponse?.error?.message || 'Something went wrong'];
  }
  
  return fields.map(field => field.message);
};

/**
 * Get field-specific errors as object
 */
export const getFieldErrors = (errorResponse) => {
  const fields = errorResponse?.error?.details?.fields;
  
  if (!fields || !Array.isArray(fields)) return {};
  
  return fields.reduce((acc, field) => {
    acc[field.field] = field.message;
    return acc;
  }, {});
};

/**
 * Main API call wrapper with error handling
 */
export const apiCall = async (url, options = {}) => {
  try {
    const response = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
    });

    const data = await response.json();

    // Check if response indicates failure
    if (!response.ok || data.success === false) {
      throw new ApiError(data);
    }

    return data;
  } catch (error) {
    // If it's already an ApiError, throw it
    if (error instanceof ApiError) {
      throw error;
    }

    // Handle network or other errors
    throw new ApiError({
      success: false,
      error: {
        code: 'NETWORK_ERROR',
        message: error.message || 'Network error occurred',
      },
    });
  }
};

/**
 * Toast/Alert helper - customize based on your UI library
 */
export const showErrorToast = (errorResponse, toastFunction) => {
  const messages = getErrorMessages(errorResponse);
  messages.forEach(msg => toastFunction?.(msg));
};

/**
 * React Hook for API calls with error handling
 */
export const useApiCall = () => {
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState(null);

  const execute = async (apiFunction, onSuccess, onError) => {
    setLoading(true);
    setError(null);

    try {
      const result = await apiFunction();
      onSuccess?.(result);
      return result;
    } catch (err) {
      setError(err instanceof ApiError ? err.fullResponse : null);
      onError?.(err);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { execute, loading, error, fieldErrors: getFieldErrors(error) };
};

// Example usage functions
export const handleApiError = (error, setFieldErrors, showToast) => {
  if (error instanceof ApiError) {
    const fieldErrors = getFieldErrors(error.fullResponse);
    const errorMessages = getErrorMessages(error.fullResponse);
    
    // Set field-specific errors for form
    if (Object.keys(fieldErrors).length > 0) {
      setFieldErrors?.(fieldErrors);
    }
    
    // Show general error toast
    if (errorMessages.length > 0 && showToast) {
      errorMessages.forEach(msg => showToast(msg, 'error'));
    }
    
    return { fieldErrors, messages: errorMessages };
  }
  
  // Handle unknown errors
  showToast?.('An unexpected error occurred', 'error');
  return { fieldErrors: {}, messages: ['An unexpected error occurred'] };
};