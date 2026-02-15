<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import {
  LayoutDashboard,
  LayoutGrid,
  Swords,
  Flag,
  MessageSquare,
  Users,
  Bell,
  Settings,
  ShieldAlert,
  Container,
  Group,
} from 'lucide-vue-next';
import AdminHeader from '../components/admin/AdminHeader.vue';

const route = useRoute();

const menuItems = [
  { name: '首页管理', path: '/admin/manage/home-manage', icon: LayoutGrid },
  { name: '训练题目', path: '/admin/manage/training', icon: Swords },
  { name: '赛事管理', path: '/admin/manage/contest', icon: Flag },
  { name: '论坛管理', path: '/admin/manage/forum', icon: MessageSquare },
  { name: '用户管理', path: '/admin/manage/user', icon: Users },
  { name: '队伍管理', path: '/admin/manage/team', icon: Group },
  { name: '容器实例', path: '/admin/manage/instance', icon: Container },
  { name: '系统公告', path: '/admin/manage/announcement', icon: Bell },
  { name: '系统日志', path: '/admin/manage/log', icon: ShieldAlert },
  { name: '全局设置', path: '/admin/manage/setting', icon: Settings },
];

const activeTab = computed(() => {
  if (route.path === '/admin/manage/home') return '/admin/manage/home';
  const item = menuItems.find((item) => route.path.startsWith(item.path));
  return item ? item.path : '';
});
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-slate-950 flex flex-col">
    <AdminHeader />
    
    <div class="flex flex-1">
      <!-- 侧边栏：紧凑型、全高度边框 -->
      <aside class="w-16 shrink-0 border-r border-slate-200 dark:border-slate-800 bg-white dark:bg-slate-900">
        <div class="sticky top-16 flex flex-col items-center py-4 gap-2">
          <!-- 顶部总览 (原 Shield 转换) -->
          <div class="mb-4 pb-4 border-b border-slate-100 dark:border-slate-800 w-full flex justify-center">
            <router-link
              to="/admin/manage/home"
              class="relative flex items-center justify-center w-12 h-12 transition-all group"
              :class="
                activeTab === '/admin/manage/home'
                  ? 'text-blue-600 dark:text-blue-400'
                  : 'text-slate-400 hover:text-slate-900 dark:hover:text-slate-100'
              "
            >
              <div 
                v-if="activeTab === '/admin/manage/home'"
                class="absolute left-0 top-0 bottom-0 w-1 bg-blue-600 dark:bg-blue-400"
              />
              <LayoutDashboard :size="22" class="transition-transform group-hover:scale-110" />
              
              <div class="absolute left-14 px-4 py-2 rounded-xl bg-slate-900 dark:bg-blue-600 text-white text-sm font-bold whitespace-nowrap opacity-0 -translate-x-3 pointer-events-none transition-all group-hover:opacity-100 group-hover:translate-x-0 z-50 shadow-2xl">
                总览 Dashboard
                <div class="absolute top-1/2 -left-1.5 -translate-y-1/2 w-3 h-3 bg-current rotate-45" />
              </div>
            </router-link>
          </div>
          
          <router-link
            v-for="item in menuItems"
            :key="item.path"
            :to="item.path"
            class="relative flex items-center justify-center w-12 h-12 transition-all group"
            :class="
              activeTab === item.path
                ? 'text-blue-600 dark:text-blue-400'
                : 'text-slate-400 hover:text-slate-900 dark:hover:text-slate-100'
            "
          >
            <!-- 选中状态的左侧指示条 -->
            <div 
              v-if="activeTab === item.path"
              class="absolute left-0 top-0 bottom-0 w-1 bg-blue-600 dark:bg-blue-400"
            />
            
            <component 
              :is="item.icon" 
              :size="20" 
              class="transition-transform group-hover:scale-110"
              :class="activeTab === item.path ? 'text-blue-600 dark:text-blue-400' : 'text-slate-400'" 
            />

            <!-- 悬浮文字提示 (Tooltip) - 更大、更清晰 -->
            <div class="absolute left-14 px-4 py-2 rounded-xl bg-slate-900 dark:bg-blue-600 text-white text-sm font-bold whitespace-nowrap opacity-0 -translate-x-3 pointer-events-none transition-all group-hover:opacity-100 group-hover:translate-x-0 z-50 shadow-2xl">
              {{ item.name }}
              <!-- 小三角 -->
              <div class="absolute top-1/2 -left-1.5 -translate-y-1/2 w-3 h-3 bg-current rotate-45" />
            </div>
          </router-link>
        </div>
      </aside>

      <!-- 内容区域：移除最大宽度限制，全屏铺满 -->
      <main class="flex-1 min-w-0 bg-slate-50 dark:bg-slate-950 p-8 overflow-y-auto">
        <router-view />
      </main>
    </div>
  </div>
</template>
