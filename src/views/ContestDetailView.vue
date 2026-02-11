<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArrowLeft } from 'lucide-vue-next'
import Vditor from 'vditor'
import 'vditor/dist/index.css'
import { contestStore, type ContestStatus } from '../stores/contestStore'

const statusLabel: Record<ContestStatus, string> = {
  ongoing: '进行中',
  upcoming: '未开始',
  finished: '已结束',
}

const route = useRoute()
const router = useRouter()

const contestId = computed(() => Number(route.params.id))
const contest = computed(() => contestStore.contests.find((c) => c.id === contestId.value))

// 判断是否是管理员页面
const isAdminPage = computed(() => route.name === 'adminContestDetail')

// 判断是否可以报名（待开始或进行中）
const canRegister = computed(() => {
  return contest.value && (contest.value.status === 'upcoming' || contest.value.status === 'ongoing')
})

// 报名状态
const isRegistered = ref(false)

// 按钮文字
const buttonText = computed(() => {
  if (!isRegistered.value) {
    return '报名'
  }
  if (contest.value?.status === 'ongoing') {
    return '进入比赛'
  }
  if (contest.value?.status === 'upcoming') {
    return '等待比赛开始'
  }
  return '报名'
})

// Markdown 预览容器
const previewContainer = ref<HTMLDivElement | null>(null)

// 检测暗色模式
const isDarkMode = computed(() => {
  return document.documentElement.classList.contains('dark')
})

// 渲染 Markdown
const renderMarkdown = async () => {
  await nextTick()
  if (!previewContainer.value || !contest.value) return

  const markdown = contest.value.description || ''
  if (!markdown) {
    previewContainer.value.innerHTML = '<p class="text-slate-400 dark:text-slate-500">暂无赛事介绍</p>'
    return
  }

  // 清空容器
  previewContainer.value.innerHTML = ''

  // 使用 Vditor 的预览功能渲染 Markdown
  // Vditor.preview(element, markdown, options)
  try {
    Vditor.preview(previewContainer.value, markdown, {
      mode: 'light',
      theme: isDarkMode.value ? 'dark' : 'classic',
      anchor: 1, // 显示标题锚点
      speech: {
        enable: false,
      },
    })
  } catch (error) {
    console.error('Vditor preview error:', error)
    // 如果 Vditor.preview 失败，使用 markdown-it 作为后备
    const MarkdownIt = (await import('markdown-it')).default
    const md = new MarkdownIt({
      html: true,
      linkify: true,
      typographer: true,
      breaks: true,
    })
    previewContainer.value.innerHTML = md.render(markdown)
  }
}

// 监听比赛变化和暗色模式变化
watch([contest, isDarkMode], () => {
  renderMarkdown()
}, { immediate: false })

onMounted(() => {
  renderMarkdown()
})

// 返回功能
const goBack = () => {
  if (isAdminPage.value) {
    router.push({ name: 'adminContest' })
  } else {
    router.push({ name: 'contest' })
  }
}

// 报名功能
const handleRegister = () => {
  if (!isRegistered.value) {
    // 报名
    // TODO: 调用 API 报名
    isRegistered.value = true
    alert('报名成功！')
  } else {
    // 已报名后的操作
    if (contest.value?.status === 'ongoing') {
      // TODO: 进入比赛
      alert('进入比赛功能待实现')
    } else if (contest.value?.status === 'upcoming') {
      // 等待比赛开始，可能不需要操作
      alert('比赛尚未开始，请等待')
    }
  }
}

</script>

<template>
  <div class="min-h-[calc(100vh-64px)]">
    <div v-if="contest" class="flex gap-6">
      <!-- 左侧：赛事介绍（Markdown） -->
      <div class="flex-[2] space-y-4 rounded-xl border border-slate-200/80 bg-white/80 p-6 shadow-sm dark:border-slate-800 dark:bg-slate-900/70">
        <!-- 返回按钮 -->
        <div class="mb-4">
          <button
            type="button"
            class="inline-flex items-center gap-2 rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="goBack"
          >
            <ArrowLeft class="h-4 w-4" />
            返回
          </button>
        </div>
        <h2 class="mb-6 text-center text-3xl font-bold text-slate-900 dark:text-slate-50">
          赛事介绍
        </h2>
        <div
          ref="previewContainer"
          class="vditor-preview"
        />
      </div>

      <!-- 右侧：比赛信息卡片 -->
      <aside class="w-96 shrink-0">
        <div
          class="space-y-4 rounded-xl border border-slate-200/80 bg-white/80 p-4 shadow-sm dark:border-slate-800 dark:bg-slate-900/70"
        >
          <div class="overflow-hidden rounded-lg bg-slate-100 dark:bg-slate-800">
            <img
              :src="contest.imageUrl"
              :alt="contest.name"
              class="h-40 w-full object-cover"
            />
          </div>

          <div class="space-y-3 text-sm">
            <div class="flex items-center justify-between gap-2">
              <h3 class="truncate text-base font-semibold text-slate-900 dark:text-slate-50">
                {{ contest.name }}
              </h3>
              <span
                class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium"
                :class="
                  contest.type === 'individual'
                    ? 'bg-purple-100 text-purple-700 dark:bg-purple-500/15 dark:text-purple-300'
                    : 'bg-indigo-100 text-indigo-700 dark:bg-indigo-500/15 dark:text-indigo-300'
                "
              >
                {{ contest.type === 'individual' ? '个人赛' : '团体赛' }}
              </span>
            </div>

            <div class="flex items-center justify-between text-xs">
              <span class="text-slate-500 dark:text-slate-400">赛事状态</span>
              <span
                class="inline-flex items-center gap-1 rounded-full px-2 py-0.5 font-medium"
                :class="{
                  'bg-red-100 text-red-700 dark:bg-red-500/15 dark:text-red-300':
                    contest.status === 'ongoing',
                  'bg-emerald-100 text-emerald-700 dark:bg-emerald-500/15 dark:text-emerald-300':
                    contest.status === 'upcoming',
                  'bg-slate-100 text-slate-700 dark:bg-slate-500/15 dark:text-slate-300':
                    contest.status === 'finished',
                }"
              >
                <span
                  class="h-2 w-2 rounded-full"
                  :class="{
                    'bg-red-500': contest.status === 'ongoing',
                    'bg-emerald-500': contest.status === 'upcoming',
                    'bg-slate-400': contest.status === 'finished',
                  }"
                />
                {{ statusLabel[contest.status] }}
              </span>
            </div>

            <div class="space-y-1 text-xs text-slate-600 dark:text-slate-300">
              <div class="flex items-center justify-between">
                <span>开始时间</span>
                <span>{{ contest.startTime }}</span>
              </div>
              <div class="flex items-center justify-between">
                <span>结束时间</span>
                <span>{{ contest.endTime }}</span>
              </div>
              <div class="flex items-center justify-between">
                <span>报名人数</span>
                <span>{{ contest.participantCount }} 人</span>
              </div>
            </div>

            <!-- 报名按钮（仅用户页面，比赛状态为待开始或进行中时显示） -->
            <div v-if="!isAdminPage && canRegister" class="pt-2">
              <button
                type="button"
                class="w-full rounded-lg bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-950"
                @click="handleRegister"
              >
                {{ buttonText }}
              </button>
            </div>
          </div>
        </div>
      </aside>
    </div>

    <div
      v-else
      class="flex min-h-[200px] items-center justify-center rounded-xl border border-dashed border-slate-300 bg-white/60 text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-900/70 dark:text-slate-400"
    >
      未找到对应的赛事信息。
    </div>
  </div>
</template>

<style scoped>
/* 使用 Vditor 的预览样式 */
:deep(.vditor-preview) {
  color: rgb(30 41 59); /* text-slate-800 */
}

.dark :deep(.vditor-preview) {
  color: rgb(226 232 240); /* text-slate-200 */
}

/* 确保 Vditor 预览样式正确应用 */
:deep(.vditor-preview__action) {
  display: none; /* 隐藏预览工具栏 */
}
</style>

