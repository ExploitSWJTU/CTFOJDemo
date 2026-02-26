<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  Trash2,
  Eye,
  Paperclip,
  Save,
  Container,
  Plus,
  X,
  ExternalLink,
} from 'lucide-vue-next'
import VChart from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent, TitleComponent } from 'echarts/components'
import { CATEGORY_MAP } from '@/constants/category'
import type { Category } from '@/types/challenge'
import {
  removeContestChallenge,
  getContestChallenge,
  updateContestChallenge,
} from '@/stores/contestChallengeStore'

use([CanvasRenderer, LineChart, GridComponent, TooltipComponent, TitleComponent])

type ChallengeType = 'static_attachment' | 'static_container' | 'dynamic_attachment' | 'dynamic_container'

const challengeTypeLabel: Record<ChallengeType, string> = {
  static_attachment: '静态附件',
  static_container: '静态容器',
  dynamic_attachment: '动态附件',
  dynamic_container: '动态容器',
}

const route = useRoute()
const router = useRouter()
const contestId = computed(() => route.params.id as string)
const challengeId = computed(() => route.params.challengeId as string)

const form = ref({
  title: 'web1',
  challengeType: 'dynamic_container' as ChallengeType,
  category: 'Web' as Category,
  description: '',
  maxSubmissions: 0,
  points: 1000,
  difficulty: 5.0,
  minScoreRatio: 0.3, // 0–1，对应 0%–100%
  firstBloodEnabled: true,
  image: 'signin',
  port: 5000,
  networkMode: 'public',
  cpuLimit: 1,
  memoryLimit: 64,
  storageLimit: 256,
})

const hints = ref<string[]>([])
const testContainerEnabled = ref(false)
const testContainerUrl = ref('http://challenge.example.com:5000')
const deleteConfirmVisible = ref(false)

const categoryOptions = computed(() =>
  Object.keys(CATEGORY_MAP)
    .filter((k) => k !== 'All')
    .map((k) => ({ value: k as Category, label: k }))
)

const flagPath = computed(() =>
  `/admin/manage/contest/${contestId.value}/challenges/${challengeId.value}/flag`
)

const isContainerType = computed(
  () =>
    form.value.challengeType === 'static_container' ||
    form.value.challengeType === 'dynamic_container'
)

// 最低分比例 0–100 显示，内部 0–1
const minScoreRatioPercent = computed({
  get: () => Math.round(form.value.minScoreRatio * 100),
  set: (v: number) => { form.value.minScoreRatio = Math.max(0, Math.min(100, v)) / 100 },
})

function addHint() {
  hints.value.push('')
}

function removeHint(index: number) {
  hints.value.splice(index, 1)
}

// 分值随解出次数：score(n) = minScore + (points - minScore) * 0.5^(n / max(0.5, difficulty))；整数解出次数对应整数分数
function scoreAtSolves(solves: number): number {
  const p = Math.floor(Number(form.value.points)) || 0
  const r = form.value.minScoreRatio
  const d = Math.max(0.1, form.value.difficulty)
  const minScore = p * r
  const decay = Math.pow(0.5, solves / d)
  return Math.round(minScore + (p - minScore) * decay)
}

const scoreChartOption = computed(() => {
  const maxSolves = 150
  const xData: number[] = []
  const yData: number[] = []
  for (let n = 0; n <= maxSolves; n++) {
    xData.push(n)
    yData.push(scoreAtSolves(n))
  }
  return {
    title: { text: '题目分值随解出次数变化', left: 'center', textStyle: { fontSize: 12 } },
    tooltip: {
      trigger: 'axis',
      formatter: (params: unknown) => {
        const p = Array.isArray(params) ? params[0] : null
        if (p && typeof p === 'object' && 'dataIndex' in p) {
          const i = p.dataIndex as number
          return `解出 ${xData[i]} 次 · 分值 ${yData[i]}`
        }
        return ''
      },
    },
    grid: { left: 48, right: 24, top: 40, bottom: 36 },
    xAxis: {
      type: 'category',
      name: '解出次数',
      nameLocation: 'middle',
      nameGap: 28,
      data: xData,
      boundaryGap: false,
      axisLabel: { interval: 14 },
    },
    yAxis: {
      type: 'value',
      name: '分值',
      nameLocation: 'middle',
      nameGap: 40,
    },
    series: [{
      type: 'line',
      data: yData,
      smooth: true,
      symbol: 'none',
      lineStyle: { width: 2 },
    }],
  }
})

function goBack() {
  router.push(`/admin/manage/contest/${contestId.value}/challenges`)
}

function openDeleteConfirm() {
  deleteConfirmVisible.value = true
}

function closeDeleteConfirm() {
  deleteConfirmVisible.value = false
}

function confirmDeleteChallenge() {
  removeContestChallenge(contestId.value, Number(challengeId.value))
  closeDeleteConfirm()
  goBack()
}

function previewChallenge() {
  // TODO: 打开预览弹窗或新页
}

function saveChallenge() {
  const cid = contestId.value
  const id = Number(challengeId.value)
  const diff = form.value.difficulty
  const difficultyLabel = diff <= 3 ? 'Easy' : diff <= 6 ? 'Medium' : 'Hard'
  updateContestChallenge(cid, id, {
    title: form.value.title.trim(),
    category: form.value.category,
    challengeType: form.value.challengeType,
    description: form.value.description,
    points: Math.max(0, Math.round(form.value.points)),
    difficulty: difficultyLabel,
    firstBloodReward: form.value.firstBloodEnabled ? '+10 积分' : '',
  })
  goBack()
}

function createTestContainer() {
  testContainerEnabled.value = true
  testContainerUrl.value = `http://challenge-${challengeId.value}.example.com:${form.value.port}`
}

function destroyTestContainer() {
  testContainerEnabled.value = false
}

function loadChallengeIntoForm() {
  const c = getContestChallenge(contestId.value, Number(challengeId.value))
  if (!c) return
  form.value.title = c.title
  form.value.challengeType = c.challengeType
  form.value.category = c.category
  form.value.description = c.description
  form.value.points = c.points
  form.value.difficulty = c.difficulty === 'Easy' ? 3 : c.difficulty === 'Medium' ? 5 : c.difficulty === 'Hard' ? 7 : 5
  form.value.firstBloodEnabled = !!c.firstBloodReward
}

onMounted(loadChallengeIntoForm)
watch([contestId, challengeId], loadChallengeIntoForm)
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6 md:p-8">
      <!-- 顶部四个按钮 -->
      <div class="mb-8 flex flex-wrap items-center gap-3">
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg border border-red-200 bg-white px-4 py-2 text-sm font-medium text-red-700 transition hover:bg-red-50 dark:border-red-900 dark:bg-slate-800 dark:text-red-300 dark:hover:bg-red-900/20"
          @click="openDeleteConfirm"
        >
          <Trash2 class="h-4 w-4" />
          删除题目
        </button>
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="previewChallenge"
        >
          <Eye class="h-4 w-4" />
          预览题目
        </button>
        <RouterLink
          :to="flagPath"
          class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
        >
          <Paperclip class="h-4 w-4" />
          编辑附件及flag
        </RouterLink>
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="saveChallenge"
        >
          <Save class="h-4 w-4" />
          保存修改
        </button>
      </div>

      <div class="space-y-8 max-w-4xl">
        <!-- 基础信息 -->
        <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            基础信息
          </h4>
          <div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
            <div class="sm:col-span-2">
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                题目标题 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.title"
                type="text"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="web1"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">题目类型</label>
              <div class="rounded-lg border border-slate-200 bg-slate-100 px-4 py-2.5 text-sm text-slate-600 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300">
                {{ challengeTypeLabel[form.challengeType] }}
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                题目类别 <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.category"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              >
                <option v-for="opt in categoryOptions" :key="opt.value" :value="opt.value">
                  {{ opt.label }}
                </option>
              </select>
            </div>
            <div class="sm:col-span-2">
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">题目描述</label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">
                支持 Markdown 语法
              </p>
              <textarea
                v-model="form.description"
                rows="6"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="题目描述..."
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">最大提交次数</label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">
                0 表示不限制
              </p>
              <input
                v-model.number="form.maxSubmissions"
                type="number"
                min="0"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div class="sm:col-span-2">
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">题目提示</label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-2">
                支持 Inline Markdown 语法
              </p>
              <div class="space-y-2">
                <div
                  v-for="(_, index) in hints"
                  :key="index"
                  class="flex items-center gap-2"
                >
                  <input
                    v-model="hints[index]"
                    type="text"
                    class="flex-1 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                    placeholder="提示内容"
                  />
                  <button
                    type="button"
                    class="shrink-0 rounded-lg p-2 text-slate-400 hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                    title="删除提示"
                    @click="removeHint(index)"
                  >
                    <X class="h-4 w-4" />
                  </button>
                </div>
                <button
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-dashed border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition hover:border-blue-400 hover:text-blue-600 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-blue-500 dark:hover:text-blue-400"
                  @click="addHint"
                >
                  <Plus class="h-4 w-4" />
                  添加提示
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- 分值设置 + 曲线图 -->
        <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            分值设置
          </h4>
          <div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                题目分值 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.points"
                type="number"
                min="0"
                step="1"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                @blur="form.points = Math.max(0, Math.round(form.points))"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                难度系数 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.difficulty"
                type="number"
                step="0.01"
                min="0"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div class="sm:col-span-2">
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                最低分比例 <span class="text-red-500">*</span>
              </label>
              <div class="flex items-center gap-4">
                <input
                  v-model.number="minScoreRatioPercent"
                  type="range"
                  min="0"
                  max="100"
                  class="h-2 flex-1 rounded-full bg-slate-200 dark:bg-slate-700 accent-blue-600"
                />
                <span class="w-12 shrink-0 text-right text-sm font-medium text-slate-700 dark:text-slate-200">
                  {{ minScoreRatioPercent }}%
                </span>
              </div>
              <p class="mt-1 text-xs text-slate-500 dark:text-slate-400">
                0%–100%
              </p>
            </div>
            <div class="sm:col-span-2 flex items-center justify-between rounded-lg border border-slate-200 bg-white px-4 py-3 dark:border-slate-700 dark:bg-slate-800">
              <div>
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">三血奖励</span>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  启用题目的三血加分滑块
                </p>
              </div>
              <label class="relative inline-flex cursor-pointer items-center">
                <input v-model="form.firstBloodEnabled" type="checkbox" class="peer sr-only" />
                <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
              </label>
            </div>
            <div class="sm:col-span-2 min-h-[260px] w-full rounded-lg border border-slate-200 bg-white p-4 dark:border-slate-700 dark:bg-slate-800">
              <VChart
                :option="scoreChartOption"
                class="h-64 w-full"
                style="height: 256px; width: 100%;"
                autoresize
              />
            </div>
          </div>
        </div>

        <!-- 容器配置（仅题目类型为静态容器/动态容器时显示） -->
        <div
          v-if="isContainerType"
          class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30"
        >
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            容器配置
          </h4>
          <div class="grid grid-cols-1 gap-6 sm:grid-cols-2">
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                容器镜像 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="form.image"
                type="text"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="signin"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                服务端口 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.port"
                type="number"
                min="1"
                max="65535"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div class="sm:col-span-2">
              <div class="flex flex-wrap items-center gap-3">
                <button
                  v-if="!testContainerEnabled"
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                  @click="createTestContainer"
                >
                  <Container class="h-4 w-4" />
                  创建测试容器
                </button>
                <template v-else>
                  <div class="flex flex-1 min-w-0 flex-wrap items-center gap-2 rounded-lg border border-emerald-200 bg-emerald-50/50 px-4 py-3 dark:border-emerald-800 dark:bg-emerald-900/20">
                    <span class="text-sm font-medium text-slate-700 dark:text-slate-200">实例入口：</span>
                    <a
                      :href="testContainerUrl"
                      target="_blank"
                      rel="noopener noreferrer"
                      class="inline-flex items-center gap-1 truncate text-sm text-blue-600 hover:underline dark:text-blue-400"
                    >
                      {{ testContainerUrl }}
                      <ExternalLink class="h-3.5 w-3.5 shrink-0" />
                    </a>
                  </div>
                  <button
                    type="button"
                    class="inline-flex items-center gap-2 rounded-lg border border-red-200 bg-white px-4 py-2 text-sm font-medium text-red-700 transition hover:bg-red-50 dark:border-red-900 dark:bg-slate-800 dark:text-red-300 dark:hover:bg-red-900/20"
                    @click="destroyTestContainer"
                  >
                    <Trash2 class="h-4 w-4" />
                    销毁测试容器
                  </button>
                </template>
              </div>
              <p v-if="!testContainerEnabled" class="mt-2 text-sm text-slate-500 dark:text-slate-400">
                测试容器未开启
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                网络模式 <span class="text-red-500">*</span>
              </label>
              <select
                v-model="form.networkMode"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              >
                <option value="public">
                  公共网络访问 - 开放
                </option>
                <option value="isolated">
                  隔离网络
                </option>
              </select>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                CPU 限制 (0.1 CPUs) <span class="text-red-500">*</span>
              </label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">
                乘以 0.1 即为 CPU 核心数
              </p>
              <input
                v-model.number="form.cpuLimit"
                type="number"
                step="0.1"
                min="0.1"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                内存限制 (MB) <span class="text-red-500">*</span>
              </label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">
                限制容器使用的 RAM
              </p>
              <input
                v-model.number="form.memoryLimit"
                type="number"
                min="1"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                存储限制 (MB) <span class="text-red-500">*</span>
              </label>
              <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">
                限制存储空间，含镜像大小
              </p>
              <input
                v-model.number="form.storageLimit"
                type="number"
                min="1"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 删除题目确认弹窗 -->
    <a-modal
      v-model:visible="deleteConfirmVisible"
      title="删除题目"
      width="400px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeDeleteConfirm"
    >
      <p class="py-2 text-sm text-slate-600 dark:text-slate-400">
        确定删除该题目？此操作不可恢复。
      </p>
      <div class="flex justify-end gap-2 pt-4">
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="closeDeleteConfirm"
        >
          取消
        </button>
        <button
          type="button"
          class="rounded-lg bg-red-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-red-700"
          @click="confirmDeleteChallenge"
        >
          确定删除
        </button>
      </div>
    </a-modal>
  </div>
</template>
