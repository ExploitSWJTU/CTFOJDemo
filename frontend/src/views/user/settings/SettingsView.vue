<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { User, Mail, Phone, Hash, UserCircle, FileText, ArrowLeft, Save, Camera, X } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()
const profile = ref({
  username: 'swjtunsa',
  email: 'swjtunsa@ctf.com',
  phone: '1',
  studentId: '1',
  realName: '1',
  description: '楼上是给',
})

const avatarUrl = ref('https://api.dicebear.com/7.x/avataaars/svg?seed=Settings')

// 更换邮箱弹窗
const showEmailModal = ref(false)
const newEmail = ref('')
const emailCode = ref('')
const emailSending = ref(false)

function openChangeEmail() {
  newEmail.value = ''
  emailCode.value = ''
  showEmailModal.value = true
}

function closeEmailModal() {
  showEmailModal.value = false
}

function sendEmailCode() {
  if (!newEmail.value.trim()) {
    alert('请先输入新邮箱')
    return
  }
  emailSending.value = true
  // TODO: 调用发送验证码接口
  setTimeout(() => {
    emailSending.value = false
    alert('验证码已发送到新邮箱')
  }, 800)
}

function submitChangeEmail() {
  const email = newEmail.value.trim()
  const code = emailCode.value.trim()
  if (!email) {
    alert('请输入新邮箱')
    return
  }
  if (!code) {
    alert('请输入验证码')
    return
  }
  // TODO: 调用更换邮箱接口
  profile.value.email = email
  closeEmailModal()
  alert('邮箱更换成功')
}

// 更换密码弹窗
const showPasswordModal = ref(false)
const oldPassword = ref('')
const newPassword = ref('')
const confirmPassword = ref('')

function openChangePassword() {
  oldPassword.value = ''
  newPassword.value = ''
  confirmPassword.value = ''
  showPasswordModal.value = true
}

function closePasswordModal() {
  showPasswordModal.value = false
}

function submitChangePassword() {
  if (!oldPassword.value) {
    alert('请输入原密码')
    return
  }
  if (!newPassword.value) {
    alert('请输入新密码')
    return
  }
  if (newPassword.value !== confirmPassword.value) {
    alert('两次输入的新密码不一致')
    return
  }
  // TODO: 调用更换密码接口
  closePasswordModal()
  alert('密码修改成功')
}

function handleBack() {
  if (route.path.startsWith('/admin')) {
    router.push('/admin/manage/dashboard')
  } else {
    router.push('/')
  }
}

function handleSave() {
  // TODO: 调用接口保存
  alert('保存成功')
}

function handleChangeEmail() {
  openChangeEmail()
}

function handleChangePassword() {
  openChangePassword()
}
</script>

<template>
  <div class="min-h-[60vh] space-y-8 pb-16">
    <!-- 顶部：返回 + 保存 -->
    <div class="flex flex-wrap items-center justify-between gap-4">
      <button
        type="button"
        class="rounded-button flex items-center gap-2 px-3 py-2 text-sm font-medium text-slate-600 transition-colors hover:bg-slate-100 hover:text-slate-900 dark:text-slate-400 dark:hover:bg-slate-800 dark:hover:text-slate-100"
        @click="handleBack"
      >
        <ArrowLeft :size="18" />
        返回
      </button>
      <button
        type="button"
        class="rounded-xl bg-blue-600 px-5 py-2.5 text-sm font-bold text-white shadow-md transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 dark:focus:ring-offset-slate-900"
        @click="handleSave"
      >
        <Save :size="16" class="mr-1.5 inline-block align-middle" />
        保存
      </button>
    </div>

    <section
      class="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-lg dark:border-slate-800 dark:bg-slate-900"
    >
      <!-- 头像区 -->
      <div
        class="border-b border-slate-100 bg-gradient-to-br from-slate-50 to-white px-8 py-8 dark:border-slate-800 dark:from-slate-900/50 dark:to-slate-900"
      >
        <div class="flex flex-wrap items-center gap-8">
          <div class="group relative shrink-0">
            <div
              class="flex h-28 w-28 overflow-hidden rounded-2xl border-2 border-slate-200 shadow-md ring-2 ring-white dark:border-slate-700 dark:ring-slate-800"
            >
              <img
                :src="avatarUrl"
                alt="头像"
                class="h-full w-full object-cover"
              />
            </div>
            <button
              type="button"
              class="absolute inset-0 flex items-center justify-center rounded-2xl bg-slate-900/50 opacity-0 transition-opacity group-hover:opacity-100"
              title="更换头像"
            >
              <Camera class="h-8 w-8 text-white" />
            </button>
          </div>
          <div class="min-w-0">
            <p class="text-sm font-medium text-slate-500 dark:text-slate-400">
              头像
            </p>
            <p class="mt-0.5 text-xs text-slate-400 dark:text-slate-500">
              点击头像区域可更换（功能待接后端）
            </p>
          </div>
        </div>
      </div>

      <!-- 个人资料表单 -->
      <div class="p-8">
        <h2 class="mb-6 flex items-center gap-2 text-lg font-semibold text-slate-800 dark:text-slate-100">
          <UserCircle :size="22" class="text-blue-500" />
          个人资料
        </h2>
        <div class="grid gap-6 sm:grid-cols-1 lg:grid-cols-2">
          <div class="space-y-2">
            <label
              for="username"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <User :size="16" />
              用户名
            </label>
            <input
              id="username"
              v-model="profile.username"
              type="text"
              placeholder="请输入用户名"
              class="input-field w-full rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/80 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400 dark:focus:bg-slate-800 dark:focus:ring-blue-400/20"
            />
          </div>
          <div class="space-y-2">
            <label
              for="email"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <Mail :size="16" />
              邮箱
            </label>
            <input
              id="email"
              :value="profile.email"
              type="email"
              disabled
              class="w-full cursor-not-allowed rounded-xl border border-slate-200 bg-slate-100 px-4 py-3 text-slate-500 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-400"
            />
            <p class="text-xs text-slate-400 dark:text-slate-500">
              通过底部「更改邮箱」可申请更换绑定邮箱
            </p>
          </div>
          <div class="space-y-2">
            <label
              for="phone"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <Phone :size="16" />
              手机号
            </label>
            <input
              id="phone"
              v-model="profile.phone"
              type="text"
              placeholder="请输入手机号"
              class="input-field w-full rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/80 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400 dark:focus:ring-blue-400/20"
            />
          </div>
          <div class="space-y-2">
            <label
              for="studentId"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <Hash :size="16" />
              学工号
            </label>
            <input
              id="studentId"
              v-model="profile.studentId"
              type="text"
              placeholder="请输入学工号"
              class="input-field w-full rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/80 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400 dark:focus:ring-blue-400/20"
            />
          </div>
          <div class="space-y-2 lg:col-span-2">
            <label
              for="realName"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <UserCircle :size="16" />
              真实姓名
            </label>
            <input
              id="realName"
              v-model="profile.realName"
              type="text"
              placeholder="请输入真实姓名"
              class="input-field w-full max-w-md rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/80 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400 dark:focus:ring-blue-400/20"
            />
          </div>
          <div class="space-y-2 lg:col-span-2">
            <label
              for="description"
              class="flex items-center gap-2 text-sm font-medium text-slate-600 dark:text-slate-400"
            >
              <FileText :size="16" />
              描述
            </label>
            <textarea
              id="description"
              v-model="profile.description"
              rows="3"
              placeholder="介绍一下自己"
              class="input-field w-full resize-y rounded-xl border border-slate-200 bg-slate-50/80 px-4 py-3 text-slate-800 outline-none transition placeholder:text-slate-400 focus:border-blue-500 focus:bg-white focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800/80 dark:text-slate-200 dark:placeholder:text-slate-500 dark:focus:border-blue-400 dark:focus:ring-blue-400/20"
            />
          </div>
        </div>

        <!-- 底部操作按钮：返回、更改邮箱、更改密码 -->
        <div class="mt-10 flex flex-wrap items-center gap-3 border-t border-slate-100 pt-8 dark:border-slate-800">
          <button
            type="button"
            class="rounded-xl border border-slate-200 bg-white px-5 py-2.5 text-sm font-medium text-slate-700 shadow-sm transition-colors hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200 dark:hover:bg-slate-700"
            @click="handleBack"
          >
            返回
          </button>
          <button
            type="button"
            class="rounded-xl border border-slate-200 bg-white px-5 py-2.5 text-sm font-medium text-slate-600 shadow-sm transition-colors hover:bg-slate-50 hover:text-slate-900 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700 dark:hover:text-slate-100"
            @click="handleChangeEmail"
          >
            更改邮箱
          </button>
          <button
            type="button"
            class="rounded-xl border border-slate-200 bg-white px-5 py-2.5 text-sm font-medium text-slate-600 shadow-sm transition-colors hover:bg-slate-50 hover:text-slate-900 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700 dark:hover:text-slate-100"
            @click="handleChangePassword"
          >
            更改密码
          </button>
        </div>
      </div>
    </section>

    <!-- 更换邮箱弹窗 -->
    <Teleport to="body">
      <div
        v-if="showEmailModal"
        class="fixed inset-0 z-[100] flex items-center justify-center p-4"
        @click.self="closeEmailModal"
      >
        <div
          class="absolute inset-0 bg-slate-900/60 backdrop-blur-sm"
          aria-hidden
        />
        <div
          class="relative w-full max-w-md rounded-2xl border border-slate-200 bg-white p-6 shadow-2xl dark:border-slate-700 dark:bg-slate-900"
          @click.stop
        >
          <div class="mb-4 flex items-center justify-between">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
              更换邮箱
            </h3>
            <button
              type="button"
              class="rounded-lg p-1.5 text-slate-400 hover:bg-slate-100 hover:text-slate-600 dark:hover:bg-slate-800 dark:hover:text-slate-300"
              @click="closeEmailModal"
            >
              <X :size="20" />
            </button>
          </div>
          <div class="space-y-4">
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-600 dark:text-slate-400">
                新邮箱 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="newEmail"
                type="email"
                placeholder="请输入新邮箱"
                class="w-full rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-slate-800 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-600 dark:text-slate-400">
                验证码 <span class="text-red-500">*</span>
              </label>
              <div class="flex gap-2">
                <input
                  v-model="emailCode"
                  type="text"
                  placeholder="请输入验证码"
                  class="min-w-0 flex-1 rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-slate-800 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
                <button
                  type="button"
                  class="shrink-0 rounded-xl border border-slate-200 bg-white px-4 py-2.5 text-sm font-medium text-slate-600 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700"
                  :disabled="emailSending"
                  @click="sendEmailCode"
                >
                  {{ emailSending ? '发送中...' : '获取验证码' }}
                </button>
              </div>
            </div>
          </div>
          <div class="mt-6 flex justify-end gap-3">
            <button
              type="button"
              class="rounded-xl border border-slate-200 px-4 py-2.5 text-sm font-medium text-slate-600 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:hover:bg-slate-700"
              @click="closeEmailModal"
            >
              取消
            </button>
            <button
              type="button"
              class="rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-bold text-white hover:bg-blue-700"
              @click="submitChangeEmail"
            >
              确定
            </button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- 更换密码弹窗 -->
    <Teleport to="body">
      <div
        v-if="showPasswordModal"
        class="fixed inset-0 z-[100] flex items-center justify-center p-4"
        @click.self="closePasswordModal"
      >
        <div
          class="absolute inset-0 bg-slate-900/60 backdrop-blur-sm"
          aria-hidden
        />
        <div
          class="relative w-full max-w-md rounded-2xl border border-slate-200 bg-white p-6 shadow-2xl dark:border-slate-700 dark:bg-slate-900"
          @click.stop
        >
          <div class="mb-4 flex items-center justify-between">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100">
              更换密码
            </h3>
            <button
              type="button"
              class="rounded-lg p-1.5 text-slate-400 hover:bg-slate-100 hover:text-slate-600 dark:hover:bg-slate-800 dark:hover:text-slate-300"
              @click="closePasswordModal"
            >
              <X :size="20" />
            </button>
          </div>
          <div class="space-y-4">
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-600 dark:text-slate-400">
                原密码 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="oldPassword"
                type="password"
                placeholder="P4ssW@rd"
                autocomplete="current-password"
                class="w-full rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-slate-800 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-600 dark:text-slate-400">
                密码 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="newPassword"
                type="password"
                placeholder="P4ssW@rd"
                autocomplete="new-password"
                class="w-full rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-slate-800 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
            <div>
              <label class="mb-1.5 block text-sm font-medium text-slate-600 dark:text-slate-400">
                确认密码 <span class="text-red-500">*</span>
              </label>
              <input
                v-model="confirmPassword"
                type="password"
                placeholder="P4ssW@rd"
                autocomplete="new-password"
                class="w-full rounded-xl border border-slate-200 bg-slate-50 px-4 py-3 text-slate-800 outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
              />
            </div>
          </div>
          <div class="mt-6 flex justify-end gap-3">
            <button
              type="button"
              class="rounded-xl border border-slate-200 px-4 py-2.5 text-sm font-medium text-slate-600 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:hover:bg-slate-700"
              @click="closePasswordModal"
            >
              取消
            </button>
            <button
              type="button"
              class="rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-bold text-white hover:bg-blue-700"
              @click="submitChangePassword"
            >
              确定
            </button>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>
