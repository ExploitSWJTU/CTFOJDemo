<script setup lang="ts">
import { ref, onBeforeUnmount } from 'vue'
import { Settings, Globe, ToggleLeft, Trophy, Cloud, UserPlus, Shield, Users, FileCheck, Upload } from 'lucide-vue-next'

type SectionId =
  | 'basic'
  | 'features'
  | 'contest'
  | 'thirdparty'
  | 'register'
  | 'session'
  | 'roles'
  | 'privacy'

const platformSections: { id: SectionId; label: string; icon: typeof Globe }[] = [
  { id: 'basic', label: '基本信息', icon: Globe },
  { id: 'features', label: '功能开关', icon: ToggleLeft },
  { id: 'contest', label: '赛事配置', icon: Trophy },
  { id: 'thirdparty', label: '第三方服务', icon: Cloud },
]

const accountSections: { id: SectionId; label: string; icon: typeof UserPlus }[] = [
  { id: 'register', label: '注册与密码', icon: UserPlus },
  { id: 'session', label: '会话安全', icon: Shield },
  { id: 'roles', label: '角色权限', icon: Users },
  { id: 'privacy', label: '隐私合规', icon: FileCheck },
]

const currentSection = ref<SectionId>('basic')

// 表单状态（示例，可接后端）
const basic = ref({
  platformName: 'SWJTU CTF OJ',
  themeColor: '#2563eb',
  logoUrl: 'https://api.dicebear.com/7.x/shapes/svg?seed=ctf',
  siteDesc: '西南交通大学 CTF 在线判题平台',
  contactEmail: 'ctf@swjtu.edu.cn',
})
const features = ref({
  allowRegister: true,
  allowGuestView: false,
  dailyCheckIn: true,
  forumEnabled: true,
})
const contest = ref({ maxTeamSize: 4, allowLateSubmit: false, scoreboardFreeze: 30 })
const thirdparty = ref({ ssoEnabled: false, ssoProvider: '', ossEndpoint: '' })
const register = ref({ minPasswordLen: 8, requireEmailVerify: true, inviteOnly: false })
const session = ref({ sessionTimeout: 24, singleSession: false, ipBinding: false })
const roles = ref({ defaultRole: 'user', adminApproval: false })
const privacy = ref({ showRank: true, showRealName: false, gdprExport: true })

// Logo 本地上传
const logoInputRef = ref<HTMLInputElement | null>(null)
const logoFile = ref<File | null>(null)
const logoPreviewUrl = ref<string>('')

function onLogoFileChange(e: Event) {
  const input = e.target as HTMLInputElement
  const file = input.files?.[0]
  if (!file) return
  if (!file.type.startsWith('image/')) {
    alert('请选择图片文件（如 PNG、JPG、SVG）')
    input.value = ''
    return
  }
  if (logoPreviewUrl.value) URL.revokeObjectURL(logoPreviewUrl.value)
  logoFile.value = file
  logoPreviewUrl.value = URL.createObjectURL(file)
  input.value = ''
}

function clearLogoFile() {
  if (logoPreviewUrl.value) URL.revokeObjectURL(logoPreviewUrl.value)
  logoFile.value = null
  logoPreviewUrl.value = ''
}

onBeforeUnmount(() => {
  if (logoPreviewUrl.value) URL.revokeObjectURL(logoPreviewUrl.value)
})

function saveSection() {
  // TODO: 若有 logoFile 则先上传获取 URL，再与 basic 一并提交
  alert('保存成功')
}
</script>

<template>
  <div class="flex flex-col gap-6">
    <div class="flex items-center justify-between border-b border-slate-200 pb-4 dark:border-slate-800">
      <div class="flex items-center gap-3">
        <Settings class="text-blue-600" :size="28" />
        <div>
          <h2 class="text-xl font-bold text-slate-900 dark:text-slate-50">
            系统设置
          </h2>
          <p class="mt-0.5 text-xs text-slate-500 dark:text-slate-400">
            平台设置与账户策略
          </p>
        </div>
      </div>
    </div>

    <div class="flex gap-6">
      <!-- 左侧导航 -->
      <aside class="w-56 shrink-0 space-y-4">
        <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 overflow-hidden">
          <div class="bg-slate-50 dark:bg-slate-800/60 px-4 py-2.5">
            <span class="text-xs font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400">平台设置</span>
          </div>
          <nav class="p-1">
            <button
              v-for="s in platformSections"
              :key="s.id"
              type="button"
              class="flex w-full items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-sm font-medium transition-colors"
              :class="currentSection === s.id
                ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                : 'text-slate-600 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800'"
              @click="currentSection = s.id"
            >
              <component :is="s.icon" :size="18" />
              {{ s.label }}
            </button>
          </nav>
        </div>
        <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 overflow-hidden">
          <div class="bg-slate-50 dark:bg-slate-800/60 px-4 py-2.5">
            <span class="text-xs font-bold uppercase tracking-wider text-slate-500 dark:text-slate-400">账户策略</span>
          </div>
          <nav class="p-1">
            <button
              v-for="s in accountSections"
              :key="s.id"
              type="button"
              class="flex w-full items-center gap-2.5 rounded-lg px-3 py-2.5 text-left text-sm font-medium transition-colors"
              :class="currentSection === s.id
                ? 'bg-blue-50 text-blue-700 dark:bg-blue-900/30 dark:text-blue-300'
                : 'text-slate-600 hover:bg-slate-50 dark:text-slate-400 dark:hover:bg-slate-800'"
              @click="currentSection = s.id"
            >
              <component :is="s.icon" :size="18" />
              {{ s.label }}
            </button>
          </nav>
        </div>
      </aside>

      <!-- 右侧内容 -->
      <main class="min-w-0 flex-1">
        <div class="rounded-xl border border-slate-200 bg-white dark:border-slate-800 dark:bg-slate-900 shadow-sm overflow-hidden">
          <!-- 基本信息 -->
          <section v-show="currentSection === 'basic'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              基本信息
            </h3>
            <div class="space-y-4 max-w-md">
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">平台名</label>
                <input
                  v-model="basic.platformName"
                  type="text"
                  placeholder="如：SWJTU CTF OJ"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">主题色</label>
                <div class="flex items-center gap-3">
                  <input v-model="basic.themeColor" type="color" class="h-10 w-14 cursor-pointer rounded-lg border border-slate-200 dark:border-slate-700 bg-white dark:bg-slate-800 p-1" />
                  <input
                    v-model="basic.themeColor"
                    type="text"
                    class="flex-1 rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 font-mono text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                    placeholder="#2563eb"
                    maxlength="7"
                  />
                </div>
                <p class="mt-1 text-xs text-slate-500 dark:text-slate-400">
                  用于导航、按钮等主色
                </p>
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">Logo</label>
                <div class="flex flex-wrap items-start gap-4">
                  <div class="h-14 w-14 shrink-0 overflow-hidden rounded-xl border-2 border-slate-200 dark:border-slate-700 bg-slate-100 dark:bg-slate-800">
                    <img
                      v-if="logoPreviewUrl || basic.logoUrl"
                      :src="logoPreviewUrl || basic.logoUrl"
                      alt="Logo"
                      class="h-full w-full object-contain"
                      @error="($event.target as HTMLImageElement).style.display = 'none'"
                    />
                    <div v-else class="flex h-full w-full items-center justify-center text-slate-400 text-xs">
                      暂无
                    </div>
                  </div>
                  <div class="min-w-0 flex-1 space-y-2">
                    <input
                      v-model="basic.logoUrl"
                      type="url"
                      placeholder="Logo 图片地址（或本地上传）"
                      class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                      :disabled="!!logoFile"
                    />
                    <div class="flex items-center gap-2">
                      <input
                        ref="logoInputRef"
                        type="file"
                        accept="image/*"
                        class="hidden"
                        @change="onLogoFileChange"
                      />
                      <button type="button" class="inline-flex items-center gap-1.5 rounded-lg border border-slate-200 bg-white px-3 py-1.5 text-sm text-slate-600 hover:bg-slate-50 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 dark:hover:bg-slate-700" @click="logoInputRef?.click()">
                        <Upload :size="16" />
                        本地上传
                      </button>
                      <span v-if="logoFile" class="text-xs text-slate-500 dark:text-slate-400">{{ logoFile.name }}</span>
                      <button
                        v-if="logoFile"
                        type="button"
                        class="text-xs text-slate-500 underline hover:text-slate-700 dark:hover:text-slate-300"
                        @click="clearLogoFile"
                      >
                        清除
                      </button>
                    </div>
                  </div>
                </div>
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">站点描述</label>
                <input v-model="basic.siteDesc" type="text" class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200" />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">联系邮箱</label>
                <input v-model="basic.contactEmail" type="email" class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200" />
              </div>
            </div>
          </section>

          <!-- 功能开关 -->
          <section v-show="currentSection === 'features'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              功能开关
            </h3>
            <div class="space-y-4 max-w-md">
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">开放注册</span>
                <input v-model="features.allowRegister" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">游客可见题目列表</span>
                <input v-model="features.allowGuestView" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">每日签到</span>
                <input v-model="features.dailyCheckIn" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">论坛功能</span>
                <input v-model="features.forumEnabled" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
            </div>
          </section>

          <!-- 赛事配置 -->
          <section v-show="currentSection === 'contest'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              赛事配置
            </h3>
            <div class="space-y-4 max-w-md">
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">队伍最大人数</label>
                <input
                  v-model.number="contest.maxTeamSize"
                  type="number"
                  min="1"
                  max="20"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">允许迟交</span>
                <input v-model="contest.allowLateSubmit" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">榜单冻结时间（分钟）</label>
                <input
                  v-model.number="contest.scoreboardFreeze"
                  type="number"
                  min="0"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
            </div>
          </section>

          <!-- 第三方服务 -->
          <section v-show="currentSection === 'thirdparty'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              第三方服务
            </h3>
            <div class="space-y-4 max-w-md">
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">启用 SSO 登录</span>
                <input v-model="thirdparty.ssoEnabled" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <div v-if="thirdparty.ssoEnabled">
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">SSO 提供商</label>
                <input
                  v-model="thirdparty.ssoProvider"
                  type="text"
                  placeholder="如：CAS / OAuth2"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">对象存储端点（可选）</label>
                <input
                  v-model="thirdparty.ossEndpoint"
                  type="text"
                  placeholder="https://..."
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
            </div>
          </section>

          <!-- 注册与密码 -->
          <section v-show="currentSection === 'register'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              注册与密码
            </h3>
            <div class="space-y-4 max-w-md">
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">密码最小长度</label>
                <input
                  v-model.number="register.minPasswordLen"
                  type="number"
                  min="6"
                  max="32"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">注册需邮箱验证</span>
                <input v-model="register.requireEmailVerify" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">仅邀请注册</span>
                <input v-model="register.inviteOnly" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
            </div>
          </section>

          <!-- 会话安全 -->
          <section v-show="currentSection === 'session'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              会话安全
            </h3>
            <div class="space-y-4 max-w-md">
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">会话超时（小时）</label>
                <input
                  v-model.number="session.sessionTimeout"
                  type="number"
                  min="1"
                  max="720"
                  class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200"
                />
              </div>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">单设备登录（踢出其他会话）</span>
                <input v-model="session.singleSession" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">IP 绑定（同 IP 才有效）</span>
                <input v-model="session.ipBinding" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
            </div>
          </section>

          <!-- 角色权限 -->
          <section v-show="currentSection === 'roles'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              角色权限
            </h3>
            <div class="space-y-4 max-w-md">
              <div>
                <label class="block text-sm font-medium text-slate-600 dark:text-slate-400 mb-1">新用户默认角色</label>
                <select v-model="roles.defaultRole" class="w-full rounded-lg border border-slate-200 bg-slate-50 px-3 py-2 text-sm dark:border-slate-700 dark:bg-slate-800 dark:text-slate-200">
                  <option value="user">
                    用户
                  </option>
                  <option value="guest">
                    访客
                  </option>
                </select>
              </div>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">管理员需审核后生效</span>
                <input v-model="roles.adminApproval" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
            </div>
          </section>

          <!-- 隐私合规 -->
          <section v-show="currentSection === 'privacy'" class="p-6">
            <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-4">
              隐私合规
            </h3>
            <div class="space-y-4 max-w-md">
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">榜单显示用户名</span>
                <input v-model="privacy.showRank" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">榜单显示真实姓名</span>
                <input v-model="privacy.showRealName" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
              <label class="flex items-center justify-between rounded-lg border border-slate-200 dark:border-slate-700 p-3 cursor-pointer hover:bg-slate-50 dark:hover:bg-slate-800/50">
                <span class="text-sm font-medium text-slate-700 dark:text-slate-200">开放数据导出（GDPR 等）</span>
                <input v-model="privacy.gdprExport" type="checkbox" class="rounded border-slate-300 text-blue-600 focus:ring-blue-500" />
              </label>
            </div>
          </section>

          <!-- 底部保存 -->
          <div class="flex justify-end gap-2 border-t border-slate-200 dark:border-slate-800 px-6 py-4 bg-slate-50/50 dark:bg-slate-800/30">
            <button
              type="button"
              class="rounded-lg bg-blue-600 px-4 py-2 text-sm font-bold text-white hover:bg-blue-700"
              @click="saveSection"
            >
              保存当前设置
            </button>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>
