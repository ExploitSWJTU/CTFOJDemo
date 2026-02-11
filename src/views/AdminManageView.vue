<script setup lang="ts">
import { ref } from 'vue'

// 标签页选项
const tabs = [
  { key: 'training', label: '训练管理' },
  { key: 'contest', label: '赛事管理' },
  { key: 'forum', label: '论坛管理' },
  { key: 'user', label: '用户管理' },
  { key: 'team', label: '队伍管理' },
  { key: 'instance', label: '实例管理' },
  { key: 'log', label: '系统日志' },
  { key: 'settings', label: '系统设置' },
]

const activeTab = ref('training')
</script>

<template>
  <div class="min-h-[calc(100vh-64px)]">
    <!-- 标签页 -->
    <div class="mb-6 flex items-center justify-end gap-2 border-b border-slate-200 dark:border-slate-800">
      <button
        v-for="tab in tabs"
        :key="tab.key"
        type="button"
        class="px-4 py-2 text-sm font-medium transition-colors"
        :class="
          activeTab === tab.key
            ? 'border-b-2 border-blue-600 text-blue-600 dark:border-blue-400 dark:text-blue-400'
            : 'text-slate-500 hover:text-slate-700 dark:text-slate-400 dark:hover:text-slate-200'
        "
        @click="activeTab = tab.key"
      >
        {{ tab.label }}
      </button>
    </div>

    <!-- 内容区域 -->
    <div class="rounded-xl border border-slate-200/80 bg-white/80 p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900/70">
      <div v-if="activeTab === 'training'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          训练管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          训练管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'contest'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          赛事管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          赛事管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'forum'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          论坛管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          论坛管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'user'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          用户管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          用户管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'team'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          队伍管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          队伍管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'instance'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          实例管理
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          实例管理功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'log'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          系统日志
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          系统日志功能待实现
        </p>
      </div>

      <div v-else-if="activeTab === 'settings'" class="space-y-4">
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          系统设置
        </h2>
        <p class="text-slate-600 dark:text-slate-400">
          系统设置功能待实现
        </p>
      </div>
    </div>
  </div>
</template>

