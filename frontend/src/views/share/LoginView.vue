<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Terminal, User, Lock } from 'lucide-vue-next'

const router = useRouter()
const username = ref('')
const password = ref('')

function handleSubmit() {
  const u = username.value.trim()
  const p = password.value
  if (!u) {
    alert('请输入用户名或学工号')
    return
  }
  if (!p) {
    alert('请输入密码')
    return
  }
  // 模拟登录成功，跳转首页
  router.push('/')
}
</script>

<template>
  <div class="flex min-h-screen items-center justify-center bg-slate-100 dark:bg-slate-950 px-4">
    <div class="w-full max-w-sm rounded-2xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-800 dark:bg-slate-900">
      <div class="mb-8 flex items-center justify-center gap-2">
        <div class="flex h-12 w-12 items-center justify-center rounded-xl bg-blue-600 text-white">
          <Terminal class="h-6 w-6" />
        </div>
        <h1 class="text-xl font-bold text-slate-900 dark:text-slate-100">
          登录
        </h1>
      </div>

      <form class="space-y-5" @submit.prevent="handleSubmit">
        <div>
          <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
            用户名 / 学工号
          </label>
          <div class="relative">
            <User class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
            <input
              v-model="username"
              type="text"
              placeholder="请输入用户名或学工号"
              autocomplete="username"
              class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
            />
          </div>
        </div>

        <div>
          <div class="mb-1.5 flex items-center justify-between">
            <label class="text-sm font-medium text-slate-700 dark:text-slate-300">
              密码
            </label>
            <router-link
              to="/reset-password"
              class="text-xs text-blue-600 hover:underline dark:text-blue-400"
            >
              忘记密码？
            </router-link>
          </div>
          <div class="relative">
            <Lock class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
            <input
              v-model="password"
              type="password"
              placeholder="请输入密码"
              autocomplete="current-password"
              class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
            />
          </div>
        </div>

        <button
          type="submit"
          class="w-full rounded-lg bg-blue-600 py-3 text-sm font-bold text-white transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900"
        >
          登录
        </button>

        <div class="relative flex items-center py-2">
          <div class="flex-1 border-t border-slate-200 dark:border-slate-700" />
          <span class="px-3 text-xs text-slate-400 dark:text-slate-500">或</span>
          <div class="flex-1 border-t border-slate-200 dark:border-slate-700" />
        </div>

        <button
          type="button"
          class="w-full rounded-lg border-2 border-slate-200 py-3 text-sm font-bold text-slate-700 transition-colors hover:border-slate-300 hover:bg-slate-50 dark:border-slate-700 dark:text-slate-300 dark:hover:border-slate-600 dark:hover:bg-slate-800/50"
        >
          统一身份认证登录
        </button>
      </form>

      <p class="mt-6 text-center text-xs text-slate-500 dark:text-slate-400">
        <router-link to="/" class="text-blue-600 hover:underline dark:text-blue-400">
          返回首页
        </router-link>
        <span class="mx-2">·</span>
        <router-link to="/register" class="text-blue-600 hover:underline dark:text-blue-400">
          注册
        </router-link>
      </p>
    </div>
  </div>
</template>
