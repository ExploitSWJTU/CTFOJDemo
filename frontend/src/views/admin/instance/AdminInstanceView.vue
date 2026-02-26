<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Container, Search, RefreshCw } from 'lucide-vue-next'

interface InstanceRow {
  id: number
  challenge: string
  user: string
  avatar: string
  ip: string
  port: number
  /** 开始时间 */
  startTime: string
  /** 结束时间 */
  endTime: string
  /** 容器 ID */
  containerId: string
}

// 模拟活跃实例数据：用户、题目、生命周期（开始时间到结束时间）、容器 ID、访问入口
const activeInstances = ref<InstanceRow[]>([
  { id: 1, challenge: 'Easy_Heap_Overflow', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.101', port: 12345, startTime: '02-25 13:00', endTime: '02-25 14:00', containerId: 'a1b2c3d4e5f6' },
  { id: 2, challenge: 'SQL_Injection_Advanced', user: 'admin', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin', ip: '10.10.10.102', port: 54321, startTime: '02-25 12:30', endTime: '02-25 14:00', containerId: 'b2c3d4e5f6a7' },
  { id: 3, challenge: 'Web_XSS_Challenge', user: 'user1', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user1', ip: '10.10.10.103', port: 8080, startTime: '02-25 14:10', endTime: '02-25 14:33', containerId: 'c3d4e5f6a7b8' },
  { id: 4, challenge: 'Pwn_Stack_Overflow', user: 'user2', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user2', ip: '10.10.10.104', port: 31337, startTime: '02-25 13:15', endTime: '02-25 14:01', containerId: 'd4e5f6a7b8c9' },
  { id: 5, challenge: 'Misc_Stegano', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.105', port: 9000, startTime: '02-25 12:45', endTime: '02-25 14:00', containerId: 'e5f6a7b8c9d0' },
  { id: 6, challenge: 'Crypto_RSA_Intro', user: 'ctfer01', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=c1', ip: '10.10.10.106', port: 2222, startTime: '02-25 13:22', endTime: '02-25 14:00', containerId: 'f6a7b8c9d0e1' },
  { id: 7, challenge: 'Reverse_Crackme', user: 'ctfer02', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=c2', ip: '10.10.10.107', port: 3333, startTime: '02-25 14:20', endTime: '02-25 14:33', containerId: 'a7b8c9d0e1f2' },
  { id: 8, challenge: 'Web_SSRF', user: 'user3', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user3', ip: '10.10.10.108', port: 4444, startTime: '02-25 13:00', endTime: '02-25 14:00', containerId: 'b8c9d0e1f2a3' },
  { id: 9, challenge: 'Pwn_Format_String', user: 'admin', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin', ip: '10.10.10.109', port: 5555, startTime: '02-25 13:05', endTime: '02-25 14:01', containerId: 'c9d0e1f2a3b4' },
  { id: 10, challenge: 'Misc_QR_Code', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.110', port: 6666, startTime: '02-25 13:32', endTime: '02-25 14:00', containerId: 'd0e1f2a3b4c5' },
  { id: 11, challenge: 'Web_File_Upload', user: 'user1', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user1', ip: '10.10.10.111', port: 7777, startTime: '02-25 13:19', endTime: '02-25 14:00', containerId: 'e1f2a3b4c5d6' },
  { id: 12, challenge: 'Crypto_AES_ECB', user: 'ctfer03', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=c3', ip: '10.10.10.112', port: 8888, startTime: '02-25 12:38', endTime: '02-25 14:01', containerId: 'f2a3b4c5d6e7' },
  { id: 13, challenge: 'Reverse_Packer', user: 'user2', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user2', ip: '10.10.10.113', port: 9999, startTime: '02-25 13:52', endTime: '02-25 14:01', containerId: 'a3b4c5d6e7f8' },
  { id: 14, challenge: 'Pwn_ROP_Chain', user: 'ctfer01', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=c1', ip: '10.10.10.114', port: 10000, startTime: '02-25 13:27', endTime: '02-25 14:00', containerId: 'b4c5d6e7f8a9' },
  { id: 15, challenge: 'Web_JWT_Weak', user: 'admin', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin', ip: '10.10.10.115', port: 10001, startTime: '02-25 12:55', endTime: '02-25 14:01', containerId: 'c5d6e7f8a9b0' },
  { id: 16, challenge: 'Misc_Network_Pcap', user: 'user3', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user3', ip: '10.10.10.116', port: 10002, startTime: '02-25 13:43', endTime: '02-25 14:01', containerId: 'd6e7f8a9b0c1' },
  { id: 17, challenge: 'SQL_Blind_Inject', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.117', port: 10003, startTime: '02-25 13:12', endTime: '02-25 14:01', containerId: 'e7f8a9b0c1d2' },
  { id: 18, challenge: 'Crypto_Hash_Collision', user: 'ctfer02', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=c2', ip: '10.10.10.118', port: 10004, startTime: '02-25 13:07', endTime: '02-25 14:02', containerId: 'f8a9b0c1d2e3' },
  { id: 19, challenge: 'Reverse_Anti_Debug', user: 'user1', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user1', ip: '10.10.10.119', port: 10005, startTime: '02-25 13:50', endTime: '02-25 14:02', containerId: 'a9b0c1d2e3f4' },
  { id: 20, challenge: 'Pwn_Shellcode', user: 'user2', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=user2', ip: '10.10.10.120', port: 10006, startTime: '02-25 12:41', endTime: '02-25 14:00', containerId: 'b0c1d2e3f4a5' },
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
      inst.ip.includes(q) ||
      inst.containerId.toLowerCase().includes(q)
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
            placeholder="搜索用户 / 题目 / IP / 容器 ID..."
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

    <!-- 实例列表：用户、题目、生命周期、容器 ID、访问入口 -->
    <div class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl overflow-hidden shadow-sm">
      <table class="w-full text-left border-collapse text-xs">
        <thead>
          <tr class="bg-slate-50 dark:bg-slate-900 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-200 dark:border-slate-800">
            <th class="px-4 py-3">
              用户
            </th>
            <th class="px-4 py-3">
              题目
            </th>
            <th class="px-4 py-3">
              生命周期
            </th>
            <th class="px-4 py-3">
              容器 ID
            </th>
            <th class="px-4 py-3">
              访问入口
            </th>
            <th class="px-4 py-3 text-right">
              操作
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
              <div class="flex items-center gap-2">
                <img
                  :src="inst.avatar"
                  class="w-7 h-7 rounded-full border border-slate-200 dark:border-slate-700 shrink-0"
                  :alt="inst.user"
                />
                <span class="font-bold text-slate-800 dark:text-slate-100">{{ inst.user }}</span>
              </div>
            </td>
            <td class="px-4 py-3 text-slate-700 dark:text-slate-200">
              {{ inst.challenge }}
            </td>
            <td class="px-4 py-3 font-mono text-xs text-slate-600 dark:text-slate-300">
              {{ inst.startTime }} 至 {{ inst.endTime }}
            </td>
            <td class="px-4 py-3 font-mono text-[10px] text-slate-500 dark:text-slate-400">
              {{ inst.containerId }}
            </td>
            <td class="px-4 py-3 font-mono text-blue-600 dark:text-blue-400 font-bold">
              {{ inst.ip }}:{{ inst.port }}
            </td>
            <td class="px-4 py-3 text-right">
              <button
                type="button"
                class="px-2 py-1 bg-rose-50 dark:bg-rose-900/20 text-rose-600 text-[10px] font-bold rounded border border-rose-100 dark:border-rose-900/30 hover:bg-rose-100 dark:hover:bg-rose-900/30 transition-colors"
                @click="destroyInstance(inst)"
              >
                销毁
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
