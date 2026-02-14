<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import {
  Terminal,
  Search,
  Activity,
  Bell,
  Sun,
  Moon,
  Settings,
  Edit,
  LogOut,
  Cpu,
  Database,
  Globe,
} from 'lucide-vue-next';

const searchText = ref('');
const isDark = ref(false);
const route = useRoute();

const toggleDarkMode = () => {
  isDark.value = !isDark.value;
  updateTheme();
};

const updateTheme = () => {
  const el = document.documentElement;
  const body = document.body;
  
  if (isDark.value) {
    el.classList.add('dark');
    body.classList.add('dark');
    el.setAttribute('arco-theme', 'dark');
    body.setAttribute('arco-theme', 'dark');
    el.style.colorScheme = 'dark';
    localStorage.setItem('theme', 'dark');
  } else {
    el.classList.remove('dark');
    body.classList.remove('dark');
    el.removeAttribute('arco-theme');
    body.removeAttribute('arco-theme');
    el.style.colorScheme = 'light';
    localStorage.setItem('theme', 'light');
  }
};

onMounted(() => {
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme === 'dark' || (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    isDark.value = true;
  } else {
    isDark.value = false;
  }
  updateTheme();
});

// Mock Server Load Data
const serverStats = ref([
  { label: 'CPU', value: 45, icon: Cpu, color: 'text-blue-500' },
  { label: 'RAM', value: 62, icon: Activity, color: 'text-emerald-500' },
  { label: 'DB', value: 12, icon: Database, color: 'text-amber-500' },
]);
</script>

<template>
  <header
    class="sticky top-0 z-50 w-full border-b border-slate-200 bg-white/95 backdrop-blur-md transition-colors dark:border-slate-800 dark:bg-slate-900/95"
  >
    <div class="mx-auto flex h-16 max-w-384 items-center justify-between px-6">
      <!-- 左侧：Logo + 管理后台标识 -->
      <div class="flex items-center gap-6">
        <router-link to="/" class="flex cursor-pointer items-center gap-2.5">
          <div
            class="flex h-8 w-8 items-center justify-center rounded-lg bg-slate-900 text-white shadow-lg dark:bg-blue-600"
          >
            <Terminal :size="20" />
          </div>
          <div class="flex flex-col">
            <span class="text-base font-black tracking-tight text-slate-800 dark:text-slate-100 leading-tight">
              SWJTU <span class="text-blue-600">Admin</span>
            </span>
            <span class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">Management</span>
          </div>
        </router-link>

        <div class="h-8 w-px bg-slate-200 dark:bg-slate-800 mx-2" />

        <!-- 管理页搜索 -->
        <div class="group relative">
          <Search class="absolute top-1/2 left-3 -translate-y-1/2 text-slate-400" :size="16" />
          <input
            v-model="searchText"
            type="text"
            placeholder="搜索管理项、用户、容器..."
            class="h-9 w-64 rounded-lg border border-slate-200 bg-slate-50 pr-4 pl-9 text-xs transition-all focus:w-80 focus:border-blue-400 focus:bg-white focus:outline-none dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:focus:border-blue-500 dark:focus:bg-slate-900"
          />
        </div>
      </div>

      <!-- 右侧：服务器负载 + 工具 + 头像 -->
      <div class="flex items-center gap-6">
        <!-- 服务器负载 -->
        <div class="hidden xl:flex items-center gap-4">
          <div v-for="stat in serverStats" :key="stat.label" class="flex items-center gap-2">
            <component :is="stat.icon" :size="14" :class="stat.color" />
            <div class="flex flex-col">
              <span class="text-[9px] font-bold text-slate-400 uppercase leading-none mb-0.5">{{ stat.label }}</span>
              <div class="flex items-center gap-1.5">
                <div class="h-1 w-12 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                  <div class="h-full bg-current transition-all" :class="stat.color" :style="{ width: `${stat.value}%` }" />
                </div>
                <span class="text-[10px] font-mono font-bold text-slate-600 dark:text-slate-400">{{ stat.value }}%</span>
              </div>
            </div>
          </div>
        </div>

        <div class="h-6 w-px bg-slate-200 dark:bg-slate-800 mx-1" />

        <div class="flex items-center gap-2">
          <!-- 通知 -->
          <button class="relative rounded-lg p-2 text-slate-400 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:hover:bg-slate-800 dark:hover:text-blue-400">
            <Bell :size="20" />
            <span class="absolute top-2 right-2 flex h-2 w-2">
              <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75" />
              <span class="relative inline-flex rounded-full h-2 w-2 bg-red-500" />
            </span>
          </button>

          <!-- 黑夜模式切换 -->
          <button
            class="rounded-lg p-2 text-slate-400 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:hover:bg-slate-800 dark:hover:text-blue-400"
            @click="toggleDarkMode"
          >
            <Sun v-if="isDark" :size="20" />
            <Moon v-else :size="20" />
          </button>

          <!-- 设置 -->
          <button class="rounded-lg p-2 text-slate-400 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:hover:bg-slate-800 dark:hover:text-blue-400">
            <Settings :size="20" />
          </button>
        </div>

        <!-- 头像 -->
        <a-dropdown trigger="click" position="br">
          <div
            class="flex cursor-pointer items-center gap-2 rounded-full border border-slate-200 p-1 pr-3 transition-all hover:bg-slate-50 dark:border-slate-700 dark:hover:bg-slate-800"
          >
            <div class="h-7 w-7 overflow-hidden rounded-full border border-slate-200 dark:border-slate-700">
              <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" alt="avatar" />
            </div>
            <span class="text-xs font-bold text-slate-700 dark:text-slate-200">Admin</span>
          </div>
          <template #content>
            <div
              class="rounded-select w-48 border border-slate-200 bg-white py-1 shadow-xl dark:border-slate-800 dark:bg-slate-900"
            >
              <div class="px-3 py-2">
                <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-1">
                  Signed in as
                </p>
                <p class="text-sm font-bold text-slate-700 dark:text-slate-200">
                  System Admin
                </p>
              </div>
              <div class="my-1 h-px bg-slate-100 dark:bg-slate-800" />
              <a
                href="#"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium text-slate-600 hover:bg-slate-50 hover:text-slate-900 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-100"
              >
                <Edit :size="16" />
                修改资料
              </a>
              <div class="my-1 h-px bg-slate-100 dark:bg-slate-800" />
              <a
                href="#"
                class="flex items-center gap-2 px-3 py-2 text-sm font-medium text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20"
              >
                <LogOut :size="16" />
                退出管理
              </a>
            </div>
          </template>
        </a-dropdown>
      </div>
    </div>
  </header>
</template>
