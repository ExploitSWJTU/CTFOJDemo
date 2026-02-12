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
    {
      path: '/admin/manage',
      redirect: '/admin/manage/training',
    },
    {
      path: '/admin/manage/training',
      name: 'adminManageTraining',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/contest',
      name: 'adminManageContest',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/forum',
      name: 'adminManageForum',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/user',
      name: 'adminManageUser',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/team',
      name: 'adminManageTeam',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/instance',
      name: 'adminManageInstance',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/log',
      name: 'adminManageLog',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/setting',
      name: 'adminManageSetting',
      component: () => import('../views/AdminManageView.vue'),
    },
  ],
});

export default router;
