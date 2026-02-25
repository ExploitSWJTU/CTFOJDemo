<script setup lang="ts">
import { ref, computed, watch, onMounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArrowLeft, MessageSquare, Send, ThumbsUp, User } from 'lucide-vue-next'
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import {
  getForumPost,
  getCommentsByPostId,
  addForumComment,
  hasUserLikedPost,
  togglePostLike,
  hasUserLikedComment,
  toggleCommentLike,
} from '../../../stores/forumStore'
import { useUserStore } from '../../../stores/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

function getOrCreateGuestId(): number {
  if (typeof sessionStorage === 'undefined') return -1
  const stored = sessionStorage.getItem('forum_guest_id')
  if (stored) return Number(stored)
  const id = -Math.abs(Math.floor(Math.random() * 1e9))
  sessionStorage.setItem('forum_guest_id', String(id))
  return id
}

const postId = computed(() => Number(route.params.id))
const rawPost = computed(() => getForumPost(postId.value))
const post = computed(() => {
  const p = rawPost.value
  if (!p) return undefined
  if ((p.status ?? 'published') !== 'published') return undefined
  return p
})
const comments = computed(() => getCommentsByPostId(postId.value))

/** 是否为当前帖子的作者（用于展示审核结果与备注） */
const isAuthor = computed(() => {
  const p = rawPost.value
  if (!p) return false
  const uid = userStore.user?.id
  const name = userStore.user?.username
  return (uid !== undefined && p.authorId === uid) || (name !== undefined && p.author === name)
})

/** 未通过审核的帖子且为作者：展示审核结果与管理员备注 */
const showReviewResult = computed(() => {
  const p = rawPost.value
  if (!p || (p.status ?? 'published') === 'published') return false
  return isAuthor.value
})

const reviewStatusLabel = computed(() => {
  const p = rawPost.value
  if (!p) return ''
  const s = p.status ?? 'published'
  return s === 'pending' ? '待审核中' : '未通过审核'
})

const commentContent = ref('')
const submitting = ref(false)
const contentPreviewRef = ref<HTMLDivElement | null>(null)

const authorName = computed(() => userStore.user?.username ?? '游客')
const effectiveUserId = computed(() => userStore.user?.id ?? getOrCreateGuestId())
const liked = computed(() =>
  post.value ? hasUserLikedPost(post.value.id, effectiveUserId.value) : false
)
const likeCount = computed(() => (post.value?.likedByUserIds?.length ?? 0))

function commentLiked(commentId: number) {
  return hasUserLikedComment(commentId, effectiveUserId.value)
}

function commentLikeCount(comment: { likedByUserIds?: number[] }) {
  return comment.likedByUserIds?.length ?? 0
}

function handleCommentLike(commentId: number) {
  toggleCommentLike(commentId, effectiveUserId.value)
}

function handleLike() {
  if (!post.value) return
  togglePostLike(post.value.id, effectiveUserId.value)
}

function goBack() {
  router.push({ path: '/forum' })
}

function submitComment() {
  const content = commentContent.value.trim()
  if (!content || !post.value) return
  submitting.value = true
  addForumComment({
    postId: post.value.id,
    content,
    author: authorName.value,
    authorId: userStore.user?.id,
  })
  commentContent.value = ''
  submitting.value = false
}

const isDarkMode = computed(() => document.documentElement.classList.contains('dark'))

function renderMarkdown() {
  if (!contentPreviewRef.value || !rawPost.value) return
  const markdown = rawPost.value.content || ''
  contentPreviewRef.value.innerHTML = ''
  if (!markdown) {
    contentPreviewRef.value.innerHTML = '<p class="text-slate-400 dark:text-slate-500">暂无内容</p>'
    return
  }
  try {
    const previewOptions = {
      mode: 'light' as const,
      anchor: 1,
      speech: { enable: false },
      theme: isDarkMode.value ? 'dark' : 'classic',
    }
    Vditor.preview(contentPreviewRef.value, markdown, previewOptions as unknown as Parameters<typeof Vditor.preview>[2])
  } catch {
    import('markdown-it').then(({ default: MarkdownIt }) => {
      const md = new MarkdownIt({ html: true, linkify: true, breaks: true })
      contentPreviewRef.value!.innerHTML = md.render(markdown)
    })
  }
}

watch([rawPost, isDarkMode, showReviewResult], () => renderMarkdown(), { immediate: false })
onMounted(() => nextTick().then(() => renderMarkdown()))

watch(postId, () => {
  commentContent.value = ''
})
</script>

<template>
  <div class="min-h-[calc(100vh-64px)]">
    <div class="mx-auto max-w-3xl px-6 py-8">
      <!-- 返回：紧贴正文区域 -->
      <button
        type="button"
        class="mb-6 inline-flex items-center gap-2 rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-medium text-slate-600 shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-slate-600"
        @click="goBack"
      >
        <ArrowLeft class="h-4 w-4" />
        返回论坛
      </button>

      <!-- 作者查看未通过/待审核时的审核结果与备注 -->
      <template v-if="showReviewResult && rawPost">
        <div class="mb-6 rounded-2xl border-2 border-amber-200 bg-amber-50/80 p-6 dark:border-amber-800 dark:bg-amber-900/20">
          <h2 class="mb-2 text-lg font-semibold text-amber-800 dark:text-amber-200">
            审核结果
          </h2>
          <p class="mb-2 text-amber-700 dark:text-amber-300">
            状态：{{ reviewStatusLabel }}
          </p>
          <p v-if="rawPost.reviewedAt" class="mb-2 text-sm text-slate-600 dark:text-slate-400">
            审核时间：{{ rawPost.reviewedAt }}
            <span v-if="rawPost.reviewedBy">，审核人：{{ rawPost.reviewedBy }}</span>
          </p>
          <p v-if="(rawPost.status ?? '') === 'rejected' && rawPost.rejectReason" class="mt-3 rounded-lg bg-white/80 p-3 text-slate-700 dark:bg-slate-800/80 dark:text-slate-300">
            <span class="font-medium">管理员备注：</span>{{ rawPost.rejectReason }}
          </p>
        </div>
        <div class="rounded-2xl border border-slate-200 bg-white p-6 dark:border-slate-800 dark:bg-slate-900">
          <h1 class="mb-4 text-2xl font-bold text-slate-900 dark:text-slate-50">
            {{ rawPost.title }}
          </h1>
          <p class="mb-4 text-sm text-slate-500 dark:text-slate-400">
            作者：{{ rawPost.author }} · {{ rawPost.createdAt }}
          </p>
          <div class="prose prose-slate max-w-none dark:prose-invert">
            <div ref="contentPreviewRef" class="vditor-preview rounded-xl bg-slate-50/80 px-4 py-4 dark:bg-slate-800/50" />
          </div>
        </div>
      </template>

      <template v-else-if="post">
        <header class="mb-10">
          <h1 class="text-3xl font-bold leading-tight tracking-tight text-slate-900 dark:text-slate-50 sm:text-4xl">
            {{ post.title }}
          </h1>
          <div class="mt-6 flex flex-wrap items-center gap-x-4 gap-y-2 border-b border-slate-200 pb-6 dark:border-slate-700">
            <span class="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-300">
              <span class="flex h-8 w-8 items-center justify-center rounded-full bg-slate-200 text-slate-500 dark:bg-slate-700 dark:text-slate-400">
                <User class="h-4 w-4" />
              </span>
              {{ post.author }}
            </span>
            <span class="text-sm text-slate-500 dark:text-slate-400">{{ post.createdAt }}</span>
            <span v-if="post.updatedAt !== post.createdAt" class="text-sm text-slate-500 dark:text-slate-400">
              更新于 {{ post.updatedAt }}
            </span>
            <span class="inline-flex items-center gap-1.5 text-sm text-slate-500 dark:text-slate-400">
              <MessageSquare class="h-4 w-4" />
              {{ post.replyCount ?? 0 }} 条回复
            </span>
            <span class="inline-flex items-center gap-1.5">
              <button
                type="button"
                :title="liked ? '取消点赞' : '点赞'"
                class="inline-flex items-center gap-1.5 rounded-lg px-2.5 py-1.5 text-sm transition"
                :class="
                  liked
                    ? 'bg-blue-100 text-blue-600 dark:bg-blue-900/40 dark:text-blue-400'
                    : 'text-slate-500 hover:bg-slate-100 hover:text-slate-700 dark:text-slate-400 dark:hover:bg-slate-700 dark:hover:text-slate-300'
                "
                @click.stop="handleLike"
              >
                <ThumbsUp class="h-4 w-4" :fill="liked ? 'currentColor' : 'none'" />
                {{ likeCount }}
              </button>
            </span>
          </div>
        </header>

        <!-- 正文 Markdown 渲染 -->
        <div class="prose prose-slate max-w-none dark:prose-invert">
          <div
            ref="contentPreviewRef"
            class="vditor-preview rounded-2xl bg-slate-50/80 px-6 py-6 dark:bg-slate-800/50"
          />
        </div>

        <section class="mt-14 border-t border-slate-200 pt-10 dark:border-slate-700">
          <h2 class="mb-6 text-xl font-semibold text-slate-900 dark:text-slate-50">
            评论
            <span class="ml-2 font-normal text-slate-500 dark:text-slate-400">({{ comments.length }})</span>
          </h2>

          <ul class="space-y-6">
            <li
              v-for="comment in comments"
              :key="comment.id"
              class="flex gap-4 rounded-2xl border border-slate-200/80 bg-white p-5 dark:border-slate-800 dark:bg-slate-900/50"
            >
              <span class="flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-slate-200 text-slate-500 dark:bg-slate-700 dark:text-slate-400">
                <User class="h-5 w-5" />
              </span>
              <div class="min-w-0 flex-1">
                <div class="mb-2 flex flex-wrap items-center gap-2">
                  <span class="font-medium text-slate-800 dark:text-slate-200">{{ comment.author }}</span>
                  <span class="text-sm text-slate-500 dark:text-slate-400">{{ comment.createdAt }}</span>
                  <button
                    type="button"
                    :title="commentLiked(comment.id) ? '取消点赞' : '点赞'"
                    class="ml-auto inline-flex items-center gap-1 rounded-lg px-2 py-1 text-xs transition"
                    :class="
                      commentLiked(comment.id)
                        ? 'bg-blue-100 text-blue-600 dark:bg-blue-900/40 dark:text-blue-400'
                        : 'text-slate-500 hover:bg-slate-100 dark:text-slate-400 dark:hover:bg-slate-700'
                    "
                    @click.stop="handleCommentLike(comment.id)"
                  >
                    <ThumbsUp class="h-3.5 w-3.5" :fill="commentLiked(comment.id) ? 'currentColor' : 'none'" />
                    {{ commentLikeCount(comment) }}
                  </button>
                </div>
                <p class="whitespace-pre-wrap text-slate-700 dark:text-slate-300">
                  {{ comment.content }}
                </p>
              </div>
            </li>
            <li
              v-if="comments.length === 0"
              class="rounded-2xl border border-dashed border-slate-200 bg-slate-50/50 py-12 text-center text-slate-500 dark:border-slate-700 dark:bg-slate-800/30 dark:text-slate-400"
            >
              暂无评论，快来抢沙发～
            </li>
          </ul>

          <div class="mt-8 rounded-2xl border border-slate-200 bg-white p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900/50">
            <label class="mb-3 block text-sm font-medium text-slate-700 dark:text-slate-300">
              发表评论
            </label>
            <textarea
              v-model="commentContent"
              rows="4"
              placeholder="写下你的评论…"
              class="mb-4 w-full resize-y rounded-xl border border-slate-200 bg-slate-50/50 px-4 py-3 text-slate-800 placeholder-slate-400 transition focus:border-blue-500 focus:bg-white focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-200 dark:placeholder-slate-500 dark:focus:bg-slate-800 dark:focus:ring-blue-500/20"
            />
            <button
              type="button"
              :disabled="!commentContent.trim() || submitting"
              class="inline-flex items-center gap-2 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-offset-slate-900"
              @click="submitComment"
            >
              <Send class="h-4 w-4" />
              发送
            </button>
          </div>
        </section>
      </template>

      <div
        v-else
        class="rounded-2xl border border-slate-200 bg-white py-16 text-center dark:border-slate-800 dark:bg-slate-900"
      >
        <p class="text-slate-500 dark:text-slate-400">
          帖子不存在或您无权限查看
        </p>
        <button
          type="button"
          class="mt-4 text-blue-600 hover:underline dark:text-blue-400"
          @click="goBack"
        >
          返回论坛
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.vditor-preview) {
  padding: 1.5rem;
}
.dark :deep(.vditor-preview) {
  background: rgba(30, 41, 59, 0.5);
}
</style>
