<script setup lang="ts">
import { ref, computed } from 'vue'
import { Users, Copy, Check, X, Edit, Save } from 'lucide-vue-next'
import { createTeam, joinTeamByInviteCode, getUserTeams, leaveTeam, updateTeam } from '../stores/teamStore'
import type { Team } from '../types/team'

// 当前用户ID（模拟，实际应该从登录状态获取）
const currentUserId = ref(2) // 假设当前用户是lisi
const currentUsername = ref('lisi')
const currentUserAvatar = ref('https://api.dicebear.com/7.x/avataaars/svg?seed=lisi')

// 标签页
const activeTab = ref<'myTeams' | 'join' | 'create'>('myTeams')

// 我的队伍列表
const myTeams = computed(() => getUserTeams(currentUserId.value))

// 加入队伍相关
const inviteCode = ref('')
const joinError = ref('')
const joinSuccess = ref(false)

// 创建队伍相关
const showCreateDialog = ref(false)
const createForm = ref({
  name: '',
  description: '',
  avatar: '',
})
const createAvatarPreview = ref('')
const createAvatarFile = ref<File | null>(null)
const createAvatarInputRef = ref<HTMLInputElement | null>(null)

// 处理创建队伍头像文件选择
const handleCreateAvatarChange = (event: Event) => {
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
    createAvatarFile.value = file
    const reader = new FileReader()
    reader.onload = (e) => {
      createAvatarPreview.value = e.target?.result as string
      createForm.value.avatar = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

// 清除创建队伍头像
const clearCreateAvatar = () => {
  createAvatarPreview.value = ''
  createAvatarFile.value = null
  createForm.value.avatar = ''
  if (createAvatarInputRef.value) {
    createAvatarInputRef.value.value = ''
  }
}

// 编辑队伍相关
const showEditDialog = ref(false)
const editingTeam = ref<Team | null>(null)
const editForm = ref({
  name: '',
  description: '',
  avatar: '',
})
const editAvatarPreview = ref('')
const editAvatarFile = ref<File | null>(null)
const editAvatarInputRef = ref<HTMLInputElement | null>(null)

// 处理编辑队伍头像文件选择
const handleEditAvatarChange = (event: Event) => {
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
    editAvatarFile.value = file
    const reader = new FileReader()
    reader.onload = (e) => {
      editAvatarPreview.value = e.target?.result as string
      editForm.value.avatar = e.target?.result as string
    }
    reader.readAsDataURL(file)
  }
}

// 清除编辑队伍头像
const clearEditAvatar = () => {
  editAvatarPreview.value = ''
  editAvatarFile.value = null
  editForm.value.avatar = ''
  if (editAvatarInputRef.value) {
    editAvatarInputRef.value.value = ''
  }
}

// 打开编辑对话框
const openEditDialog = (team: Team) => {
  editingTeam.value = team
  editForm.value = {
    name: team.name,
    description: team.description,
    avatar: team.avatar || '',
  }
  editAvatarPreview.value = team.avatar || ''
  editAvatarFile.value = null
  showEditDialog.value = true
}

// 关闭编辑对话框
const closeEditDialog = () => {
  showEditDialog.value = false
  editingTeam.value = null
  editForm.value = {
    name: '',
    description: '',
    avatar: '',
  }
  editAvatarPreview.value = ''
  editAvatarFile.value = null
}

// 保存编辑
const handleSaveEdit = () => {
  if (!editingTeam.value) return

  if (!editForm.value.name.trim()) {
    alert('请填写队伍名称')
    return
  }

  const success = updateTeam(editingTeam.value.id, {
    name: editForm.value.name,
    description: editForm.value.description,
    avatar: editForm.value.avatar || editingTeam.value.avatar,
  })

  if (success) {
    closeEditDialog()
    alert('队伍信息已更新')
  } else {
    alert('更新失败')
  }
}

// 加入队伍
const handleJoinTeam = () => {
  joinError.value = ''
  joinSuccess.value = false

  if (!inviteCode.value.trim()) {
    joinError.value = '请输入邀请码'
    return
  }

  const success = joinTeamByInviteCode(
    inviteCode.value.trim().toUpperCase(),
    currentUserId.value,
    currentUsername.value,
    currentUserAvatar.value
  )

  if (success) {
    joinSuccess.value = true
    inviteCode.value = ''
    setTimeout(() => {
      joinSuccess.value = false
      activeTab.value = 'myTeams'
    }, 2000)
  } else {
    joinError.value = '邀请码无效或您已经是该队伍的成员'
  }
}

// 创建队伍
const handleCreateTeam = () => {
  if (!createForm.value.name.trim()) {
    alert('请填写队伍名称')
    return
  }

  const newTeam = createTeam({
    name: createForm.value.name,
    description: createForm.value.description,
    avatar: createForm.value.avatar,
    creatorId: currentUserId.value,
    creatorUsername: currentUsername.value,
    creatorAvatar: currentUserAvatar.value,
  })

  if (newTeam) {
    showCreateDialog.value = false
    createForm.value = {
      name: '',
      description: '',
      avatar: '',
    }
    createAvatarPreview.value = ''
    createAvatarFile.value = null
    activeTab.value = 'myTeams'
    alert(`队伍创建成功！邀请码：${newTeam.inviteCode}`)
  }
}

// 离开队伍
const handleLeaveTeam = (team: Team) => {
  const isCreator = team.creatorId === currentUserId.value
  const message = isCreator
    ? `确定要解散队伍 "${team.name}" 吗？解散后所有成员将被移除。`
    : `确定要离开队伍 "${team.name}" 吗？`
  
  if (confirm(message)) {
    const success = leaveTeam(team.id, currentUserId.value)
    if (success) {
      alert(isCreator ? '队伍已解散' : '已成功离开队伍')
    } else {
      alert('操作失败')
    }
  }
}

// 复制邀请码
const copyInviteCode = (code: string) => {
  navigator.clipboard.writeText(code)
  alert('邀请码已复制到剪贴板')
}
</script>

<template>
  <div class="min-h-[calc(100vh-64px)] p-6">
    <div class="mx-auto max-w-6xl">
      <!-- 标签页 -->
      <div class="mb-6 flex gap-2 border-b border-slate-200 dark:border-slate-800">
        <button
          class="px-4 py-2 text-sm font-medium transition-colors"
          :class="
            activeTab === 'myTeams'
              ? 'border-b-2 border-blue-600 text-blue-600 dark:border-blue-400'
              : 'text-slate-500 hover:text-slate-700 dark:text-slate-400 dark:hover:text-slate-200'
          "
          @click="activeTab = 'myTeams'"
        >
          我的队伍
        </button>
        <button
          class="px-4 py-2 text-sm font-medium transition-colors"
          :class="
            activeTab === 'join'
              ? 'border-b-2 border-blue-600 text-blue-600 dark:border-blue-400'
              : 'text-slate-500 hover:text-slate-700 dark:text-slate-400 dark:hover:text-slate-200'
          "
          @click="activeTab = 'join'"
        >
          加入队伍
        </button>
        <button
          class="px-4 py-2 text-sm font-medium transition-colors"
          :class="
            activeTab === 'create'
              ? 'border-b-2 border-blue-600 text-blue-600 dark:border-blue-400'
              : 'text-slate-500 hover:text-slate-700 dark:text-slate-400 dark:hover:text-slate-200'
          "
          @click="activeTab = 'create'"
        >
          新增队伍
        </button>
      </div>

      <!-- 我的队伍 -->
      <div v-if="activeTab === 'myTeams'" class="space-y-4">
        <div v-if="myTeams.length === 0" class="rounded-xl border border-slate-200 bg-white p-12 text-center dark:border-slate-800 dark:bg-slate-900">
          <Users class="mx-auto mb-4 h-12 w-12 text-slate-400" />
          <p class="text-slate-500 dark:text-slate-400">
            您还没有加入任何队伍
          </p>
        </div>

        <div v-else class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <div
            v-for="team in myTeams"
            :key="team.id"
            class="flex flex-col rounded-xl border border-slate-200 bg-white p-4 shadow-sm transition-shadow hover:shadow-md dark:border-slate-800 dark:bg-slate-900"
          >
            <!-- 队伍名称和简介部分（可伸缩） -->
            <div class="flex-1">
              <div class="mb-3 flex items-center gap-3">
                <img
                  :src="team.avatar || `https://api.dicebear.com/7.x/shapes/svg?seed=${team.name}`"
                  :alt="team.name"
                  class="h-12 w-12 rounded-full object-cover"
                />
                <div class="flex-1">
                  <h3 class="font-bold text-slate-900 dark:text-slate-50">
                    {{ team.name }}
                  </h3>
                  <p class="text-xs text-slate-500 dark:text-slate-400">
                    {{ team.members.length }} 名成员
                  </p>
                </div>
              </div>

              <p class="mb-3 line-clamp-2 text-sm text-slate-600 dark:text-slate-300">
                {{ team.description }}
              </p>
            </div>

            <!-- 底部固定内容：队员头像和邀请码 -->
            <div class="mt-auto space-y-3">
              <div class="flex -space-x-2">
                <img
                  v-for="member in team.members.slice(0, 5)"
                  :key="member.id"
                  :src="member.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${member.username}`"
                  :alt="member.username"
                  class="h-6 w-6 rounded-full border-2 border-white object-cover dark:border-slate-900"
                  :title="member.username"
                />
                <div
                  v-if="team.members.length > 5"
                  class="flex h-6 w-6 items-center justify-center rounded-full border-2 border-white bg-slate-100 text-xs font-medium text-slate-600 dark:border-slate-900 dark:bg-slate-800 dark:text-slate-300"
                >
                  +{{ team.members.length - 5 }}
                </div>
              </div>

              <div class="flex items-center justify-between gap-2">
                <div class="flex items-center gap-2">
                  <span class="text-xs text-slate-500 dark:text-slate-400">邀请码：</span>
                  <code class="rounded bg-slate-100 px-2 py-1 text-xs font-mono font-bold text-slate-700 dark:bg-slate-800 dark:text-slate-300">
                    {{ team.inviteCode }}
                  </code>
                  <button
                    class="rounded p-1 text-slate-400 hover:text-blue-600 dark:hover:text-blue-400"
                    title="复制邀请码"
                    @click="copyInviteCode(team.inviteCode || '')"
                  >
                    <Copy class="h-3 w-3" />
                  </button>
                </div>
                <div class="flex items-center gap-2">
                  <button
                    class="rounded-lg border border-blue-200 bg-blue-50 px-3 py-1 text-xs font-medium text-blue-600 transition-colors hover:bg-blue-100 dark:border-blue-800 dark:bg-blue-900/20 dark:text-blue-400 dark:hover:bg-blue-900/30"
                    title="编辑队伍"
                    @click="openEditDialog(team)"
                  >
                    <Edit class="h-3 w-3 inline" />
                  </button>
                  <button
                    class="rounded-lg border border-red-200 bg-red-50 px-3 py-1 text-xs font-medium text-red-600 transition-colors hover:bg-red-100 dark:border-red-800 dark:bg-red-900/20 dark:text-red-400 dark:hover:bg-red-900/30"
                    @click="handleLeaveTeam(team)"
                  >
                    {{ team.creatorId === currentUserId ? '解散' : '离开' }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 加入队伍 -->
      <div v-if="activeTab === 'join'" class="rounded-xl border border-slate-200 bg-white p-6 dark:border-slate-800 dark:bg-slate-900">
        <h2 class="mb-4 text-xl font-bold text-slate-900 dark:text-slate-50">
          通过邀请码加入队伍
        </h2>
        <div class="space-y-4">
          <div>
            <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              邀请码
            </label>
            <input
              v-model="inviteCode"
              type="text"
              placeholder="请输入邀请码"
              class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm uppercase outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              @keyup.enter="handleJoinTeam"
            />
            <p v-if="joinError" class="mt-2 text-sm text-red-600 dark:text-red-400">
              {{ joinError }}
            </p>
            <p v-if="joinSuccess" class="mt-2 flex items-center gap-2 text-sm text-green-600 dark:text-green-400">
              <Check class="h-4 w-4" />
              加入成功！
            </p>
          </div>
          <button
            class="w-full rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="handleJoinTeam"
          >
            加入队伍
          </button>
        </div>
      </div>

      <!-- 新增队伍 -->
      <div v-if="activeTab === 'create'" class="rounded-xl border border-slate-200 bg-white p-6 dark:border-slate-800 dark:bg-slate-900">
        <h2 class="mb-4 text-xl font-bold text-slate-900 dark:text-slate-50">
          创建新队伍
        </h2>
        <div class="space-y-4">
          <div>
            <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              队伍名称 <span class="text-red-500">*</span>
            </label>
            <input
              v-model="createForm.name"
              type="text"
              placeholder="请输入队伍名称"
              class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            />
          </div>

          <div>
            <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              队伍简介
            </label>
            <textarea
              v-model="createForm.description"
              rows="3"
              placeholder="请输入队伍简介"
              class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            />
          </div>

          <div>
            <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              队伍头像
            </label>
            <div class="space-y-3">
              <div v-if="createAvatarPreview" class="flex items-center gap-4">
                <img
                  :src="createAvatarPreview"
                  alt="头像预览"
                  class="h-20 w-20 rounded-full object-cover border-2 border-slate-200 dark:border-slate-700"
                />
                <button
                  type="button"
                  class="rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                  @click="clearCreateAvatar"
                >
                  清除
                </button>
              </div>
              <label
                class="flex cursor-pointer items-center justify-center rounded-lg border-2 border-dashed border-slate-300 bg-slate-50 px-4 py-3 text-sm font-medium text-slate-600 transition-colors hover:border-blue-500 hover:bg-blue-50 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-400 dark:hover:border-blue-500 dark:hover:bg-slate-700"
              >
                <input
                  ref="createAvatarInputRef"
                  type="file"
                  accept="image/*"
                  class="hidden"
                  @change="handleCreateAvatarChange"
                />
                <span>{{ createAvatarPreview ? '更换头像' : '上传头像' }}</span>
              </label>
              <p class="text-xs text-slate-500 dark:text-slate-400">
                支持 JPG、PNG、GIF 格式，大小不超过 5MB
              </p>
            </div>
          </div>

          <button
            class="w-full rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="handleCreateTeam"
          >
            创建队伍
          </button>
        </div>
      </div>
    </div>

    <!-- 编辑队伍对话框 -->
    <div
      v-if="showEditDialog"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="closeEditDialog"
    >
      <div
        class="w-full max-w-2xl rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900"
      >
        <!-- 对话框头部 -->
        <div class="flex items-center justify-between border-b border-slate-200 px-6 py-4 dark:border-slate-800">
          <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100">
            编辑队伍信息
          </h3>
          <button
            class="rounded-lg p-1 text-slate-400 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:hover:bg-slate-800 dark:hover:text-slate-100"
            @click="closeEditDialog"
          >
            <X class="h-5 w-5" />
          </button>
        </div>

        <!-- 对话框内容 -->
        <div class="px-6 py-4">
          <div class="space-y-4">
            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队伍名称 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="editForm.name"
                type="text"
                placeholder="请输入队伍名称"
                class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队伍简介
              </label>
              <textarea
                v-model="editForm.description"
                rows="3"
                placeholder="请输入队伍简介"
                class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队伍头像
              </label>
              <div class="space-y-3">
                <div v-if="editAvatarPreview" class="flex items-center gap-4">
                  <img
                    :src="editAvatarPreview"
                    alt="头像预览"
                    class="h-20 w-20 rounded-full object-cover border-2 border-slate-200 dark:border-slate-700"
                  />
                  <button
                    type="button"
                    class="rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                    @click="clearEditAvatar"
                  >
                    清除
                  </button>
                </div>
                <label
                  class="flex cursor-pointer items-center justify-center rounded-lg border-2 border-dashed border-slate-300 bg-slate-50 px-4 py-3 text-sm font-medium text-slate-600 transition-colors hover:border-blue-500 hover:bg-blue-50 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-400 dark:hover:border-blue-500 dark:hover:bg-slate-700"
                >
                  <input
                    ref="editAvatarInputRef"
                    type="file"
                    accept="image/*"
                    class="hidden"
                    @change="handleEditAvatarChange"
                  />
                  <span>{{ editAvatarPreview ? '更换头像' : '上传头像' }}</span>
                </label>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  支持 JPG、PNG、GIF 格式，大小不超过 5MB
                </p>
              </div>
            </div>

            <div v-if="editingTeam">
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                邀请码
              </label>
              <div class="flex items-center gap-2">
                <code class="rounded bg-slate-100 px-3 py-2 text-sm font-mono font-bold text-slate-700 dark:bg-slate-800 dark:text-slate-300">
                  {{ editingTeam.inviteCode }}
                </code>
                <button
                  class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                  title="复制邀请码"
                  @click="copyInviteCode(editingTeam.inviteCode || '')"
                >
                  <Copy class="h-4 w-4" />
                </button>
              </div>
            </div>

            <div v-if="editingTeam">
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                队员列表 ({{ editingTeam.members.length }})
              </label>
              <div class="max-h-60 space-y-2 overflow-y-auto rounded-lg border border-slate-200 bg-slate-50 p-3 dark:border-slate-700 dark:bg-slate-800">
                <div
                  v-for="member in editingTeam.members"
                  :key="member.id"
                  class="flex items-center gap-3 rounded-lg border border-slate-200 bg-white p-2 dark:border-slate-700 dark:bg-slate-900"
                >
                  <img
                    :src="member.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${member.username}`"
                    :alt="member.username"
                    class="h-10 w-10 rounded-full object-cover"
                  />
                  <div class="flex-1">
                    <p class="text-sm font-medium text-slate-900 dark:text-slate-50">
                      {{ member.username }}
                    </p>
                    <p
                      v-if="editingTeam.creatorId === member.id"
                      class="text-xs text-blue-600 dark:text-blue-400"
                    >
                      创建者
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 对话框底部 -->
        <div class="flex items-center justify-end gap-3 border-t border-slate-200 px-6 py-4 dark:border-slate-800">
          <button
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
            @click="closeEditDialog"
          >
            取消
          </button>
          <button
            class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="handleSaveEdit"
          >
            <Save class="h-4 w-4" />
            保存
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

