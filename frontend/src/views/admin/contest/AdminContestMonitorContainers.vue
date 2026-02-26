<script setup lang="ts">
import { ref, computed } from 'vue'
import { RefreshCw } from 'lucide-vue-next'

const userOrTeams = ['Team_Alpha', 'PwnMaster', 'CryptoKing', 'WebNinja', 'MiscHunter', 'ReverseLab', 'FlagSeeker', 'Newbie_01', 'CTF_Fan', 'ByteDancer']
const containerChallenges = ['pwn_stack', 'web_container', 'Reverse_ELF', 'Misc_Container']
const containerUsages = ref(
  Array.from({ length: 22 }, (_, i) => ({
    id: i + 1,
    userOrTeam: userOrTeams[i % userOrTeams.length],
    challengeTitle: containerChallenges[i % containerChallenges.length],
    status: 'running' as const,
    startedAt: `2025-03-15 ${String(10 + Math.floor(i / 6)).padStart(2, '0')}:${String((i * 7 + 20) % 60).padStart(2, '0')}:00`,
    timeLeft: `${String(Math.floor((60 - i * 2) % 60)).padStart(2, '0')}:${String((55 - i * 3 + 60) % 60).padStart(2, '0')}`,
  }))
)
const containerPageSize = 10
const containerPage = ref(1)
const containerTotalPages = computed(() => Math.max(1, Math.ceil(containerUsages.value.length / containerPageSize)))
const paginatedContainers = computed(() => {
  const start = (containerPage.value - 1) * containerPageSize
  return containerUsages.value.slice(start, start + containerPageSize)
})

function refresh() {
  containerPage.value = 1
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-4 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          容器使用情况
        </h3>
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="refresh"
        >
          <RefreshCw class="h-4 w-4" />
          刷新
        </button>
      </div>
      <div class="overflow-x-auto rounded-lg border border-slate-200 dark:border-slate-700">
        <table class="w-full min-w-[600px] text-sm">
          <thead>
            <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-700 dark:bg-slate-800/60">
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                用户 / 队伍
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                题目
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                状态
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                开启时间
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                剩余时间
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="row in paginatedContainers"
              :key="row.id"
              class="border-b border-slate-100 dark:border-slate-700/80 hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="px-4 py-3 text-slate-700 dark:text-slate-200">
                {{ row.userOrTeam }}
              </td>
              <td class="px-4 py-3 text-slate-700 dark:text-slate-200">
                {{ row.challengeTitle }}
              </td>
              <td class="px-4 py-3">
                <span class="rounded-full bg-green-100 px-2 py-0.5 text-xs font-medium text-green-700 dark:bg-green-900/30 dark:text-green-300">
                  运行中
                </span>
              </td>
              <td class="px-4 py-3 font-mono text-xs text-slate-600 dark:text-slate-400">
                {{ row.startedAt }}
              </td>
              <td class="px-4 py-3 font-mono text-xs text-slate-600 dark:text-slate-400">
                {{ row.timeLeft }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="containerTotalPages > 1" class="mt-3 flex items-center justify-end gap-2 text-sm text-slate-500 dark:text-slate-400">
        <button
          type="button"
          class="rounded px-2 py-1 hover:bg-slate-100 dark:hover:bg-slate-800 disabled:opacity-50"
          :disabled="containerPage <= 1"
          @click="containerPage--"
        >
          上一页
        </button>
        <span>{{ containerPage }} / {{ containerTotalPages }}</span>
        <button
          type="button"
          class="rounded px-2 py-1 hover:bg-slate-100 dark:hover:bg-slate-800 disabled:opacity-50"
          :disabled="containerPage >= containerTotalPages"
          @click="containerPage++"
        >
          下一页
        </button>
      </div>
    </section>
  </div>
</template>
