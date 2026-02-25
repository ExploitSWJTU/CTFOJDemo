<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { Save, ArrowLeft } from 'lucide-vue-next'
import { createForumPost } from '../../../stores/forumStore'

const router = useRouter()

const title = ref('')
const vditorRef = ref<InstanceType<typeof Vditor> | null>(null)
const editorContainer = ref<HTMLDivElement | null>(null)

const isDarkMode = computed(() => document.documentElement.classList.contains('dark'))

onMounted(() => {
  if (editorContainer.value) {
    vditorRef.value = new Vditor(editorContainer.value, {
      height: Math.min(450, window.innerHeight - 320),
      placeholder: '输入 Markdown 格式的帖子内容…',
      toolbarConfig: { pin: true },
      cache: { enable: false },
      mode: 'sv',
      preview: { mode: 'both', actions: ['desktop', 'tablet', 'mobile'] },
      value: '',
      theme: isDarkMode.value ? 'dark' : 'classic',
      upload: { accept: 'image/*', handler: () => Promise.resolve('') },
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
  router.push('/admin/manage/forum')
}

function save() {
  const t = title.value.trim()
  if (!t) {
    alert('请填写标题')
    return
  }
  const content = vditorRef.value?.getValue() ?? ''
  const newPost = createForumPost({
    title: t,
    content,
    author: 'admin',
    status: 'published',
  })
  if (newPost) {
    router.push('/admin/manage/forum')
  }
}
</script>

<template>
  <div class="space-y-4">
    <div class="flex items-center justify-between gap-4">
      <button
        type="button"
        class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
        @click="goBack"
      >
        <ArrowLeft class="h-4 w-4" />
        返回论坛管理
      </button>
    </div>

    <div class="rounded-xl border border-slate-200 bg-white p-6 dark:border-slate-800 dark:bg-slate-900">
      <div class="mb-4">
        <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
          标题 <span class="text-red-500">*</span>
        </label>
        <input
          v-model="title"
          type="text"
          placeholder="帖子标题"
          class="h-10 w-full rounded-lg border border-slate-200 bg-slate-50 px-3 text-sm outline-none focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
        />
      </div>
      <div>
        <label class="mb-2 block text-sm font-medium text-slate-700 dark:text-slate-300">
          内容（Markdown）
        </label>
        <div ref="editorContainer" class="forum-edit-editor" />
      </div>
      <div class="mt-6 flex gap-3">
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="save"
        >
          <Save class="h-4 w-4" />
          发布
        </button>
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
          @click="goBack"
        >
          取消
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
:deep(.forum-edit-editor .vditor) {
  border-radius: 0.5rem;
  overflow: hidden;
}
</style>
