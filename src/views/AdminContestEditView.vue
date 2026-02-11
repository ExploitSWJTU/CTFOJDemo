<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { Save, ArrowLeft } from 'lucide-vue-next'
import { contestStore, updateContest, type ContestType } from '../stores/contestStore'

const route = useRoute()
const router = useRouter()

const contestId = computed(() => Number(route.params.id))
const contest = computed(() => contestStore.contests.find((c) => c.id === contestId.value))

// 编辑表单数据
const formData = ref({
  name: '',
  brief: '',
  startTime: '',
  endTime: '',
  type: 'individual' as ContestType,
  imageUrl: '',
})

// 图片上传相关
const imageFile = ref<File | null>(null)
const imagePreview = ref<string>('')

// Vditor 实例
const vditorRef = ref<InstanceType<typeof Vditor> | null>(null)
const editorContainer = ref<HTMLDivElement | null>(null)

// 检测暗色模式
const isDarkMode = computed(() => {
  return document.documentElement.classList.contains('dark')
})

// 处理图片选择
const handleImageSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0]
  if (file) {
    imageFile.value = file
    // 创建预览
    const reader = new FileReader()
    reader.onload = (e) => {
      imagePreview.value = e.target?.result as string
      formData.value.imageUrl = imagePreview.value // 使用 base64 作为临时 URL
    }
    reader.readAsDataURL(file)
  }
}

// 初始化表单数据和 Vditor
onMounted(() => {
  if (contest.value) {
    // 初始化表单数据
    formData.value = {
      name: contest.value.name || '',
      brief: contest.value.brief || '',
      startTime: contest.value.startTime || '',
      endTime: contest.value.endTime || '',
      type: contest.value.type || 'individual',
      imageUrl: contest.value.imageUrl || '',
    }
    // 初始化图片预览
    imagePreview.value = contest.value.imageUrl || ''

    // 初始化 Vditor
    if (editorContainer.value) {
      vditorRef.value = new Vditor(editorContainer.value, {
        height: window.innerHeight - 400,
        placeholder: '输入 Markdown 格式的赛事介绍...',
        toolbarConfig: {
          pin: true,
        },
        cache: {
          enable: false,
        },
        mode: 'sv', // 分屏模式：左侧编辑，右侧预览
        preview: {
          mode: 'both', // 显示编辑和预览
          actions: ['desktop', 'tablet', 'mobile'],
        },
        value: contest.value.description || '',
        theme: isDarkMode.value ? 'dark' : 'classic',
        upload: {
          accept: 'image/*',
          handler: () => {
            // 可以在这里实现图片上传功能
            return Promise.resolve('')
          },
        },
      })
    }
  }
})

// 清理 Vditor 实例
onBeforeUnmount(() => {
  if (vditorRef.value) {
    vditorRef.value.destroy()
    vditorRef.value = null
  }
})

// 保存
const saveDescription = () => {
  if (!contest.value || !vditorRef.value) return

  if (confirm('确定要保存修改吗？')) {
    // 获取 Vditor 的内容
    const content = vditorRef.value.getValue()
    // 更新共享数据
    const success = updateContest(contestId.value, {
      name: formData.value.name,
      brief: formData.value.brief,
      description: content,
      startTime: formData.value.startTime,
      endTime: formData.value.endTime,
      type: formData.value.type,
      imageUrl: formData.value.imageUrl,
    })
    if (success) {
      alert('保存成功！')
      // 保存后停留在当前编辑页面
    } else {
      alert('保存失败，未找到对应的比赛')
    }
  }
}

// 返回
const goBack = () => {
  router.push({ name: 'adminContest' })
}
</script>

<template>
  <div class="min-h-[calc(100vh-64px)] space-y-4">
    <div v-if="contest" class="space-y-4">
      <!-- 头部操作栏 -->
      <div class="flex items-center justify-between">
        <div class="space-y-2">
          <h1 class="text-2xl font-bold text-slate-900 dark:text-slate-50">
            编辑赛事介绍
          </h1>
          <div class="inline-flex items-center gap-2 rounded-lg bg-slate-100 px-3 py-1.5 dark:bg-slate-800">
            <span class="text-xs font-medium text-slate-500 dark:text-slate-400">赛事名称：</span>
            <span class="text-sm font-semibold text-slate-700 dark:text-slate-200">
              {{ contest.name }}
            </span>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <button
            type="button"
            class="inline-flex items-center gap-2 rounded-lg border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="goBack"
          >
            <ArrowLeft class="h-4 w-4" />
            返回
          </button>
          <button
            type="button"
            class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-950"
            @click="saveDescription"
          >
            <Save class="h-4 w-4" />
            保存
          </button>
        </div>
      </div>

      <!-- 编辑区域 -->
      <div class="space-y-4">
        <!-- 基本信息编辑 -->
        <div class="grid grid-cols-1 gap-4 rounded-xl border border-slate-200/80 bg-white/80 p-4 shadow-sm dark:border-slate-800 dark:bg-slate-900/70">
          <div class="space-y-4">
            <!-- 比赛标题 -->
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                比赛标题
              </label>
              <input
                v-model="formData.name"
                type="text"
                class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400"
                placeholder="请输入比赛标题"
              />
            </div>

            <!-- 简介 -->
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                简介
              </label>
              <textarea
                v-model="formData.brief"
                rows="3"
                class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400"
                placeholder="请输入比赛简介"
              />
            </div>

            <!-- 开始时间和结束时间 -->
            <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
              <div>
                <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                  开始时间
                </label>
                <input
                  v-model="formData.startTime"
                  type="text"
                  class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400"
                  placeholder="例如：2025-03-10 19:00"
                />
              </div>
              <div>
                <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                  结束时间
                </label>
                <input
                  v-model="formData.endTime"
                  type="text"
                  class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 placeholder:text-slate-400 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400"
                  placeholder="例如：2025-03-10 22:00"
                />
              </div>
            </div>

            <!-- 赛制 -->
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                赛制
              </label>
              <select
                v-model="formData.type"
                class="w-full rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm text-slate-900 focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:focus:border-blue-400"
              >
                <option value="individual">
                  个人赛
                </option>
                <option value="team">
                  团体赛
                </option>
              </select>
            </div>

            <!-- 比赛海报 -->
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
                比赛海报
              </label>
              <div class="space-y-3">
                <input
                  type="file"
                  accept="image/*"
                  class="block w-full text-sm text-slate-500 file:mr-4 file:rounded-lg file:border-0 file:bg-blue-50 file:px-4 file:py-2 file:text-sm file:font-semibold file:text-blue-700 hover:file:bg-blue-100 dark:file:bg-blue-900/20 dark:file:text-blue-300"
                  @change="handleImageSelect"
                />
                <div
                  v-if="imagePreview"
                  class="relative overflow-hidden rounded-lg border border-slate-300 dark:border-slate-700"
                >
                  <img
                    :src="imagePreview"
                    alt="比赛海报预览"
                    class="h-48 w-full object-cover"
                  />
                </div>
                <div
                  v-else
                  class="flex h-48 items-center justify-center rounded-lg border-2 border-dashed border-slate-300 bg-slate-50 text-sm text-slate-400 dark:border-slate-700 dark:bg-slate-800"
                >
                  预览将在此显示
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Markdown 编辑器 -->
        <div>
          <div
            ref="editorContainer"
            class="rounded-lg border border-slate-300 dark:border-slate-700 overflow-hidden"
          />
        </div>
      </div>
    </div>

    <!-- 未找到比赛 -->
    <div
      v-else
      class="flex min-h-[200px] items-center justify-center rounded-xl border border-dashed border-slate-300 bg-white/60 text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-900/70 dark:text-slate-400"
    >
      未找到对应的赛事信息。
    </div>
  </div>
</template>

<style scoped>
/* Vditor 样式已通过 import 引入 */
/* 如果需要自定义样式，可以在这里添加 */
:deep(.vditor) {
  border: none;
}

:deep(.vditor-content) {
  background-color: white;
}

.dark :deep(.vditor-content) {
  background-color: #1e293b;
  color: #e2e8f0;
}
</style>

