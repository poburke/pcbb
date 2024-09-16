import { createRouter, createWebHistory } from 'vue-router'; // Vue 3
import HomePage from '@/views/HomePage.vue';
import keycloak from '@/services/keycloak';  // Import Keycloak instance

const routes = [
  { path: '/', name: 'Home', component: HomePage },
  // You can add protected routes here in the future, like:
  // { path: '/protected', name: 'ProtectedPage', component: ProtectedPage, meta: { requiresAuth: true } }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

// Add navigation guard for protected routes
router.beforeEach((to, from, next) => {
  // Check if the route requires authentication
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!keycloak.authenticated) {
      // If not authenticated, redirect to Keycloak login
      keycloak.login();
    } else {
      // User is authenticated, allow access
      next();
    }
  } else {
    // Route doesn't require authentication, proceed as normal
    next();
  }
});

export default router;
