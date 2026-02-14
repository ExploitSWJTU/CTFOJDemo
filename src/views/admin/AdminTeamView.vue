<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Search, Plus, Edit, Trash2, X, Save, ChevronLeft, ChevronRight } from 'lucide-vue-next'
import { teamStore, updateTeam, deleteTeam, createTeam } from '../../stores/teamStore'
import type { Team } from '../../types/team'
import { users as mockUsers } from '../../mock/users'
import type { User } from '../../types/user'

// ========== 队伍管理相关逻辑 ==========
// 使用共享的队伍数据
const teamList = computed(() => teamStore.teams)
const userList = ref<User[]>([...mockUsers])

// 队伍搜索
const teamSearchQuery = ref('')

// 队伍分页
const teamPageSize = ref(10)
const teamCurrentPage = ref(1)

// 编辑/创建队伍对话框
const showTeamDialog = ref(false)
const editingTeam = ref<Team | null>(null)
const isTeamEditMode = computed(() => editingTeam.value !== null)

// 队伍表单数据
const teamFormData = ref<Partial<Team>>({
  name: '',
  avatar: '',
  description: '',
  members: [],
})

// 队伍头像文件预览
const teamAvatarPreview = ref<string>('')
const teamAvatarFile = ref<File | null>(null)
const teamAvatarInputRef = ref<HTMLInputElement | null>(null)

// 可用用户列表（用于选择队员）
const availableUsers = computed(() => userList.value)

// 过滤后的队伍列表
const filteredTeams = computed(() => {
  let result = teamList.value
  if (teamSearchQuery.value.trim()) {
    const query = teamSearchQuery.value.toLowerCase().trim()
    result = result.filter(
      (team) =>
        team.name.toLowerCase().includes(query) ||
        team.description.toLowerCase().includes(query) ||
        team.members.some((m) => m.username.toLowerCase().includes(query))
    )
  }
  return result
})

// 分页后的队伍列表
const paginatedTeams = computed(() => {
  const start = (teamCurrentPage.value - 1) * teamPageSize.value
  const end = start + teamPageSize.value
  return filteredTeams.value.slice(start, end)
})

// 队伍总页数
const teamTotalPages = computed(() => {
  return Math.ceil(filteredTeams.value.length / teamPageSize.value)
})

// 监听队伍搜索，重置页码
watch(teamSearchQuery, () => {
  teamCurrentPage.value = 1
})

// 处理队伍头像文件选择
const handleTeamAvatarChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    if (!file.type.startsWith('image/')) {
      alert('请选择图片文件')
      return
    }
    if (file.size > 5 * 1024 * 1024) {
      alert('图片大小不能超过5MB')
      return
    }
    teamAvatarFile.value = file
    const reader = new FileReader()
    reader.onload = (e) => {
      teamAvatarPreview.value = e.target?.result as string
      teamFormData.value.avatar = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

// 打开创建队伍对话框
const openCreateTeamDialog = () => {
  editingTeam.value = null
  teamFormData.value = {
    name: '',
    avatar: '',
    description: '',
    members: [],
  }
  teamAvatarPreview.value = ''
  teamAvatarFile.value = null
  showTeamDialog.value = true
}

// 打开编辑队伍对话框
const openEditTeamDialog = (team: Team) => {
  editingTeam.value = team
  teamFormData.value = {
    name: team.name,
    avatar: team.avatar,
    description: team.description,
    members: [...team.members],
  }
  teamAvatarPreview.value = team.avatar || ''
  teamAvatarFile.value = null
  showTeamDialog.value = true
}

// 清除队伍头像
const clearTeamAvatar = () => {
  teamAvatarPreview.value = ''
  teamAvatarFile.value = null
  teamFormData.value.avatar = ''
  if (teamAvatarInputRef.value) {
    teamAvatarInputRef.value.value = ''
  }
}

// 关闭队伍对话框
const closeTeamDialog = () => {
  showTeamDialog.value = false
  editingTeam.value = null
  teamFormData.value = {
    name: '',
    avatar: '',
    description: '',
    members: [],
  }
  teamAvatarPreview.value = ''
  teamAvatarFile.value = null
}

// 添加队员
const addTeamMember = (userId: number) => {
  const user = availableUsers.value.find((u) => u.id === userId)
  if (user && teamFormData.value.members) {
    const exists = teamFormData.value.members.some((m) => m.id === user.id)
    if (!exists) {
      teamFormData.value.members.push({
        id: user.id,
        username: user.username,
        avatar: user.avatar,
      })
    }
  }
}

// 移除队员
const removeTeamMember = (memberId: number) => {
  if (teamFormData.value.members) {
    const index = teamFormData.value.members.findIndex((m) => m.id === memberId)
    if (index !== -1) {
      teamFormData.value.members.splice(index, 1)
    }
  }
}

// 保存队伍
const saveTeam = () => {
  if (!teamFormData.value.name) {
    alert('请填写队伍名称')
    return
  }

  if (isTeamEditMode.value && editingTeam.value) {
    // 更新队伍
    updateTeam(editingTeam.value.id, {
      name: teamFormData.value.name,
      avatar: teamFormData.value.avatar,
      description: teamFormData.value.description,
      members: teamFormData.value.members,
    })
  } else {
    // 创建新队伍
    const firstUser = userList.value[0]
    if (firstUser) {
      const newTeam = createTeam({
        name: teamFormData.value.name!,
        avatar: teamFormData.value.avatar,
        description: teamFormData.value.description || '',
        creatorId: firstUser.id,
        creatorUsername: firstUser.username,
        creatorAvatar: firstUser.avatar,
      })
      if (newTeam) {
        alert(`队伍创建成功！邀请码：${newTeam.inviteCode}`)
      }
    }
  }
  closeTeamDialog()
}

// 删除队伍
const handleDeleteTeam = (team: Team) => {
  if (confirm(`确定要删除队伍 "${team.name}" 吗？`)) {
    deleteTeam(team.id)
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
          v-model="teamSearchQuery"
          type="text"
          placeholder="搜索队伍名、简介或队员..."
          class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 pr-4 pl-9 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
        />
      </div>
      <button
        class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
        @click="openCreateTeamDialog"
      >
        <Plus class="h-4 w-4" />
        新建队伍
      </button>
    </div>

    <!-- 队伍列表 -->
    <div class="overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-slate-50 dark:bg-slate-800/50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                序号
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                头像
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                队伍名
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                队员
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                队伍简介
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                邀请码
              </th>
              <th class="px-6 py-3 text-right text-xs font-bold uppercase tracking-wider text-slate-500">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200 bg-white dark:divide-slate-800 dark:bg-slate-900">
            <tr
              v-for="(team, index) in paginatedTeams"
              :key="team.id"
              class="transition-colors hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-900 dark:text-slate-100">
                {{ (teamCurrentPage - 1) * teamPageSize + index + 1 }}
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <img
                  :src="team.avatar || `https://api.dicebear.com/7.x/shapes/svg?seed=${team.name}`"
                  :alt="team.name"
                  class="h-10 w-10 rounded-full object-cover"
                />
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm font-medium text-slate-900 dark:text-slate-100">
                {{ team.name }}
              </td>
              <td class="px-6 py-4">
                <div class="flex -space-x-2">
                  <img
                    v-for="member in team.members.slice(0, 3)"
                    :key="member.id"
                    :src="member.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${member.username}`"
                    :alt="member.username"
                    class="h-8 w-8 rounded-full border-2 border-white object-cover dark:border-slate-900"
                    :title="member.username"
                  />
                  <div
                    v-if="team.members.length > 3"
                    class="flex h-8 w-8 items-center justify-center rounded-full border-2 border-white bg-slate-100 text-xs font-medium text-slate-600 dark:border-slate-900 dark:bg-slate-800 dark:text-slate-300"
                  >
                    +{{ team.members.length - 3 }}
                  </div>
                </div>
                <div v-if="team.members.length === 0" class="text-xs text-slate-400">
                  暂无队员
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                <div class="max-w-md truncate" :title="team.description">
                  {{ team.description || '暂无简介' }}
                </div>
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <code class="rounded bg-slate-100 px-2 py-1 text-xs font-mono font-bold text-slate-700 dark:bg-slate-800 dark:text-slate-300">
                  {{ team.inviteCode || 'N/A' }}
                </code>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-2">
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
                    title="编辑"
                    @click="openEditTeamDialog(team)"
                  >
                    <Edit class="h-4 w-4" />
                  </button>
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                    title="删除"
                    @click="handleDeleteTeam(team)"
                  >
                    <Trash2 class="h-4 w-4" />
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="paginatedTeams.length === 0">
              <td colspan="7" class="px-6 py-12 text-center text-sm text-slate-500">
                暂无队伍数据
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 队伍分页 -->
    <div v-if="teamTotalPages > 1" class="mt-4 flex items-center justify-between">
      <div class="text-sm text-slate-600 dark:text-slate-400">
        共 {{ filteredTeams.length }} 条记录，第 {{ teamCurrentPage }} / {{ teamTotalPages }} 页
      </div>
      <div class="flex items-center gap-2">
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="teamCurrentPage === 1"
          @click="teamCurrentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>
        <span class="px-3 py-2 text-sm font-medium text-slate-700 dark:text-slate-300">
          {{ teamCurrentPage }} / {{ teamTotalPages }}
        </span>
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="teamCurrentPage === teamTotalPages"
          @click="teamCurrentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>

    <!-- 创建/编辑队伍对话框 -->
    <div
      v-if="showTeamDialog"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="closeTeamDialog"
    >
      <div
        class="w-full max-w-2xl rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900"
      >
        <!-- 对话框头部 -->
        <div class="flex items-center justify-between border-b border-slate-200 px-6 py-4 dark:border-slate-800">
          <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100">
            {{ isTeamEditMode ? '编辑队伍' : '新建队伍' }}
          </h3>
          <button
            class="rounded-lg p-1 text-slate-400 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:hover:bg-slate-800 dark:hover:text-slate-100"
            @click="closeTeamDialog"
          >
            <X class="h-5 w-5" />
          </button>
        </div>

        <!-- 对话框内容 -->
        <div class="px-6 py-4">
          <div class="space-y-4">
            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队伍名 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="teamFormData.name"
                type="text"
                class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="请输入队伍名"
              />
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                头像
              </label>
              <div class="space-y-3">
                <div v-if="teamAvatarPreview" class="flex items-center gap-4">
                  <img
                    :src="teamAvatarPreview"
                    alt="头像预览"
                    class="h-20 w-20 rounded-full object-cover border-2 border-slate-200 dark:border-slate-700"
                  />
                  <button
                    type="button"
                    class="rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                    @click="clearTeamAvatar"
                  >
                    清除
                  </button>
                </div>
                <label
                  class="flex cursor-pointer items-center justify-center rounded-lg border-2 border-dashed border-slate-300 bg-slate-50 px-4 py-3 text-sm font-medium text-slate-600 transition-colors hover:border-blue-500 hover:bg-blue-50 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-400 dark:hover:border-blue-500 dark:hover:bg-slate-700"
                >
                  <input
                    ref="teamAvatarInputRef"
                    type="file"
                    accept="image/*"
                    class="hidden"
                    @change="handleTeamAvatarChange"
                  />
                  <span>{{ teamAvatarPreview ? '更换头像' : '上传头像' }}</span>
                </label>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  支持 JPG、PNG、GIF 格式，大小不超过 5MB
                </p>
              </div>
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队伍简介
              </label>
              <textarea
                v-model="teamFormData.description"
                rows="3"
                class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="请输入队伍简介"
              />
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                添加队员
              </label>
              <select
                class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                @change="(e) => {
                  const target = e.target as HTMLSelectElement
                  if (target.value) {
                    addTeamMember(Number(target.value))
                    target.value = ''
                  }
                }"
              >
                <option value="">
                  选择用户添加到队伍
                </option>
                <option
                  v-for="user in availableUsers"
                  :key="user.id"
                  :value="user.id"
                  :disabled="teamFormData.members?.some(m => m.id === user.id)"
                >
                  {{ user.username }} ({{ user.realName || user.email }})
                </option>
              </select>
            </div>

            <div v-if="teamFormData.members && teamFormData.members.length > 0">
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                当前队员 ({{ teamFormData.members.length }})
              </label>
              <div class="space-y-2">
                <div
                  v-for="member in teamFormData.members"
                  :key="member.id"
                  class="flex items-center justify-between rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 dark:border-slate-700 dark:bg-slate-800"
                >
                  <div class="flex items-center gap-2">
                    <img
                      :src="member.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${member.username}`"
                      :alt="member.username"
                      class="h-6 w-6 rounded-full object-cover"
                    />
                    <span class="text-sm font-medium text-slate-700 dark:text-slate-300">{{ member.username }}</span>
                  </div>
                  <button
                    class="rounded-lg p-1 text-slate-400 transition-colors hover:bg-slate-200 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                    title="移除"
                    @click="removeTeamMember(member.id)"
                  >
                    <X class="h-4 w-4" />
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 对话框底部 -->
        <div class="flex items-center justify-end gap-3 border-t border-slate-200 px-6 py-4 dark:border-slate-800">
          <button
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
            @click="closeTeamDialog"
          >
            取消
          </button>
          <button
            class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="saveTeam"
          >
            <Save class="h-4 w-4" />
            保存
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
