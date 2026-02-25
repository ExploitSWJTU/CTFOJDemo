<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Terminal, Mail, KeyRound, Lock } from 'lucide-vue-next'

const router = useRouter()
const step = ref(1)

const email = ref('')
const verifyCode = ref('')
const sending = ref(false)
const countdown = ref(0)

const newPassword = ref('')
const confirmPassword = ref('')

// 模拟验证码（实际应后端校验）
const MOCK_CODE = '123456'

function sendCode() {
  const e = email.value.trim()
  if (!e) {
    alert('请先输入邮箱地址')
    return
  }
  sending.value = true
  countdown.value = 60
  const timer = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(timer)
      sending.value = false
    }
  }, 1000)
  // 模拟发送
  alert('验证码已发送至您的邮箱（演示：验证码 123456）')
}

function nextStep() {
  const e = email.value.trim()
  const code = verifyCode.value.trim()
  if (!e) {
    alert('请输入邮箱地址')
    return
  }
  if (!code) {
    alert('请输入验证码')
    return
  }
  if (code !== MOCK_CODE) {
    alert('验证码错误')
    return
  }
  step.value = 2
}

function handleReset() {
  const p = newPassword.value
  const cp = confirmPassword.value
  if (!p) {
    alert('请输入新密码')
    return
  }
  if (p.length < 6) {
    alert('密码长度至少 6 位')
    return
  }
  if (p !== cp) {
    alert('两次输入的密码不一致')
    return
  }
  alert('密码重置成功，请登录')
  router.push('/login')
}
</script>

<template>
  <div class="flex min-h-screen items-center justify-center bg-slate-100 dark:bg-slate-950 px-4 py-8">
    <div class="w-full max-w-sm rounded-2xl border border-slate-200 bg-white p-8 shadow-xl dark:border-slate-800 dark:bg-slate-900">
      <div class="mb-6 flex items-center justify-center gap-2">
        <div class="flex h-12 w-12 items-center justify-center rounded-xl bg-blue-600 text-white">
          <Terminal class="h-6 w-6" />
        </div>
        <h1 class="text-xl font-bold text-slate-900 dark:text-slate-100">
          重置密码
        </h1>
      </div>

      <!-- 第一步：邮箱 + 验证码 -->
      <template v-if="step === 1">
        <div class="space-y-4">
          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
              邮箱
            </label>
            <div class="relative">
              <Mail class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
              <input
                v-model="email"
                type="email"
                placeholder="请输入邮箱地址"
                autocomplete="email"
                class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
              />
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
              验证码
            </label>
            <div class="flex gap-2">
              <div class="relative flex-1">
                <KeyRound class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
                <input
                  v-model="verifyCode"
                  type="text"
                  placeholder="请输入验证码"
                  maxlength="6"
                  class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
                />
              </div>
              <button
                type="button"
                :disabled="sending || countdown > 0"
                class="shrink-0 rounded-lg border border-slate-200 bg-slate-50 px-4 text-sm font-medium text-slate-700 transition-colors hover:bg-slate-100 disabled:opacity-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                @click="sendCode"
              >
                {{ countdown > 0 ? `${countdown}s 后重发` : '发送验证码' }}
              </button>
            </div>
          </div>

          <button
            type="button"
            class="w-full rounded-lg bg-blue-600 py-3 text-sm font-bold text-white transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900"
            @click="nextStep"
          >
            下一步
          </button>
        </div>
      </template>

      <!-- 第二步：新密码 + 确认密码 -->
      <template v-else>
        <div class="space-y-4">
          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
              新密码
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
              <input
                v-model="newPassword"
                type="password"
                placeholder="请输入新密码"
                autocomplete="new-password"
                class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
              />
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-sm font-medium text-slate-700 dark:text-slate-300">
              确认密码
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
              <input
                v-model="confirmPassword"
                type="password"
                placeholder="请再次输入新密码"
                autocomplete="new-password"
                class="h-11 w-full rounded-lg border border-slate-200 bg-slate-50 pl-10 pr-4 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:placeholder:text-slate-500"
              />
            </div>
          </div>

          <button
            type="button"
            class="w-full rounded-lg bg-blue-600 py-3 text-sm font-bold text-white transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900"
            @click="handleReset"
          >
            确认重置
          </button>

          <button
            type="button"
            class="w-full text-center text-sm text-slate-500 hover:text-slate-700 dark:hover:text-slate-300"
            @click="step = 1"
          >
            返回上一步
          </button>
        </div>
      </template>

      <p class="mt-6 text-center text-xs text-slate-500 dark:text-slate-400">
        <router-link to="/" class="text-blue-600 hover:underline dark:text-blue-400">
          返回首页
        </router-link>
        <span class="mx-2">·</span>
        <router-link to="/login" class="text-blue-600 hover:underline dark:text-blue-400">
          去登录
        </router-link>
      </p>
    </div>
  </div>
</template>
