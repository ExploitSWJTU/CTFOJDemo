<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'

type LogStatus = 'success' | 'warning' | 'error'

interface LogEntry {
  time: string
  user: string
  message: string
  status: LogStatus
}

const logs = ref<LogEntry[]>([
  { time: '2026-02-25 21:30:12', user: 'admin', message: '登录管理后台', status: 'success' },
  { time: '2026-02-25 21:35:47', user: 'admin', message: '创建训练题目「pwn_stack_overflow」', status: 'success' },
  { time: '2026-02-25 21:40:03', user: 'ctfer01', message: '多次输入错误密码，账户被临时锁定', status: 'warning' },
  { time: '2026-02-25 21:45:19', user: 'system', message: '容器实例销毁失败，Docker API 不可用', status: 'error' },
  { time: '2026-02-25 21:50:32', user: 'admin', message: '修改全局配置：开启注册功能', status: 'success' },
  { time: '2026-02-25 21:52:10', user: 'admin', message: '编辑用户 ctfer01 的真实姓名与学工号', status: 'success' },
  { time: '2026-02-25 21:55:00', user: 'rikka', message: '登录用户端', status: 'success' },
  { time: '2026-02-25 21:56:22', user: 'admin', message: '新建比赛「二月月赛」', status: 'success' },
  { time: '2026-02-25 21:58:45', user: 'system', message: '定时任务：清理过期容器完成', status: 'success' },
  { time: '2026-02-25 22:00:11', user: 'ctfer02', message: '提交题目 flag 错误次数过多，触发限流', status: 'warning' },
  { time: '2026-02-25 22:02:33', user: 'admin', message: '发布系统公告「本周六维护通知」', status: 'success' },
  { time: '2026-02-25 22:05:07', user: 'user1', message: '注册新账号', status: 'success' },
  { time: '2026-02-25 22:08:19', user: 'admin', message: '重置用户 user3 的登录密码', status: 'success' },
  { time: '2026-02-25 22:10:00', user: 'system', message: '数据库备份任务执行失败，磁盘空间不足', status: 'error' },
  { time: '2026-02-25 22:12:44', user: 'admin', message: '删除训练题目「Deprecated_Challenge」', status: 'success' },
  { time: '2026-02-25 22:15:30', user: 'ctfer01', message: '首次解出题目「Web_XSS_Challenge」', status: 'success' },
  { time: '2026-02-25 22:18:02', user: 'admin', message: '编辑论坛帖子「环境搭建指南」', status: 'success' },
  { time: '2026-02-25 22:20:55', user: 'rikka', message: '申请延长容器时间', status: 'success' },
  { time: '2026-02-25 22:23:17', user: 'system', message: '证书即将过期提醒（30 天后）', status: 'warning' },
  { time: '2026-02-25 22:25:40', user: 'admin', message: '修改首页轮播图顺序', status: 'success' },
  { time: '2026-02-25 22:28:11', user: 'user2', message: '登录用户端', status: 'success' },
  { time: '2026-02-25 22:30:00', user: 'admin', message: '导出用户列表 CSV', status: 'success' },
  { time: '2026-02-25 22:32:45', user: 'ctfer03', message: '多次尝试访问未授权管理接口', status: 'warning' },
  { time: '2026-02-25 22:35:20', user: 'admin', message: '更新全局设置：每日签到开关', status: 'success' },
  { time: '2026-02-25 22:38:03', user: 'system', message: '实例节点 10.10.10.105 心跳超时', status: 'error' },
  { time: '2026-02-25 22:40:50', user: 'admin', message: '审核通过队伍「SWJTU_Team_A」', status: 'success' },
  { time: '2026-02-25 22:43:22', user: 'user1', message: '修改个人资料', status: 'success' },
  { time: '2026-02-25 22:45:00', user: 'admin', message: '查看系统日志', status: 'success' },
])

const statusTextMap: Record<LogStatus, string> = {
  success: '成功',
  warning: '警告',
  error: '失败',
}

const statusClassMap: Record<LogStatus, string> = {
  success: 'bg-emerald-50 text-emerald-700 border-emerald-100 dark:bg-emerald-900/40 dark:text-emerald-200 dark:border-emerald-800',
  warning: 'bg-amber-50 text-amber-700 border-amber-100 dark:bg-amber-900/40 dark:text-amber-200 dark:border-amber-800',
  error: 'bg-red-50 text-red-700 border-red-100 dark:bg-red-900/40 dark:text-red-200 dark:border-red-800',
}

const logPageSize = ref(10)
const logCurrentPage = ref(1)

const logTotalPages = computed(() =>
  Math.max(1, Math.ceil(logs.value.length / logPageSize.value))
)

const paginatedLogs = computed(() => {
  const start = (logCurrentPage.value - 1) * logPageSize.value
  return logs.value.slice(start, start + logPageSize.value)
})

watch(logTotalPages, (tp) => {
  if (logCurrentPage.value > tp && tp > 0) logCurrentPage.value = tp
})
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between border-b border-slate-200 pb-3 dark:border-slate-800">
      <div>
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          系统日志
        </h2>
        <p class="mt-1 text-xs text-slate-500 dark:text-slate-400">
          展示最近的后台操作记录，方便追踪问题和审计。
        </p>
      </div>
    </div>

    <div
      class="overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900"
    >
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200 dark:divide-slate-800">
          <thead class="bg-slate-50 dark:bg-slate-900/60">
            <tr>
              <th
                class="px-6 py-3 text-left text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400"
              >
                时间
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400"
              >
                用户
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400"
              >
                信息
              </th>
              <th
                class="px-6 py-3 text-left text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400"
              >
                状态
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200 bg-white text-sm dark:divide-slate-800 dark:bg-slate-900">
            <tr
              v-for="log in paginatedLogs"
              :key="`${log.time}-${log.user}-${log.message}`"
              class="hover:bg-slate-50 dark:hover:bg-slate-800/60"
            >
              <td class="whitespace-nowrap px-6 py-3 font-mono text-xs text-slate-500 dark:text-slate-400">
                {{ log.time }}
              </td>
              <td class="whitespace-nowrap px-6 py-3 text-slate-800 dark:text-slate-100">
                {{ log.user }}
              </td>
              <td class="px-6 py-3 text-slate-700 dark:text-slate-200">
                {{ log.message }}
              </td>
              <td class="whitespace-nowrap px-6 py-3">
                <span
                  class="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold"
                  :class="statusClassMap[log.status]"
                >
                  <span
                    class="mr-1 inline-block h-1.5 w-1.5 rounded-full bg-current opacity-70"
                  />
                  {{ statusTextMap[log.status] }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 分页 -->
    <div
      v-if="logTotalPages > 1"
      class="mt-4 flex items-center justify-between border-t border-slate-200 pt-4 dark:border-slate-800"
    >
      <span class="text-sm text-slate-500 dark:text-slate-400">
        共 {{ logs.length }} 条，第 {{ logCurrentPage }} / {{ logTotalPages }} 页
      </span>
      <div class="flex items-center gap-2">
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="logCurrentPage === 1"
          @click="logCurrentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>
        <span class="px-3 py-2 text-sm font-medium text-slate-700 dark:text-slate-300">
          {{ logCurrentPage }} / {{ logTotalPages }}
        </span>
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="logCurrentPage === logTotalPages"
          @click="logCurrentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>

