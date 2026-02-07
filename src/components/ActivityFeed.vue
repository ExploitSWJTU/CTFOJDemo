<script setup lang="ts">
import {
  Activity,
  Plus,
  Compass,
  Tag,
  BookOpen,
  ThumbsUp,
  MessageSquare,
  MoreHorizontal,
} from 'lucide-vue-next';

const feeds = [
  {
    id: 1,
    type: 'new_challenge',
    icon: Plus,
    iconColor: 'text-emerald-500',
    iconBg: 'bg-white',
    time: '2 小时前',
    title: 'Kernel Pwn 101: Ret2User',
    desc: '新增一道适合初学者的内核 Pwn 题目。学习如何绕过 SMEP/SMAP 并提权。',
    stats: [
      { icon: Compass, text: '45 次解出', type: 'text' },
      { icon: Tag, text: 'kernel, pwn', type: 'mono' },
    ],
    avatars: ['A', 'B'],
  },
  {
    id: 2,
    type: 'writeup',
    icon: BookOpen,
    iconColor: 'text-blue-500',
    iconBg: 'bg-white',
    user: 'Rikka',
    action: '发布了题解',
    time: '5 小时前',
    title: '"Easy Heap OverFlow" 详解',
    codeSnippet: true,
    stats: [
      { icon: ThumbsUp, text: '24', type: 'btn' },
      { icon: MessageSquare, text: '8', type: 'btn' },
    ],
  },
  {
    id: 3,
    type: 'discussion',
    icon: MessageSquare,
    iconColor: 'text-purple-500',
    iconBg: 'bg-white',
    user: 'YYM',
    action: '发起讨论',
    time: '8 小时前',
    title: '楼上是给',
    tags: ['#给'],
    stats: [{ icon: MessageSquare, text: '15 条评论', type: 'btn' }],
  },
];
</script>

<template>
  <div class="col-span-3">
    <h3
      class="mb-6 flex items-center gap-2 px-2 text-lg font-bold text-slate-800 dark:text-slate-100"
    >
      <Activity class="text-blue-600 dark:text-blue-400" :size="20" />
      最新动态
    </h3>

    <div
      class="relative space-y-6 before:absolute before:top-2 before:bottom-2 before:left-4 before:w-0.5 before:bg-slate-200 dark:before:bg-slate-800"
    >
      <div v-for="feed in feeds" :key="feed.id" class="relative pl-12">
        <!-- 图标 -->
        <div
          class="absolute top-0 left-0 z-10 flex h-8 w-8 items-center justify-center rounded-full border-2 border-slate-200 bg-white transition-colors dark:border-slate-800 dark:bg-slate-900"
        >
          <component :is="feed.icon" :class="['h-4 w-4', feed.iconColor]" />
        </div>

        <!-- 卡片 -->
        <div
          class="group cursor-pointer rounded-2xl border border-slate-200 bg-white p-5 shadow-sm transition-all hover:border-blue-300 dark:border-slate-800 dark:bg-slate-900 dark:hover:border-blue-700"
        >
          <!-- 头部 -->
          <div class="mb-2 flex items-start justify-between">
            <div class="flex items-center gap-2">
              <template v-if="feed.type === 'new_challenge'">
                <span
                  class="rounded-full bg-emerald-100 px-2 py-0.5 text-[10px] font-bold text-emerald-700 transition-colors dark:bg-emerald-900/30 dark:text-emerald-400"
                  >新题上线</span
                >
              </template>
              <template v-else>
                <img
                  :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${feed.user}`"
                  class="h-5 w-5 rounded-full bg-slate-100 dark:bg-slate-800"
                />
                <span class="text-xs font-bold text-slate-700 dark:text-slate-200">{{
                  feed.user
                }}</span>
                <span class="text-xs text-slate-400 dark:text-slate-500">{{ feed.action }}</span>
                <span class="mx-1 text-xs text-slate-400">·</span>
              </template>
              <span class="text-xs text-slate-400 dark:text-slate-500">{{ feed.time }}</span>
            </div>
            <button
              v-if="feed.type === 'new_challenge'"
              class="text-slate-300 transition-colors group-hover:text-blue-500 dark:group-hover:text-blue-400"
            >
              <MoreHorizontal :size="16" />
            </button>
          </div>

          <!-- 内容 -->
          <h4
            class="mb-2 text-base font-bold text-slate-800 transition-colors group-hover:text-blue-600 dark:text-slate-100 dark:group-hover:text-blue-400"
          >
            {{ feed.title }}
          </h4>
          <p v-if="feed.desc" class="mb-4 line-clamp-2 text-sm text-slate-500 dark:text-slate-400">
            {{ feed.desc }}
          </p>

          <div
            v-if="feed.codeSnippet"
            class="mb-3 rounded-xl border border-slate-100 bg-slate-50 p-3 font-mono text-xs text-slate-600 transition-colors dark:border-slate-800 dark:bg-slate-800/50 dark:text-slate-400"
          >
            <span class="text-blue-600 dark:text-blue-400">from</span> pwn
            <span class="text-blue-600 dark:text-blue-400">import</span> *<br />
            p = process(<span class="text-green-600 dark:text-green-400">"./pwn"</span>)<br />
            <span class="text-slate-400 dark:text-slate-600"># 泄露堆基地址...</span>
          </div>

          <div v-if="feed.tags" class="mb-3 flex flex-wrap gap-2">
            <span
              v-for="tag in feed.tags"
              :key="tag"
              class="rounded bg-slate-100 px-2 py-0.5 text-[10px] font-bold text-slate-500 transition-colors dark:bg-slate-800 dark:text-slate-400"
              >{{ tag }}</span
            >
          </div>

          <!-- 底部/统计 -->
          <div class="flex items-center gap-4">
            <div v-for="(stat, idx) in feed.stats" :key="idx" class="flex items-center gap-1.5">
              <component :is="stat.icon" class="h-3.5 w-3.5 text-slate-400 dark:text-slate-600" />
              <span
                v-if="stat.type !== 'btn'"
                :class="[
                  'text-xs',
                  stat.type === 'mono'
                    ? 'font-mono text-slate-500 dark:text-slate-400'
                    : 'font-bold text-slate-600 transition-colors dark:text-slate-300',
                ]"
                >{{ stat.text }}</span
              >
              <span
                v-else
                class="cursor-pointer text-xs font-bold text-slate-500 transition-colors hover:text-blue-600 dark:text-slate-400 dark:hover:text-blue-400"
                >{{ stat.text }}</span
              >
            </div>

            <div v-if="feed.avatars" class="ml-auto flex -space-x-2">
              <img
                v-for="seed in feed.avatars"
                :key="seed"
                class="h-6 w-6 rounded-full border-2 border-white transition-colors dark:border-slate-900"
                :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${seed}`"
              />
              <div
                class="flex h-6 w-6 items-center justify-center rounded-full border-2 border-white bg-slate-100 text-[8px] font-bold text-slate-500 transition-colors dark:border-slate-900 dark:bg-slate-800 dark:text-slate-400"
              >
                +12
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <button
      class="mt-12 w-full rounded-xl border border-slate-200 bg-white py-3 text-sm font-bold text-slate-500 shadow-sm transition-all hover:bg-slate-50 hover:text-blue-600 dark:border-slate-800 dark:bg-slate-900 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
    >
      加载更多动态
    </button>
  </div>
</template>
