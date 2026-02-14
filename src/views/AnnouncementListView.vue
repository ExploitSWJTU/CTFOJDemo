<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ChevronLeft, ChevronRight } from 'lucide-vue-next'
import { getPublishedAnnouncements } from '../stores/announcementStore'

const route = useRoute()
const router = useRouter()

// 获取所有已发布的公告
const allAnnouncements = computed(() => getPublishedAnnouncements())

// 分页
const pageSize = ref(10)
const currentPage = ref(Number(route.query.page) || 1)

// 分页后的公告列表
const paginatedAnnouncements = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return allAnnouncements.value.slice(start, end)
})

// 总页数
const totalPages = computed(() => {
  return Math.ceil(allAnnouncements.value.length / pageSize.value)
})

// 监听页码变化，更新 URL
watch(currentPage, (newPage) => {
  router.push({
    query: {
      ...route.query,
      page: newPage,
    },
  })
})

// 监听总页数变化，确保当前页码在有效范围内
watch(totalPages, (newTotalPages) => {
  if (currentPage.value > newTotalPages && newTotalPages > 0) {
    currentPage.value = newTotalPages
  }
})

// 组件挂载时从 URL 读取页码
onMounted(() => {
  const pageFromUrl = Number(route.query.page)
  if (pageFromUrl && pageFromUrl >= 1) {
    if (totalPages.value > 0) {
      currentPage.value = Math.min(pageFromUrl, totalPages.value)
    } else {
      currentPage.value = 1
    }
  } else {
    currentPage.value = 1
  }
})

// 监听路由变化，同步页码
watch(
  () => route.query.page,
  (newPage) => {
    const pageNum = Number(newPage)
    if (pageNum && pageNum >= 1 && pageNum !== currentPage.value) {
      if (totalPages.value > 0) {
        currentPage.value = Math.min(pageNum, totalPages.value)
      } else {
        currentPage.value = 1
      }
    } else if (!newPage && currentPage.value !== 1) {
      currentPage.value = 1
    }
  }
)

// 跳转到公告详情
const goToAnnouncementDetail = (id: number) => {
  router.push({ path: `/announcement/${id}` })
}
</script>

<template>
  <div class="min-h-[calc(100vh-64px)] p-6">
    <div class="mx-auto max-w-4xl">
      <h1 class="mb-6 text-3xl font-bold text-slate-900 dark:text-slate-50">
        全站公告
      </h1>

      <!-- 公告列表 -->
      <div class="space-y-4">
        <div
          v-for="announcement in paginatedAnnouncements"
          :key="announcement.id"
          class="group cursor-pointer rounded-xl border border-slate-200 bg-white p-4 shadow-sm transition-all hover:border-blue-300 hover:shadow-md dark:border-slate-800 dark:bg-slate-900 dark:hover:border-blue-500/60"
          @click="goToAnnouncementDetail(announcement.id)"
        >
          <div class="flex items-start justify-between gap-4">
            <div class="flex-1">
              <div class="mb-2 flex items-center gap-2">
                <h3 class="text-lg font-semibold text-slate-900 dark:text-slate-50">
                  {{ announcement.title }}
                </h3>
                <span
                  v-if="announcement.isPinned"
                  class="inline-flex items-center gap-1 rounded-full bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-700 dark:bg-blue-900/30 dark:text-blue-400"
                >
                  📌 置顶
                </span>
              </div>
              <div class="flex items-center gap-4 text-sm text-slate-500 dark:text-slate-400">
                <span>{{ announcement.author }}</span>
                <span>{{ announcement.createdAt }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- 空状态 -->
        <div
          v-if="paginatedAnnouncements.length === 0"
          class="rounded-xl border border-slate-200 bg-white p-12 text-center dark:border-slate-800 dark:bg-slate-900"
        >
          <p class="text-slate-500 dark:text-slate-400">
            暂无公告
          </p>
        </div>
      </div>

      <!-- 分页 -->
      <div
        v-if="totalPages > 0"
        class="mt-6 flex items-center justify-center gap-2"
      >
        <button
          type="button"
          :disabled="currentPage === 1"
          class="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="currentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
        </button>

        <div class="flex items-center gap-1">
          <span class="px-3 py-2 text-sm text-slate-600 dark:text-slate-400">
            共 {{ allAnnouncements.length }} 条，第 {{ currentPage }} / {{ totalPages }} 页
          </span>
          <template v-if="totalPages > 1">
            <button
              v-for="page in totalPages"
              :key="page"
              type="button"
              class="rounded-lg px-3 py-2 text-sm font-medium transition"
              :class="
                currentPage === page
                  ? 'bg-blue-600 text-white'
                  : 'border border-slate-300 bg-white text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700'
              "
              @click="currentPage = page"
            >
              {{ page }}
            </button>
          </template>
        </div>

        <button
          type="button"
          :disabled="currentPage === totalPages"
          class="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="currentPage++"
        >
          <ChevronRight class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>

