import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/training',
      name: 'training',
      component: () => import('../views/TrainingView.vue'),
    },
    {
      path: '/contest',
      name: 'contest',
      component: () => import('../views/ContestView.vue'),
    },
    {
      path: '/forum',
      name: 'forum',
      component: () => import('../views/ForumView.vue'),
    },
  ],
});

export default router;
