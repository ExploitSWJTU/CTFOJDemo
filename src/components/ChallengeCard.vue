<script setup lang="ts">
import { computed } from 'vue';
import { User, Trophy, Eye, CircleCheckBig } from 'lucide-vue-next';
import type { Challenge } from '../types/challenge';
import { CATEGORY_MAP } from '../constants/category';

const props = defineProps<{
  challenge: Challenge;
}>();

const emit = defineEmits<{
  (e: 'view-details', id: number): void;
}>();

const categoryMeta = computed(() => {
  return CATEGORY_MAP[props.challenge.category] || CATEGORY_MAP.Misc;
});

const categoryIcon = computed(() => categoryMeta.value.icon);
const categoryColor = computed(() => categoryMeta.value.cardClass);

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
    :class="{ 'opacity-50 grayscale-[0.6]': challenge.status === 'solved' }"
  >
    <!-- 状态标签 (新版：大图标溢出) -->
    <div
      v-if="challenge.status === 'solved'"
      class="pointer-events-none absolute -top-4 -right-4 z-0 h-2/3 opacity-75 transition-transform duration-500 group-hover:scale-110"
    >
      <CircleCheckBig class="h-full w-auto text-green-500" :stroke-width="1.8" />
    </div>

    <div class="p-5">
      <!-- 头部：图标 + 标题 -->
      <div class="mb-4 flex items-start gap-3">
        <div
          class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg"
          :class="categoryColor"
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
      class="absolute inset-0 z-20 flex items-center justify-center rounded-xl bg-white/80 opacity-0 backdrop-blur-sm transition-opacity duration-300 group-hover:opacity-100 dark:bg-slate-900/80"
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
