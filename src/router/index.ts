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
      path: '/contest/:id',
      name: 'contestDetail',
      component: () => import('../views/ContestDetailView.vue'),
    },
    {
      path: '/forum',
      name: 'forum',
      component: () => import('../views/ForumView.vue'),
    },
    {
      path: '/admin/contest',
      name: 'adminContest',
      component: () => import('../views/AdminContestView.vue'),
    },
    {
      path: '/admin/contest/:id',
      name: 'adminContestDetail',
      component: () => import('../views/ContestDetailView.vue'),
    },
    {
      path: '/admin/contest/edit/:id',
      name: 'adminContestEdit',
      component: () => import('../views/AdminContestEditView.vue'),
    },
    {
      path: '/admin/contest/create',
      name: 'adminContestCreate',
      component: () => import('../views/AdminContestCreateView.vue'),
    },
  ],
});

export default router;
