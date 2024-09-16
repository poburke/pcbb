// src/main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import { initKeycloak } from './services/keycloak';

initKeycloak().then(() => {
  createApp(App)
    .use(router)
    .mount('#app');
});
