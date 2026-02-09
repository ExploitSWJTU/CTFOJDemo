<script setup lang="ts">
import { ref, computed } from 'vue';
import {
  Search,
  Filter,
  EyeOff,
  Play,
  Copy,
  Trash2,
  RefreshCw,
  Send,
  Loader2,
} from 'lucide-vue-next';
import MarkdownIt from 'markdown-it';
import { challenges as mockChallenges } from '../mock/challenges';
import type { Challenge, Category, Difficulty } from '../types/challenge';
import ChallengeCard from '../components/ChallengeCard.vue';

// Initialize Markdown parser
const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true,
});

// State
const challenges = ref<Challenge[]>(mockChallenges);
const selectedCategory = ref<Category | 'All'>('All');
const searchQuery = ref('');
const selectedDifficulty = ref<Difficulty | 'All'>('All');
const hideSolved = ref(false);
const detailVisible = ref(false);
const currentChallenge = ref<Challenge | null>(null);
const flagInput = ref('');
const submitting = ref(false);

// Constants
const categories: (Category | 'All')[] = ['All', 'Web', 'Pwn', 'Crypto', 'Misc'];
const difficulties: (Difficulty | 'All')[] = ['All', 'Easy', 'Medium', 'Hard'];

// Filtering
const filteredChallenges = computed(() => {
  return challenges.value.filter((c) => {
    const matchCategory = selectedCategory.value === 'All' || c.category === selectedCategory.value;
    const matchSearch =
      c.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      c.description.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchDifficulty =
      selectedDifficulty.value === 'All' || c.difficulty === selectedDifficulty.value;
    const matchSolved = hideSolved.value ? c.status !== 'solved' : true;

    return matchCategory && matchSearch && matchDifficulty && matchSolved;
  });
});

// Actions
const openDetail = (id: number) => {
  const challenge = challenges.value.find((c) => c.id === id);
  if (challenge) {
    currentChallenge.value = challenge;
    detailVisible.value = true;
    flagInput.value = '';
  }
};

const startContainer = () => {
  if (!currentChallenge.value) return;
  currentChallenge.value.containerState = 'loading';
  // Mock API call
  setTimeout(() => {
    if (currentChallenge.value) {
      currentChallenge.value.containerState = 'running';
      currentChallenge.value.containerInfo = {
        ip: '10.10.10.101',
        port: Math.floor(Math.random() * 50000) + 10000,
        timeLeft: '01:00:00',
      };
    }
  }, 2000);
};

const destroyContainer = () => {
  if (!currentChallenge.value) return;
  currentChallenge.value.containerState = 'idle';
  currentChallenge.value.containerInfo = undefined;
};

const extendTime = () => {
  if (currentChallenge.value?.containerInfo) {
    currentChallenge.value.containerInfo.timeLeft = '01:00:00';
  }
};

const submitFlag = () => {
  if (!flagInput.value) return;
  submitting.value = true;
  setTimeout(() => {
    submitting.value = false;
    alert('Flag submitted! (Mock)');
  }, 1000);
};

const copyToClipboard = (text: string) => {
  navigator.clipboard.writeText(text);
  // In a real app, show a toast notification here
};

const renderedDescription = computed(() => {
  return currentChallenge.value ? md.render(currentChallenge.value.description) : '';
});
</script>

<template>
  <div class="flex min-h-[calc(100vh-64px)] gap-6">
    <!-- Left Sidebar -->
    <aside class="hidden w-16 shrink-0 flex-col gap-4 lg:flex">
      <div
        class="sticky top-24 flex flex-col gap-2 rounded-2xl bg-white p-2 shadow-sm dark:bg-slate-900"
      >
        <button
          v-for="cat in categories"
          :key="cat"
          class="group relative flex h-10 w-10 items-center justify-center rounded-xl font-medium transition-all"
          :class="
            selectedCategory === cat
              ? 'bg-blue-600 text-white shadow-lg shadow-blue-200 dark:shadow-none'
              : 'text-slate-500 hover:bg-slate-100 hover:text-slate-900 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-100'
          "
          @click="selectedCategory = cat"
        >
          <span class="text-xs font-bold">{{ cat === 'All' ? 'All' : cat[0] }}</span>
          <!-- Tooltip -->
          <span
            class="absolute left-14 z-50 rounded bg-slate-800 px-2 py-1 text-xs whitespace-nowrap text-white opacity-0 transition-opacity group-hover:opacity-100"
          >
            {{ cat }}
          </span>
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="min-w-0 flex-1 pb-10">
      <!-- Top Filter Bar -->
      <div
        class="sticky top-16 z-30 -mx-4 mb-6 bg-slate-50/80 px-4 py-4 backdrop-blur-md lg:static lg:mx-0 lg:bg-transparent lg:p-0 lg:backdrop-blur-none dark:bg-slate-950/80 lg:dark:bg-transparent"
      >
        <div
          class="flex flex-wrap items-center justify-between gap-4 rounded-2xl bg-white p-4 shadow-sm dark:bg-slate-900"
        >
          <div class="flex min-w-[200px] flex-1 items-center gap-4">
            <div class="relative max-w-md flex-1">
              <Search class="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-400" />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="搜索题目..."
                class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 pr-4 pl-9 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
          </div>

          <div class="flex items-center gap-4">
            <div class="flex items-center gap-2">
              <Filter class="h-4 w-4 text-slate-400" />
              <select
                v-model="selectedDifficulty"
                class="h-10 rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              >
                <option v-for="diff in difficulties" :key="diff" :value="diff">
                  {{ diff }}
                </option>
              </select>
            </div>

            <div class="h-6 w-px bg-slate-200 dark:bg-slate-700" />

            <label class="flex cursor-pointer items-center gap-2 select-none">
              <div class="relative">
                <input v-model="hideSolved" type="checkbox" class="peer sr-only" />
                <div
                  class="h-6 w-11 rounded-full bg-slate-200 peer-checked:bg-blue-600 peer-focus:ring-4 peer-focus:ring-blue-300 peer-focus:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:border after:border-gray-300 after:bg-white after:transition-all after:content-[''] peer-checked:after:translate-x-full peer-checked:after:border-white dark:border-gray-600 dark:bg-gray-700 dark:peer-focus:ring-blue-800"
                />
              </div>
              <span
                class="flex items-center gap-1 text-sm font-medium text-slate-600 dark:text-slate-300"
              >
                <EyeOff class="h-4 w-4" /> 隐藏已解出
              </span>
            </label>
          </div>
        </div>
      </div>

      <!-- Challenge Grid -->
      <div class="grid grid-cols-1 gap-6 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4">
        <ChallengeCard
          v-for="challenge in filteredChallenges"
          :key="challenge.id"
          :challenge="challenge"
          @view-details="openDetail"
        />
      </div>

      <!-- Detail Modal (Drawer Style) -->
      <a-drawer
        v-model:visible="detailVisible"
        width="600px"
        :footer="false"
        class="challenge-drawer"
        :unmount-on-close="true"
      >
        <template #title>
          <div class="flex items-center gap-3">
            <span class="text-lg font-bold">{{ currentChallenge?.title }}</span>
            <a-tag
              v-if="currentChallenge"
              :color="currentChallenge.category === 'Web' ? 'arcoblue' : 'gray'"
            >
              {{ currentChallenge.category }}
            </a-tag>
          </div>
        </template>

        <div v-if="currentChallenge" class="flex h-full flex-col gap-6">
          <!-- Description -->
          <div
            class="prose prose-sm dark:prose-invert max-w-none rounded-xl bg-slate-50 p-4 dark:bg-slate-800/50"
          >
            <!-- eslint-disable-next-line vue/no-v-html -->
            <div class="markdown-body" v-html="renderedDescription" />
          </div>

          <!-- Container Operations -->
          <div class="rounded-xl border border-slate-200 p-5 dark:border-slate-700">
            <h4 class="mb-4 text-sm font-bold tracking-wider text-slate-500 uppercase">环境操作</h4>

            <!-- State: Idle -->
            <div v-if="currentChallenge.containerState === 'idle'" class="text-center">
              <button
                class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-6 py-3 font-bold text-white transition-transform hover:scale-105 hover:bg-blue-700 active:scale-95"
                @click="startContainer"
              >
                <Play class="h-5 w-5" /> 开启靶机环境
              </button>
              <p class="mt-2 text-xs text-slate-500">点击开启后将为您分配独立的 Docker 容器</p>
            </div>

            <!-- State: Loading -->
            <div
              v-else-if="currentChallenge.containerState === 'loading'"
              class="flex flex-col items-center py-4"
            >
              <Loader2 class="mb-3 h-8 w-8 animate-spin text-blue-600" />
              <span class="text-sm font-medium text-slate-600 dark:text-slate-300"
                >容器部署中，请稍后...</span
              >
              <div
                class="mt-4 h-2 w-64 overflow-hidden rounded-full bg-slate-100 dark:bg-slate-800"
              >
                <div class="animate-progress h-full bg-blue-500" />
              </div>
            </div>

            <!-- State: Running -->
            <div
              v-else-if="
                currentChallenge.containerState === 'running' && currentChallenge.containerInfo
              "
              class="space-y-4"
            >
              <div
                class="flex items-center justify-between rounded-lg bg-slate-100 p-3 dark:bg-slate-800"
              >
                <div class="flex flex-col">
                  <span class="text-xs text-slate-500">连接地址</span>
                  <div
                    class="flex items-center gap-2 font-mono text-base font-bold text-slate-800 dark:text-slate-100"
                  >
                    {{ currentChallenge.containerInfo.ip }}:{{
                      currentChallenge.containerInfo.port
                    }}
                    <button
                      class="text-slate-400 hover:text-blue-600"
                      @click="
                        copyToClipboard(
                          `${currentChallenge.containerInfo.ip}:${currentChallenge.containerInfo.port}`,
                        )
                      "
                    >
                      <Copy class="h-4 w-4" />
                    </button>
                  </div>
                </div>
                <div class="text-right">
                  <span class="text-xs text-slate-500">剩余时间</span>
                  <div class="font-mono text-lg font-bold text-red-500">
                    {{ currentChallenge.containerInfo.timeLeft }}
                  </div>
                </div>
              </div>

              <div class="flex gap-3">
                <button
                  class="flex flex-1 items-center justify-center gap-2 rounded-lg border border-slate-200 bg-white py-2 text-sm font-bold text-slate-700 hover:bg-slate-50 hover:text-blue-600 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:text-blue-400"
                  @click="extendTime"
                >
                  <RefreshCw class="h-4 w-4" /> 延时
                </button>
                <button
                  class="flex flex-1 items-center justify-center gap-2 rounded-lg border border-slate-200 bg-white py-2 text-sm font-bold text-slate-700 hover:border-red-500 hover:bg-red-50 hover:text-red-600 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-red-500 dark:hover:bg-red-900/20 dark:hover:text-red-400"
                  @click="destroyContainer"
                >
                  <Trash2 class="h-4 w-4" /> 销毁
                </button>
              </div>
            </div>
          </div>

          <!-- Submission -->
          <div class="mt-auto">
            <h4 class="mb-2 text-sm font-bold tracking-wider text-slate-500 uppercase">
              提交 Flag
            </h4>
            <div class="flex gap-2">
              <input
                v-model="flagInput"
                type="text"
                placeholder="flag{...}"
                class="flex-1 rounded-lg border border-slate-200 bg-slate-50 px-4 py-2 outline-none focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-200 dark:border-slate-700 dark:bg-slate-800 dark:focus:border-blue-500 dark:focus:bg-slate-900 dark:focus:ring-blue-900/30"
                @keyup.enter="submitFlag"
              />
              <button
                class="flex items-center gap-2 rounded-lg bg-slate-900 px-6 font-bold text-white transition-colors hover:bg-blue-600 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-slate-700 dark:hover:bg-blue-600"
                :disabled="!flagInput || submitting"
                @click="submitFlag"
              >
                <Loader2 v-if="submitting" class="h-4 w-4 animate-spin" />
                <Send v-else class="h-4 w-4" />
                提交
              </button>
            </div>
          </div>
        </div>
      </a-drawer>
    </div>
  </div>
</template>

<style scoped>
@import 'github-markdown-css/github-markdown-light.css';

/* 简单的进度条动画 */
@keyframes progress {
  0% {
    width: 0%;
  }
  50% {
    width: 70%;
  }
  100% {
    width: 90%;
  }
}
.animate-progress {
  animation: progress 2s ease-in-out infinite;
}

/* 适配暗色模式的 Markdown 样式 */
.dark .markdown-body {
  background-color: transparent;
  color: #e2e8f0;
}
.dark .markdown-body pre {
  background-color: #1e293b;
}
</style>
