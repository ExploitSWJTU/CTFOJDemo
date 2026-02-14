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
      component: () => import('../views/AdminManageView.vue'),
      children: [
        {
          path: '',
          redirect: '/admin/manage/training',
        },
        {
          path: 'home',
          name: 'adminManageHome',
          component: () => import('../views/admin/AdminHomeView.vue'),
        },
        {
          path: 'training',
          name: 'adminManageTraining',
          component: () => import('../views/admin/AdminTrainingView.vue'),
        },
        {
          path: 'contest',
          name: 'adminManageContest',
          component: () => import('../views/admin/AdminContestView.vue'),
        },
        {
          path: 'contest/create',
          name: 'adminManageContestCreate',
          component: () => import('../views/admin/AdminContestCreateView.vue'),
        },
        {
          path: 'contest/edit/:id',
          name: 'adminManageContestEdit',
          component: () => import('../views/admin/AdminContestEditView.vue'),
        },
        {
          path: 'forum',
          name: 'adminManageForum',
          component: () => import('../views/admin/AdminForumView.vue'),
        },
        {
          path: 'user',
          name: 'adminManageUser',
          component: () => import('../views/admin/AdminUserView.vue'),
        },
        {
          path: 'team',
          name: 'adminManageTeam',
          component: () => import('../views/admin/AdminTeamView.vue'),
        },
        {
          path: 'instance',
          name: 'adminManageInstance',
          component: () => import('../views/admin/AdminInstanceView.vue'),
        },
        {
          path: 'announcement',
          name: 'adminManageAnnouncement',
          component: () => import('../views/admin/AdminAnnouncementView.vue'),
        },
        {
          path: 'log',
          name: 'adminManageLog',
          component: () => import('../views/admin/AdminLogView.vue'),
        },
        {
          path: 'setting',
          name: 'adminManageSetting',
          component: () => import('../views/admin/AdminSettingView.vue'),
        },
      ],
    },
  ],
});

export default router;