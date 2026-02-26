<script setup lang="ts">
import { ref, computed } from 'vue'
import { Download, Eye, ChevronLeft, ChevronRight } from 'lucide-vue-next'

const userNames = ['Team_Alpha', 'PwnMaster', 'CryptoKing', 'WebNinja', 'MiscHunter', 'ReverseLab', 'FlagSeeker', 'Newbie_01', 'CTF_Fan', 'ByteDancer', 'ShellHunter', 'NetRunner', 'CodeBreaker', 'HashCracker', 'KeyFinder']
const writeupFileNames = ['writeup.pdf', 'writeup.md', 'solution.pdf', '题解.docx', 'writeup.zip', 'solution.md', 'writeup.docx', '题解.pdf']

const writeupList = computed(() =>
  Array.from({ length: 58 }, (_, i) => ({
    id: i + 1,
    userOrTeam: userNames[i % userNames.length],
    fileName: writeupFileNames[i % writeupFileNames.length],
    submitTime: `2025-03-${String(10 + (i % 20)).padStart(2, '0')} ${String(14 + (i % 6)).padStart(2, '0')}:${String(i % 60).padStart(2, '0')}`,
  }))
)

const writeupPageSize = 10
const writeupPage = ref(1)
const writeupTotalPages = computed(() => Math.ceil(writeupList.value.length / writeupPageSize))
const paginatedWriteups = computed(() => {
  const start = (writeupPage.value - 1) * writeupPageSize
  return writeupList.value.slice(start, start + writeupPageSize)
})

function downloadWriteup(item: { fileName?: string; userOrTeam?: string }) {
  alert(`下载：${item.userOrTeam ?? ''} - ${item.fileName ?? ''}`)
}

function downloadAllWriteups() {
  alert(`下载全部 ${writeupList.value.length} 个 Writeup 文件`)
}

function viewWriteup(item: { fileName?: string; userOrTeam?: string }) {
  alert(`查看：${item.userOrTeam ?? ''} - ${item.fileName ?? ''}`)
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-4 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          Writeup 列表
        </h3>
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="downloadAllWriteups"
        >
          <Download class="h-4 w-4" />
          下载全部文件
        </button>
      </div>
      <div class="overflow-x-auto rounded-lg border border-slate-200 dark:border-slate-700">
        <table class="w-full min-w-[600px] text-sm">
          <thead>
            <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-700 dark:bg-slate-800/60">
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                序号
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                用户 / 队伍
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                文件名
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                提交时间
              </th>
              <th class="px-4 py-3 text-right font-bold text-slate-600 dark:text-slate-300">
                操作
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(item, idx) in paginatedWriteups"
              :key="item.id"
              class="border-b border-slate-100 transition-colors hover:bg-slate-50 dark:border-slate-800 dark:hover:bg-slate-800/50"
            >
              <td class="px-4 py-3 text-slate-600 dark:text-slate-400">
                {{ (writeupPage - 1) * writeupPageSize + idx + 1 }}
              </td>
              <td class="px-4 py-3 font-medium text-slate-800 dark:text-slate-100">
                {{ item.userOrTeam }}
              </td>
              <td class="px-4 py-3 text-slate-600 dark:text-slate-300">
                {{ item.fileName }}
              </td>
              <td class="px-4 py-3 text-slate-500 dark:text-slate-400">
                {{ item.submitTime }}
              </td>
              <td class="px-4 py-3 text-right">
                <div class="flex items-center justify-end gap-2">
                  <button
                    type="button"
                    class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-2.5 py-1.5 text-xs font-medium text-slate-600 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                    @click="viewWriteup(item)"
                  >
                    <Eye class="h-3.5 w-3.5" /> 查看
                  </button>
                  <button
                    type="button"
                    class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-2.5 py-1.5 text-xs font-medium text-slate-600 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                    @click="downloadWriteup(item)"
                  >
                    <Download class="h-3.5 w-3.5" /> 下载
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div v-if="writeupTotalPages > 1" class="mt-4 flex items-center justify-between">
        <span class="text-sm text-slate-500 dark:text-slate-400">
          共 {{ writeupList.length }} 条，第 {{ writeupPage }} / {{ writeupTotalPages }} 页
        </span>
        <div class="flex items-center gap-2">
          <button
            type="button"
            class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            :disabled="writeupPage <= 1"
            @click="writeupPage--"
          >
            <ChevronLeft class="h-4 w-4" /> 上一页
          </button>
          <button
            type="button"
            class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            :disabled="writeupPage >= writeupTotalPages"
            @click="writeupPage++"
          >
            下一页 <ChevronRight class="h-4 w-4" />
          </button>
        </div>
      </div>
    </section>
  </div>
</template>
