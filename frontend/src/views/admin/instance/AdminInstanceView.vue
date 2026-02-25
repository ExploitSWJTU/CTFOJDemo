<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Container, Cpu, Layers, Search, RefreshCw } from 'lucide-vue-next'

interface InstanceRow {
  id: number
  challenge: string
  user: string
  avatar: string
  ip: string
  port: number
  cpu: number
  ram: number
}

// 模拟活跃实例数据（与 training 实例监控一致的数据结构）
const activeInstances = ref<InstanceRow[]>([
  { id: 1, challenge: 'Easy_Heap_Overflow', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.101', port: 12345, cpu: 12, ram: 24 },
  { id: 2, challenge: 'SQL_Injection_Advanced', user: 'admin', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin', ip: '10.10.10.102', port: 54321, cpu: 45, ram: 62 },
  { id: 3, challenge: 'Web_XSS_Challenge', user: 'user1', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user1', ip: '10.10.10.103', port: 8080, cpu: 28, ram: 128 },
  { id: 4, challenge: 'Pwn_Stack_Overflow', user: 'user2', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user2', ip: '10.10.10.104', port: 31337, cpu: 8, ram: 64 },
  { id: 5, challenge: 'Misc_Stegano', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.105', port: 9000, cpu: 5, ram: 32 },
])

const searchQuery = ref('')
const pageSize = ref(10)
const currentPage = ref(1)

const filteredInstances = computed(() => {
  if (!searchQuery.value.trim()) return activeInstances.value
  const q = searchQuery.value.toLowerCase().trim()
  return activeInstances.value.filter(
    (inst) =>
      inst.user.toLowerCase().includes(q) ||
      inst.challenge.toLowerCase().includes(q) ||
      inst.ip.includes(q)
  )
})

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredInstances.value.length / pageSize.value))
)

const paginatedInstances = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  return filteredInstances.value.slice(start, start + pageSize.value)
})

watch(searchQuery, () => { currentPage.value = 1 })
watch(totalPages, (tp) => {
  if (currentPage.value > tp && tp > 0) currentPage.value = tp
})

function destroyInstance(inst: InstanceRow) {
  if (!confirm(`确定销毁 @${inst.user} 的实例「${inst.challenge}」？`)) return
  const idx = activeInstances.value.findIndex((i) => i.id === inst.id)
  if (idx !== -1) activeInstances.value.splice(idx, 1)
}

function refreshList() {
  // 模拟刷新，实际可对接后端
  searchQuery.value = ''
}
</script>

<template>
  <div class="flex flex-col gap-6">
    <!-- 顶部标题与操作 -->
    <div class="flex items-center justify-between border-b border-slate-200 dark:border-slate-800 pb-4">
      <div class="flex items-center gap-4">
        <h2 class="text-xl font-black text-slate-900 dark:text-white flex items-center gap-2">
          <Container :size="24" class="text-blue-600" />
          实例监控
        </h2>
        <p class="text-xs text-slate-500 dark:text-slate-400">
          当前运行中的题目实例，可查看端点与资源占用并销毁
        </p>
      </div>
      <div class="flex items-center gap-2">
        <div class="relative max-w-xs">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" :size="14" />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索用户 / 题目 / IP..."
            class="w-full h-9 pl-9 pr-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs outline-none focus:border-blue-500 transition-all"
          />
        </div>
        <button
          type="button"
          class="flex items-center gap-2 px-3 py-1.5 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-slate-800 rounded-lg text-xs font-bold transition-all"
          @click="refreshList"
        >
          <RefreshCw :size="14" />
          刷新
        </button>
      </div>
    </div>

    <!-- 实例列表表格（与 training 实例监控同款样式） -->
    <div class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl overflow-hidden shadow-sm">
      <table class="w-full text-left border-collapse text-xs">
        <thead>
          <tr class="bg-slate-50 dark:bg-slate-900 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-200 dark:border-slate-800">
            <th class="px-4 py-3">
              User / Challenge
            </th>
            <th class="px-4 py-3">
              Endpoint
            </th>
            <th class="px-4 py-3">
              Resources
            </th>
            <th class="px-4 py-3 text-right">
              Actions
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
          <tr
            v-for="inst in paginatedInstances"
            :key="inst.id"
            class="hover:bg-slate-50/50 dark:hover:bg-slate-900/30 transition-colors"
          >
            <td class="px-4 py-3">
              <div class="flex items-center gap-3">
                <img
                  :src="inst.avatar"
                  class="w-7 h-7 rounded-full border border-slate-200 dark:border-slate-700"
                  :alt="inst.user"
                />
                <div>
                  <div class="text-xs font-bold text-slate-800 dark:text-slate-100">
                    @{{ inst.user }}
                  </div>
                  <div class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter">
                    {{ inst.challenge }}
                  </div>
                </div>
              </div>
            </td>
            <td class="px-4 py-3 font-mono text-blue-600 dark:text-blue-400 font-bold">
              {{ inst.ip }}:{{ inst.port }}
            </td>
            <td class="px-4 py-3">
              <div class="flex items-center gap-2">
                <Cpu :size="12" class="text-slate-400" />
                <span class="font-mono">{{ inst.cpu }}%</span>
                <Layers :size="12" class="text-slate-400 ml-2" />
                <span class="font-mono">{{ inst.ram }}M</span>
              </div>
            </td>
            <td class="px-4 py-3 text-right">
              <button
                type="button"
                class="px-2 py-1 bg-rose-50 dark:bg-rose-900/20 text-rose-600 text-[10px] font-bold rounded border border-rose-100 dark:border-rose-900/30 hover:bg-rose-100 dark:hover:bg-rose-900/30 transition-colors"
                @click="destroyInstance(inst)"
              >
                Destroy
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div
        v-if="filteredInstances.length === 0"
        class="px-4 py-12 text-center text-xs text-slate-500 dark:text-slate-400"
      >
        {{ searchQuery ? '没有匹配的实例' : '暂无运行中的实例' }}
      </div>
    </div>

    <!-- 分页 -->
    <div
      v-if="totalPages > 1"
      class="mt-4 flex items-center justify-between border-t border-slate-200 dark:border-slate-800 pt-4"
    >
      <span class="text-xs text-slate-500 dark:text-slate-400">
        共 {{ filteredInstances.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
      </span>
      <div class="flex items-center gap-2">
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-xs font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="currentPage <= 1"
          @click="currentPage = Math.max(1, currentPage - 1)"
        >
          上一页
        </button>
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-xs font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="currentPage >= totalPages"
          @click="currentPage = Math.min(totalPages, currentPage + 1)"
        >
          下一页
        </button>
      </div>
    </div>
  </div>
</template>
