<script setup lang="ts">
import { ref, computed } from 'vue'
import { RefreshCw } from 'lucide-vue-next'

interface SubmitterWithOrder {
  name: string
  order: number
  submittedAt: string
}
interface SameFlagItem {
  challengeTitle: string
  challengeType: 'dynamic_container' | 'dynamic_attachment'
  flag: string
  submitters: SubmitterWithOrder[]
}

const userOrTeams = ['Team_Alpha', 'PwnMaster', 'CryptoKing', 'WebNinja', 'MiscHunter', 'ReverseLab', 'FlagSeeker', 'Newbie_01', 'CTF_Fan', 'ByteDancer']
const sameFlagTitles = ['pwn_dynamic', 'Web_附件题', 'Misc_动态容器', 'Reverse_动态附件', 'SQLi_多实例', 'Crypto_容器']
const sameFlagList = ref<SameFlagItem[]>(
  Array.from({ length: 18 }, (_, i) => {
    const n = 2 + (i % 3)
    const names = userOrTeams.slice(i % 6, (i % 6) + n)
    if (names.length < n) names.push(...userOrTeams.slice(0, n - names.length))
    const title = sameFlagTitles[i % sameFlagTitles.length]
    return {
      challengeTitle: title ?? '未知题目',
      challengeType: (i % 2 === 0 ? 'dynamic_container' : 'dynamic_attachment') as 'dynamic_container' | 'dynamic_attachment',
      flag: 'flag{***}',
      submitters: names.map((name, j) => ({
        name,
        order: j + 1,
        submittedAt: `2025-03-15 14:${String(20 + j * 5 + (i % 3)).padStart(2, '0')}:${String((i + j) * 7 % 60).padStart(2, '0')}`,
      })),
    }
  })
)
const sameFlagPageSize = 8
const sameFlagPage = ref(1)
const sameFlagTotalPages = computed(() => Math.max(1, Math.ceil(sameFlagList.value.length / sameFlagPageSize)))
const paginatedSameFlag = computed(() => {
  const start = (sameFlagPage.value - 1) * sameFlagPageSize
  return sameFlagList.value.slice(start, start + sameFlagPageSize)
})

function refresh() {
  sameFlagPage.value = 1
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-4 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          作弊信息
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
      <p class="mb-4 text-xs text-slate-500 dark:text-slate-400">
        以下题目为「一题多 flag」或「一队一 flag」，若多支队伍提交了相同 flag 可能存在共享或泄露。列表按提交先后顺序标明。
      </p>
      <div class="overflow-x-auto rounded-lg border border-slate-200 dark:border-slate-700">
        <table class="w-full min-w-[640px] text-sm">
          <thead>
            <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-700 dark:bg-slate-800/60">
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                题目
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                题目类型
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                Flag
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                提交该 Flag 的用户/队伍（按提交先后）
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(row, index) in paginatedSameFlag"
              :key="index"
              class="border-b border-slate-100 dark:border-slate-700/80 hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="whitespace-nowrap px-4 py-3 text-slate-700 dark:text-slate-200">
                {{ row.challengeTitle }}
              </td>
              <td class="whitespace-nowrap px-4 py-3">
                <span
                  :class="row.challengeType === 'dynamic_container'
                    ? 'rounded-full bg-amber-100 px-2 py-0.5 text-xs font-medium text-amber-700 dark:bg-amber-900/30 dark:text-amber-300'
                    : 'rounded-full bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'"
                >
                  {{ row.challengeType === 'dynamic_container' ? '动态容器' : '动态附件' }}
                </span>
              </td>
              <td class="px-4 py-3 font-mono text-xs text-slate-500 dark:text-slate-400">
                {{ row.flag }}
              </td>
              <td class="whitespace-nowrap px-4 py-3 text-slate-700 dark:text-slate-200">
                <span
                  v-for="(s, i) in row.submitters"
                  :key="s.name + s.submittedAt"
                  class="mr-2 inline-flex items-baseline gap-1"
                >
                  <span class="shrink-0 rounded bg-slate-200 px-1.5 py-0.5 text-[10px] font-bold text-slate-600 dark:bg-slate-600 dark:text-slate-300">第{{ s.order }}位</span>
                  <span>{{ s.name }}</span>
                  <span class="font-mono text-xs text-slate-500 dark:text-slate-400">{{ s.submittedAt }}</span>
                  <template v-if="i < row.submitters.length - 1">、</template>
                </span>
                <span v-if="row.submitters.length > 1" class="ml-1 text-amber-600 dark:text-amber-400">(共 {{ row.submitters.length }} 个)</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="sameFlagTotalPages > 1" class="mt-3 flex items-center justify-end gap-2 text-sm text-slate-500 dark:text-slate-400">
        <button
          type="button"
          class="rounded px-2 py-1 hover:bg-slate-100 dark:hover:bg-slate-800 disabled:opacity-50"
          :disabled="sameFlagPage <= 1"
          @click="sameFlagPage--"
        >
          上一页
        </button>
        <span>{{ sameFlagPage }} / {{ sameFlagTotalPages }}</span>
        <button
          type="button"
          class="rounded px-2 py-1 hover:bg-slate-100 dark:hover:bg-slate-800 disabled:opacity-50"
          :disabled="sameFlagPage >= sameFlagTotalPages"
          @click="sameFlagPage++"
        >
          下一页
        </button>
      </div>
    </section>
  </div>
</template>
