<script setup lang="ts">
import { ref, computed } from 'vue'
import { ChevronLeft, ChevronRight, Edit, Trash2, Pin } from 'lucide-vue-next'

type AnnounceStatus = 'published' | 'draft'

interface AnnounceItem {
  id: number
  title: string
  content: string
  status: AnnounceStatus
  pinned: boolean
  author: string
  createdAt: string
  updatedAt: string
}

const statusLabel: Record<AnnounceStatus, string> = {
  published: '已发布',
  draft: '草稿',
}

const statusClass: Record<AnnounceStatus, string> = {
  published: 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300',
  draft: 'bg-slate-100 text-slate-600 dark:bg-slate-700 dark:text-slate-300',
}

const titles = [
  '比赛规则说明',
  '题目环境已就绪',
  'Web 题 Hint 发布',
  '提交须知',
  '比赛时间调整通知',
  '答疑方式说明',
  '禁止多账号参赛',
  'Writeup 提交截止提醒',
  '积分榜冻结说明',
  '颁奖与证书发放',
  '题目纠错说明',
  '加时赛安排',
  '违规处理公示',
  '比赛圆满结束感谢',
  '赛后题解发布通知',
]
const authors = ['管理员', '系统', '裁判组', '出题组', '运营']

function buildInitialList(): AnnounceItem[] {
  return Array.from({ length: 48 }, (_, i) => ({
    id: i + 1,
    title: titles[i % titles.length] + (i >= titles.length ? `（${Math.floor(i / titles.length)}）` : ''),
    content: `此为公告「${titles[i % titles.length]}」的正文内容，支持 Markdown。`,
    status: (i % 5 === 0 ? 'draft' : 'published') as AnnounceStatus,
    pinned: i % 4 === 0,
    author: authors[i % authors.length] ?? '管理员',
    createdAt: `2025-03-${String(1 + (i % 15)).padStart(2, '0')} ${String(9 + (i % 10)).padStart(2, '0')}:${String((i * 11) % 60).padStart(2, '0')}`,
    updatedAt: `2025-03-${String(5 + (i % 18)).padStart(2, '0')} ${String(14 + (i % 6)).padStart(2, '0')}:${String((i * 13) % 60).padStart(2, '0')}`,
  }))
}

const announcementList = ref<AnnounceItem[]>(buildInitialList())

const pageSize = 10
const currentPage = ref(1)
const totalPages = computed(() => Math.ceil(announcementList.value.length / pageSize))
const paginatedList = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return announcementList.value.slice(start, start + pageSize)
})

// 编辑弹窗
const editModalVisible = ref(false)
const editForm = ref({ title: '', content: '', status: 'published' as AnnounceStatus, pinned: false })
let editingId: number | null = null

function openEdit(item: AnnounceItem) {
  editingId = item.id
  editForm.value = {
    title: item.title,
    content: item.content,
    status: item.status,
    pinned: item.pinned,
  }
  editModalVisible.value = true
}

function closeEditModal() {
  editModalVisible.value = false
  editingId = null
}

function formatNow() {
  const d = new Date()
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}`
}

function saveAnnounce() {
  if (!editingId) return
  const item = announcementList.value.find((a) => a.id === editingId)
  if (!item) return
  item.title = editForm.value.title.trim() || item.title
  item.content = editForm.value.content
  item.status = editForm.value.status
  item.pinned = editForm.value.pinned
  item.updatedAt = formatNow()
  closeEditModal()
}

function editAnnounce(item: AnnounceItem) {
  openEdit(item)
}

function deleteAnnounce(item: AnnounceItem) {
  if (!confirm(`确定删除公告「${item.title}」？`)) return
  const idx = announcementList.value.findIndex((a) => a.id === item.id)
  if (idx === -1) return
  announcementList.value.splice(idx, 1)
  const maxPage = Math.ceil(announcementList.value.length / pageSize) || 1
  if (currentPage.value > maxPage) currentPage.value = maxPage
}

// 发布公告弹窗
const publishModalVisible = ref(false)
const publishForm = ref({ title: '', content: '', status: 'published' as AnnounceStatus, pinned: false })

function openPublishModal() {
  publishForm.value = { title: '', content: '', status: 'published', pinned: false }
  publishModalVisible.value = true
}

function closePublishModal() {
  publishModalVisible.value = false
}

function getNextId(): number {
  const max = announcementList.value.reduce((m, a) => Math.max(m, a.id), 0)
  return max + 1
}

function submitPublish() {
  const title = publishForm.value.title.trim()
  if (!title) return
  const now = formatNow()
  const newItem: AnnounceItem = {
    id: getNextId(),
    title,
    content: publishForm.value.content,
    status: publishForm.value.status,
    pinned: publishForm.value.pinned,
    author: '管理员',
    createdAt: now,
    updatedAt: now,
  }
  announcementList.value.unshift(newItem)
  closePublishModal()
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-4 flex items-center justify-between">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          比赛公告
        </h3>
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="openPublishModal"
        >
          <Edit class="h-4 w-4" />
          发布公告
        </button>
      </div>

      <div class="overflow-x-auto rounded-lg border border-slate-200 dark:border-slate-700">
        <table class="w-full min-w-[800px] text-sm">
          <thead>
            <tr class="border-b border-slate-200 bg-slate-50 dark:border-slate-700 dark:bg-slate-800/60">
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 w-14 whitespace-nowrap">
                序号
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 whitespace-nowrap">
                标题
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 w-20 whitespace-nowrap">
                状态
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 w-16 whitespace-nowrap">
                置顶
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 w-20 whitespace-nowrap">
                作者
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 whitespace-nowrap">
                创建时间
              </th>
              <th class="px-4 py-3 text-left font-bold text-slate-600 dark:text-slate-300 whitespace-nowrap">
                更新时间
              </th>
              <th class="px-4 py-3 text-right font-bold text-slate-600 dark:text-slate-300 w-28 whitespace-nowrap">
                操作
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(item, idx) in paginatedList"
              :key="item.id"
              class="border-b border-slate-100 transition-colors hover:bg-slate-50 dark:border-slate-800 dark:hover:bg-slate-800/50"
            >
              <td class="px-4 py-3 text-slate-600 dark:text-slate-400 whitespace-nowrap">
                {{ (currentPage - 1) * pageSize + idx + 1 }}
              </td>
              <td class="max-w-xs px-4 py-3 font-medium text-slate-800 dark:text-slate-100 truncate whitespace-nowrap" :title="item.title">
                {{ item.title }}
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                <span
                  class="inline-flex rounded-full px-2 py-0.5 text-xs font-medium"
                  :class="statusClass[item.status]"
                >
                  {{ statusLabel[item.status] }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap">
                <span v-if="item.pinned" class="inline-flex items-center gap-1 text-amber-600 dark:text-amber-400">
                  <Pin class="h-3.5 w-3.5 shrink-0" /> 是
                </span>
                <span v-else class="text-slate-400 dark:text-slate-500">否</span>
              </td>
              <td class="px-4 py-3 text-slate-600 dark:text-slate-300 whitespace-nowrap">
                {{ item.author }}
              </td>
              <td class="px-4 py-3 text-slate-500 dark:text-slate-400 whitespace-nowrap">
                {{ item.createdAt }}
              </td>
              <td class="px-4 py-3 text-slate-500 dark:text-slate-400 whitespace-nowrap">
                {{ item.updatedAt }}
              </td>
              <td class="px-4 py-3 text-right whitespace-nowrap">
                <div class="flex items-center justify-end gap-1">
                  <button
                    type="button"
                    class="rounded-lg p-2 text-slate-500 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:hover:bg-slate-700 dark:hover:text-blue-400"
                    title="编辑"
                    @click="editAnnounce(item)"
                  >
                    <Edit class="h-4 w-4" />
                  </button>
                  <button
                    type="button"
                    class="rounded-lg p-2 text-slate-500 transition-colors hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                    title="删除"
                    @click="deleteAnnounce(item)"
                  >
                    <Trash2 class="h-4 w-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="totalPages > 1" class="mt-4 flex items-center justify-between">
        <span class="text-sm text-slate-500 dark:text-slate-400">
          共 {{ announcementList.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
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

    <!-- 发布公告弹窗 -->
    <a-modal
      v-model:visible="publishModalVisible"
      title="发布公告"
      width="520px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closePublishModal"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">标题</label>
          <input
            v-model="publishForm.title"
            type="text"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            placeholder="公告标题"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">内容</label>
          <textarea
            v-model="publishForm.content"
            rows="5"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            placeholder="公告正文，支持 Markdown"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">状态</label>
          <select
            v-model="publishForm.status"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          >
            <option value="published">
              已发布
            </option>
            <option value="draft">
              草稿
            </option>
          </select>
        </div>
        <div class="flex items-center justify-between rounded-lg border border-slate-200 bg-slate-50 px-4 py-3 dark:border-slate-700 dark:bg-slate-800">
          <span class="text-sm font-medium text-slate-700 dark:text-slate-200">置顶</span>
          <label class="relative inline-flex cursor-pointer items-center">
            <input v-model="publishForm.pinned" type="checkbox" class="peer sr-only" />
            <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
          </label>
        </div>
        <div class="flex justify-end gap-2 pt-2">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="closePublishModal"
          >
            取消
          </button>
          <button
            type="button"
            class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="submitPublish"
          >
            发布
          </button>
        </div>
      </div>
    </a-modal>

    <!-- 编辑公告弹窗 -->
    <a-modal
      v-model:visible="editModalVisible"
      title="编辑公告"
      width="520px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeEditModal"
    >
      <div class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">标题</label>
          <input
            v-model="editForm.title"
            type="text"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            placeholder="公告标题"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">内容</label>
          <textarea
            v-model="editForm.content"
            rows="5"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            placeholder="公告正文，支持 Markdown"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">状态</label>
          <select
            v-model="editForm.status"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          >
            <option value="published">
              已发布
            </option>
            <option value="draft">
              草稿
            </option>
          </select>
        </div>
        <div class="flex items-center justify-between rounded-lg border border-slate-200 bg-slate-50 px-4 py-3 dark:border-slate-700 dark:bg-slate-800">
          <span class="text-sm font-medium text-slate-700 dark:text-slate-200">置顶</span>
          <label class="relative inline-flex cursor-pointer items-center">
            <input v-model="editForm.pinned" type="checkbox" class="peer sr-only" />
            <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
          </label>
        </div>
        <div class="flex justify-end gap-2 pt-2">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="closeEditModal"
          >
            取消
          </button>
          <button
            type="button"
            class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="saveAnnounce"
          >
            保存
          </button>
        </div>
      </div>
    </a-modal>
  </div>
</template>
