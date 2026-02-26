<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ArrowLeft, Settings, Megaphone, ListChecks, Activity, ChevronDown, ChevronRight, UserCheck, FileText } from 'lucide-vue-next'
import { contestStore } from '../../../stores/contestStore'

const route = useRoute()
const router = useRouter()

const contestId = computed(() => route.params.id as string)
const contest = computed(() => contestStore.contests.find((c) => String(c.id) === contestId.value))

const sectionsBeforeMonitor = [
  { id: 'info', label: '信息编辑', icon: Settings, path: 'info' },
  { id: 'announcement', label: '比赛公告', icon: Megaphone, path: 'announcement' },
  { id: 'challenges', label: '题目管理', icon: ListChecks, path: 'challenges' },
] as const

const monitorSubSections = [
  { path: 'submissions', label: 'Flag 提交记录' },
  { path: 'containers', label: '容器使用情况' },
  { path: 'same-flag', label: '作弊信息' },
] as const

const sectionsAfterMonitor = [
  { id: 'teamReview', label: '队伍审核', icon: UserCheck, path: 'team-review' },
  { id: 'writeups', label: 'Writeups', icon: FileText, path: 'writeups' },
] as const

const currentSectionPath = computed(() => {
  const p = route.path
  const base = `/admin/manage/contest/${contestId.value}`
  if (p === base || p === base + '/') return 'info'
  const rest = p.slice(base.length).replace(/^\//, '')
  const first = rest.split('/')[0]
  return first || 'info'
})

/** 当前是否在比赛监控的某个子页 */
const isMonitorSection = computed(() => currentSectionPath.value === 'monitor')

/** 当前监控子路径：submissions | containers | same-flag */
const currentMonitorSubPath = computed(() => {
  if (!isMonitorSection.value) return ''
  const p = route.path
  const base = `/admin/manage/contest/${contestId.value}/monitor`
  if (p === base || p === base + '/') return 'submissions'
  const rest = p.slice(base.length).replace(/^\//, '')
  return rest.split('/')[0] || 'submissions'
})

/** 比赛监控折叠展开：可手动切换，进入监控子页时自动展开 */
const monitorExpanded = ref(false)
watch(isMonitorSection, (on) => {
  if (on) monitorExpanded.value = true
}, { immediate: true })

function sectionTo(sectionPath: string) {
  return { path: `/admin/manage/contest/${contestId.value}/${sectionPath}` }
}

function goBack() {
  if (window.history.length > 1) {
    router.back()
  } else {
    router.push({ path: '/admin/manage/contest' })
  }
}
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between border-b border-slate-200 pb-4 dark:border-slate-800">
      <div class="flex items-center gap-4">
        <button
          type="button"
          class="inline-flex items-center gap-2 rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
          @click="goBack"
        >
          <ArrowLeft class="h-4 w-4" />
          返回
        </button>
        <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
          {{ contest?.name ?? '比赛管理' }}
        </h2>
      </div>
    </div>

    <div v-if="contest" class="flex gap-6">
      <aside class="w-56 shrink-0">
        <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 overflow-hidden">
          <nav class="p-1">
            <RouterLink
              v-for="s in sectionsBeforeMonitor"
              :key="s.id"
              :to="sectionTo(s.path)"
              class="flex w-full items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-sm font-medium transition-colors"
              :class="currentSectionPath === s.path
                ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                : 'text-slate-600 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800'"
            >
              <component :is="s.icon" :size="18" />
              {{ s.label }}
            </RouterLink>

            <!-- 比赛监控（点击整条展开/折叠，折叠图标在「比赛监控」字样右侧） -->
            <div class="py-0.5">
              <button
                type="button"
                class="flex w-full items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-sm font-medium transition-colors hover:bg-slate-50 dark:hover:bg-slate-800"
                :class="isMonitorSection
                  ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                  : 'text-slate-600 dark:text-slate-400'"
                @click="monitorExpanded = !monitorExpanded"
              >
                <Activity :size="18" class="shrink-0" />
                <span class="min-w-0 flex-1 truncate">比赛监控</span>
                <component :is="monitorExpanded ? ChevronDown : ChevronRight" class="h-4 w-4 shrink-0" />
              </button>
              <template v-if="monitorExpanded">
                <RouterLink
                  v-for="sub in monitorSubSections"
                  :key="sub.path"
                  :to="sectionTo('monitor/' + sub.path)"
                  class="flex w-full items-center gap-2.5 rounded-lg pl-10 pr-3 py-2 text-left text-sm font-medium transition-colors"
                  :class="currentMonitorSubPath === sub.path
                    ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                    : 'text-slate-600 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800'"
                >
                  {{ sub.label }}
                </RouterLink>
              </template>
            </div>

            <RouterLink
              v-for="s in sectionsAfterMonitor"
              :key="s.id"
              :to="sectionTo(s.path)"
              class="flex w-full items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-sm font-medium transition-colors"
              :class="currentSectionPath === s.path
                ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                : 'text-slate-600 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800'"
            >
              <component :is="s.icon" :size="18" />
              {{ s.label }}
            </RouterLink>
          </nav>
        </div>
      </aside>

      <main class="min-w-0 flex-1">
        <router-view />
      </main>
    </div>

    <div
      v-else
      class="rounded-xl border border-dashed border-slate-300 bg-white/60 py-12 text-center text-sm text-slate-500 dark:border-slate-700 dark:bg-slate-900/70 dark:text-slate-400"
    >
      未找到该比赛
    </div>
  </div>
</template>
