<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { MessageSquare, ThumbsUp, Plus, ChevronLeft, ChevronRight, Search } from 'lucide-vue-next'
import { searchForumPosts } from '../../../stores/forumStore'

const router = useRouter()
const route = useRoute()

const searchKeyword = ref((route.query.q as string) || '')
const pageSize = ref(10)
const currentPage = ref(Number(route.query.page) || 1)

const filteredPosts = computed(() => searchForumPosts(searchKeyword.value))
const totalPages = computed(() =>
  Math.max(1, Math.ceil(filteredPosts.value.length / pageSize.value))
)
const paginatedPosts = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  return filteredPosts.value.slice(start, start + pageSize.value)
})

function doSearch() {
  currentPage.value = 1
  router.replace({
    query: {
      ...route.query,
      q: searchKeyword.value || undefined,
      page: currentPage.value > 1 ? currentPage.value : undefined,
    },
  })
}

watch(currentPage, (page) => {
  router.replace({
    query: {
      ...route.query,
      page: page > 1 ? page : undefined,
    },
  })
})

onMounted(() => {
  const pageFromUrl = Number(route.query.page)
  if (pageFromUrl >= 1) currentPage.value = Math.min(pageFromUrl, totalPages.value)
  const q = route.query.q
  if (typeof q === 'string') searchKeyword.value = q
})

watch(
  () => [route.query.page, route.query.q],
  ([page, q]) => {
    const p = Number(page)
    if (p >= 1 && p !== currentPage.value) currentPage.value = Math.min(p, totalPages.value)
    if (typeof q === 'string' && q !== searchKeyword.value) searchKeyword.value = q
  }
)

function goToPost(id: number) {
  router.push({ path: `/forum/${id}` })
}

function goToCreate() {
  router.push({ path: '/forum/create' })
}

function excerpt(text: string, maxLen: number) {
  const t = (text || '').replace(/\n/g, ' ').trim()
  return t.length <= maxLen ? t : t.slice(0, maxLen) + '…'
}
</script>

<template>
  <div class="min-h-[calc(100vh-64px)]">
    <!-- 顶部横幅 -->
    <section class="border-b border-slate-200 bg-gradient-to-b from-slate-50 to-white px-6 py-10 dark:border-slate-800 dark:from-slate-900/50 dark:to-slate-900">
      <div class="mx-auto max-w-4xl">
        <div class="flex flex-col gap-6 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 class="text-3xl font-bold tracking-tight text-slate-900 dark:text-slate-50 sm:text-4xl">
              交流论坛
            </h1>
            <p class="mt-2 text-slate-500 dark:text-slate-400">
              分享思路、讨论题目、交流经验
            </p>
          </div>
          <button
            type="button"
            class="inline-flex shrink-0 items-center justify-center gap-2 rounded-xl bg-blue-600 px-5 py-2.5 text-sm font-medium text-white shadow-sm transition hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-offset-slate-900"
            @click="goToCreate"
          >
            <Plus class="h-5 w-5" />
            发布帖子
          </button>
        </div>
      </div>
    </section>

    <div class="mx-auto max-w-4xl px-6 py-8">
      <!-- 搜索 -->
      <div class="mb-8">
        <div class="flex gap-3">
          <div class="relative flex-1">
            <Search
              class="absolute left-4 top-1/2 h-5 w-5 -translate-y-1/2 text-slate-400"
              aria-hidden="true"
            />
            <input
              v-model="searchKeyword"
              type="search"
              placeholder="搜索帖子标题或内容…"
              class="w-full rounded-xl border border-slate-200 bg-white py-3 pl-12 pr-4 text-slate-800 placeholder-slate-400 shadow-sm transition focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder-slate-500"
              @keydown.enter="doSearch"
            />
          </div>
          <button
            type="button"
            class="rounded-xl border border-slate-200 bg-white px-5 py-3 text-sm font-medium text-slate-700 shadow-sm transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="doSearch"
          >
            搜索
          </button>
        </div>
      </div>

      <!-- 帖子列表 -->
      <div class="space-y-1">
        <article
          v-for="(post, i) in paginatedPosts"
          :key="post.id"
          class="group relative flex cursor-pointer gap-4 rounded-2xl border border-slate-200/80 bg-white p-5 transition-all hover:border-blue-200 hover:shadow-md dark:border-slate-800 dark:bg-slate-900/50 dark:hover:border-blue-900/60 dark:hover:bg-slate-900"
          @click="goToPost(post.id)"
        >
          <div
            class="absolute left-0 top-6 h-12 w-1 rounded-r-full bg-blue-500 opacity-0 transition-opacity group-hover:opacity-100"
            aria-hidden="true"
          />
          <div class="flex min-w-0 flex-1 flex-col gap-3">
            <h2 class="line-clamp-2 text-lg font-semibold leading-snug text-slate-900 transition group-hover:text-blue-600 dark:text-slate-50 dark:group-hover:text-blue-400">
              {{ post.title }}
            </h2>
            <p class="line-clamp-2 text-sm leading-relaxed text-slate-500 dark:text-slate-400">
              {{ excerpt(post.content, 120) }}
            </p>
            <footer class="flex flex-wrap items-center gap-x-4 gap-y-1 text-xs text-slate-500 dark:text-slate-400">
              <span class="font-medium text-slate-600 dark:text-slate-300">{{ post.author }}</span>
              <span>{{ post.createdAt }}</span>
              <span class="inline-flex items-center gap-1">
                <MessageSquare class="h-3.5 w-3.5" />
                {{ post.replyCount ?? 0 }}
              </span>
              <span class="inline-flex items-center gap-1">
                <ThumbsUp class="h-3.5 w-3.5" />
                {{ (post.likedByUserIds?.length ?? 0) }}
              </span>
            </footer>
          </div>
          <span class="hidden shrink-0 text-2xl font-bold tabular-nums text-slate-200 dark:text-slate-600 sm:block">
            {{ (currentPage - 1) * pageSize + i + 1 }}
          </span>
        </article>

        <div
          v-if="paginatedPosts.length === 0"
          class="rounded-2xl border border-dashed border-slate-200 bg-slate-50/50 py-16 text-center dark:border-slate-700 dark:bg-slate-800/30"
        >
          <p class="text-slate-500 dark:text-slate-400">
            {{ searchKeyword ? '没有找到相关帖子' : '暂无帖子，快来发布第一篇吧' }}
          </p>
          <button
            v-if="!searchKeyword"
            type="button"
            class="mt-4 text-blue-600 hover:underline dark:text-blue-400"
            @click="goToCreate"
          >
            发布帖子
          </button>
        </div>
      </div>

      <!-- 分页 -->
      <nav
        v-if="totalPages > 0"
        class="mt-10 flex flex-wrap items-center justify-center gap-2"
        aria-label="分页"
      >
        <button
          type="button"
          :disabled="currentPage === 1"
          class="inline-flex h-10 items-center rounded-xl border border-slate-200 bg-white px-4 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="currentPage--"
        >
          <ChevronLeft class="h-4 w-4" />
          上一页
        </button>
        <div class="flex items-center gap-1">
          <span class="px-3 text-sm text-slate-500 dark:text-slate-400">
            共 {{ filteredPosts.length }} 条
          </span>
          <template v-if="totalPages > 1">
            <button
              v-for="page in totalPages"
              :key="page"
              type="button"
              class="inline-flex h-10 min-w-10 items-center justify-center rounded-xl px-3 text-sm font-medium transition"
              :class="
                currentPage === page
                  ? 'bg-blue-600 text-white shadow-sm'
                  : 'border border-slate-200 bg-white text-slate-700 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700'
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
          class="inline-flex h-10 items-center rounded-xl border border-slate-200 bg-white px-4 text-sm font-medium text-slate-700 transition hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="currentPage++"
        >
          下一页
          <ChevronRight class="h-4 w-4" />
        </button>
      </nav>
    </div>
  </div>
</template>
