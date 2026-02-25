<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft } from 'lucide-vue-next'
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { createForumPost } from '../../../stores/forumStore'
import { useUserStore } from '../../../stores/user'

const router = useRouter()
const userStore = useUserStore()

const title = ref('')
const submitting = ref(false)
const vditorRef = ref<InstanceType<typeof Vditor> | null>(null)
const editorContainer = ref<HTMLDivElement | null>(null)

const authorName = computed(() => userStore.user?.username ?? '游客')
const isDarkMode = computed(() => document.documentElement.classList.contains('dark'))

onMounted(() => {
  if (editorContainer.value) {
    vditorRef.value = new Vditor(editorContainer.value, {
      height: Math.min(400, window.innerHeight - 320),
      placeholder: '支持 Markdown，输入帖子内容…',
      toolbarConfig: { pin: true },
      cache: { enable: false },
      mode: 'sv',
      preview: { mode: 'both', actions: ['desktop', 'tablet', 'mobile'] },
      value: '',
      theme: isDarkMode.value ? 'dark' : 'classic',
      upload: {
        accept: 'image/*',
        handler: () => Promise.resolve(''),
      },
    })
  }
})

onBeforeUnmount(() => {
  if (vditorRef.value) {
    vditorRef.value.destroy()
    vditorRef.value = null
  }
})

function goBack() {
  router.push({ path: '/forum' })
}

function submit() {
  const t = title.value.trim()
  if (!t) return
  const content = vditorRef.value?.getValue()?.trim() ?? ''
  submitting.value = true
  const post = createForumPost({
    title: t,
    content,
    author: authorName.value,
    authorId: userStore.user?.id,
    status: 'pending',
  })
  submitting.value = false
  if (post) {
    alert('帖子已提交，等待审核通过后将在论坛展示。')
    router.push({ path: '/forum' })
  } else {
    router.push({ path: '/forum' })
  }
}
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

      <div class="rounded-2xl border border-slate-200 bg-white shadow-sm dark:border-slate-800 dark:bg-slate-900">
        <div class="border-b border-slate-200 px-8 py-6 dark:border-slate-700">
          <h1 class="text-2xl font-bold text-slate-900 dark:text-slate-50">
            发布帖子
          </h1>
          <p class="mt-1 text-sm text-slate-500 dark:text-slate-400">
            分享你的思路、经验或提问，支持 Markdown 编辑与预览
          </p>
        </div>

        <form class="p-8" @submit.prevent="submit">
          <div class="mb-6">
            <label for="post-title" class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              标题
            </label>
            <input
              id="post-title"
              v-model="title"
              type="text"
              placeholder="请输入帖子标题"
              class="w-full rounded-xl border border-slate-200 bg-slate-50/50 px-4 py-3 text-slate-800 placeholder-slate-400 transition focus:border-blue-500 focus:bg-white focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-200 dark:placeholder-slate-500 dark:focus:bg-slate-800 dark:focus:ring-blue-500/20"
            />
          </div>

          <div class="mb-8">
            <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
              内容（Markdown）
            </label>
            <div ref="editorContainer" class="forum-editor" />
          </div>

          <div class="flex flex-wrap gap-3">
            <button
              type="submit"
              :disabled="!title.trim() || submitting"
              class="inline-flex items-center rounded-xl bg-blue-600 px-5 py-2.5 text-sm font-medium text-white shadow-sm transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-offset-slate-900"
            >
              发布
            </button>
            <button
              type="button"
              class="rounded-xl border border-slate-200 bg-white px-5 py-2.5 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
              @click="goBack"
            >
              取消
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.forum-editor .vditor) {
  border-radius: 0.75rem;
  overflow: hidden;
}
:deep(.forum-editor .vditor-content) {
  border-color: rgb(226 232 240);
}
.dark :deep(.forum-editor .vditor-content) {
  border-color: rgb(51 65 85);
}
</style>
