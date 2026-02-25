<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { Search, Plus, Edit, Trash2, X, ChevronLeft, ChevronRight, MessageSquare, ThumbsUp, CheckCircle, XCircle, Eye } from 'lucide-vue-next'
import {
  forumStore,
  getForumPostsForAdmin,
  deleteForumPost,
  setPostStatus,
} from '../../../stores/forumStore'
import type { ForumPost } from '../../../types/forum'
import type { PostStatus } from '../../../types/forum'

const router = useRouter()

const postList = computed(() => getForumPostsForAdmin())

const searchQuery = ref('')
const statusFilter = ref<'all' | PostStatus>('all')

const pageSize = ref(10)
const currentPage = ref(1)

const pendingCount = computed(() =>
  postList.value.filter((p) => (p.status ?? 'published') === 'pending').length,
)
const publishedCount = computed(() =>
  postList.value.filter((p) => (p.status ?? 'published') === 'published').length,
)
const rejectedCount = computed(() =>
  postList.value.filter((p) => (p.status ?? 'published') === 'rejected').length,
)

const filteredPosts = computed(() => {
  let result = postList.value
  if (statusFilter.value !== 'all') {
    const s = statusFilter.value
    result = result.filter((p) => (p.status ?? 'published') === s)
  }
  if (searchQuery.value.trim()) {
    const q = searchQuery.value.toLowerCase().trim()
    result = result.filter(
      (p) =>
        p.title.toLowerCase().includes(q) ||
        p.content.toLowerCase().includes(q) ||
        p.author.toLowerCase().includes(q),
    )
  }
  return result
})

const paginatedPosts = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  return filteredPosts.value.slice(start, start + pageSize.value)
})

const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredPosts.value.length / pageSize.value)),
)

watch([searchQuery, statusFilter], () => {
  currentPage.value = 1
})

watch(totalPages, (tp) => {
  if (currentPage.value > tp && tp > 0) currentPage.value = tp
})

function statusLabel(s: PostStatus | undefined): string {
  const v = s ?? 'published'
  return { pending: '待审核', published: '已通过', rejected: '已拒绝' }[v]
}

function setStatusFilter(v: 'all' | PostStatus) {
  statusFilter.value = v
}

function openCreate() {
  router.push('/admin/manage/forum/create')
}

function goToEdit(post: ForumPost) {
  router.push(`/admin/manage/forum/edit/${post.id}`)
}

function goToCommentReview(post: ForumPost) {
  router.push(`/admin/manage/forum/comments/${post.id}`)
}

function handleApprove(post: ForumPost) {
  setPostStatus(post.id, 'published', { reviewedBy: '管理员' })
}

const showViewDialog = ref(false)
const viewPost = ref<ForumPost | null>(null)

function openView(post: ForumPost) {
  viewPost.value = post
  showViewDialog.value = true
}

function closeViewDialog() {
  showViewDialog.value = false
  viewPost.value = null
}

const showRejectDialog = ref(false)
const rejectPost = ref<ForumPost | null>(null)
const rejectReasonInput = ref('')

function openRejectDialog(post: ForumPost) {
  rejectPost.value = post
  rejectReasonInput.value = ''
  showRejectDialog.value = true
}

function confirmReject() {
  if (rejectPost.value) {
    setPostStatus(rejectPost.value.id, 'rejected', { reviewedBy: '管理员', rejectReason: rejectReasonInput.value.trim() })
    showRejectDialog.value = false
    rejectPost.value = null
    rejectReasonInput.value = ''
  }
}

function handleReject(post: ForumPost) {
  openRejectDialog(post)
}

function handleDelete(post: ForumPost) {
  if (!confirm(`确定要删除帖子「${post.title}」吗？该帖下的所有评论将一并删除。`)) return
  const ok = deleteForumPost(post.id)
  if (ok && paginatedPosts.value.length <= 1 && currentPage.value > 1) {
    currentPage.value--
  }
}

function replyCount(postId: number) {
  return forumStore.comments.filter((c) => c.postId === postId).length
}

/** 某帖子待审核评论数（用于红点提示） */
function pendingCommentCount(postId: number) {
  return forumStore.comments.filter(
    (c) => c.postId === postId && (c.status ?? 'published') === 'pending'
  ).length
}

function likeCount(post: ForumPost) {
  return post.likedByUserIds?.length ?? 0
}
</script>

<template>
  <div class="space-y-4">
    <!-- 审核概览：体现审核过程 -->
    <div class="rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 dark:border-slate-700 dark:bg-slate-800/50">
      <p class="text-sm text-slate-600 dark:text-slate-400">
        <span class="font-medium text-slate-700 dark:text-slate-300">审核概览：</span>
        待审核 <span class="font-semibold text-amber-600 dark:text-amber-400">{{ pendingCount }}</span> 条，
        已通过 <span class="font-semibold text-green-600 dark:text-green-400">{{ publishedCount }}</span> 条，
        已拒绝 <span class="font-semibold text-red-600 dark:text-red-400">{{ rejectedCount }}</span> 条
        <span v-if="pendingCount > 0" class="ml-2 text-amber-600 dark:text-amber-400">
          · 请切换至「待审核」处理
        </span>
      </p>
    </div>

    <div class="flex flex-wrap items-center justify-between gap-4">
      <div class="flex flex-wrap items-center gap-3">
        <div class="relative max-w-md">
          <Search class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索帖子标题、内容或作者..."
            class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 pl-9 pr-4 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          />
        </div>
        <!-- 状态筛选：体现审核流程 -->
        <div class="flex rounded-lg border border-slate-200 bg-white p-0.5 dark:border-slate-700 dark:bg-slate-800">
          <button
            type="button"
            class="rounded-md px-3 py-2 text-sm font-medium transition-colors"
            :class="
              statusFilter === 'all'
                ? 'bg-slate-200 text-slate-900 dark:bg-slate-600 dark:text-slate-100'
                : 'text-slate-600 hover:bg-slate-100 dark:text-slate-400 dark:hover:bg-slate-700'
            "
            @click="setStatusFilter('all')"
          >
            全部
          </button>
          <button
            type="button"
            class="flex items-center gap-1.5 rounded-md px-3 py-2 text-sm font-medium transition-colors"
            :class="
              statusFilter === 'pending'
                ? 'bg-amber-100 text-amber-800 dark:bg-amber-500/30 dark:text-amber-200'
                : 'text-slate-600 hover:bg-slate-100 dark:text-slate-400 dark:hover:bg-slate-700'
            "
            @click="setStatusFilter('pending')"
          >
            待审核
            <span
              v-if="pendingCount > 0"
              class="min-w-[1.25rem] rounded-full bg-amber-500/20 px-1.5 py-0.5 text-xs font-semibold text-amber-700 dark:bg-amber-400/20 dark:text-amber-300"
            >
              {{ pendingCount }}
            </span>
          </button>
          <button
            type="button"
            class="rounded-md px-3 py-2 text-sm font-medium transition-colors"
            :class="
              statusFilter === 'published'
                ? 'bg-green-100 text-green-800 dark:bg-green-500/30 dark:text-green-200'
                : 'text-slate-600 hover:bg-slate-100 dark:text-slate-400 dark:hover:bg-slate-700'
            "
            @click="setStatusFilter('published')"
          >
            已通过
          </button>
          <button
            type="button"
            class="rounded-md px-3 py-2 text-sm font-medium transition-colors"
            :class="
              statusFilter === 'rejected'
                ? 'bg-red-100 text-red-800 dark:bg-red-500/30 dark:text-red-200'
                : 'text-slate-600 hover:bg-slate-100 dark:text-slate-400 dark:hover:bg-slate-700'
            "
            @click="setStatusFilter('rejected')"
          >
            已拒绝
          </button>
        </div>
      </div>
      <button
        type="button"
        class="flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
        @click="openCreate"
      >
        <Plus class="h-4 w-4" />
        新建帖子
      </button>
    </div>

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
                作者
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                状态
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                回复
              </th>
              <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                点赞
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
              v-for="(post, index) in paginatedPosts"
              :key="post.id"
              class="transition-colors hover:bg-slate-50 dark:hover:bg-slate-800/50"
            >
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-900 dark:text-slate-100">
                {{ (currentPage - 1) * pageSize + index + 1 }}
              </td>
              <td class="px-6 py-4 text-sm font-medium text-slate-900 dark:text-slate-100">
                <div class="max-w-xs truncate" :title="post.title">
                  {{ post.title }}
                </div>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ post.author }}
              </td>
              <td class="whitespace-nowrap px-6 py-4">
                <span
                  class="inline-flex rounded-full px-2.5 py-0.5 text-xs font-semibold"
                  :class="{
                    'bg-amber-100 text-amber-700 dark:bg-amber-500/20 dark:text-amber-300': (post.status ?? 'published') === 'pending',
                    'bg-green-100 text-green-700 dark:bg-green-500/20 dark:text-green-300': (post.status ?? 'published') === 'published',
                    'bg-red-100 text-red-700 dark:bg-red-500/20 dark:text-red-300': (post.status ?? 'published') === 'rejected',
                  }"
                >
                  {{ statusLabel(post.status) }}
                </span>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                <span class="inline-flex items-center gap-1">
                  <MessageSquare class="h-4 w-4" />
                  {{ post.replyCount ?? replyCount(post.id) }}
                </span>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                <span class="inline-flex items-center gap-1">
                  <ThumbsUp class="h-4 w-4" />
                  {{ likeCount(post) }}
                </span>
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ post.createdAt }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                {{ post.updatedAt }}
              </td>
              <td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
                <div class="flex flex-wrap items-center justify-end gap-1">
                  <button
                    type="button"
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
                    title="查看"
                    @click="openView(post)"
                  >
                    <Eye class="h-4 w-4" />
                  </button>
                  <template v-if="(post.status ?? 'published') === 'published' || (post.status ?? 'published') === 'rejected'">
                    <button
                      type="button"
                      class="relative rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-indigo-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-indigo-400"
                      title="评论审核"
                      @click="goToCommentReview(post)"
                    >
                      <MessageSquare class="h-4 w-4" />
                      <span
                        v-if="pendingCommentCount(post.id) > 0"
                        class="absolute right-0.5 top-0.5 h-2 w-2 rounded-full bg-red-500"
                        aria-hidden="true"
                      />
                    </button>
                    <button
                      type="button"
                      class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-blue-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-blue-400"
                      title="编辑"
                      @click="goToEdit(post)"
                    >
                      <Edit class="h-4 w-4" />
                    </button>
                  </template>
                  <template v-if="(post.status ?? 'published') === 'pending'">
                    <button
                      type="button"
                      class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-green-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-green-400"
                      title="通过审核"
                      @click="handleApprove(post)"
                    >
                      <CheckCircle class="h-4 w-4" />
                    </button>
                    <button
                      type="button"
                      class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                      title="拒绝"
                      @click="handleReject(post)"
                    >
                      <XCircle class="h-4 w-4" />
                    </button>
                  </template>
                  <button
                    type="button"
                    class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                    title="删除"
                    @click="handleDelete(post)"
                  >
                    <Trash2 class="h-4 w-4" />
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="paginatedPosts.length === 0">
              <td colspan="9" class="px-6 py-12 text-center text-sm text-slate-500">
                暂无帖子数据
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <div v-if="totalPages > 1" class="mt-4 flex items-center justify-between">
      <div class="text-sm text-slate-600 dark:text-slate-400">
        共 {{ filteredPosts.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
      </div>
      <div class="flex items-center gap-2">
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="currentPage === 1"
          @click="currentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>
        <span class="px-3 py-2 text-sm font-medium text-slate-700 dark:text-slate-300">
          {{ currentPage }} / {{ totalPages }}
        </span>
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          :disabled="currentPage === totalPages"
          @click="currentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>

    <!-- 查看帖子（只读）弹窗 -->
    <div
      v-if="showViewDialog && viewPost"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="closeViewDialog"
    >
      <div
        class="w-full max-w-3xl max-h-[90vh] flex flex-col rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900"
      >
        <div class="flex shrink-0 items-center justify-between border-b border-slate-200 px-6 py-4 dark:border-slate-800">
          <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100">
            查看帖子
          </h3>
          <button
            type="button"
            class="rounded-lg p-1 text-slate-400 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:hover:bg-slate-800 dark:hover:text-slate-100"
            @click="closeViewDialog"
          >
            <X class="h-5 w-5" />
          </button>
        </div>
        <div class="flex-1 overflow-y-auto px-6 py-4 space-y-4">
          <div>
            <span class="text-sm font-medium text-slate-500 dark:text-slate-400">标题</span>
            <p class="mt-1 text-slate-900 dark:text-slate-100">
              {{ viewPost.title }}
            </p>
          </div>
          <div class="flex flex-wrap gap-4 text-sm">
            <div>
              <span class="text-slate-500 dark:text-slate-400">作者：</span>
              <span class="text-slate-700 dark:text-slate-300">{{ viewPost.author }}</span>
            </div>
            <div>
              <span class="text-slate-500 dark:text-slate-400">状态：</span>
              <span class="font-medium">{{ statusLabel(viewPost.status) }}</span>
            </div>
            <div>
              <span class="text-slate-500 dark:text-slate-400">创建：</span>
              <span class="text-slate-700 dark:text-slate-300">{{ viewPost.createdAt }}</span>
            </div>
            <div>
              <span class="text-slate-500 dark:text-slate-400">更新：</span>
              <span class="text-slate-700 dark:text-slate-300">{{ viewPost.updatedAt }}</span>
            </div>
            <template v-if="viewPost.reviewedAt || viewPost.reviewedBy">
              <div>
                <span class="text-slate-500 dark:text-slate-400">审核时间：</span>
                <span class="text-slate-700 dark:text-slate-300">{{ viewPost.reviewedAt || '—' }}</span>
              </div>
              <div>
                <span class="text-slate-500 dark:text-slate-400">审核人：</span>
                <span class="text-slate-700 dark:text-slate-300">{{ viewPost.reviewedBy || '—' }}</span>
              </div>
            </template>
            <div v-if="viewPost.rejectReason" class="w-full">
              <span class="text-slate-500 dark:text-slate-400">拒绝原因/备注：</span>
              <p class="mt-1 rounded-lg bg-red-50 px-3 py-2 text-sm text-red-800 dark:bg-red-900/30 dark:text-red-200">
                {{ viewPost.rejectReason }}
              </p>
            </div>
          </div>
          <div>
            <span class="text-sm font-medium text-slate-500 dark:text-slate-400">正文</span>
            <pre class="mt-2 max-h-80 overflow-auto rounded-lg border border-slate-200 bg-slate-50 p-4 text-sm whitespace-pre-wrap font-sans text-slate-800 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200">{{ viewPost.content }}</pre>
          </div>
          <div v-if="(viewPost.status ?? 'published') === 'pending'" class="flex gap-2 pt-2">
            <button
              type="button"
              class="flex items-center gap-2 rounded-lg bg-green-600 px-4 py-2 text-sm font-medium text-white hover:bg-green-700"
              @click="handleApprove(viewPost); closeViewDialog()"
            >
              <CheckCircle class="h-4 w-4" />
              通过审核
            </button>
            <button
              type="button"
              class="flex items-center gap-2 rounded-lg border border-red-300 bg-white px-4 py-2 text-sm font-medium text-red-600 hover:bg-red-50 dark:border-red-700 dark:bg-slate-800 dark:hover:bg-red-900/20"
              @click="openRejectDialog(viewPost); closeViewDialog()"
            >
              <XCircle class="h-4 w-4" />
              拒绝
            </button>
          </div>
        </div>
        <div class="flex shrink-0 justify-end border-t border-slate-200 px-6 py-4 dark:border-slate-800">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
            @click="closeViewDialog"
          >
            关闭
          </button>
        </div>
      </div>
    </div>

    <!-- 拒绝审核弹窗（填写备注） -->
    <div
      v-if="showRejectDialog && rejectPost"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="showRejectDialog = false"
    >
      <div class="w-full max-w-md rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900 p-6">
        <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100 mb-2">
          拒绝帖子
        </h3>
        <p class="text-sm text-slate-600 dark:text-slate-400 mb-4">
          请填写拒绝原因/备注，作者将看到此内容。
        </p>
        <textarea
          v-model="rejectReasonInput"
          rows="4"
          placeholder="请输入拒绝原因或给作者的备注..."
          class="mb-4 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
        />
        <div class="flex justify-end gap-2">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
            @click="showRejectDialog = false; rejectPost = null; rejectReasonInput = ''"
          >
            取消
          </button>
          <button
            type="button"
            class="rounded-lg bg-red-600 px-4 py-2 text-sm font-medium text-white hover:bg-red-700"
            @click="confirmReject"
          >
            确认拒绝
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
