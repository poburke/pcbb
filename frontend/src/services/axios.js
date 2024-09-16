// src/services/axios.js
import axios from 'axios';
import keycloak from './keycloak';

// Create an Axios instance
const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api',  // Backend API base URL
  timeout: 1000,
});

// Add a request interceptor
apiClient.interceptors.request.use((config) => {
  // Public routes that don't require authentication
  const publicRoutes = ['/public', '/another-public-route']; // Example public routes

  // Check if the route is public
  if (!publicRoutes.includes(config.url)) {
    // Add the token for protected routes
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
  }

  return config;
}, (error) => {
  return Promise.reject(error);
});

export default apiClient;
