<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { ArrowLeft, Trophy, ListChecks, ChevronLeft, ChevronRight } from 'lucide-vue-next'
import { contestStore } from '@/stores/contestStore'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, TitleComponent, LegendComponent } from 'echarts/components'

use([CanvasRenderer, LineChart, GridComponent, TooltipComponent, TitleComponent, LegendComponent])

const route = useRoute()
const contestId = computed(() => String(route.params.id))
const contest = computed(() => contestStore.contests.find((c) => String(c.id) === contestId.value))

// 积分榜全量 mock 数据（增多样本用于分页）
const names = [
  'Team_Alpha', 'PwnMaster', 'CryptoKing', 'WebNinja', 'MiscHunter', 'ReverseLab', 'FlagSeeker', 'Newbie_01', 'CTF_Fan', 'Learner',
  'ByteDancer', 'ShellHunter', 'NetRunner', 'CodeBreaker', 'HashCracker', 'KeyFinder', 'BugHunter', 'PatchMaster', 'StackSmash', 'HeapLeak',
  'FormatString', 'RaceCondition', 'LogicBomb', 'SideChannel', 'TimingAttack', 'SQLNinja', 'XSSMaster', 'CSRFGuard', 'SSRFPro', 'RCEKing',
  'CipherLover', 'StegoArt', 'ForensicLab', 'OSINT_Pro', 'Reverser_X', 'Disasm_King', 'DebugMaster', 'Emulator', 'Fuzzer_99', 'ExploitDev',
  'RedTeam_A', 'BlueTeam_B', 'Pentest_C', 'SecResearcher', 'VulnHunter', 'ZeroDay_1', 'ScriptKid', 'Elite_Hacker', 'WhiteHat_42', 'GrayHat_7',
]
const fullScoreboard = computed(() =>
  names.map((name, i) => ({
    name,
    score: Math.max(100, 3500 - i * 68),
    solved: Math.max(0, 12 - Math.floor(i / 4)),
  }))
    .sort((a, b) => b.score - a.score)
    .map((r, i) => ({ ...r, rank: i + 1 }))
)

// 分页
const pageSize = 10
const currentPage = ref(1)
const totalPages = computed(() => Math.ceil(fullScoreboard.value.length / pageSize))
const paginatedScoreboard = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return fullScoreboard.value.slice(start, start + pageSize)
})

function goPrev() {
  if (currentPage.value > 1) currentPage.value--
}
function goNext() {
  if (currentPage.value < totalPages.value) currentPage.value++
}

// 前 10 名选手随时间积分变化（折线图）mock：横轴以天为单位
const timePoints = ['第1天', '第2天', '第3天', '第4天', '第5天', '第6天', '第7天', '第8天', '第9天', '第10天']
const top10 = computed(() => fullScoreboard.value.slice(0, 10))
const lineChartColors = ['#3b82f6', '#8b5cf6', '#06b6d4', '#10b981', '#f59e0b', '#ef4444', '#ec4899', '#6366f1', '#14b8a6', '#84cc16']

const lineChartOption = computed(() => {
  const series = top10.value.map((row, idx) => {
    const base = Math.round(row.score * 0.08)
    const step = (row.score - base) / (timePoints.length - 1)
    const data = timePoints.map((_, i) => Math.round(base + step * i))
    return {
      name: row.name,
      type: 'line',
      smooth: true,
      symbol: 'circle',
      symbolSize: 6,
      data,
      lineStyle: { width: 2 },
      itemStyle: { color: lineChartColors[idx] },
    }
  })
  return {
    tooltip: { trigger: 'axis' },
    legend: { type: 'scroll', bottom: 0, textStyle: { fontSize: 10 } },
    grid: { left: '10%', right: '8%', top: '8%', bottom: '20%' },
    xAxis: { type: 'category', data: timePoints, boundaryGap: false },
    yAxis: { type: 'value', name: '积分' },
    series,
  }
})
</script>

<template>
  <div class="min-h-[calc(100vh-64px)]">
    <div v-if="contest" class="flex gap-6">
      <!-- 主内容区 -->
      <div class="min-w-0 flex-1 space-y-6">
        <div class="flex items-center gap-4">
          <router-link
            :to="`/contest/${contestId}/challenges`"
            class="inline-flex items-center gap-2 rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          >
            <ArrowLeft class="h-4 w-4" />
            返回题目
          </router-link>
          <span class="text-slate-500 dark:text-slate-400">|</span>
          <h1 class="text-xl font-bold text-slate-900 dark:text-slate-50">
            {{ contest.name }} · 积分排行
          </h1>
        </div>

        <!-- 上图：前十选手积分随时间变化折线图 -->
        <div class="rounded-xl border border-slate-200 bg-white p-4 shadow-sm dark:border-slate-800 dark:bg-slate-900">
          <div class="h-80 w-full">
            <v-chart class="h-full w-full" :option="lineChartOption" autoresize />
          </div>
        </div>

        <!-- 下图：积分榜表格 + 分页 -->
        <div class="rounded-xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900 overflow-hidden">
          <div class="overflow-x-auto">
            <table class="w-full min-w-[400px] text-sm">
              <thead>
                <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-800 dark:bg-slate-800/60">
                  <th class="w-14 shrink-0 px-4 py-3 text-center font-bold text-slate-600 dark:text-slate-300 whitespace-nowrap">
                    排名
                  </th>
                  <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                    用户 / 队伍
                  </th>
                  <th class="px-4 py-3 text-right font-bold text-slate-600 dark:text-slate-300">
                    总分
                  </th>
                  <th class="px-4 py-3 text-right font-bold text-slate-600 dark:text-slate-300">
                    解出题数
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="row in paginatedScoreboard"
                  :key="row.rank"
                  class="border-b border-slate-100 transition-colors hover:bg-slate-50 dark:border-slate-800 dark:hover:bg-slate-800/50"
                  :class="{ 'bg-blue-50/50 dark:bg-blue-900/10': row.rank <= 3 }"
                >
                  <td class="w-14 shrink-0 px-4 py-3 text-center font-bold text-slate-800 dark:text-slate-100 tabular-nums whitespace-nowrap">
                    <span
                      v-if="row.rank <= 3"
                      class="inline-flex h-7 w-7 flex-shrink-0 items-center justify-center rounded-full text-xs font-black text-white"
                      :class="{
                        'bg-amber-500': row.rank === 1,
                        'bg-slate-400': row.rank === 2,
                      }"
                      :style="row.rank === 3 ? { backgroundColor: '#b87333' } : undefined"
                    >
                      {{ row.rank }}
                    </span>
                    <span v-else>{{ row.rank }}</span>
                  </td>
                  <td class="px-4 py-3 font-medium text-slate-800 dark:text-slate-100">
                    {{ row.name }}
                  </td>
                  <td class="px-4 py-3 text-right font-bold text-blue-600 dark:text-blue-400">
                    {{ row.score }}
                  </td>
                  <td class="px-4 py-3 text-right text-slate-600 dark:text-slate-300">
                    {{ row.solved }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="flex items-center justify-between border-t border-slate-200 px-4 py-3 dark:border-slate-800">
            <span class="text-sm text-slate-500 dark:text-slate-400">
              共 {{ fullScoreboard.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
            </span>
            <div class="flex items-center gap-2">
              <button
                type="button"
                class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                :disabled="currentPage <= 1"
                @click="goPrev"
              >
                <ChevronLeft :size="16" />
                上一页
              </button>
              <button
                type="button"
                class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                :disabled="currentPage >= totalPages"
                @click="goNext"
              >
                下一页
                <ChevronRight :size="16" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧栏：与比赛题目页相同的尺寸与位置（w-72、sticky top-16、同一套标签样式） -->
      <aside class="flex w-72 shrink-0 flex-col gap-4 transition-all duration-300">
        <div class="sticky top-16 space-y-2">
          <div class="flex gap-1 rounded-xl border border-slate-100 bg-white p-1 shadow-sm dark:border-slate-800 dark:bg-slate-900">
            <RouterLink
              :to="`/contest/${contestId}/challenges`"
              class="flex flex-1 items-center justify-center gap-2 rounded-lg px-4 py-2.5 text-sm font-medium text-slate-600 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-200"
            >
              <ListChecks :size="18" />
              比赛题目
            </RouterLink>
            <span
              class="flex flex-1 items-center justify-center gap-2 rounded-lg bg-blue-50 px-4 py-2.5 text-sm font-bold text-blue-700 dark:bg-blue-900/30 dark:text-blue-300"
            >
              <Trophy :size="18" />
              积分排行
            </span>
          </div>
        </div>
      </aside>
    </div>

    <div
      v-else
      class="flex min-h-[200px] items-center justify-center rounded-xl border border-dashed border-slate-300 bg-white/60 text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-900/70 dark:text-slate-400"
    >
      未找到对应的赛事信息。
    </div>
  </div>
</template>
