<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArrowLeft, CheckCircle, XCircle, Trash2 } from 'lucide-vue-next'
import {
  getForumPost,
  getCommentsByPostIdForAdmin,
  setCommentStatus,
  deleteForumComment,
} from '../../../stores/forumStore'
import type { ForumComment } from '../../../types/forum'
import type { CommentStatus } from '../../../types/forum'

const route = useRoute()
const router = useRouter()

const postId = computed(() => Number(route.params.postId))
const post = computed(() => getForumPost(postId.value))
const comments = computed(() => getCommentsByPostIdForAdmin(postId.value))

function statusLabel(s: CommentStatus | undefined): string {
  const v = s ?? 'published'
  return { pending: '待审核', published: '已通过', rejected: '已拒绝' }[v]
}

const showRejectDialog = ref(false)
const rejectComment = ref<ForumComment | null>(null)
const rejectReasonInput = ref('')

function openRejectDialog(comment: ForumComment) {
  rejectComment.value = comment
  rejectReasonInput.value = ''
  showRejectDialog.value = true
}

function confirmReject() {
  if (rejectComment.value) {
    setCommentStatus(rejectComment.value.id, 'rejected', {
      reviewedBy: '管理员',
      rejectReason: rejectReasonInput.value.trim(),
    })
    showRejectDialog.value = false
    rejectComment.value = null
    rejectReasonInput.value = ''
  }
}

function handleApprove(comment: ForumComment) {
  setCommentStatus(comment.id, 'published', { reviewedBy: '管理员' })
}

function handleReject(comment: ForumComment) {
  openRejectDialog(comment)
}

function handleDelete(comment: ForumComment) {
  if (!confirm(`确定删除该评论？`)) return
  deleteForumComment(comment.id)
}

function goBack() {
  router.push({ path: '/admin/manage/forum' })
}

const pendingCount = computed(
  () => comments.value.filter((c) => (c.status ?? 'published') === 'pending').length
)
</script>

<template>
  <div class="space-y-4">
    <div class="flex flex-wrap items-center justify-between gap-4">
      <button
        type="button"
        class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
        @click="goBack"
      >
        <ArrowLeft class="h-4 w-4" />
        返回论坛管理
      </button>
    </div>

    <div
      v-if="!post"
      class="rounded-xl border border-slate-200 bg-white py-12 text-center text-slate-500 dark:border-slate-800 dark:bg-slate-900 dark:text-slate-400"
    >
      帖子不存在
      <button
        type="button"
        class="mt-2 text-blue-600 hover:underline dark:text-blue-400"
        @click="goBack"
      >
        返回论坛
      </button>
    </div>

    <template v-else>
      <div class="rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 dark:border-slate-700 dark:bg-slate-800/50">
        <h1 class="text-lg font-bold text-slate-900 dark:text-slate-100">
          {{ post.title }}
        </h1>
        <p class="mt-1 text-sm text-slate-600 dark:text-slate-400">
          评论管理 · 共 {{ comments.length }} 条
          <span v-if="pendingCount > 0" class="ml-2 font-medium text-amber-600 dark:text-amber-400">
            待审核 {{ pendingCount }} 条
          </span>
        </p>
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
                  内容
                </th>
                <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                  作者
                </th>
                <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                  状态
                </th>
                <th class="px-6 py-3 text-left text-xs font-bold uppercase tracking-wider text-slate-500">
                  时间
                </th>
                <th class="px-6 py-3 text-right text-xs font-bold uppercase tracking-wider text-slate-500">
                  操作
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-200 bg-white dark:divide-slate-800 dark:bg-slate-900">
              <tr
                v-for="(comment, index) in comments"
                :key="comment.id"
                class="transition-colors hover:bg-slate-50 dark:hover:bg-slate-800/50"
              >
                <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-900 dark:text-slate-100">
                  {{ index + 1 }}
                </td>
                <td class="max-w-md px-6 py-4 text-sm text-slate-700 dark:text-slate-300">
                  <div class="line-clamp-3 whitespace-pre-wrap" :title="comment.content">
                    {{ comment.content }}
                  </div>
                  <p
                    v-if="(comment.status ?? 'published') === 'rejected' && comment.rejectReason"
                    class="mt-2 rounded bg-red-50 px-2 py-1 text-xs text-red-700 dark:bg-red-900/30 dark:text-red-200"
                  >
                    拒绝原因：{{ comment.rejectReason }}
                  </p>
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                  {{ comment.author }}
                </td>
                <td class="whitespace-nowrap px-6 py-4">
                  <span
                    class="inline-flex rounded-full px-2.5 py-0.5 text-xs font-semibold"
                    :class="{
                      'bg-amber-100 text-amber-700 dark:bg-amber-500/20 dark:text-amber-300': (comment.status ?? 'published') === 'pending',
                      'bg-green-100 text-green-700 dark:bg-green-500/20 dark:text-green-300': (comment.status ?? 'published') === 'published',
                      'bg-red-100 text-red-700 dark:bg-red-500/20 dark:text-red-300': (comment.status ?? 'published') === 'rejected',
                    }"
                  >
                    {{ statusLabel(comment.status) }}
                  </span>
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-sm text-slate-600 dark:text-slate-300">
                  {{ comment.createdAt }}
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-right text-sm font-medium">
                  <div class="flex flex-wrap items-center justify-end gap-1">
                    <template v-if="(comment.status ?? 'published') === 'pending'">
                      <button
                        type="button"
                        class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-green-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-green-400"
                        title="通过"
                        @click="handleApprove(comment)"
                      >
                        <CheckCircle class="h-4 w-4" />
                      </button>
                      <button
                        type="button"
                        class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                        title="拒绝"
                        @click="handleReject(comment)"
                      >
                        <XCircle class="h-4 w-4" />
                      </button>
                    </template>
                    <button
                      type="button"
                      class="rounded-lg p-2 text-slate-600 transition-colors hover:bg-slate-100 hover:text-red-600 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-red-400"
                      title="删除"
                      @click="handleDelete(comment)"
                    >
                      <Trash2 class="h-4 w-4" />
                    </button>
                  </div>
                </td>
              </tr>
              <tr v-if="comments.length === 0">
                <td colspan="6" class="px-6 py-12 text-center text-sm text-slate-500">
                  该帖子暂无评论
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- 拒绝评论弹窗 -->
      <div
        v-if="showRejectDialog && rejectComment"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
        @click.self="showRejectDialog = false"
      >
        <div class="w-full max-w-md rounded-xl border border-slate-200 bg-white shadow-xl dark:border-slate-800 dark:bg-slate-900 p-6">
          <h3 class="text-lg font-bold text-slate-900 dark:text-slate-100 mb-2">
            拒绝评论
          </h3>
          <p class="text-sm text-slate-600 dark:text-slate-400 mb-4">
            可选填写拒绝原因/备注。
          </p>
          <textarea
            v-model="rejectReasonInput"
            rows="3"
            placeholder="请输入拒绝原因..."
            class="mb-4 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          />
          <div class="flex justify-end gap-2">
            <button
              type="button"
              class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
              @click="showRejectDialog = false; rejectComment = null; rejectReasonInput = ''"
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
    </template>
  </div>
</template>
