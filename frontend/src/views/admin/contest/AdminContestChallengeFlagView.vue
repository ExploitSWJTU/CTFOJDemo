<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Paperclip, Plus, Upload, X, Save } from 'lucide-vue-next'
import { getContestChallenge } from '@/stores/contestChallengeStore'
import type { ChallengeType } from '@/stores/contestChallengeStore'

const route = useRoute()
const contestId = computed(() => route.params.id as string)
const challengeId = computed(() => route.params.challengeId as string)

const challenge = ref<ReturnType<typeof getContestChallenge>>(undefined)
const challengeType = computed<ChallengeType | undefined>(() => challenge.value?.challengeType)

const isStaticAttachment = computed(() => challengeType.value === 'static_attachment')
const isStaticContainer = computed(() => challengeType.value === 'static_container')
const isDynamicAttachment = computed(() => challengeType.value === 'dynamic_attachment')
const isDynamicContainer = computed(() => challengeType.value === 'dynamic_container')

const useStaticAttachmentForm = computed(() => isStaticAttachment.value || isStaticContainer.value)

// 动态附件：多对「附件列表 + 对应 Flag」，每对分发给一个队/用户；每组支持批量上传多个附件
interface AttachmentFlagPair {
  attachments: { name: string; size: number }[]
  flag: string
}
const dynamicPairs = ref<AttachmentFlagPair[]>([])

function addDynamicPair() {
  dynamicPairs.value.push({ attachments: [], flag: '' })
}

function removeDynamicPair(index: number) {
  dynamicPairs.value.splice(index, 1)
}

function handleDynamicBatchUpload(index: number) {
  const input = document.createElement('input')
  input.type = 'file'
  input.multiple = true
  input.onchange = () => {
    if (!input.files?.length) return
    const pair = dynamicPairs.value[index]
    if (!pair) return
    for (let i = 0; i < input.files.length; i++) {
      const f = input.files[i]
      if (!f) continue
      pair.attachments.push({ name: f.name, size: f.size })
    }
  }
  input.click()
}

function removeDynamicAttachmentFile(pairIndex: number, fileIndex: number) {
  const pair = dynamicPairs.value[pairIndex]
  if (pair) pair.attachments.splice(fileIndex, 1)
}

function clearDynamicAttachments(pairIndex: number) {
  const pair = dynamicPairs.value[pairIndex]
  if (pair) pair.attachments = []
}

/** 批量上传几组：一次选择多个文件，每个文件生成一组（一组一个附件 + 一个 Flag） */
function handleBatchAddGroups() {
  const input = document.createElement('input')
  input.type = 'file'
  input.multiple = true
  input.onchange = () => {
    if (!input.files?.length) return
    for (let i = 0; i < input.files.length; i++) {
      const f = input.files[i]
      if (!f) continue
      dynamicPairs.value.push({
        attachments: [{ name: f.name, size: f.size }],
        flag: '',
      })
    }
  }
  input.click()
}

// 静态附件：附件管理
type AttachmentType = 'none' | 'local'
const attachmentType = ref<AttachmentType>('none')
const uploadedFiles = ref<{ name: string; size: number }[]>([])

function handleUploadClick() {
  const input = document.createElement('input')
  input.type = 'file'
  input.multiple = true
  input.onchange = () => {
    if (!input.files) return
    for (let i = 0; i < input.files.length; i++) {
      const f = input.files[i]
      if (!f) continue
      uploadedFiles.value.push({ name: f.name, size: f.size })
    }
  }
  input.click()
}

function removeFile(index: number) {
  uploadedFiles.value.splice(index, 1)
}

// 静态附件：Flag 管理（与添加提示方式一致）
const flags = ref<string[]>([])

// 动态容器：附件管理与静态相同，Flag 为模版字符串（留空 = 随机 GUID）
const flagTemplate = ref('')

function addFlag() {
  flags.value.push('')
}

function removeFlag(index: number) {
  flags.value.splice(index, 1)
}

function loadChallenge() {
  challenge.value = getContestChallenge(contestId.value, Number(challengeId.value))
  // TODO: 从后端或本地持久化加载附件类型、已上传文件、flags
}

const saving = ref(false)

async function saveFlagConfig() {
  if (saving.value) return
  saving.value = true
  try {
    // TODO: 提交附件与 Flag 配置到后端
    await new Promise((r) => setTimeout(r, 500))
    alert('保存成功')
  } catch {
    alert('保存失败，请重试')
  } finally {
    saving.value = false
  }
}

onMounted(loadChallenge)
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6 md:p-8">
      <div class="mb-6 flex flex-wrap items-center justify-between gap-4">
        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 flex items-center gap-2">
          <Paperclip class="h-5 w-5" />
          编辑附件及 Flag
        </h3>
        <button
          v-if="challenge"
          type="button"
          class="inline-flex items-center gap-2 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700 disabled:opacity-60 disabled:cursor-not-allowed"
          :disabled="saving"
          @click="saveFlagConfig"
        >
          <Save class="h-4 w-4" />
          {{ saving ? '保存中…' : '保存' }}
        </button>
      </div>

      <template v-if="!challenge">
        <p class="text-slate-500 dark:text-slate-400">
          未找到该题目。
        </p>
      </template>

      <!-- 静态附件 / 静态容器：同一套附件管理 + Flag 管理 -->
      <template v-else-if="useStaticAttachmentForm">
        <div class="space-y-8 max-w-4xl">
          <!-- 附件管理 -->
          <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
            <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
              附件管理
            </h4>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  附件类型 <span class="text-red-500">*</span>
                </label>
                <div class="flex flex-wrap gap-6">
                  <label class="flex cursor-pointer items-center gap-2">
                    <input
                      v-model="attachmentType"
                      type="radio"
                      value="none"
                      class="rounded-full border-slate-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="text-sm text-slate-700 dark:text-slate-200">无附件</span>
                  </label>
                  <label class="flex cursor-pointer items-center gap-2">
                    <input
                      v-model="attachmentType"
                      type="radio"
                      value="local"
                      class="rounded-full border-slate-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="text-sm text-slate-700 dark:text-slate-200">本地附件</span>
                  </label>
                </div>
              </div>
              <div v-if="attachmentType === 'local'" class="space-y-2">
                <button
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                  @click="handleUploadClick"
                >
                  <Upload class="h-4 w-4" />
                  上传附件
                </button>
                <div v-if="uploadedFiles.length" class="mt-2 space-y-2">
                  <div
                    v-for="(f, index) in uploadedFiles"
                    :key="index"
                    class="flex items-center justify-between rounded-lg border border-slate-200 bg-white px-4 py-2 dark:border-slate-700 dark:bg-slate-800"
                  >
                    <span class="text-sm text-slate-700 dark:text-slate-200">{{ f.name }}</span>
                    <span class="text-xs text-slate-500 dark:text-slate-400">
                      {{ (f.size / 1024).toFixed(1) }} KB
                    </span>
                    <button
                      type="button"
                      class="rounded p-1.5 text-slate-400 hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                      title="删除"
                      @click="removeFile(index)"
                    >
                      <X class="h-4 w-4" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Flag 管理 -->
          <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
            <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
              Flag 管理
            </h4>
            <div class="space-y-2">
              <div
                v-for="(_, index) in flags"
                :key="index"
                class="flex items-center gap-2"
              >
                <input
                  v-model="flags[index]"
                  type="text"
                  class="flex-1 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                  placeholder="flag{...}"
                />
                <button
                  type="button"
                  class="shrink-0 rounded-lg p-2 text-slate-400 hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                  title="删除"
                  @click="removeFlag(index)"
                >
                  <X class="h-4 w-4" />
                </button>
              </div>
              <button
                type="button"
                class="inline-flex items-center gap-2 rounded-lg border border-dashed border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition hover:border-blue-400 hover:text-blue-600 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-blue-500 dark:hover:text-blue-400"
                @click="addFlag"
              >
                <Plus class="h-4 w-4" />
                添加 Flag
              </button>
            </div>
          </div>
        </div>
      </template>

      <!-- 动态附件：多对「附件 + Flag」，每对分发给一个队/用户 -->
      <template v-else-if="isDynamicAttachment">
        <div class="space-y-8 max-w-4xl">
          <p class="text-sm text-slate-500 dark:text-slate-400">
            可添加多组「附件 + 对应 Flag」。支持「批量上传几组」：一次选择多个文件，每个文件自动生成一组，再为每组填写对应 Flag。比赛时每个队/用户将随机或按序获得其中一组附件及对应 Flag。
          </p>
          <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
            <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
              附件与 Flag 配对
            </h4>
            <div class="space-y-4">
              <div
                v-for="(pair, index) in dynamicPairs"
                :key="index"
                class="flex flex-col gap-3 rounded-lg border border-slate-200 bg-white p-4 dark:border-slate-700 dark:bg-slate-800"
              >
                <div class="flex items-center justify-between gap-2">
                  <span class="text-xs font-medium text-slate-500 dark:text-slate-400">第 {{ index + 1 }} 组</span>
                  <button
                    type="button"
                    class="rounded p-1.5 text-slate-400 hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                    title="删除该组"
                    @click="removeDynamicPair(index)"
                  >
                    <X class="h-4 w-4" />
                  </button>
                </div>
                <div class="grid gap-3 sm:grid-cols-2">
                  <div>
                    <label class="mb-1 block text-xs font-medium text-slate-600 dark:text-slate-400">附件（可批量上传）</label>
                    <div class="space-y-2">
                      <button
                        type="button"
                        class="inline-flex items-center gap-1.5 rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm text-slate-700 transition hover:bg-slate-100 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                        @click="handleDynamicBatchUpload(index)"
                      >
                        <Upload class="h-3.5 w-3.5" />
                        批量上传附件
                      </button>
                      <div v-if="pair.attachments.length" class="space-y-1.5">
                        <div
                          v-for="(file, fileIndex) in pair.attachments"
                          :key="fileIndex"
                          class="flex items-center justify-between gap-2 rounded-lg border border-slate-200 bg-slate-50/50 px-3 py-2 dark:border-slate-600 dark:bg-slate-800/50"
                        >
                          <span class="min-w-0 truncate text-sm text-slate-600 dark:text-slate-300">{{ file.name }}</span>
                          <span class="shrink-0 text-xs text-slate-400">({{ (file.size / 1024).toFixed(1) }} KB)</span>
                          <button
                            type="button"
                            class="shrink-0 rounded p-1 text-slate-400 hover:text-red-600"
                            title="移除"
                            @click="removeDynamicAttachmentFile(index, fileIndex)"
                          >
                            <X class="h-3.5 w-3.5" />
                          </button>
                        </div>
                        <button
                          type="button"
                          class="text-xs text-slate-500 hover:text-red-600"
                          @click="clearDynamicAttachments(index)"
                        >
                          清空本组附件
                        </button>
                      </div>
                    </div>
                  </div>
                  <div>
                    <label class="mb-1 block text-xs font-medium text-slate-600 dark:text-slate-400">对应 Flag</label>
                    <input
                      v-model="pair.flag"
                      type="text"
                      class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm transition focus:border-blue-500 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-200"
                      placeholder="flag{...}"
                    />
                  </div>
                </div>
              </div>
              <div class="flex flex-wrap gap-2">
                <button
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-dashed border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition hover:border-blue-400 hover:text-blue-600 dark:border-slate-600 dark:bg-slate-800 dark:text-slate-300 dark:hover:border-blue-500 dark:hover:text-blue-400"
                  @click="addDynamicPair"
                >
                  <Plus class="h-4 w-4" />
                  添加附件及 Flag
                </button>
                <button
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-600 transition hover:border-slate-800 hover:text-blue-600 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700 dark:hover:text-blue-400"
                  @click="handleBatchAddGroups"
                >
                  <Upload class="h-4 w-4" />
                  批量上传几组
                </button>
              </div>
            </div>
          </div>
        </div>
      </template>

      <!-- 动态容器：附件管理与静态相同，Flag 为模版（[GUID]/[TEAM_HASH]/[LEET]/[CLEET]） -->
      <template v-else-if="isDynamicContainer">
        <div class="space-y-8 max-w-4xl">
          <!-- 附件管理（与静态容器相同） -->
          <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
            <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
              附件管理
            </h4>
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                  附件类型 <span class="text-red-500">*</span>
                </label>
                <div class="flex flex-wrap gap-6">
                  <label class="flex cursor-pointer items-center gap-2">
                    <input
                      v-model="attachmentType"
                      type="radio"
                      value="none"
                      class="rounded-full border-slate-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="text-sm text-slate-700 dark:text-slate-200">无附件</span>
                  </label>
                  <label class="flex cursor-pointer items-center gap-2">
                    <input
                      v-model="attachmentType"
                      type="radio"
                      value="local"
                      class="rounded-full border-slate-300 text-blue-600 focus:ring-blue-500"
                    />
                    <span class="text-sm text-slate-700 dark:text-slate-200">本地附件</span>
                  </label>
                </div>
              </div>
              <div v-if="attachmentType === 'local'" class="space-y-2">
                <button
                  type="button"
                  class="inline-flex items-center gap-2 rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
                  @click="handleUploadClick"
                >
                  <Upload class="h-4 w-4" />
                  上传附件
                </button>
                <div v-if="uploadedFiles.length" class="mt-2 space-y-2">
                  <div
                    v-for="(f, index) in uploadedFiles"
                    :key="index"
                    class="flex items-center justify-between rounded-lg border border-slate-200 bg-white px-4 py-2 dark:border-slate-700 dark:bg-slate-800"
                  >
                    <span class="text-sm text-slate-700 dark:text-slate-200">{{ f.name }}</span>
                    <span class="text-xs text-slate-500 dark:text-slate-400">
                      {{ (f.size / 1024).toFixed(1) }} KB
                    </span>
                    <button
                      type="button"
                      class="rounded p-1.5 text-slate-400 hover:bg-slate-100 hover:text-red-600 dark:hover:bg-slate-700 dark:hover:text-red-400"
                      title="删除"
                      @click="removeFile(index)"
                    >
                      <X class="h-4 w-4" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Flag 模版 -->
          <div class="rounded-xl border border-slate-200 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
            <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
              Flag 模版
            </h4>
            <textarea
              v-model="flagTemplate"
              rows="3"
              class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2 text-sm transition focus:border-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              placeholder="请输入 flag 模版字符串，留空以生成随机 GUID 作为 flag"
            />
            <div class="mt-4 rounded-lg border border-slate-200 bg-white p-4 text-xs text-slate-600 dark:border-slate-700 dark:bg-slate-800/50 dark:text-slate-400 space-y-2">
              <p class="font-medium text-slate-700 dark:text-slate-300">
                规则说明：
              </p>
              <ul class="list-disc list-inside space-y-1">
                <li>留空：将生成随机 GUID 作为 flag（如 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;1bab71b8-117f-4dea-a047-340b72101d7b&#125;</code>）。</li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[GUID]</code>：仅替换该占位符为随机 GUID。</li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[TEAM_HASH]</code>：替换为队伍 Token 与比赛信息生成的哈希值。</li>
                <li>未指定 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[TEAM_HASH]</code> 时启用 Leet 字符串（对花括号内内容变换），需确保模版熵足够高。</li>
                <li>指定 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[TEAM_HASH]</code> 且需 Leet：在模版<strong>开头</strong>加 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[LEET]</code>，此时不检查熵。</li>
                <li>启用特殊字符：用 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[CLEET]</code> 代替 <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[LEET]</code>（可能带来注入风险）。</li>
              </ul>
              <p class="font-medium text-slate-700 dark:text-slate-300 pt-2">
                示例：
              </p>
              <ul class="list-disc list-inside space-y-1">
                <li>留空 → <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;1bab71b8-117f-4dea-a047-340b72101d7b&#125;</code></li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;hello world&#125;</code> → <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;He1lo_w0r1d&#125;</code></li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[CLEET]flag&#123;hello sara&#125;</code> → <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;He1!o_$@rA&#125;</code></li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;hello_world_[TEAM_HASH]&#125;</code> → <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;hello_world_5418ce4d815c&#125;</code></li>
                <li><code class="rounded bg-slate-100 px-1 dark:bg-slate-700">[LEET]flag&#123;hello world [TEAM_HASH]&#125;</code> → <code class="rounded bg-slate-100 px-1 dark:bg-slate-700">flag&#123;He1lo_w0r1d_5418ce4d815c&#125;</code></li>
              </ul>
            </div>
          </div>
        </div>
      </template>

      <!-- 其他题目类型占位 -->
      <template v-else>
        <p class="text-slate-500 dark:text-slate-400">
          该题目类型暂无编辑附件及 Flag 配置。
        </p>
      </template>
    </section>
  </div>
</template>
