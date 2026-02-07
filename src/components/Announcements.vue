<script setup lang="ts">
import { Megaphone } from 'lucide-vue-next';

const announcements = [
  {
    id: 1,
    type: '赛事公告',
    typeColor: 'text-blue-400',
    title: '春季联赛 (Spring Cup) 即将报名',
    date: '02-06',
  },
  {
    id: 2,
    type: '版本更新',
    typeColor: 'text-emerald-400',
    title: 'V2.4.0 更新：新增容器续时功能',
    date: '02-05',
  },
  {
    id: 3,
    type: '系统通知',
    typeColor: 'text-amber-400',
    title: '关于寒假期间服务器维护的通知',
    date: '02-03',
  },
];
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between px-2">
      <h3 class="flex items-center gap-2 text-lg font-bold text-slate-800 dark:text-slate-100">
        <Megaphone class="text-blue-600 dark:text-blue-400" :size="20" />
        全站公告
      </h3>
    </div>

    <div
      class="relative flex h-65 flex-col overflow-hidden rounded-3xl border border-slate-800 bg-slate-900 p-5 shadow-2xl transition-colors"
    >
      <!-- 背景装饰 -->
      <div class="absolute top-0 right-0 h-24 w-24 bg-blue-500/10 blur-3xl" />

      <div class="custom-scrollbar relative mb-4 flex-1 space-y-4 overflow-y-auto pr-1">
        <div
          v-for="(item, index) in announcements"
          :key="item.id"
          :class="['group cursor-pointer', index !== 0 ? 'border-t border-slate-800 pt-3' : '']"
        >
          <div class="flex items-start justify-between">
            <div class="flex flex-col">
              <span
                :class="`mb-0.5 text-[9px] font-black tracking-tighter uppercase ${item.typeColor}`"
              >{{ item.type }}</span>
              <p
                :class="`text-xs font-bold text-slate-200 hover:${item.typeColor} line-clamp-1 transition-colors`"
              >
                {{ item.title }}
              </p>
            </div>
            <span class="ml-2 font-mono text-[9px] text-slate-500">{{ item.date }}</span>
          </div>
        </div>
      </div>

      <button
        class="relative mt-auto w-full shrink-0 rounded-xl border border-slate-700 bg-slate-800 py-2 text-[10px] font-black text-slate-300 transition-all hover:bg-slate-700 hover:text-white"
      >
        查看全部历史公告
      </button>
    </div>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #334155;
  border-radius: 2px;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #475569;
}
</style>
