<script setup lang="ts">
import { Hash } from 'lucide-vue-next';
const tags = ['#pwn', '#writeup', '#ctf2026', '#python', '#docker'];
</script>

<template>
  <div class="col-span-1 space-y-6">
    <div
      class="rounded-3xl border border-slate-200 bg-white p-5 shadow-sm transition-colors dark:border-slate-800 dark:bg-slate-900"
    >
      <h4 class="mb-4 flex items-center gap-2 text-sm font-bold text-slate-800 dark:text-slate-100">
        <Hash class="text-blue-500" :size="16" /> 热门标签
      </h4>
      <div class="flex flex-wrap gap-2">
        <span
          v-for="tag in tags"
          :key="tag"
          class="cursor-pointer rounded-lg bg-blue-50 px-2.5 py-1 text-[11px] font-bold text-blue-600 transition-colors hover:bg-blue-100 dark:bg-blue-900/30 dark:text-blue-400 dark:hover:bg-blue-900/50"
        >
          {{ tag }}
        </span>
      </div>
    </div>
  </div>
</template>
