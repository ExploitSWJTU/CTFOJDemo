<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArrowLeft } from 'lucide-vue-next'
import { getAnnouncement } from '../stores/announcementStore'

const route = useRoute()
const router = useRouter()

const announcementId = computed(() => Number(route.params.id))
const announcement = computed(() => getAnnouncement(announcementId.value))

// 返回
const goBack = () => {
  router.push({ path: '/announcement' })
}
</script>

<template>
  <div class="min-h-[calc(100vh-64px)] p-6">
    <div class="mx-auto max-w-4xl">
      <!-- 返回按钮卡片（左上角，在通知卡片外部） -->
      <div class="mb-4">
        <button
          class="flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-600 shadow-sm transition-all hover:border-blue-300 hover:bg-blue-50 hover:text-blue-600 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-blue-500/60 dark:hover:bg-slate-700 dark:hover:text-blue-400"
          @click="goBack"
        >
          <ArrowLeft class="h-4 w-4" />
          返回公告列表
        </button>
      </div>

      <div
        v-if="announcement"
        class="rounded-xl border border-slate-200 bg-white p-8 shadow-sm dark:border-slate-800 dark:bg-slate-900"
      >
        <!-- 标题 -->
        <div class="mb-4 flex items-start justify-between gap-4">
          <div class="flex-1">
            <div class="mb-2 flex items-center gap-2">
              <h1 class="text-2xl font-bold text-slate-900 dark:text-slate-50">
                {{ announcement.title }}
              </h1>
              <span
                v-if="announcement.isPinned"
                class="inline-flex items-center gap-1 rounded-full bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700 dark:bg-blue-900/30 dark:text-blue-400"
              >
                📌 置顶
              </span>
            </div>
            <div class="flex items-center gap-4 text-sm text-slate-500 dark:text-slate-400">
              <span>作者：{{ announcement.author }}</span>
              <span>发布时间：{{ announcement.createdAt }}</span>
              <span v-if="announcement.updatedAt !== announcement.createdAt">
                更新时间：{{ announcement.updatedAt }}
              </span>
            </div>
          </div>
        </div>

        <!-- 内容 -->
        <div class="prose prose-slate max-w-none dark:prose-invert">
          <div class="whitespace-pre-wrap text-slate-700 dark:text-slate-300">
            {{ announcement.content }}
          </div>
        </div>
      </div>

      <!-- 未找到公告 -->
      <div
        v-else
        class="rounded-xl border border-slate-200 bg-white p-12 text-center dark:border-slate-800 dark:bg-slate-900"
      >
        <p class="text-slate-500 dark:text-slate-400">
          公告不存在
        </p>
      </div>
    </div>
  </div>
</template>

