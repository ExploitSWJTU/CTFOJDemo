<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Search } from 'lucide-vue-next'

type ContestStatus = 'ongoing' | 'upcoming' | 'finished'

type ContestType = 'individual' | 'team'

interface Contest {
  id: number
  name: string
  description: string
  startTime: string
  endTime: string
  status: ContestStatus
  imageUrl: string
  countdownText?: string
  participantCount: number
  type: ContestType
}

const allContests: Contest[] = [
  {
    id: 1,
    name: '第八届西南交通大学 CTF 新秀杯',
    description: '面向零基础新生的入门赛，涵盖基础题目。',
    startTime: '2025-03-10 19:00',
    endTime: '2025-03-10 22:00',
    status: 'ongoing',
    imageUrl: 'https://dummyimage.com/640x360/0f172a/38bdf8&text=Newbie+CTF',
    participantCount: 156,
    type: 'individual',
  },
  {
    id: 2,
    name: '校内常规训练赛',
    description: '每周一次的校内训练赛，用于巩固日常练习内容。',
    startTime: '2025-03-15 19:00',
    endTime: '2025-03-15 23:00',
    status: 'upcoming',
    imageUrl: 'https://dummyimage.com/640x360/020617/a5b4fc&text=Weekly+Training',
    countdownText: '2天后开始',
    participantCount: 89,
    type: 'team',
  },
  {
    id: 3,
    name: 'SWJTU CTF 校赛',
    description: '正式选拔赛，成绩将作为集训队选拔的重要参考。',
    startTime: '2025-02-01 09:00',
    endTime: '2025-02-01 17:00',
    status: 'finished',
    imageUrl: 'https://dummyimage.com/640x360/111827/f97316&text=SWJTU+CTF+Final',
    participantCount: 234,
    type: 'team',
  },
  {
    id: 4,
    name: '春季 CTF 挑战赛',
    description: '春季学期大型 CTF 比赛，包含 Web、Crypto、Pwn 等多个方向。',
    startTime: '2025-03-20 10:00',
    endTime: '2025-03-20 18:00',
    status: 'upcoming',
    imageUrl: 'https://dummyimage.com/640x360/1e293b/10b981&text=Spring+CTF',
    countdownText: '7天后开始',
    participantCount: 201,
    type: 'individual',
  },
  {
    id: 5,
    name: '网络安全周 CTF 竞赛',
    description: '配合网络安全宣传周举办的 CTF 竞赛，提升网络安全意识。',
    startTime: '2025-04-15 09:00',
    endTime: '2025-04-15 21:00',
    status: 'upcoming',
    imageUrl: 'https://dummyimage.com/640x360/0f172a/f59e0b&text=Cyber+Security',
    countdownText: '33天后开始',
    participantCount: 178,
    type: 'team',
  },
  {
    id: 6,
    name: 'CTF 新人训练营',
    description: '面向新手的训练营，提供详细讲解和指导。',
    startTime: '2025-01-15 14:00',
    endTime: '2025-01-15 17:00',
    status: 'finished',
    imageUrl: 'https://dummyimage.com/640x360/111827/8b5cf6&text=Training+Camp',
    participantCount: 95,
    type: 'individual',
  },
  {
    id: 7,
    name: 'CTF 月度挑战',
    description: '每月一次的挑战赛，题目难度适中，适合日常练习。',
    startTime: '2025-03-25 19:00',
    endTime: '2025-03-25 23:00',
    status: 'upcoming',
    imageUrl: 'https://dummyimage.com/640x360/020617/06b6d4&text=Monthly+Challenge',
    countdownText: '12天后开始',
    participantCount: 142,
    type: 'individual',
  },
  {
    id: 8,
    name: 'CTF 团队对抗赛',
    description: '团队形式的对抗赛，考验团队协作和综合能力。',
    startTime: '2025-01-20 10:00',
    endTime: '2025-01-20 18:00',
    status: 'finished',
    imageUrl: 'https://dummyimage.com/640x360/1e293b/ef4444&text=Team+Battle',
    participantCount: 167,
    type: 'team',
  },
]

const statusLabel: Record<ContestStatus, string> = {
  ongoing: '进行中',
  upcoming: '未开始',
  finished: '已结束',
}

// 历史参赛记录
interface ContestHistory {
  contestId: number
  contestName: string
  startTime: string
  endTime: string
  rank: number
  solvedCount: number
  totalScore: number
}

const contestHistory: ContestHistory[] = [
  {
    contestId: 3,
    contestName: 'SWJTU CTF 校赛',
    startTime: '2025-02-01 09:00',
    endTime: '2025-02-01 17:00',
    rank: 12,
    solvedCount: 8,
    totalScore: 1250,
  },
]

// 搜索和筛选
const searchQuery = ref('')
const selectedStatus = ref<ContestStatus | 'all'>('all')

// 分页
const currentPage = ref(1)
const pageSize = ref(4)

const filteredContests = computed(() => {
  let result = allContests

  // 状态筛选
  if (selectedStatus.value !== 'all') {
    result = result.filter((contest) => contest.status === selectedStatus.value)
  }

  // 搜索筛选
  if (searchQuery.value.trim()) {
    const query = searchQuery.value.toLowerCase().trim()
    result = result.filter(
      (contest) =>
        contest.name.toLowerCase().includes(query) ||
        contest.description.toLowerCase().includes(query)
    )
  }

  return result
})

// 分页后的比赛列表
const paginatedContests = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredContests.value.slice(start, end)
})

// 总页数
const totalPages = computed(() => {
  return Math.ceil(filteredContests.value.length / pageSize.value)
})

// 当筛选条件改变时，重置到第一页
const resetPage = () => {
  currentPage.value = 1
}

// 监听筛选条件变化，重置页码
watch([searchQuery, selectedStatus], () => {
  resetPage()
})
</script>

<template>
  <div class="space-y-4">
    <!-- 状态颜色说明 -->
    <div class="flex justify-end text-xs text-slate-500 dark:text-slate-400">
      <div class="flex items-center gap-3">
        <span class="inline-flex items-center gap-1">
          <span class="inline-block h-2 w-2 rounded-full bg-emerald-500" />
          待开始
        </span>
        <span class="inline-flex items-center gap-1">
          <span class="inline-block h-2 w-2 rounded-full bg-red-500" />
          进行中
        </span>
        <span class="inline-flex items-center gap-1">
          <span class="inline-block h-2 w-2 rounded-full bg-slate-400 dark:bg-slate-500" />
          已结束
        </span>
      </div>
    </div>

    <!-- 主要内容区域：左右布局 -->
    <div class="flex gap-6">
      <!-- 左侧：搜索和比赛列表 -->
      <div class="flex-1 space-y-4">
        <!-- 搜索和筛选 -->
        <div class="flex w-full flex-col gap-4 rounded-xl border border-slate-200/80 bg-white/80 p-4 shadow-sm dark:border-slate-800 dark:bg-slate-900/70">
          <div class="flex flex-col gap-4 sm:flex-row sm:items-center">
            <!-- 搜索框 -->
            <div class="relative flex-1">
              <Search
                class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400"
              />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="搜索比赛名称或描述..."
                class="w-full rounded-lg border border-slate-300 bg-white py-2 pl-10 pr-4 text-sm placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400"
              />
            </div>

            <!-- 状态筛选 -->
            <div class="flex items-center gap-2">
              <label class="text-sm font-medium text-slate-600 dark:text-slate-300">
                状态：
              </label>
              <select
                v-model="selectedStatus"
                class="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-700 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:focus:border-blue-400"
              >
                <option value="all">
                  全部
                </option>
                <option value="ongoing">
                  进行中
                </option>
                <option value="upcoming">
                  未开始
                </option>
                <option value="finished">
                  已结束
                </option>
              </select>
            </div>
          </div>
        </div>

        <!-- 比赛列表 -->
        <div class="flex flex-col gap-6" aria-label="赛事列表">
          <article
            v-for="contest in paginatedContests"
            :key="contest.id"
            class="flex w-full flex-col rounded-xl border border-slate-200/80 bg-white/80 p-6 shadow-sm transition hover:-translate-y-0.5 hover:border-blue-300 hover:shadow-md dark:border-slate-800 dark:bg-slate-900/70 dark:hover:border-blue-500/60"
          >
            <header class="mb-4 flex items-start gap-4">
              <div class="h-36 w-52 overflow-hidden rounded-lg bg-slate-100 dark:bg-slate-800">
                <img
                  :src="contest.imageUrl"
                  :alt="contest.name"
                  class="h-full w-full object-cover"
                  loading="lazy"
                />
              </div>

              <div class="flex flex-1 items-start justify-between gap-4">
                <div class="min-w-0 space-y-2">
                  <div class="flex items-center gap-2">
                    <h3 class="truncate text-lg font-semibold text-slate-900 dark:text-slate-50">
                      {{ contest.name }}
                    </h3>
                    <span
                      class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium"
                      :class="
                        contest.type === 'individual'
                          ? 'bg-purple-100 text-purple-700 dark:bg-purple-500/15 dark:text-purple-300'
                          : 'bg-indigo-100 text-indigo-700 dark:bg-indigo-500/15 dark:text-indigo-300'
                      "
                    >
                      {{ contest.type === 'individual' ? '个人赛' : '团体赛' }}
                    </span>
                  </div>
                  <p class="line-clamp-2 text-sm text-slate-500 dark:text-slate-400">
                    {{ contest.description }}
                  </p>
                  <dl class="grid grid-cols-1 gap-1.5 text-sm text-slate-500 dark:text-slate-400">
                    <div class="flex items-center gap-2">
                      <dt class="shrink-0 font-medium text-slate-600 dark:text-slate-300">
                        开始时间
                      </dt>
                      <dd class="truncate">
                        {{ contest.startTime }}
                      </dd>
                    </div>
                    <div class="flex items-center gap-2">
                      <dt class="shrink-0 font-medium text-slate-600 dark:text-slate-300">
                        结束时间
                      </dt>
                      <dd class="truncate">
                        {{ contest.endTime }}
                      </dd>
                    </div>
                    <div class="flex items-center gap-2">
                      <dt class="shrink-0 font-medium text-slate-600 dark:text-slate-300">
                        报名人数
                      </dt>
                      <dd class="truncate">
                        {{ contest.participantCount }} 人
                      </dd>
                    </div>
                  </dl>
                </div>

                <div class="flex shrink-0 flex-col items-end gap-1.5 text-sm">
                  <!-- 状态色标 -->
                  <span
                    v-if="contest.status === 'ongoing'"
                    class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-red-600 dark:text-red-400"
                  >
                    <span class="inline-block h-2.5 w-2.5 rounded-full bg-red-500" />
                    {{ statusLabel.ongoing }}
                  </span>
                  <span
                    v-else-if="contest.status === 'upcoming'"
                    class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-emerald-600 dark:text-emerald-300"
                  >
                    <span class="inline-block h-2.5 w-2.5 rounded-full bg-emerald-500" />
                    {{ contest.countdownText || '即将开始' }}
                  </span>
                  <span
                    v-else
                    class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-slate-500 dark:text-slate-400"
                  >
                    <span class="inline-block h-2.5 w-2.5 rounded-full bg-slate-400 dark:bg-slate-500" />
                    {{ statusLabel.finished }}
                  </span>
                </div>
              </div>
            </header>
            <div class="-mt-13 flex justify-end">
              <button
                type="button"
                class="inline-flex h-8 items-center justify-center whitespace-nowrap rounded-lg px-4 text-sm font-medium text-white transition focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2 dark:focus-visible:ring-offset-slate-950"
                :class="
                  contest.status === 'ongoing'
                    ? 'bg-blue-600 hover:bg-blue-700'
                    : contest.status === 'upcoming'
                      ? 'bg-slate-800 hover:bg-slate-900 dark:bg-slate-100 dark:text-slate-900'
                      : 'bg-slate-400 hover:bg-slate-500 cursor-not-allowed'
                "
              >
                {{
                  contest.status === 'ongoing'
                    ? '进入比赛'
                    : contest.status === 'upcoming'
                      ? '查看详情'
                      : '已结束'
                }}
              </button>
            </div>
          </article>
        </div>

        <!-- 分页 -->
        <div
          v-if="totalPages > 1"
          class="flex items-center justify-center gap-2 pt-4"
        >
          <button
            type="button"
            :disabled="currentPage === 1"
            class="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="currentPage--"
          >
            上一页
          </button>

          <div class="flex items-center gap-1">
            <button
              v-for="page in totalPages"
              :key="page"
              type="button"
              class="rounded-lg px-3 py-2 text-sm font-medium transition"
              :class="
                currentPage === page
                  ? 'bg-blue-600 text-white'
                  : 'border border-slate-300 bg-white text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700'
              "
              @click="currentPage = page"
            >
              {{ page }}
            </button>
          </div>

          <button
            type="button"
            :disabled="currentPage === totalPages"
            class="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="currentPage++"
          >
            下一页
          </button>
        </div>
      </div>

      <!-- 右侧：历史参赛栏 -->
      <div class="w-80 shrink-0">
        <div class="rounded-xl border border-slate-200/80 bg-white/80 p-4 shadow-sm dark:border-slate-800 dark:bg-slate-900/70">
          <h3 class="mb-4 text-lg font-semibold text-slate-900 dark:text-slate-50">
            历史参赛
          </h3>
          <div class="space-y-3">
            <div
              v-for="history in contestHistory"
              :key="history.contestId"
              class="cursor-pointer rounded-lg border border-slate-200/80 p-3 transition hover:border-blue-300 hover:bg-slate-50 dark:border-slate-700 dark:hover:bg-slate-800/50 dark:hover:border-blue-500/60"
            >
              <div class="text-sm font-medium text-slate-900 dark:text-slate-50">
                {{ history.contestName }}
              </div>
              <div class="mt-2 space-y-1.5 text-xs">
                <div class="flex items-center justify-between">
                  <span class="text-slate-500 dark:text-slate-400">名次</span>
                  <span class="font-medium text-slate-700 dark:text-slate-200">
                    第 {{ history.rank }} 名
                  </span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-slate-500 dark:text-slate-400">解题数</span>
                  <span class="font-medium text-slate-700 dark:text-slate-200">
                    {{ history.solvedCount }} 题
                  </span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-slate-500 dark:text-slate-400">总积分</span>
                  <span class="font-medium text-slate-700 dark:text-slate-200">
                    {{ history.totalScore }} 分
                  </span>
                </div>
              </div>
              <div class="mt-2 text-xs text-slate-400 dark:text-slate-500">
                {{ history.startTime }} - {{ history.endTime }}
              </div>
            </div>
            <div
              v-if="contestHistory.length === 0"
              class="py-4 text-center text-sm text-slate-500 dark:text-slate-400"
            >
              暂无历史参赛记录
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
