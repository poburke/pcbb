import { createRouter, createWebHistory } from 'vue-router'; // Vue 3
import HomePage from '@/views/HomePage.vue';

const routes = [
  { path: '/', name: 'Home', component: HomePage }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
