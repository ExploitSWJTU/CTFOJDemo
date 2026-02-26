<script setup lang="ts">
import { ref, computed } from 'vue'
import { Eye, ChevronLeft, ChevronRight, UserCheck, UserX } from 'lucide-vue-next'

type TeamStatus = 'pending' | 'approved' | 'rejected'

interface TeamItem {
  id: number
  name: string
  captain: string
  memberCount: number
  members: string[]
  status: TeamStatus
  applyTime: string
}

const statusLabel: Record<TeamStatus, string> = {
  pending: '待审核',
  approved: '已通过',
  rejected: '已拒绝',
}

const statusClass: Record<TeamStatus, string> = {
  pending: 'bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-300',
  approved: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300',
  rejected: 'bg-slate-100 text-slate-600 dark:bg-slate-700 dark:text-slate-300',
}

const teamNames = [
  'Team_Alpha', 'PwnMaster', 'CryptoKing', 'WebNinja', 'MiscHunter', 'ReverseLab', 'FlagSeeker',
  'Newbie_01', 'CTF_Fan', 'ByteDancer', 'ShellHunter', 'NetRunner', 'CodeBreaker', 'HashCracker',
  'KeyFinder', 'BugHunter', 'PatchMaster', 'StackSmash', 'HeapLeak', 'FormatString', 'RaceCondition',
  'LogicBomb', 'SideChannel', 'TimingAttack', 'SQLNinja', 'XSSMaster', 'CSRFGuard', 'SSRFPro',
]
const captainNames = ['张三', '李四', '王五', '赵六', '钱七', '孙八', '周九', '吴十', '郑一', '王二', '陈三', '刘四', '杨五', '黄六', '林七', '何八', '高九', '罗十', '梁一', '宋二', '唐三', '许四', '韩五', '冯六', '邓七', '曹八', '彭九', '曾十']

const teamList = computed<TeamItem[]>(() =>
  Array.from({ length: 56 }, (_, i) => {
    const memberCount = 3 + (i % 4)
    const captain = captainNames[i % captainNames.length] ?? '未知'
    const members = Array.from({ length: memberCount }, (_, j) => (j === 0 ? captain : `成员${j}`))
    const statuses: TeamStatus[] = ['pending', 'approved', 'rejected']
    const status = statuses[i % 3]
    return {
      id: i + 1,
      name: (teamNames[i % teamNames.length] ?? 'Team') + (i >= teamNames.length ? `_${Math.floor(i / teamNames.length)}` : ''),
      captain,
      memberCount,
      members,
      status: status ?? 'pending',
      applyTime: `2025-03-${String(5 + (i % 20)).padStart(2, '0')} ${String(10 + (i % 8)).padStart(2, '0')}:${String((i * 7) % 60).padStart(2, '0')}`,
    }
  })
)

const pageSize = 10
const currentPage = ref(1)
const totalPages = computed(() => Math.ceil(teamList.value.length / pageSize))
const paginatedTeams = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return teamList.value.slice(start, start + pageSize)
})

const detailVisible = ref(false)
const selectedTeam = ref<TeamItem | null>(null)

function openDetail(team: TeamItem) {
  selectedTeam.value = team
  detailVisible.value = true
}

function closeDetail() {
  detailVisible.value = false
  selectedTeam.value = null
}

function approveTeam() {
  if (selectedTeam.value) selectedTeam.value.status = 'approved'
  closeDetail()
}

function rejectTeam() {
  if (selectedTeam.value) selectedTeam.value.status = 'rejected'
  closeDetail()
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-4 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          队伍审核列表
        </h3>
      </div>

      <div class="overflow-x-auto rounded-lg border border-slate-200 dark:border-slate-700">
        <table class="w-full min-w-[640px] text-sm">
          <thead>
            <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-700 dark:bg-slate-800/60">
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                序号
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                队伍名称
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                队长
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                人数
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                状态
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300">
                申请时间
              </th>
              <th class="px-4 py-3 text-right font-bold text-slate-600 dark:text-slate-300">
                操作
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(team, idx) in paginatedTeams"
              :key="team.id"
              class="border-b border-slate-100 transition-colors hover:bg-slate-50 dark:border-slate-800 dark:hover:bg-slate-800/50 cursor-pointer"
              @click="openDetail(team)"
            >
              <td class="px-4 py-3 text-slate-600 dark:text-slate-400">
                {{ (currentPage - 1) * pageSize + idx + 1 }}
              </td>
              <td class="px-4 py-3 font-medium text-slate-800 dark:text-slate-100">
                {{ team.name }}
              </td>
              <td class="px-4 py-3 text-slate-700 dark:text-slate-200">
                {{ team.captain }}
              </td>
              <td class="px-4 py-3 text-slate-600 dark:text-slate-400">
                {{ team.memberCount }} 人
              </td>
              <td class="px-4 py-3">
                <span
                  class="inline-flex rounded-full px-2.5 py-0.5 text-xs font-medium"
                  :class="statusClass[team.status]"
                >
                  {{ statusLabel[team.status] }}
                </span>
              </td>
              <td class="px-4 py-3 text-slate-500 dark:text-slate-400">
                {{ team.applyTime }}
              </td>
              <td class="px-4 py-3 text-right">
                <button
                  type="button"
                  class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-2.5 py-1.5 text-xs font-medium text-slate-600 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                  @click.stop="openDetail(team)"
                >
                  <Eye class="h-3.5 w-3.5" /> 查看
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="totalPages > 1" class="mt-4 flex items-center justify-between">
        <span class="text-sm text-slate-500 dark:text-slate-400">
          共 {{ teamList.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
        </span>
        <div class="flex items-center gap-2">
          <button
            type="button"
            class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            :disabled="currentPage <= 1"
            @click="currentPage--"
          >
            <ChevronLeft class="h-4 w-4" /> 上一页
          </button>
          <button
            type="button"
            class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            :disabled="currentPage >= totalPages"
            @click="currentPage++"
          >
            下一页 <ChevronRight class="h-4 w-4" />
          </button>
        </div>
      </div>
    </section>

    <!-- 队伍信息弹窗 -->
    <a-modal
      v-model:visible="detailVisible"
      title="队伍信息"
      width="480px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeDetail"
    >
      <div v-if="selectedTeam" class="space-y-4">
        <div class="grid grid-cols-2 gap-3 text-sm">
          <div>
            <span class="text-slate-500 dark:text-slate-400">队伍名称</span>
            <p class="mt-0.5 font-medium text-slate-800 dark:text-slate-100">
              {{ selectedTeam.name }}
            </p>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">队长</span>
            <p class="mt-0.5 font-medium text-slate-800 dark:text-slate-100">
              {{ selectedTeam.captain }}
            </p>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">人数</span>
            <p class="mt-0.5 font-medium text-slate-800 dark:text-slate-100">
              {{ selectedTeam.memberCount }} 人
            </p>
          </div>
          <div>
            <span class="text-slate-500 dark:text-slate-400">状态</span>
            <p class="mt-0.5">
              <span
                class="inline-flex rounded-full px-2 py-0.5 text-xs font-medium"
                :class="statusClass[selectedTeam.status]"
              >
                {{ statusLabel[selectedTeam.status] }}
              </span>
            </p>
          </div>
          <div class="col-span-2">
            <span class="text-slate-500 dark:text-slate-400">申请时间</span>
            <p class="mt-0.5 font-medium text-slate-800 dark:text-slate-100">
              {{ selectedTeam.applyTime }}
            </p>
          </div>
        </div>
        <div>
          <span class="text-slate-500 dark:text-slate-400 text-sm">成员列表</span>
          <ul class="mt-1.5 rounded-lg border border-slate-200 bg-slate-50 py-2 dark:border-slate-700 dark:bg-slate-800/50">
            <li
              v-for="(member, i) in selectedTeam.members"
              :key="i"
              class="flex items-center gap-2 px-3 py-1.5 text-sm text-slate-700 dark:text-slate-200"
            >
              <span class="text-slate-400 dark:text-slate-500">{{ i + 1 }}.</span>
              {{ member }}
            </li>
          </ul>
        </div>
        <div v-if="selectedTeam.status === 'pending'" class="flex gap-3 pt-2 border-t border-slate-200 dark:border-slate-700">
          <button
            type="button"
            class="flex flex-1 items-center justify-center gap-2 rounded-lg bg-emerald-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-emerald-700"
            @click="approveTeam"
          >
            <UserCheck class="h-4 w-4" /> 通过
          </button>
          <button
            type="button"
            class="flex flex-1 items-center justify-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm font-medium text-slate-700 transition-colors hover:bg-red-50 hover:text-red-600 hover:border-red-200 dark:border-slate-700 dark:bg-slate-800 dark:hover:bg-slate-700 dark:hover:border-red-900 dark:hover:bg-red-900/20"
            @click="rejectTeam"
          >
            <UserX class="h-4 w-4" /> 拒绝
          </button>
        </div>
      </div>
    </a-modal>
  </div>
</template>
