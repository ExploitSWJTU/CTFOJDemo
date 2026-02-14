<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  LayoutDashboard,
  Swords,
  Flag,
  MessageSquare,
  Users,
  Shield,
  Box,
  Megaphone,
  ScrollText,
  Settings,
} from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()

// 标签页选项
const tabs = [
  { key: 'home', label: '首页管理', path: '/admin/manage/home', icon: LayoutDashboard },
  { key: 'training', label: '训练管理', path: '/admin/manage/training', icon: Swords },
  { key: 'contest', label: '赛事管理', path: '/admin/manage/contest', icon: Flag },
  { key: 'forum', label: '论坛管理', path: '/admin/manage/forum', icon: MessageSquare },
  { key: 'user', label: '用户管理', path: '/admin/manage/user', icon: Users },
  { key: 'team', label: '队伍管理', path: '/admin/manage/team', icon: Shield },
  { key: 'instance', label: '实例管理', path: '/admin/manage/instance', icon: Box },
  { key: 'announcement', label: '公告管理', path: '/admin/manage/announcement', icon: Megaphone },
  { key: 'log', label: '系统日志', path: '/admin/manage/log', icon: ScrollText },
  { key: 'setting', label: '系统设置', path: '/admin/manage/setting', icon: Settings },
]

// 根据当前路由确定激活的标签页
const activeTab = computed(() => {
  const path = route.path
  const tab = tabs.find((t) => path.startsWith(t.path))
  return tab?.key || 'training'
})

const handleTabClick = (path: string) => {
  router.push(path)
}
</script>

<template>
  <div class="flex gap-6 min-h-[calc(100vh-64px)]">
    <!-- 管理页侧边栏 -->
    <aside class="w-52 shrink-0">
      <div
        class="sticky top-20 flex flex-col gap-1 rounded-2xl bg-white/80 p-3 shadow-sm border border-slate-200 dark:border-slate-800 dark:bg-slate-900/70"
      >
        <div class="px-3 py-2 mb-2 border-b border-slate-100 dark:border-slate-800">
          <h1 class="text-base font-black text-slate-800 dark:text-slate-100 tracking-tight">
            管理后台
          </h1>
          <p class="text-[9px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">
            Admin Dashboard
          </p>
        </div>
        <button
          v-for="tab in tabs"
          :key="tab.key"
          type="button"
          class="flex items-center gap-3 px-4 py-3.5 text-base font-bold transition-all rounded-xl group"
          :class="
            activeTab === tab.key
              ? 'bg-blue-600 text-white shadow-lg shadow-blue-200 dark:shadow-none'
              : 'text-slate-500 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-200'
          "
          @click="handleTabClick(tab.path)"
        >
          <component
            :is="tab.icon"
            :size="20"
            :class="
              activeTab === tab.key
                ? 'text-white'
                : 'text-slate-400 group-hover:text-slate-600 dark:group-hover:text-slate-200'
            "
          />
          {{ tab.label }}
        </button>
      </div>
    </aside>

    <!-- 内容区域 -->
    <div class="flex-1 min-w-0">
      <div
        class="rounded-2xl border border-slate-200/80 bg-white/80 p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900/70 min-h-full"
      >
        <router-view />
      </div>
    </div>
  </div>
</template>
