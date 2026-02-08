<script setup lang="ts">
import { Flag } from 'lucide-vue-next';
</script>

<template>
  <div class="flex flex-col items-center justify-center py-32 text-center">
    <div
      class="rounded-card mb-6 flex h-20 w-20 items-center justify-center bg-blue-50 text-blue-600 dark:bg-blue-900/20 dark:text-blue-400"
    >
      <Flag :size="40" />
    </div>
    <h2 class="mb-2 text-3xl font-black text-slate-800 dark:text-slate-100">
      赛事列表
    </h2>
    <p class="text-slate-500 dark:text-slate-400">
      模块正在建设中，敬请期待...
    </p>
  </div>
</template>
