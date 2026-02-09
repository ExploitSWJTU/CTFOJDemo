<script setup lang="ts">
import { computed } from 'vue';
import { Globe, Terminal, Binary, FileQuestion, User, Trophy, Eye } from 'lucide-vue-next';
import type { Challenge } from '../types/challenge';

const props = defineProps<{
  challenge: Challenge;
}>();

const emit = defineEmits<{
  (e: 'view-details', id: number): void;
}>();

const categoryIcon = computed(() => {
  switch (props.challenge.category) {
    case 'Web':
      return Globe;
    case 'Pwn':
      return Terminal;
    case 'Crypto':
      return Binary;
    case 'Misc':
      return FileQuestion;
    default:
      return FileQuestion;
  }
});

const difficultyColor = computed(() => {
  switch (props.challenge.difficulty) {
    case 'Easy':
      return 'text-green-600 bg-green-50 dark:text-green-400 dark:bg-green-900/20';
    case 'Medium':
      return 'text-orange-600 bg-orange-50 dark:text-orange-400 dark:bg-orange-900/20';
    case 'Hard':
      return 'text-red-600 bg-red-50 dark:text-red-400 dark:bg-red-900/20';
    default:
      return 'text-slate-600 bg-slate-50';
  }
});
</script>

<template>
  <div
    class="group relative flex h-full flex-col overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm transition-all duration-300 hover:-translate-y-1 hover:shadow-xl dark:border-slate-800 dark:bg-slate-900 dark:shadow-none"
  >
    <!-- 状态标签 -->
    <div
      v-if="challenge.status === 'solved'"
      class="absolute top-0 right-0 z-10 rounded-bl-xl bg-green-500 px-3 py-1 text-xs font-bold text-white shadow-sm"
    >
      已解出
    </div>

    <div class="p-5">
      <!-- 头部：图标 + 标题 -->
      <div class="mb-4 flex items-start gap-3">
        <div
          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-blue-50 text-blue-600 dark:bg-blue-900/20 dark:text-blue-400"
        >
          <component :is="categoryIcon" :size="20" />
        </div>
        <div>
          <h3 class="line-clamp-1 text-base font-bold text-slate-800 dark:text-slate-100">
            {{ challenge.title }}
          </h3>
          <p class="mt-1 line-clamp-2 text-xs text-slate-500 dark:text-slate-400">
            {{ challenge.description.slice(0, 60) }}...
          </p>
        </div>
      </div>

      <!-- 底部信息 -->
      <div
        class="mt-auto flex items-center justify-between border-t border-slate-100 pt-4 dark:border-slate-800"
      >
        <div class="flex items-center gap-3 text-xs text-slate-500 dark:text-slate-400">
          <div class="flex items-center gap-1">
            <User :size="12" />
            <span>{{ challenge.solvedCount }}</span>
          </div>
          <div class="flex items-center gap-1">
            <Trophy :size="12" />
            <span>{{ challenge.points }} pts</span>
          </div>
        </div>
        <span class="rounded-full px-2 py-0.5 text-xs font-bold" :class="difficultyColor">
          {{ challenge.difficulty }}
        </span>
      </div>
    </div>

    <!-- 悬停遮罩层 -->
    <div
      class="absolute inset-0 flex items-center justify-center bg-white/80 opacity-0 backdrop-blur-sm transition-opacity duration-300 group-hover:opacity-100 dark:bg-slate-900/80"
    >
      <button
        class="rounded-full bg-blue-600 px-6 py-2 text-sm font-bold text-white shadow-lg shadow-blue-200 transition-transform duration-300 hover:scale-105 hover:bg-blue-700 dark:shadow-none dark:hover:bg-blue-500"
        @click="emit('view-details', challenge.id)"
      >
        <div class="flex items-center gap-2">
          <Eye :size="16" />
          查看详情
        </div>
      </button>
    </div>
  </div>
</template>
