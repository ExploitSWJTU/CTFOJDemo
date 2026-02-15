import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // 用户端路由
    {
      path: '/',
      component: () => import('../layouts/UserLayout.vue'),
      children: [
        {
          path: '',
          name: 'home',
          component: () => import('../views/user/home/HomeView.vue'),
        },
        {
          path: 'training',
          name: 'training',
          component: () => import('../views/user/training/TrainingView.vue'),
        },
        {
          path: 'contest',
          name: 'contest',
          component: () => import('../views/user/contest/ContestListView.vue'),
        },
        {
          path: 'contest/:id',
          name: 'contestDetail',
          component: () => import('../views/user/contest/ContestDetailView.vue'),
        },
        {
          path: 'team',
          name: 'team',
          component: () => import('../views/user/team/TeamView.vue'),
        },
        {
          path: 'forum',
          name: 'forum',
          component: () => import('../views/user/forum/ForumView.vue'),
        },
        {
          path: 'announcement',
          name: 'announcement',
          component: () => import('../views/user/announcement/AnnouncementListView.vue'),
        },
        {
          path: 'announcement/:id',
          name: 'announcementDetail',
          component: () => import('../views/user/announcement/AnnouncementDetailView.vue'),
        },
      ],
    },
    // 管理端路由
    {
      path: '/admin/manage',
      component: () => import('../layouts/AdminLayout.vue'),
      children: [
        {
          path: '',
          redirect: '/admin/manage/home',
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