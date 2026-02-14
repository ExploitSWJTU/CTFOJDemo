<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Search, Plus, Edit, Trash2, ChevronLeft, ChevronRight } from 'lucide-vue-next'
import { contestStore, deleteContest, updateContest } from '../../stores/contestStore'

const router = useRouter()

// ========== 比赛管理相关逻辑 ==========
// 使用共享的比赛数据
const allContests = computed(() => contestStore.contests)

// 比赛搜索
const contestSearchQuery = ref('')

// 比赛分页
const contestPageSize = ref(10)
const contestCurrentPage = ref(1)

// 过滤后的比赛列表
const filteredContests = computed(() => {
  let result = allContests.value

  // 搜索筛选
  if (contestSearchQuery.value.trim()) {
    const query = contestSearchQuery.value.toLowerCase().trim()
    result = result.filter(
      (contest) =>
        contest.name.toLowerCase().includes(query) ||
        contest.brief.toLowerCase().includes(query) ||
        contest.description.toLowerCase().includes(query)
    )
  }

  return result
})

// 分页后的比赛列表
const paginatedContests = computed(() => {
  const start = (contestCurrentPage.value - 1) * contestPageSize.value
  const end = start + contestPageSize.value
  return filteredContests.value.slice(start, end)
})

// 总页数
const contestTotalPages = computed(() => {
  return Math.ceil(filteredContests.value.length / contestPageSize.value)
})

// 当搜索条件改变时，重置到第一页
watch(contestSearchQuery, () => {
  contestCurrentPage.value = 1
})

// 监听总页数变化，确保当前页码在有效范围内
watch(contestTotalPages, (newTotalPages) => {
  if (contestCurrentPage.value > newTotalPages && newTotalPages > 0) {
    contestCurrentPage.value = newTotalPages
  }
})

// 新建比赛
const createContest = () => {
  router.push({ path: '/admin/manage/contest/create' })
}

// 编辑比赛
const editContest = (id: number) => {
  router.push({ path: `/admin/manage/contest/edit/${id}` })
}

// 删除比赛
const handleDeleteContest = (id: number) => {
  if (confirm('确定要删除这个比赛吗？')) {
    const success = deleteContest(id)
    if (success) {
      alert('删除成功')
      // 如果当前页没有数据了，回到上一页
      if (paginatedContests.value.length === 0 && contestCurrentPage.value > 1) {
        contestCurrentPage.value--
      }
    } else {
      alert('删除失败')
    }
  }
}

// 切换比赛激活状态
const toggleContestActive = (contest: { id: number; isActive?: boolean }) => {
  const newActiveState = !(contest.isActive ?? false)
  const success = updateContest(contest.id, { isActive: newActiveState })
  if (!success) {
    alert('更新失败')
  }
}
</script>

<template>
  <div class="space-y-4">
    <!-- 搜索框和操作按钮 -->
    <div class="flex items-center justify-between gap-4">
      <div class="relative flex-1 max-w-md">
        <Search class="absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2 text-slate-400" />
        <input
          v-model="contestSearchQuery"
          type="text"
          placeholder="搜索比赛名称或描述..."
          class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 pr-4 pl-9 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
        />
      </div>
      <button
        class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
        @click="createContest"
      >
        <Plus class="h-4 w-4" />
        新建比赛
      </button>
    </div>

    <!-- 比赛列表 -->
    <div class="overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-slate-50 dark:bg-slate-800/50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                序号
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                比赛名称
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                状态
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                时间
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                简介
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                激活
              </th>
              <th class="px-6 py-3 text-right text-xs font-bold uppercase tracking-wider text-slate-500">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200 bg-white dark:divide-slate-800 dark:bg-slate-900">
            <tr
              v-for="(contest, index) in paginatedContests"
              :key="contest.id"
              class="transition-colors hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-900 dark:text-slate-100">
                {{ (contestCurrentPage - 1) * contestPageSize + index + 1 }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm font-medium text-slate-900 dark:text-slate-100">
                {{ contest.name }}
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <span
                  v-if="contest.status === 'ongoing'"
                  class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-red-600 dark:text-red-400"
                >
                  <span class="inline-block h-2 w-2 rounded-full bg-red-500" />
                  进行中
                </span>
                <span
                  v-else-if="contest.status === 'upcoming'"
                  class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-emerald-600 dark:text-emerald-300"
                >
                  <span class="inline-block h-2 w-2 rounded-full bg-emerald-500" />
                  待开始
                </span>
                <span
                  v-else
                  class="inline-flex items-center gap-1.5 whitespace-nowrap font-medium text-slate-500 dark:text-slate-400"
                >
                  <span class="inline-block h-2 w-2 rounded-full bg-slate-400 dark:bg-slate-500" />
                  已结束
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                <div class="space-y-1">
                  <div class="text-xs">
                    <span class="font-medium">开始：</span>{{ contest.startTime }}
                  </div>
                  <div class="text-xs">
                    <span class="font-medium">结束：</span>{{ contest.endTime }}
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                <div class="max-w-md truncate" :title="contest.brief">
                  {{ contest.brief || '暂无简介' }}
                </div>
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <label class="relative inline-flex cursor-pointer items-center">
                  <input
                    type="checkbox"
                    :checked="contest.isActive ?? false"
                    class="peer sr-only"
                    @change="toggleContestActive(contest)"
                  />
                  <div
                    class="peer h-6 w-11 rounded-full bg-slate-200 transition-colors after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all after:content-[''] peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-300 dark:border-slate-600 dark:bg-slate-700 dark:peer-focus:ring-blue-800"
                  />
                </label>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-2">
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
                    title="编辑"
                    @click="editContest(contest.id)"
                  >
                    <Edit class="h-4 w-4" />
                  </button>
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                    title="删除"
                    @click="handleDeleteContest(contest.id)"
                  >
                    <Trash2 class="h-4 w-4" />
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="paginatedContests.length === 0">
              <td colspan="7" class="px-6 py-12 text-center text-sm text-slate-500">
                暂无比赛数据
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 比赛分页 -->
    <div v-if="contestTotalPages > 1" class="mt-4 flex items-center justify-between">
      <div class="text-sm text-slate-600 dark:text-slate-400">
        共 {{ filteredContests.length }} 条记录，第 {{ contestCurrentPage }} / {{ contestTotalPages }} 页
      </div>
      <div class="flex items-center gap-2">
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="contestCurrentPage === 1"
          @click="contestCurrentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>
        <span class="px-3 py-2 text-sm font-medium text-slate-700 dark:text-slate-300">
          {{ contestCurrentPage }} / {{ contestTotalPages }}
        </span>
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="contestCurrentPage === contestTotalPages"
          @click="contestCurrentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>
