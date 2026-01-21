import { useState, useEffect, useContext, createContext } from 'react';
import { authService } from '../services/authService';

const AuthContext = createContext();

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const checkAuthStatus = async () => {
      try {
        // Try to get user info to verify if token is still valid
        const response = await authService.getMe();
        if (response) {
          setUser(response);
          setIsAuthenticated(true);
        }
      } catch (error) {
        console.error('Auth check failed:', error);
        setUser(null);
        setIsAuthenticated(false);
      } finally {
        setLoading(false);
      }
    };

    checkAuthStatus();
  }, []);

  const login = async (credentials, setErrorCallback) => {
    try {
      // credentials should have email and password properties
      const response = await authService.login(credentials);
      if (response?.user || response?.success) {
        // Update user state
        const userData = response.user || response;
        setUser(userData);
        setIsAuthenticated(true);
        return { success: true, data: response };
      } else {
        return { success: false, error: 'Login failed: Invalid response' };
      }
    } catch (error) {
      console.error('Login error:', error);
      if (setErrorCallback && typeof setErrorCallback === 'function') {
        setErrorCallback(error.message || 'Login failed');
      }
      return { success: false, error: error.message || 'Login failed' };
    }
  };

  const logout = async () => {
    try {
      await authService.logout();
      setUser(null);
      setIsAuthenticated(false);
    } catch (error) {
      console.error('Logout error:', error);
      // Even if logout API fails, clear local state
      setUser(null);
      setIsAuthenticated(false);
    }
  };

  const register = async (userData, setErrorCallback) => {
    try {
      const response = await authService.register(userData);
      // After successful registration, update the auth state to reflect the new user
      if (response?.user) {
        setUser(response.user);
        setIsAuthenticated(true);
      }
      return { success: true, data: response };
    } catch (error) {
      console.error('Registration error:', error);
      if (setErrorCallback && typeof setErrorCallback === 'function') {
        setErrorCallback(error.message || 'Registration failed');
      }
      return { success: false, error: error.message || 'Registration failed' };
    }
  };

  const value = {
    user,
    loading,
    isAuthenticated,
    login,
    logout,
    register,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};