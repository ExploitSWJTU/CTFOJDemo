<script setup lang="ts">
import { ref } from 'vue'

const form = ref({
  teamSizeLimit: 0,
  teamContainerLimit: 0,
  teamRegisterNoReview: false,
  practiceMode: false,
  requireWriteup: false,
  writeupDeadlineHours: 72,
  writeupDescription: '',
})

function saveInfo() {
  alert('保存成功')
}
</script>

<template>
  <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
    <section class="p-6 md:p-8">
      <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-6">
        信息编辑
      </h3>

      <div class="space-y-8">
        <!-- 队伍与容器 -->
        <div class="rounded-xl border border-slate-100 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            队伍与容器
          </h4>
          <div class="grid gap-6 sm:grid-cols-2">
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                队伍人数限制 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.teamSizeLimit"
                type="number"
                min="0"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
              <p class="mt-1.5 text-xs text-slate-500 dark:text-slate-400">
                0 表示不限制队伍人数
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                队伍容器限制 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.teamContainerLimit"
                type="number"
                min="0"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
              <p class="mt-1.5 text-xs text-slate-500 dark:text-slate-400">
                队伍共享的容器数量 (0 表示不限制)
              </p>
            </div>
          </div>
        </div>

        <!-- 比赛选项 -->
        <div class="rounded-xl border border-slate-100 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            比赛选项
          </h4>
          <div class="space-y-4">
            <div class="flex items-center justify-between gap-4 rounded-xl border border-slate-200 bg-white px-5 py-4 shadow-sm dark:border-slate-700 dark:bg-slate-800/50">
              <div>
                <p class="text-sm font-medium text-slate-800 dark:text-slate-100">
                  队伍报名免审核
                </p>
              </div>
              <label class="relative inline-flex cursor-pointer items-center shrink-0">
                <input v-model="form.teamRegisterNoReview" type="checkbox" class="peer sr-only" />
                <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
              </label>
            </div>
            <div class="flex items-center justify-between gap-4 rounded-xl border border-slate-200 bg-white px-5 py-4 shadow-sm dark:border-slate-700 dark:bg-slate-800/50">
              <div>
                <p class="text-sm font-medium text-slate-800 dark:text-slate-100">
                  练习模式
                </p>
                <p class="mt-0.5 text-xs text-slate-500 dark:text-slate-400">
                  比赛结束后仍然可以查看题目和提交
                </p>
              </div>
              <label class="relative inline-flex cursor-pointer items-center shrink-0">
                <input v-model="form.practiceMode" type="checkbox" class="peer sr-only" />
                <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
              </label>
            </div>
            <div class="flex items-center justify-between gap-4 rounded-xl border border-slate-200 bg-white px-5 py-4 shadow-sm dark:border-slate-700 dark:bg-slate-800/50">
              <div>
                <p class="text-sm font-medium text-slate-800 dark:text-slate-100">
                  需要 Writeup
                </p>
                <p class="mt-0.5 text-xs text-slate-500 dark:text-slate-400">
                  是否启用比赛的 Writeup 提交功能
                </p>
              </div>
              <label class="relative inline-flex cursor-pointer items-center shrink-0">
                <input v-model="form.requireWriteup" type="checkbox" class="peer sr-only" />
                <div class="h-6 w-11 rounded-full bg-slate-200 after:absolute after:left-0.5 after:top-0.5 after:h-5 after:w-5 after:rounded-full after:border after:border-slate-300 after:bg-white after:transition-all peer-checked:bg-blue-600 peer-checked:after:translate-x-full peer-checked:after:border-white dark:bg-slate-700 dark:after:border-slate-600" />
              </label>
            </div>
          </div>
        </div>

        <!-- Writeup 设置 -->
        <div class="rounded-xl border border-slate-100 bg-slate-50/50 p-6 dark:border-slate-800 dark:bg-slate-800/30">
          <h4 class="text-sm font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400 mb-4">
            Writeup 设置
          </h4>
          <div class="grid gap-6 lg:grid-cols-1">
            <div class="max-w-sm">
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">
                Writeup 提交时限 <span class="text-red-500">*</span>
              </label>
              <input
                v-model.number="form.writeupDeadlineHours"
                type="number"
                min="0"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
              <p class="mt-1.5 text-xs text-slate-500 dark:text-slate-400">
                比赛结束后允许提交 Writeup 的小时数
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2">Writeup 提交说明</label>
              <textarea
                v-model="form.writeupDescription"
                rows="4"
                class="w-full rounded-lg border border-slate-200 bg-white px-4 py-2.5 text-sm shadow-sm transition focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                placeholder="可选：说明 Writeup 格式、要求等"
              />
              <p class="mt-1.5 text-xs text-slate-500 dark:text-slate-400">
                支持 Markdown 语法
              </p>
            </div>
          </div>
        </div>

        <div class="flex flex-wrap items-center gap-3 border-t border-slate-200 pt-6 dark:border-slate-700">
          <button
            type="button"
            class="rounded-lg bg-blue-600 px-5 py-2.5 text-sm font-medium text-white shadow-sm transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900"
            @click="saveInfo"
          >
            保存当前设置
          </button>
          <span class="text-xs text-slate-500 dark:text-slate-400">修改后将应用于本场比赛</span>
        </div>
      </div>
    </section>
  </div>
</template>
