<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import {
  Home,
  Swords,
  Flag,
  MessageSquare,
  Users,
  Box,
  Bell,
  Settings,
  ShieldAlert,
} from 'lucide-vue-next';
import AdminHeader from '../components/admin/AdminHeader.vue';

const route = useRoute();

const menuItems = [
  { name: '管理首页', path: '/admin/manage/home', icon: Home },
  { name: '训练题目', path: '/admin/manage/training', icon: Swords },
  { name: '赛事管理', path: '/admin/manage/contest', icon: Flag },
  { name: '论坛管理', path: '/admin/manage/forum', icon: MessageSquare },
  { name: '用户管理', path: '/admin/manage/user', icon: Users },
  { name: '队伍管理', path: '/admin/manage/team', icon: Users },
  { name: '容器实例', path: '/admin/manage/instance', icon: Box },
  { name: '系统公告', path: '/admin/manage/announcement', icon: Bell },
  { name: '系统日志', path: '/admin/manage/log', icon: ShieldAlert },
  { name: '全局设置', path: '/admin/manage/setting', icon: Settings },
];

const activeTab = computed(() => {
  const item = menuItems.find((item) => route.path.startsWith(item.path));
  return item ? item.path : '/admin/manage/home';
});
</script>

<template>
  <div class="min-h-screen bg-slate-50 dark:bg-slate-950">
    <AdminHeader />
    
    <div class="mx-auto flex max-w-384 gap-6 p-6">
      <!-- 侧边栏 -->
      <aside class="w-52 shrink-0">
        <div class="sticky top-20 flex flex-col gap-1 rounded-2xl bg-white/80 p-3 shadow-sm border border-slate-200 dark:border-slate-800 dark:bg-slate-900/70">
          <div class="px-3 py-2 mb-2 border-b border-slate-100 dark:border-slate-800">
            <h1 class="text-base font-black text-slate-800 dark:text-slate-100 tracking-tight">
              管理控制台
            </h1>
            <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">
              Control Panel
            </p>
          </div>
          <router-link
            v-for="item in menuItems"
            :key="item.path"
            :to="item.path"
            class="flex items-center gap-3 px-4 py-3 text-sm font-bold transition-all rounded-xl group"
            :class="
              activeTab === item.path
                ? 'bg-blue-600 text-white shadow-lg shadow-blue-200 dark:shadow-none'
                : 'text-slate-500 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-200'
            "
          >
            <component :is="item.icon" :size="18" :class="activeTab === item.path ? 'text-white' : 'text-slate-400 group-hover:text-slate-600 dark:group-hover:text-slate-200'" />
            {{ item.name }}
          </router-link>
        </div>
      </aside>

      <!-- 内容区域 -->
      <main class="flex-1 min-w-0">
        <router-view />
      </main>
    </div>
  </div>
</template>
