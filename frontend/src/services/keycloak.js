// src/services/keycloak.js
import Keycloak from "keycloak-js";

const initOptions = {
  url: "http://localhost:8080/auth",  // Keycloak URL
  realm: "<pcbb>",
  clientId: "vue-spa-client",
  onLoad: "check-sso"
};

const keycloak = Keycloak(initOptions);

// Function to initialize Keycloak
export function initKeycloak() {
    return keycloak.init({ onLoad: initOptions.onLoad })
      .then(authenticated => {
        if (authenticated) {
          // User is authenticated, start token refresh
          setupTokenRefresh(60);
        }
      });
  }
  

export function login() {
  return keycloak.login();
}

export function logout() {
  return keycloak.logout();
}

export function refreshToken() {
  return keycloak.updateToken(5)  // Refresh token if it will expire in 5 seconds
    .then(refreshed => {
      if (refreshed) {
        console.log("Token refreshed");
        localStorage.setItem("token", keycloak.token);
      }
    })
    .catch(() => {
      console.error("Failed to refresh token");
    });
}

export function setupTokenRefresh(interval = 60) {
    setInterval(() => {
      if (keycloak.authenticated) {
        keycloak.updateToken(30)
          .then(refreshed => {
            if (refreshed) {
              localStorage.setItem("token", keycloak.token);
            }
          })
          .catch(() => {
            console.log("Failed to refresh token. Will prompt login when accessing a protected route.");
          });
      }
    }, interval * 1000);  // Check every `interval` seconds
  }
  

export default keycloak;
