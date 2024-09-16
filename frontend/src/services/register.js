// src/services/auth.js
import axios from 'axios';

export function registerUser(username, password) {
  return axios.post('/api/register', {
    username,
    password
  }).then(response => {
    return response.data;
  }).catch(error => {
    console.error("Registration failed", error);
    throw error;
  });
}
