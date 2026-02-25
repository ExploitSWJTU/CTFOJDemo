<script setup lang="ts">
import { ref, computed, reactive } from 'vue'
import {
  LayoutGrid,
  Image,
  Activity,
  Hash,
  CalendarCheck2,
  Plus,
  Trash2,
  ChevronUp,
  ChevronDown,
} from 'lucide-vue-next'
import {
  getCarouselSlides,
  addCarouselSlide,
  updateCarouselSlide,
  removeCarouselSlide,
  moveCarouselSlide,
  getFeedCandidatesBySource,
  getSelectedFeedIds,
  setSelectedFeedIds,
  getHotTags,
  addHotTag,
  removeHotTag,
  getTagCounts,
  getCheckInConfig,
  updateCheckInConfig,
} from '../../../stores/homeStore'
import type { CarouselSlide as CarouselSlideType, CheckInConfig } from '../../../stores/homeStore'

const FEED_SOURCE_LABELS = ['全站公告', '训练公告', '比赛公告', '论坛帖子'] as const
const FEED_PAGE_SIZE = 5

type SectionId = 'carousel' | 'feed' | 'tags' | 'checkin'
const sectionTabs: { id: SectionId; label: string; icon: typeof Image }[] = [
  { id: 'carousel', label: '轮播图管理', icon: Image },
  { id: 'feed', label: '最新动态', icon: Activity },
  { id: 'tags', label: '热门标签', icon: Hash },
  { id: 'checkin', label: '每日签到', icon: CalendarCheck2 },
]

const activeSection = ref<SectionId>('carousel')

// ---------- 轮播图管理 ----------
const carouselSlides = computed(() => getCarouselSlides())
const showCarouselForm = ref(false)
const editingSlide = ref<CarouselSlideType | null>(null)
const formSlide = ref({
  imageUrl: '',
  title: '',
  subtitle: '',
  badge: '',
})

function openAddSlide() {
  editingSlide.value = null
  formSlide.value = { imageUrl: '', title: '', subtitle: '', badge: '' }
  showCarouselForm.value = true
}

function openEditSlide(slide: CarouselSlideType) {
  editingSlide.value = slide
  formSlide.value = {
    imageUrl: slide.imageUrl,
    title: slide.title,
    subtitle: slide.subtitle,
    badge: slide.badge ?? '',
  }
  showCarouselForm.value = true
}

function saveCarouselSlide() {
  const { imageUrl, title, subtitle, badge } = formSlide.value
  if (!imageUrl.trim() || !title.trim()) {
    alert('请填写图片地址和标题')
    return
  }
  if (editingSlide.value) {
    updateCarouselSlide(editingSlide.value.id, { imageUrl, title, subtitle, badge: badge || undefined })
  } else {
    addCarouselSlide({ imageUrl, title, subtitle, badge: badge || undefined })
  }
  showCarouselForm.value = false
}

function deleteSlide(slide: CarouselSlideType) {
  if (!confirm('确定删除该轮播图？')) return
  removeCarouselSlide(slide.id)
}

function moveSlide(slide: CarouselSlideType, dir: 'up' | 'down') {
  moveCarouselSlide(slide.id, dir)
}

const carouselFileInput = ref<HTMLInputElement | null>(null)

function triggerCarouselUpload() {
  carouselFileInput.value?.click()
}

function onCarouselFileChange(e: Event) {
  const input = e.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file || !file.type.startsWith('image/')) return
  const reader = new FileReader()
  reader.onload = () => {
    formSlide.value.imageUrl = reader.result as string
  }
  reader.readAsDataURL(file)
  input.value = ''
}

// ---------- 最新动态管理（四列 + 分页） ----------
const feedBySource = computed(() => getFeedCandidatesBySource())
const selectedFeedIds = computed(() => getSelectedFeedIds())
const feedPage = reactive<Record<string, number>>({
  全站公告: 1,
  训练公告: 1,
  比赛公告: 1,
  论坛帖子: 1,
})

function feedListFor(source: (typeof FEED_SOURCE_LABELS)[number]) {
  return feedBySource.value[source] ?? []
}

function feedPaginated(source: (typeof FEED_SOURCE_LABELS)[number]) {
  const list = feedListFor(source)
  const page = feedPage[source] ?? 1
  const start = (page - 1) * FEED_PAGE_SIZE
  return list.slice(start, start + FEED_PAGE_SIZE)
}

function feedTotalPages(source: (typeof FEED_SOURCE_LABELS)[number]) {
  const list = feedListFor(source)
  return Math.max(1, Math.ceil(list.length / FEED_PAGE_SIZE))
}

function setFeedPage(source: (typeof FEED_SOURCE_LABELS)[number], page: number) {
  feedPage[source] = Math.max(1, Math.min(page, feedTotalPages(source)))
}

function toggleFeedSelection(id: string) {
  const ids = [...selectedFeedIds.value]
  const idx = ids.indexOf(id)
  if (idx === -1) ids.push(id)
  else ids.splice(idx, 1)
  setSelectedFeedIds(ids)
}

function isFeedSelected(id: string) {
  return selectedFeedIds.value.includes(id)
}

// ---------- 热门标签统计管理 ----------
const tagCounts = computed(() => getTagCounts())
const hotTags = computed(() => getHotTags())
const newTagInput = ref('')

function addTag() {
  const t = newTagInput.value.trim()
  if (t) {
    addHotTag(t)
    newTagInput.value = ''
  }
}

function isTagOnHome(tag: string) {
  return hotTags.value.includes(tag)
}

function toggleTagOnHome(tag: string) {
  if (isTagOnHome(tag)) removeHotTag(tag)
  else addHotTag(tag.startsWith('#') ? tag : `#${tag}`)
}

// ---------- 每日签到管理 ----------
const checkInForm = ref<CheckInConfig>(getCheckInConfig())

function loadCheckInConfig() {
  checkInForm.value = getCheckInConfig()
}

function saveCheckInConfig() {
  updateCheckInConfig(checkInForm.value)
  alert('签到配置已保存')
}
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex items-center justify-between border-b border-slate-200 dark:border-slate-800 pb-4 shrink-0">
      <h2 class="text-xl font-black text-slate-900 dark:text-white flex items-center gap-2">
        <LayoutGrid class="text-blue-600" :size="24" />
        首页管理
      </h2>
    </div>

    <!-- 区块导航 -->
    <nav class="flex flex-wrap gap-2">
      <button
        v-for="s in sectionTabs"
        :key="s.id"
        type="button"
        class="flex items-center gap-2 px-4 py-2 rounded-xl text-sm font-bold transition-all"
        :class="
          activeSection === s.id
            ? 'bg-blue-600 text-white shadow-lg'
            : 'bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 text-slate-600 dark:text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800'
        "
        @click="activeSection = s.id; if (s.id === 'checkin') loadCheckInConfig()"
      >
        <component :is="s.icon" :size="18" />
        {{ s.label }}
      </button>
    </nav>

    <!-- 1. 轮播图图片管理 -->
    <section
      v-show="activeSection === 'carousel'"
      class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl p-6 shadow-sm"
    >
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-sm font-black text-slate-800 dark:text-slate-100 uppercase tracking-wider">
          轮播图图片管理
        </h3>
        <button
          type="button"
          class="flex items-center gap-2 px-3 py-1.5 bg-blue-600 text-white rounded-lg text-xs font-bold hover:bg-blue-700"
          @click="openAddSlide"
        >
          <Plus :size="14" />
          添加轮播图
        </button>
      </div>
      <div class="space-y-3">
        <div
          v-for="(slide, index) in carouselSlides"
          :key="slide.id"
          class="flex items-center gap-3 rounded-lg border border-slate-200 dark:border-slate-800 p-3"
        >
          <div class="flex flex-col gap-0.5">
            <button
              type="button"
              class="p-0.5 text-slate-400 hover:text-blue-600 disabled:opacity-30"
              :disabled="index === 0"
              @click="moveSlide(slide, 'up')"
            >
              <ChevronUp :size="16" />
            </button>
            <button
              type="button"
              class="p-0.5 text-slate-400 hover:text-blue-600 disabled:opacity-30"
              :disabled="index === carouselSlides.length - 1"
              @click="moveSlide(slide, 'down')"
            >
              <ChevronDown :size="16" />
            </button>
          </div>
          <img
            :src="slide.imageUrl"
            class="h-14 w-24 rounded object-cover bg-slate-100"
            alt=""
            @error="($event.target as HTMLImageElement).src = 'data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 width=%22100%22 height=%22100%22><rect fill=%22%23e2e8f0%22 width=%22100%22 height=%22100%22/><text x=%2250%22 y=%2255%22 fill=%22%2394a3b8%22 text-anchor=%22middle%22 dy=%22.3em%22 font-size=%2212%22>暂无</text></svg>'"
          />
          <div class="flex-1 min-w-0">
            <div class="text-xs font-bold text-slate-800 dark:text-slate-100 truncate">
              {{ slide.title }}
            </div>
            <div class="text-[10px] text-slate-500 dark:text-slate-400 truncate">
              {{ slide.subtitle }}
            </div>
            <span
              v-if="slide.badge"
              class="inline-block mt-1 px-1.5 py-0.5 rounded text-[9px] font-bold bg-blue-100 text-blue-600 dark:bg-blue-900/30 dark:text-blue-400"
            >
              {{ slide.badge }}
            </span>
          </div>
          <div class="flex items-center gap-1">
            <button
              type="button"
              class="p-2 text-slate-400 hover:text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/20 rounded-lg"
              @click="openEditSlide(slide)"
            >
              编辑
            </button>
            <button
              type="button"
              class="p-2 text-slate-400 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg"
              @click="deleteSlide(slide)"
            >
              <Trash2 :size="14" />
            </button>
          </div>
        </div>
        <p v-if="carouselSlides.length === 0" class="text-sm text-slate-500 dark:text-slate-400 py-4 text-center">
          暂无轮播图，点击「添加轮播图」新增
        </p>
      </div>
    </section>

    <!-- 2. 最新动态内容管理（四列分列 + 分页） -->
    <section
      v-show="activeSection === 'feed'"
      class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl p-6 shadow-sm"
    >
      <h3 class="text-sm font-black text-slate-800 dark:text-slate-100 uppercase tracking-wider mb-2">
        最新动态内容管理
      </h3>
      <p class="text-xs text-slate-500 dark:text-slate-400 mb-4">
        勾选要在首页「最新动态」中展示的条目；首页将按时间顺序展示。每列分页显示。
      </p>
      <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-4">
        <div
          v-for="source in FEED_SOURCE_LABELS"
          :key="source"
          class="flex flex-col rounded-lg border border-slate-200 dark:border-slate-800 overflow-hidden"
        >
          <div class="px-3 py-2 bg-slate-50 dark:bg-slate-900 text-xs font-bold text-slate-700 dark:text-slate-300 border-b border-slate-200 dark:border-slate-800">
            {{ source }}
          </div>
          <div class="flex-1 min-h-0 overflow-y-auto max-h-64 space-y-2 p-2">
            <div
              v-for="c in feedPaginated(source)"
              :key="c.id"
              class="flex items-center gap-2 rounded-lg border p-2 transition-colors"
              :class="
                isFeedSelected(c.id)
                  ? 'border-blue-300 dark:border-blue-700 bg-blue-50/50 dark:bg-blue-900/20'
                  : 'border-slate-200 dark:border-slate-800'
              "
            >
              <input
                type="checkbox"
                :checked="isFeedSelected(c.id)"
                class="rounded border-slate-300 text-blue-600 shrink-0"
                @change="toggleFeedSelection(c.id)"
              />
              <div class="flex-1 min-w-0">
                <div class="text-xs font-bold text-slate-800 dark:text-slate-100 truncate">
                  {{ c.title }}
                </div>
                <div class="text-[10px] text-slate-500 dark:text-slate-400">
                  {{ c.time }}
                </div>
              </div>
            </div>
            <p v-if="feedListFor(source).length === 0" class="text-[10px] text-slate-500 py-2 text-center">
              暂无
            </p>
          </div>
          <div v-if="feedTotalPages(source) > 1" class="flex items-center justify-between px-2 py-1.5 border-t border-slate-200 dark:border-slate-800 text-[10px]">
            <span class="text-slate-500">{{ feedPage[source] }} / {{ feedTotalPages(source) }}</span>
            <div class="flex gap-1">
              <button
                type="button"
                class="px-2 py-0.5 rounded border border-slate-200 dark:border-slate-700 disabled:opacity-40"
                :disabled="(feedPage[source] ?? 1) <= 1"
                @click="setFeedPage(source, (feedPage[source] ?? 1) - 1)"
              >
                上一页
              </button>
              <button
                type="button"
                class="px-2 py-0.5 rounded border border-slate-200 dark:border-slate-700 disabled:opacity-40"
                :disabled="(feedPage[source] ?? 1) >= feedTotalPages(source)"
                @click="setFeedPage(source, (feedPage[source] ?? 1) + 1)"
              >
                下一页
              </button>
            </div>
          </div>
        </div>
      </div>
      <p class="mt-3 text-[10px] text-slate-400">
        已选 {{ selectedFeedIds.length }} 条，首页将按时间顺序展示
      </p>
    </section>

    <!-- 3. 热门标签统计管理（出现次数统计 + 选择展示） -->
    <section
      v-show="activeSection === 'tags'"
      class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl p-6 shadow-sm"
    >
      <h3 class="text-sm font-black text-slate-800 dark:text-slate-100 uppercase tracking-wider mb-2">
        热门标签统计管理
      </h3>
      <p class="text-xs text-slate-500 dark:text-slate-400 mb-4">
        统计全站公告与论坛帖子中 #标签 出现次数，勾选「展示在首页」的标签将显示在首页侧栏。
      </p>
      <div class="overflow-x-auto max-h-80 overflow-y-auto border border-slate-200 dark:border-slate-800 rounded-lg">
        <table class="w-full text-left text-xs">
          <thead class="bg-slate-50 dark:bg-slate-900 sticky top-0">
            <tr>
              <th class="px-3 py-2 font-bold text-slate-500">
                标签
              </th>
              <th class="px-3 py-2 font-bold text-slate-500">
                出现次数
              </th>
              <th class="px-3 py-2 font-bold text-slate-500">
                展示在首页
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr
              v-for="tc in tagCounts"
              :key="tc.tag"
              class="hover:bg-slate-50/50 dark:hover:bg-slate-900/30"
            >
              <td class="px-3 py-2 font-mono font-bold text-slate-800 dark:text-slate-200">
                {{ tc.tag }}
              </td>
              <td class="px-3 py-2 text-slate-600 dark:text-slate-400">
                {{ tc.count }}
              </td>
              <td class="px-3 py-2">
                <input
                  type="checkbox"
                  :checked="isTagOnHome(tc.tag)"
                  class="rounded border-slate-300 text-blue-600"
                  @change="toggleTagOnHome(tc.tag)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <p v-if="tagCounts.length === 0" class="py-4 text-center text-sm text-slate-500">
        暂无 #标签 统计（请在公告或论坛内容中使用 #xxx 格式）
      </p>
      <div class="mt-4 flex items-center gap-2">
        <span class="text-[10px] text-slate-500">手动添加首页展示标签：</span>
        <input
          v-model="newTagInput"
          type="text"
          placeholder="如 pwn 或 #pwn"
          class="w-40 h-8 rounded-lg border border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 px-2 text-sm"
          @keydown.enter="addTag"
        />
        <button
          type="button"
          class="flex items-center gap-1 px-3 py-1.5 bg-blue-600 text-white rounded-lg text-xs font-bold hover:bg-blue-700"
          @click="addTag"
        >
          <Plus :size="12" />
          添加
        </button>
      </div>
      <p class="mt-2 text-[10px] text-slate-400">
        当前首页展示：{{ hotTags.join(' ') || '无' }}
      </p>
    </section>

    <!-- 4. 每日签到管理 -->
    <section
      v-show="activeSection === 'checkin'"
      class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl p-6 shadow-sm"
    >
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-sm font-black text-slate-800 dark:text-slate-100 uppercase tracking-wider">
          每日签到管理
        </h3>
        <button
          type="button"
          class="flex items-center gap-2 px-4 py-2 rounded-lg bg-blue-600 text-white text-sm font-bold hover:bg-blue-700"
          @click="saveCheckInConfig"
        >
          保存配置
        </button>
      </div>
      <p class="text-slate-500 dark:text-slate-400 py-8 text-center">
        内容待定
      </p>
    </section>

    <!-- 轮播图编辑弹窗 -->
    <div
      v-if="showCarouselForm"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4"
      @click.self="showCarouselForm = false"
    >
      <div class="w-full max-w-lg rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 p-6 shadow-xl">
        <h4 class="text-lg font-bold text-slate-900 dark:text-slate-100 mb-4">
          {{ editingSlide ? '编辑轮播图' : '添加轮播图' }}
        </h4>
        <div class="space-y-4">
          <div>
            <label class="block text-xs font-bold text-slate-500 uppercase mb-1">图片</label>
            <div class="flex gap-2 items-start">
              <input
                ref="carouselFileInput"
                type="file"
                accept="image/*"
                class="hidden"
                @change="onCarouselFileChange"
              />
              <button
                type="button"
                class="shrink-0 px-3 py-2 rounded-lg border border-slate-200 dark:border-slate-700 text-xs font-bold hover:bg-slate-50 dark:hover:bg-slate-800"
                @click="triggerCarouselUpload"
              >
                上传图片
              </button>
              <input
                v-model="formSlide.imageUrl"
                type="url"
                placeholder="或填写图片地址"
                class="flex-1 h-10 rounded-lg border border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 px-3 text-sm"
              />
            </div>
            <img
              v-if="formSlide.imageUrl"
              :src="formSlide.imageUrl"
              class="mt-2 h-24 rounded-lg object-cover bg-slate-100 max-w-xs"
              alt="预览"
            />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-500 uppercase mb-1">标题</label>
            <input
              v-model="formSlide.title"
              type="text"
              class="w-full h-10 rounded-lg border border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 px-3 text-sm"
            />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-500 uppercase mb-1">副标题 / 描述</label>
            <input
              v-model="formSlide.subtitle"
              type="text"
              class="w-full h-10 rounded-lg border border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 px-3 text-sm"
            />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-500 uppercase mb-1">角标文字（可选）</label>
            <input
              v-model="formSlide.badge"
              type="text"
              class="w-full h-10 rounded-lg border border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-900 px-3 text-sm"
              placeholder="如：正在进行、热门"
            />
          </div>
        </div>
        <div class="mt-6 flex justify-end gap-2">
          <button
            type="button"
            class="px-4 py-2 border border-slate-200 dark:border-slate-800 rounded-lg text-sm font-bold"
            @click="showCarouselForm = false"
          >
            取消
          </button>
          <button
            type="button"
            class="px-4 py-2 bg-blue-600 text-white rounded-lg text-sm font-bold hover:bg-blue-700"
            @click="saveCarouselSlide"
          >
            保存
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
