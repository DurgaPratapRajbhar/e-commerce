// Environment Configuration
export const ENV = {
  // Current environment
  NODE_ENV: process.env.NODE_ENV || 'development',
  
  // API Configuration
  API_BASE_URL: process.env.REACT_APP_API_BASE_URL || 'http://localhost:8001'
  API_VERSION: process.env.REACT_APP_API_VERSION || 'v1',
  API_TIMEOUT: parseInt(process.env.REACT_APP_API_TIMEOUT) || 30000,
  
  // Authentication
  JWT_COOKIE_NAME: process.env.REACT_APP_JWT_COOKIE_NAME || 'auth_token',
  
  // Feature flags
  ENABLE_ANALYTICS: process.env.REACT_APP_ENABLE_ANALYTICS === 'true',
  ENABLE_LOGGING: process.env.REACT_APP_ENABLE_LOGGING !== 'false',
  
  // Third-party services
  SENTRY_DSN: process.env.REACT_APP_SENTRY_DSN || null,
  GOOGLE_ANALYTICS_ID: process.env.REACT_APP_GOOGLE_ANALYTICS_ID || null,
  
  // Development flags
  MOCK_API: process.env.REACT_APP_MOCK_API === 'true',
  DEBUG_MODE: process.env.NODE_ENV === 'development',
};

// Environment-specific configurations
export const ENV_CONFIG = {
  development: {
    API_BASE_URL: 'http://localhost:8081',
    ENABLE_LOGGING: true,
    DEBUG_MODE: true,
  },
  
  production: {
    API_BASE_URL: process.env.REACT_APP_API_BASE_URL || 'https://api.yourdomain.com',
    ENABLE_LOGGING: true,
    DEBUG_MODE: false,
  },
  
  test: {
    API_BASE_URL: 'http://localhost:8081',
    ENABLE_LOGGING: false,
    DEBUG_MODE: true,
  },
};

// Get environment-specific config
export const getEnvConfig = () => {
  const env = ENV.NODE_ENV;
  return {
    ...ENV,
    ...(ENV_CONFIG[env] || {}),
  };
};

// Check if running in specific environment
export const isDevelopment = () => ENV.NODE_ENV === 'development';
export const isProduction = () => ENV.NODE_ENV === 'production';
export const isTest = () => ENV.NODE_ENV === 'test';