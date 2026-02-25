import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // 登录页（独立，无布局）
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/share/LoginView.vue'),
    },
    // 注册页（独立，无布局）
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/share/RegisterView.vue'),
    },
    // 重置密码页（独立，无布局）
    {
      path: '/reset-password',
      name: 'resetPassword',
      component: () => import('../views/share/ResetPasswordView.vue'),
    },
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
          path: 'forum/create',
          name: 'forumCreate',
          component: () => import('../views/user/forum/ForumPostCreateView.vue'),
        },
        {
          path: 'forum/:id',
          name: 'forumPostDetail',
          component: () => import('../views/user/forum/ForumPostDetailView.vue'),
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
        {
          path: 'settings',
          name: 'settings',
          component: () => import('../views/user/settings/SettingsView.vue'),
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
          redirect: '/admin/manage/dashboard',
        },
        {
          path: 'dashboard',
          name: 'adminManageDashboard',
          component: () => import('../views/admin/dashboard/AdminHomeView.vue'),
        },
        {
          path: 'home',
          name: 'adminManageHome',
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
          path: 'forum/create',
          name: 'adminManageForumCreate',
          component: () => import('../views/admin/forum/AdminForumCreateView.vue'),
        },
        {
          path: 'forum/edit/:id',
          name: 'adminManageForumEdit',
          component: () => import('../views/admin/forum/AdminForumEditView.vue'),
        },
        {
          path: 'forum/comments/:postId',
          name: 'adminManageForumComments',
          component: () => import('../views/admin/forum/AdminForumCommentsView.vue'),
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
    // 管理员个人设置（与 /settings 内容相同）
    {
      path: '/admin/settings',
      component: () => import('../layouts/AdminLayout.vue'),
      children: [
        {
          path: '',
          name: 'adminSettings',
          component: () => import('../views/user/settings/SettingsView.vue'),
        },
      ],
    },
  ],
});

export default router;