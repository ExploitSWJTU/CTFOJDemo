<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  Plus,
  Trophy,
  Edit,
  Paperclip,
  Database,
  FilePlus,
} from 'lucide-vue-next'
import { challenges as trainingChallenges } from '@/mock/challenges'
import { CATEGORY_MAP } from '@/constants/category'
import type { Category } from '@/types/challenge'
import {
  getContestChallenges,
  addContestChallenge,
  getNextChallengeId,
  type ContestChallenge,
  type ChallengeType,
} from '@/stores/contestChallengeStore'

const route = useRoute()
const router = useRouter()
const contestId = computed(() => route.params.id as string)

const challengeTypeLabel: Record<ChallengeType, string> = {
  static_attachment: '静态附件',
  static_container: '静态容器',
  dynamic_attachment: '动态附件',
  dynamic_container: '动态容器',
}

// 从 store 获取当前比赛的题目列表（与编辑页删除联动）
const contestChallenges = computed(() => getContestChallenges(contestId.value))

function editChallenge(item: ContestChallenge) {
  router.push(`/admin/manage/contest/${contestId.value}/challenges/${item.id}`)
}

function editAttachmentAndFlag(item: ContestChallenge) {
  router.push(`/admin/manage/contest/${contestId.value}/challenges/${item.id}/flag`)
}

// ---------- 三血奖励配置弹窗 ----------
const firstBloodModalVisible = ref(false)
const firstBloodForm = ref({
  first: 5.0,
  second: 3.0,
  third: 1.0,
})

function openFirstBloodSettings() {
  firstBloodModalVisible.value = true
}

function closeFirstBloodModal() {
  firstBloodModalVisible.value = false
}

function saveFirstBloodSettings() {
  // TODO: 提交到后端
  closeFirstBloodModal()
}

// ---------- 新建题目：来源选择 ----------
const addSourceModalVisible = ref(false)

function openAddChallenge() {
  addSourceModalVisible.value = true
}

function closeAddSourceModal() {
  addSourceModalVisible.value = false
}

function openFromBank() {
  closeAddSourceModal()
  fromBankModalVisible.value = true
}

function openNewChallengeForm() {
  closeAddSourceModal()
  newChallengeFormVisible.value = true
}

// ---------- 从题库选择 ----------
const fromBankModalVisible = ref(false)
const selectedBankIds = ref<number[]>([])

const bankList = computed(() => trainingChallenges)

function toggleBankSelect(id: number) {
  const idx = selectedBankIds.value.indexOf(id)
  if (idx === -1) selectedBankIds.value.push(id)
  else selectedBankIds.value.splice(idx, 1)
}

function isBankSelected(id: number) {
  return selectedBankIds.value.includes(id)
}

function confirmAddFromBank() {
  const cid = contestId.value
  trainingChallenges.forEach((c) => {
    if (!selectedBankIds.value.includes(c.id)) return
    const existing = contestChallenges.value.some((x) => x.title === c.title && x.category === c.category)
    if (existing) return
    addContestChallenge(cid, {
      id: getNextChallengeId(cid),
      title: c.title,
      category: c.category,
      challengeType: 'static_attachment',
      description: c.description,
      points: c.points,
      solvedCount: 0,
      difficulty: c.difficulty,
      enabled: true,
      firstBloodReward: '',
    })
  })
  selectedBankIds.value = []
  fromBankModalVisible.value = false
}

function closeFromBankModal() {
  fromBankModalVisible.value = false
  selectedBankIds.value = []
}

// ---------- 新建题目表单 ----------
const newChallengeFormVisible = ref(false)
const newForm = ref<{ title: string; category: Category; challengeType: ChallengeType }>({
  title: '',
  category: 'Misc',
  challengeType: 'static_attachment',
})

const categoryOptions = computed(() =>
  Object.keys(CATEGORY_MAP)
    .filter((k) => k !== 'All')
    .map((k) => ({ value: k as Category, label: k }))
)

function submitNewChallenge() {
  const title = newForm.value.title.trim()
  if (!title) return
  const cid = contestId.value
  addContestChallenge(cid, {
    id: getNextChallengeId(cid),
    title,
    category: newForm.value.category,
    challengeType: newForm.value.challengeType,
    description: '',
    points: 100,
    solvedCount: 0,
    difficulty: 'Easy',
    enabled: true,
    firstBloodReward: '',
  })
  newForm.value = { title: '', category: 'Misc', challengeType: 'static_attachment' as ChallengeType }
  newChallengeFormVisible.value = false
}

function closeNewChallengeForm() {
  newChallengeFormVisible.value = false
  newForm.value = { title: '', category: 'Misc', challengeType: 'static_attachment' as ChallengeType }
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6">
      <div class="mb-6 flex flex-wrap items-center justify-between gap-4">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
          题目管理
        </h3>
        <div class="flex items-center gap-2">
          <button
            type="button"
            class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="openFirstBloodSettings"
          >
            <Trophy class="h-4 w-4" />
            三血奖励
          </button>
          <button
            type="button"
            class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="openAddChallenge"
          >
            <Plus class="h-4 w-4" />
            新建题目
          </button>
        </div>
      </div>

      <!-- 题目卡片网格 -->
      <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <div
          v-for="item in contestChallenges"
          :key="item.id"
          class="relative rounded-xl border border-slate-200 bg-slate-50/50 dark:border-slate-700 dark:bg-slate-800/30 overflow-hidden transition-all"
          :class="{ 'opacity-70': !item.enabled }"
        >
          <!-- 激活滑块（右上） -->
          <div class="absolute top-3 right-3 z-10 flex items-center gap-2">
            <span class="text-xs font-medium text-slate-500 dark:text-slate-400">启用</span>
            <label class="relative inline-flex cursor-pointer items-center">
              <input
                v-model="item.enabled"
                type="checkbox"
                class="peer sr-only"
              />
              <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
            </label>
          </div>

          <div class="p-4 pt-10">
            <div class="mb-3 flex items-center gap-2">
              <div
                class="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg"
                :class="(CATEGORY_MAP[item.category] || CATEGORY_MAP.Misc).cardClass"
              >
                <component :is="(CATEGORY_MAP[item.category] || CATEGORY_MAP.Misc).icon" class="h-4 w-4" />
              </div>
              <div class="min-w-0 flex-1">
                <h4 class="truncate font-semibold text-slate-800 dark:text-slate-100">
                  {{ item.title }}
                </h4>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  {{ item.category }} · {{ challengeTypeLabel[item.challengeType] }} · {{ item.points }} pts
                </p>
              </div>
            </div>

            <div class="flex flex-wrap items-center gap-2">
              <button
                type="button"
                class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-xs font-medium text-slate-700 transition hover:bg-slate-100 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                @click="editChallenge(item)"
              >
                <Edit class="h-3.5 w-3.5" />
                编辑题目
              </button>
              <button
                type="button"
                class="inline-flex items-center gap-1 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-xs font-medium text-slate-700 transition hover:bg-slate-100 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                @click="editAttachmentAndFlag(item)"
              >
                <Paperclip class="h-3.5 w-3.5" />
                编辑附件及 flag
              </button>
            </div>
          </div>
        </div>
      </div>

      <p v-if="contestChallenges.length === 0" class="py-8 text-center text-sm text-slate-500 dark:text-slate-400">
        暂无题目，点击「新建题目」添加。
      </p>
    </section>

    <!-- 三血奖励配置弹窗 -->
    <a-modal
      v-model:visible="firstBloodModalVisible"
      title="三血奖励"
      width="480px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeFirstBloodModal"
    >
      <p class="mb-5 text-sm leading-relaxed text-slate-600 dark:text-slate-400">
        三血奖励加成是指当一个题目被前三个队伍解出时，每个队伍可以得到的分值奖励。三血的奖励基于题目的当前分值，并以一个固定百分比的形式累加至该队伍的得分中。
      </p>
      <div class="space-y-4">
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">一血奖励 (%)</label>
          <input
            v-model.number="firstBloodForm.first"
            type="number"
            step="0.1"
            min="0"
            max="100"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          />
        </div>
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">二血奖励 (%)</label>
          <input
            v-model.number="firstBloodForm.second"
            type="number"
            step="0.1"
            min="0"
            max="100"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          />
        </div>
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">三血奖励 (%)</label>
          <input
            v-model.number="firstBloodForm.third"
            type="number"
            step="0.1"
            min="0"
            max="100"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          />
        </div>
        <div class="flex justify-end gap-2 pt-2">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="closeFirstBloodModal"
          >
            取消
          </button>
          <button
            type="button"
            class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="saveFirstBloodSettings"
          >
            保存
          </button>
        </div>
      </div>
    </a-modal>

    <!-- 新建题目：选择来源 -->
    <a-modal
      v-model:visible="addSourceModalVisible"
      title="新建题目"
      width="400px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeAddSourceModal"
    >
      <div class="space-y-3 py-2">
        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-left transition hover:bg-slate-100 dark:border-slate-700 dark:bg-slate-800 dark:hover:bg-slate-700"
          @click="openFromBank"
        >
          <Database class="h-5 w-5 text-blue-600" />
          <div>
            <span class="font-medium text-slate-800 dark:text-slate-100">从题库选择</span>
            <span class="block text-xs text-slate-500 dark:text-slate-400">从 /admin/manage/training 题库中勾选题目加入本比赛</span>
          </div>
        </button>
        <button
          type="button"
          class="flex w-full items-center gap-3 rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-left transition hover:bg-slate-100 dark:border-slate-700 dark:bg-slate-800 dark:hover:bg-slate-700"
          @click="openNewChallengeForm"
        >
          <FilePlus class="h-5 w-5 text-emerald-600" />
          <div>
            <span class="font-medium text-slate-800 dark:text-slate-100">新建题目</span>
            <span class="block text-xs text-slate-500 dark:text-slate-400">创建一道新题目（填写标题、类别、类型）</span>
          </div>
        </button>
      </div>
    </a-modal>

    <!-- 从题库选择：勾选列表 -->
    <a-modal
      v-model:visible="fromBankModalVisible"
      title="从题库选择题目"
      width="560px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeFromBankModal"
    >
      <p class="mb-4 text-sm text-slate-500 dark:text-slate-400">
        勾选要加入本比赛的题目，点击「确定添加」完成。
      </p>
      <div class="max-h-80 space-y-2 overflow-y-auto rounded-lg border border-slate-200 dark:border-slate-700 p-2">
        <label
          v-for="c in bankList"
          :key="c.id"
          class="flex cursor-pointer items-center gap-3 rounded-lg border px-3 py-2 transition"
          :class="isBankSelected(c.id) ? 'border-blue-300 bg-blue-50 dark:border-blue-700 dark:bg-blue-900/20' : 'border-slate-200 dark:border-slate-700 hover:bg-slate-50 dark:hover:bg-slate-800/50'"
        >
          <input
            type="checkbox"
            :checked="isBankSelected(c.id)"
            class="rounded border-slate-300"
            @change="toggleBankSelect(c.id)"
          />
          <div class="flex flex-1 items-center gap-2">
            <span class="font-medium text-slate-800 dark:text-slate-100">{{ c.title }}</span>
            <span class="text-xs text-slate-500 dark:text-slate-400">{{ c.category }} · {{ c.points }} pts</span>
          </div>
        </label>
      </div>
      <div class="mt-4 flex justify-end gap-2">
        <button
          type="button"
          class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="closeFromBankModal"
        >
          取消
        </button>
        <button
          type="button"
          class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
          @click="confirmAddFromBank"
        >
          确定添加
        </button>
      </div>
    </a-modal>

    <!-- 新建题目：表单 -->
    <a-modal
      v-model:visible="newChallengeFormVisible"
      title="新建题目"
      width="480px"
      :footer="false"
      :unmount-on-close="true"
      modal-class="rounded-xl overflow-hidden"
      @cancel="closeNewChallengeForm"
    >
      <div class="space-y-4">
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">
            题目标题 <span class="text-red-500">*</span>
          </label>
          <input
            v-model="newForm.title"
            type="text"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
            placeholder="请输入题目标题"
          />
        </div>
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">
            题目类别 <span class="text-red-500">*</span>
          </label>
          <select
            v-model="newForm.category"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          >
            <option v-for="opt in categoryOptions" :key="opt.value" :value="opt.value">
              {{ opt.label }}
            </option>
          </select>
        </div>
        <div>
          <label class="mb-1 block text-sm font-medium text-slate-600 dark:text-slate-400">
            题目类型 <span class="text-red-500">*</span>
          </label>
          <select
            v-model="newForm.challengeType"
            class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
          >
            <option value="static_attachment">
              静态附件
            </option>
            <option value="static_container">
              静态容器
            </option>
            <option value="dynamic_attachment">
              动态附件
            </option>
            <option value="dynamic_container">
              动态容器
            </option>
          </select>
        </div>
        <div class="flex justify-end gap-2 pt-2">
          <button
            type="button"
            class="rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="closeNewChallengeForm"
          >
            取消
          </button>
          <button
            type="button"
            class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700 disabled:opacity-50"
            :disabled="!newForm.title.trim()"
            @click="submitNewChallenge"
          >
            创建
          </button>
        </div>
      </div>
    </a-modal>
  </div>
</template>
