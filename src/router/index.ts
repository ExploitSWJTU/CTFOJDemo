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
      path: '/team',
      name: 'team',
      component: () => import('../views/TeamView.vue'),
    },
    {
      path: '/forum',
      name: 'forum',
      component: () => import('../views/ForumView.vue'),
    },
    {
      path: '/announcement',
      name: 'announcement',
      component: () => import('../views/AnnouncementListView.vue'),
    },
    {
      path: '/announcement/:id',
      name: 'announcementDetail',
      component: () => import('../views/AnnouncementDetailView.vue'),
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
      path: '/admin/manage/home',
      name: 'adminManageHome',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/contest',
      name: 'adminManageContest',
      component: () => import('../views/AdminManageView.vue'),
    },
    {
      path: '/admin/manage/contest/create',
      name: 'adminManageContestCreate',
      component: () => import('../views/AdminContestCreateView.vue'),
    },
    {
      path: '/admin/manage/contest/edit/:id',
      name: 'adminManageContestEdit',
      component: () => import('../views/AdminContestEditView.vue'),
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
      path: '/admin/manage/announcement',
      name: 'adminManageAnnouncement',
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
