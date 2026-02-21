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
          component: () => import('../views/admin/dashboard/AdminHomeView.vue'),
        },
        {
          path: 'home-manage',
          name: 'adminManageHomeManage',
          component: () => import('../views/admin/home/AdminHomeManageView.vue'),
        },
        {
          path: 'training',
          name: 'adminManageTraining',
          component: () => import('../views/admin/training/AdminTrainingView.vue'),
        },
        {
          path: 'contest',
          name: 'adminManageContest',
          component: () => import('../views/admin/contest/AdminContestView.vue'),
        },
        {
          path: 'contest/create',
          name: 'adminManageContestCreate',
          component: () => import('../views/admin/contest/AdminContestCreateView.vue'),
        },
        {
          path: 'contest/edit/:id',
          name: 'adminManageContestEdit',
          component: () => import('../views/admin/contest/AdminContestEditView.vue'),
        },
        {
          path: 'forum',
          name: 'adminManageForum',
          component: () => import('../views/admin/forum/AdminForumView.vue'),
        },
        {
          path: 'user',
          name: 'adminManageUser',
          component: () => import('../views/admin/user/AdminUserView.vue'),
        },
        {
          path: 'team',
          name: 'adminManageTeam',
          component: () => import('../views/admin/team/AdminTeamView.vue'),
        },
        {
          path: 'instance',
          name: 'adminManageInstance',
          component: () => import('../views/admin/instance/AdminInstanceView.vue'),
        },
        {
          path: 'announcement',
          name: 'adminManageAnnouncement',
          component: () => import('../views/admin/announcement/AdminAnnouncementView.vue'),
        },
        {
          path: 'log',
          name: 'adminManageLog',
          component: () => import('../views/admin/log/AdminLogView.vue'),
        },
        {
          path: 'setting',
          name: 'adminManageSetting',
          component: () => import('../views/admin/setting/AdminSettingView.vue'),
        },
      ],
    },
  ],
});

export default router;