<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { Search, Plus, Edit, Trash2, X, Save, ChevronLeft, ChevronRight } from 'lucide-vue-next'
import {
  announcementStore,
  updateAnnouncement,
  createAnnouncement,
  deleteAnnouncement,
} from '../../../stores/announcementStore'
import type { Announcement } from '../../../types/announcement'

// ========== 公告管理相关逻辑 ==========
// 使用共享的公告数据
const announcementList = computed(() => announcementStore.announcements)

// 公告搜索
const announcementSearchQuery = ref('')

// 公告分页
const announcementPageSize = ref(10)
const announcementCurrentPage = ref(1)

// 过滤后的公告列表
const filteredAnnouncements = computed(() => {
  let result = announcementList.value
  if (announcementSearchQuery.value.trim()) {
    const query = announcementSearchQuery.value.toLowerCase().trim()
    result = result.filter(
      (announcement) =>
        announcement.title.toLowerCase().includes(query) ||
        announcement.content.toLowerCase().includes(query) ||
        announcement.author.toLowerCase().includes(query),
    )
  }
  return result
})

// 分页后的公告列表
const paginatedAnnouncements = computed(() => {
  const start = (announcementCurrentPage.value - 1) * announcementPageSize.value
  const end = start + announcementPageSize.value
  return filteredAnnouncements.value.slice(start, end)
})

// 总页数
const announcementTotalPages = computed(() => {
  return Math.ceil(filteredAnnouncements.value.length / announcementPageSize.value)
})

// 当搜索条件改变时，重置到第一页
watch(announcementSearchQuery, () => {
  announcementCurrentPage.value = 1
})

// 监听总页数变化，确保当前页码在有效范围内
watch(announcementTotalPages, (newTotalPages) => {
  if (announcementCurrentPage.value > newTotalPages && newTotalPages > 0) {
    announcementCurrentPage.value = newTotalPages
  }
})

// 编辑/创建公告对话框
const showAnnouncementDialog = ref(false)
const editingAnnouncement = ref<Announcement | null>(null)
const isAnnouncementEditMode = computed(() => editingAnnouncement.value !== null)

// 公告表单数据
const announcementFormData = ref({
  title: '',
  content: '',
  isPinned: false,
  status: 'published' as 'published' | 'draft',
})

// 打开创建公告对话框
const openCreateAnnouncementDialog = () => {
  editingAnnouncement.value = null
  announcementFormData.value = {
    title: '',
    content: '',
    isPinned: false,
    status: 'published',
  }
  showAnnouncementDialog.value = true
}

// 打开编辑公告对话框
const openEditAnnouncementDialog = (announcement: Announcement) => {
  editingAnnouncement.value = announcement
  announcementFormData.value = {
    title: announcement.title,
    content: announcement.content,
    isPinned: announcement.isPinned,
    status: announcement.status,
  }
  showAnnouncementDialog.value = true
}

// 关闭公告对话框
const closeAnnouncementDialog = () => {
  showAnnouncementDialog.value = false
  editingAnnouncement.value = null
  announcementFormData.value = {
    title: '',
    content: '',
    isPinned: false,
    status: 'published',
  }
}

// 保存公告
const saveAnnouncement = () => {
  if (!announcementFormData.value.title.trim()) {
    alert('请填写公告标题')
    return
  }
  if (!announcementFormData.value.content.trim()) {
    alert('请填写公告内容')
    return
  }

  if (isAnnouncementEditMode.value && editingAnnouncement.value) {
    // 更新公告
    updateAnnouncement(editingAnnouncement.value.id, {
      title: announcementFormData.value.title,
      content: announcementFormData.value.content,
      isPinned: announcementFormData.value.isPinned,
      status: announcementFormData.value.status,
    })
  } else {
    // 创建新公告
    createAnnouncement({
      title: announcementFormData.value.title,
      content: announcementFormData.value.content,
      isPinned: announcementFormData.value.isPinned,
      status: announcementFormData.value.status,
    })
  }
  closeAnnouncementDialog()
}

// 删除公告
const handleDeleteAnnouncement = (announcement: Announcement) => {
  if (confirm(`确定要删除公告 "${announcement.title}" 吗？`)) {
    const success = deleteAnnouncement(announcement.id)
    if (success) {
      // 如果当前页没有数据了，回到上一页
      if (paginatedAnnouncements.value.length === 0 && announcementCurrentPage.value > 1) {
        announcementCurrentPage.value--
      }
    }
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
          v-model="announcementSearchQuery"
          type="text"
          placeholder="搜索公告标题、内容或作者..."
          class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 pr-4 pl-9 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
        />
      </div>
      <button
        class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
        @click="openCreateAnnouncementDialog"
      >
        <Plus class="h-4 w-4" />
        新建公告
      </button>
    </div>

    <!-- 公告列表 -->
    <div class="overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-slate-50 dark:bg-slate-800/50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                序号
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                标题
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                状态
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                置顶
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                作者
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                创建时间
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                更新时间
              </th>
              <th class="px-6 py-3 text-right text-xs font-bold uppercase tracking-wider text-slate-500">
                操作
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-200 bg-white dark:divide-slate-800 dark:bg-slate-900">
            <tr
              v-for="(announcement, index) in paginatedAnnouncements"
              :key="announcement.id"
              class="transition-colors hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-900 dark:text-slate-100">
                {{ (announcementCurrentPage - 1) * announcementPageSize + index + 1 }}
              </td>
              <td class="px-6 py-4 text-sm font-medium text-slate-900 dark:text-slate-100">
                <div class="max-w-md truncate" :title="announcement.title">
                  {{ announcement.title }}
                </div>
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <span
                  class="inline-flex rounded-full px-2.5 py-0.5 text-xs font-semibold"
                  :class="
                    announcement.status === 'published'
                      ? 'bg-green-100 text-green-700 dark:bg-green-500/15 dark:text-green-300'
                      : 'bg-yellow-100 text-yellow-700 dark:bg-yellow-500/15 dark:text-yellow-300'
                  "
                >
                  {{ announcement.status === 'published' ? '已发布' : '草稿' }}
                </span>
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <span
                  v-if="announcement.isPinned"
                  class="inline-flex items-center gap-1 text-xs font-medium text-blue-600 dark:text-blue-400"
                >
                  <span>📌</span>
                  置顶
                </span>
                <span v-else class="text-xs text-slate-400">-</span>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ announcement.author }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ announcement.createdAt }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ announcement.updatedAt }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-2">
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
                    title="编辑"
                    @click="openEditAnnouncementDialog(announcement)"
                  >
                    <Edit class="h-4 w-4" />
                  </button>
                  <button
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                    title="删除"
                    @click="handleDeleteAnnouncement(announcement)"
                  >
                    <Trash2 class="h-4 w-4" />
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="paginatedAnnouncements.length === 0">
              <td colspan="8" class="px-6 py-12 text-center text-sm text-slate-500">
                暂无公告数据
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 公告分页 -->
    <div v-if="announcementTotalPages > 1" class="mt-4 flex items-center justify-between">
      <div class="text-sm text-slate-600 dark:text-slate-400">
        共 {{ filteredAnnouncements.length }} 条记录，第 {{ announcementCurrentPage }} / {{ announcementTotalPages }} 页
      </div>
      <div class="flex items-center gap-2">
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="announcementCurrentPage === 1"
          @click="announcementCurrentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>
        <span class="px-3 py-2 text-sm font-medium text-slate-700 dark:text-slate-300">
          {{ announcementCurrentPage }} / {{ announcementTotalPages }}
        </span>
        <button
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="announcementCurrentPage === announcementTotalPages"
          @click="announcementCurrentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>

    <!-- 创建/编辑公告对话框 -->
    <div
      v-if="showAnnouncementDialog"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="closeAnnouncementDialog"
    >
      <div
        class="w-full max-w-3xl rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900"
      >
        <!-- 对话框头部 -->
        <div class="flex items-center justify-between border-b border-slate-200 px-6 py-4 dark:border-slate-800">
          <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100">
            {{ isAnnouncementEditMode ? '编辑公告' : '新建公告' }}
          </h3>
          <button
            class="rounded-lg p-1 text-slate-400 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:hover:bg-slate-800 dark:hover:text-slate-100"
            @click="closeAnnouncementDialog"
          >
            <X class="h-5 w-5" />
          </button>
        </div>

        <!-- 对话框内容 -->
        <div class="px-6 py-4">
          <div class="space-y-4">
            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                公告标题 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="announcementFormData.title"
                type="text"
                placeholder="请输入公告标题"
                class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>

            <div>
              <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                公告内容 <span class="text-red-500">*</span>
              </label>
              <textarea
                v-model="announcementFormData.content"
                rows="6"
                placeholder="请输入公告内容"
                class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>

            <div class="flex items-center gap-6">
              <div class="flex items-center gap-2">
                <input
                  v-model="announcementFormData.isPinned"
                  type="checkbox"
                  class="h-4 w-4 rounded border-slate-300 text-blue-600 focus:ring-blue-500 dark:border-slate-600"
                />
                <label class="text-sm font-medium text-slate-700 dark:text-slate-300">
                  置顶
                </label>
              </div>

              <div>
                <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
                  状态
                </label>
                <select
                  v-model="announcementFormData.status"
                  class="h-10 rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                >
                  <option value="published">
                    已发布
                  </option>
                  <option value="draft">
                    草稿
                  </option>
                </select>
              </div>
            </div>
          </div>
        </div>

        <!-- 对话框底部 -->
        <div class="flex items-center justify-end gap-3 border-t border-slate-200 px-6 py-4 dark:border-slate-800">
          <button
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
            @click="closeAnnouncementDialog"
          >
            取消
          </button>
          <button
            class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="saveAnnouncement"
          >
            <Save class="h-4 w-4" />
            保存
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
