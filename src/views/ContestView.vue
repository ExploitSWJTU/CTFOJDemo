<script setup lang="ts">
import { useRoute } from 'vue-router';
import { Flag, Users } from 'lucide-vue-next';

const route = useRoute();

const sidebarItems = [
  { name: '赛事列表', path: '/contest', icon: Flag },
  { name: '我的队伍', path: '/contest/team', icon: Users },
];
</script>

<template>
  <div class="flex gap-6 min-h-[calc(100vh-64px)]">
    <!-- 侧边栏 -->
    <aside class="w-52 shrink-0">
      <div class="sticky top-20 flex flex-col gap-1 rounded-2xl bg-white/80 p-3 shadow-sm border border-slate-200 dark:border-slate-800 dark:bg-slate-900/70">
        <div class="px-3 py-2 mb-2 border-b border-slate-100 dark:border-slate-800">
          <h1 class="text-base font-black text-slate-800 dark:text-slate-100 tracking-tight">
            赛事中心
          </h1>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">
            Contest Center
          </p>
        </div>
        <router-link
          v-for="item in sidebarItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-4 py-3 text-sm font-bold transition-all rounded-xl group"
          :class="
            route.path === item.path
              ? 'bg-blue-600 text-white shadow-lg shadow-blue-200 dark:shadow-none'
              : 'text-slate-500 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-200'
          "
        >
          <component :is="item.icon" :size="18" :class="route.path === item.path ? 'text-white' : 'text-slate-400 group-hover:text-slate-600 dark:group-hover:text-slate-200'" />
          {{ item.name }}
        </router-link>
      </div>
    </aside>

    <!-- 内容区域 -->
    <div class="flex-1 min-w-0">
      <router-view />
    </div>
  </div>
</template>