<script setup lang="ts">
import { ref, computed } from 'vue';
import {
  Plus,
  Search,
  Database,
  Container,
  History,
  Trash2,
  Edit,
  Eye,
  EyeOff,
  Zap,
  Trophy,
  BarChart3,
  PieChart,
  Cpu,
  Layers,
} from 'lucide-vue-next';
import { challenges as mockChallenges } from '../../../mock/challenges';
import { CATEGORY_MAP, CATEGORIES } from '../../../constants/category';
import type { Challenge, Category } from '../../../types/challenge';

// 定义管理端特有的题目类型
interface AdminChallenge extends Challenge {
  isPublic: boolean;
  isDynamic: boolean;
  solveRate: number;
  activeContainers: number;
  totalSubmissions: number;
}

// 页签管理
const activeTab = ref('bank'); 

// --- 题库管理状态 ---
const searchQuery = ref('');
const selectedCategory = ref('All');
// 扩展 Mock 数据以匹配新需求
const challenges = ref<AdminChallenge[]>(mockChallenges.map(c => ({
  ...c,
  isPublic: Math.random() > 0.2,
  isDynamic: c.category === 'Web' || c.category === 'Pwn',
  solveRate: Math.floor(Math.random() * 100),
  activeContainers: c.category === 'Web' || c.category === 'Pwn' ? Math.floor(Math.random() * 15) : 0,
  totalSubmissions: Math.floor(c.solvedCount * (1.2 + Math.random())),
})));

// 分页逻辑
const currentPage = ref(1);
const pageSize = ref(10);

const filteredChallenges = computed(() => {
  return challenges.value.filter(c => {
    const matchSearch = c.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    const matchCat = selectedCategory.value === 'All' || c.category === selectedCategory.value;
    return matchSearch && matchCat;
  });
});

const totalPages = computed(() => Math.ceil(filteredChallenges.value.length / pageSize.value));
const paginatedChallenges = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value;
  return filteredChallenges.value.slice(start, start + pageSize.value);
});

// --- 全站统计数据 ---
const totalSolves = computed(() => challenges.value.reduce((acc, c) => acc + c.solvedCount, 0));
const avgSolveRate = computed(() => {
  const sum = challenges.value.reduce((acc, c) => acc + c.solveRate, 0);
  return Math.round(sum / challenges.value.length);
});

// 模拟活跃容器数据
const activeInstances = ref([
  { id: 1, challenge: 'Easy_Heap_Overflow', user: 'rikka', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=rikka', ip: '10.10.10.101', port: 12345, cpu: 12, ram: 24 },
  { id: 2, challenge: 'SQL_Injection_Advanced', user: 'admin', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin', ip: '10.10.10.102', port: 54321, cpu: 45, ram: 62 },
]);

// 编辑 Modal 状态
const editModalVisible = ref(false);
const editingChallenge = ref<AdminChallenge | null>(null);

const getCategoryMeta = (label: string) => CATEGORY_MAP[label as Category] || CATEGORY_MAP['Misc'];

// 操作函数
const toggleStatus = (id: number) => {
  const c = challenges.value.find(item => item.id === id);
  if (c) c.isPublic = !c.isPublic;
};

const openEditModal = (challenge: AdminChallenge) => {
  editingChallenge.value = { ...challenge };
  editModalVisible.value = true;
};
</script>

<template>
  <div class="flex flex-col gap-6">
    <!-- 顶部导航 -->
    <div class="flex items-center justify-between border-b border-slate-200 dark:border-slate-800 pb-4">
      <div class="flex items-center gap-6">
        <h2 class="text-xl font-black text-slate-900 dark:text-white flex items-center gap-2">
          <Database :size="24" class="text-blue-600" />
          训练运维中心
        </h2>
        <nav class="flex items-center gap-1 bg-slate-100 dark:bg-slate-800/50 p-1 rounded-lg">
          <button 
            v-for="tab in [
              { id: 'bank', label: '题库管理', icon: Database },
              { id: 'instances', label: '实例监控', icon: Container },
              { id: 'logs', label: '审计审计', icon: History }
            ]"
            :key="tab.id"
            class="flex items-center gap-2 px-3 py-1.5 text-xs font-bold rounded-md transition-all"
            :class="activeTab === tab.id ? 'bg-white dark:bg-slate-700 text-blue-600 shadow-sm' : 'text-slate-500 hover:text-slate-700 dark:hover:text-slate-300'"
            @click="activeTab = tab.id"
          >
            <component :is="tab.icon" :size="14" />
            {{ tab.label }}
          </button>
        </nav>
      </div>
      <button class="flex items-center gap-2 px-3 py-1.5 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-xs font-bold transition-all shadow-lg shadow-blue-200 dark:shadow-none">
        <Plus :size="14" /> 发布新题目
      </button>
    </div>

    <!-- 1. 题库管理 (双栏布局) -->
    <div v-if="activeTab === 'bank'" class="flex gap-6">
      <!-- 左侧：列表与分页 -->
      <div class="flex-1 min-w-0 space-y-4">
        <div class="flex items-center gap-3">
          <div class="relative flex-1 max-w-sm">
            <Search class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400" :size="14" />
            <input 
              v-model="searchQuery"
              type="text" 
              placeholder="搜索题目名称..." 
              class="w-full h-9 pl-9 pr-4 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs outline-none focus:border-blue-500 transition-all"
            />
          </div>
          <select 
            v-model="selectedCategory"
            class="h-9 px-2 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs font-bold outline-none"
          >
            <option value="All">
              全部分类
            </option>
            <option v-for="(_cfg, label) in CATEGORY_MAP" :key="label" :value="label">
              {{ label }}
            </option>
          </select>
        </div>

        <div class="border border-slate-200 dark:border-slate-800 rounded-xl overflow-hidden bg-white dark:bg-slate-950 shadow-sm">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="bg-slate-50 dark:bg-slate-900 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-200 dark:border-slate-800">
                <th class="px-4 py-3">
                  题目详情
                </th>
                <th class="px-4 py-3 text-center">
                  状态控制
                </th>
                <th class="px-4 py-3 text-center">
                  解出 / 提交
                </th>
                <th class="px-4 py-3 text-right">
                  快速操作
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-100 dark:divide-slate-800 text-xs">
              <tr v-for="c in paginatedChallenges" :key="c.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-900/30 transition-colors group">
                <td class="px-4 py-3">
                  <div class="flex items-center gap-4">
                    <!-- 分类图标替代标号 -->
                    <div 
                      class="w-10 h-10 rounded-xl flex items-center justify-center border shadow-sm transition-transform group-hover:scale-110"
                      :class="getCategoryMeta(c.category).cardClass"
                    >
                      <component :is="getCategoryMeta(c.category).icon" :size="20" />
                    </div>
                    <div class="flex flex-col min-w-0">
                      <div class="flex items-center gap-2 mb-0.5">
                        <span class="font-bold text-slate-800 dark:text-slate-100 truncate text-sm">{{ c.title }}</span>
                        <span v-if="c.isDynamic" class="flex items-center gap-0.5 px-1.5 py-0.5 rounded bg-blue-50 dark:bg-blue-900/30 text-blue-600 dark:text-blue-400 text-[8px] font-black uppercase border border-blue-100 dark:border-blue-800">
                          <Zap :size="8" /> Dynamic
                        </span>
                      </div>
                      <div class="flex items-center gap-3 text-[10px] font-bold text-slate-400 uppercase tracking-tight">
                        <span class="font-mono text-blue-600 dark:text-blue-400 text-sm font-black">{{ c.points }} Pts</span>
                        <span>•</span>
                        <span class="flex items-center gap-1">
                          <Container :size="10" /> {{ c.activeContainers }} 运行中
                        </span>
                      </div>
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3 text-center">
                  <button class="px-2.5 py-1 rounded-full font-black text-[9px] uppercase border transition-all shadow-sm" :class="c.isPublic ? 'bg-emerald-50 text-emerald-600 border-emerald-100 dark:bg-emerald-900/20 dark:border-emerald-800' : 'bg-slate-50 text-slate-400 border-slate-200 dark:bg-slate-800 dark:border-slate-700'" @click="toggleStatus(c.id)">
                    <component :is="c.isPublic ? Eye : EyeOff" :size="10" class="inline mr-1" />
                    {{ c.isPublic ? 'Public' : 'Hidden' }}
                  </button>
                </td>
                <td class="px-4 py-3 text-center">
                  <div class="flex flex-col items-center gap-1">
                    <span class="font-mono font-bold text-slate-700 dark:text-slate-200 text-sm">{{ c.solvedCount }} / {{ c.totalSubmissions }}</span>
                    <div class="w-16 h-1 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                      <div class="h-full bg-emerald-500" :style="{ width: `${c.solveRate}%` }" />
                    </div>
                  </div>
                </td>
                <td class="px-4 py-3 text-right">
                  <div class="flex items-center justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                    <button class="p-1.5 text-slate-400 hover:text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/20 rounded-md" @click="openEditModal(c)">
                      <Edit :size="14" />
                    </button>
                    <button class="p-1.5 text-slate-400 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-md">
                      <Trash2 :size="14" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- 分页控制器 -->
        <div class="flex items-center justify-between px-2 pt-2">
          <span class="text-[10px] font-bold text-slate-400 uppercase">Page {{ currentPage }} of {{ totalPages }}</span>
          <div class="flex gap-1">
            <button 
              :disabled="currentPage === 1" 
              class="px-3 py-1 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs font-bold disabled:opacity-30"
              @click="currentPage = Math.max(1, currentPage - 1)"
            >
              Prev
            </button>
            <button 
              :disabled="currentPage === totalPages" 
              class="px-3 py-1 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-lg text-xs font-bold disabled:opacity-30"
              @click="currentPage = Math.min(totalPages, currentPage + 1)"
            >
              Next
            </button>
          </div>
        </div>
      </div>

      <!-- 右侧：全站总览统计 -->
      <aside class="w-72 shrink-0 space-y-4">
        <div class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl p-5 shadow-sm">
          <div class="flex items-center gap-2 mb-6">
            <BarChart3 :size="18" class="text-blue-600" />
            <h3 class="text-sm font-black text-slate-800 dark:text-slate-100 uppercase tracking-wider">
              全站总览
            </h3>
          </div>

          <div class="space-y-6">
            <!-- 核心指标 -->
            <div class="grid grid-cols-2 gap-4">
              <div class="flex flex-col">
                <span class="text-[9px] font-bold text-slate-400 uppercase tracking-tighter mb-1 flex items-center gap-1"><Database :size="10" /> 总题目数</span>
                <span class="text-xl font-black text-slate-800 dark:text-white">{{ challenges.length }}</span>
              </div>
              <div class="flex flex-col">
                <span class="text-[9px] font-bold text-slate-400 uppercase tracking-tighter mb-1 flex items-center gap-1"><Trophy :size="10" /> 总解出数</span>
                <span class="text-xl font-black text-emerald-500">{{ totalSolves }}</span>
              </div>
            </div>

            <div class="h-px bg-slate-100 dark:bg-slate-800" />

            <!-- 统计详情 -->
            <div class="space-y-4">
              <div class="flex flex-col gap-2">
                <div class="flex justify-between items-center text-[10px] font-bold uppercase">
                  <span class="text-slate-500">全站平均通过率</span>
                  <span class="text-blue-600">{{ avgSolveRate }}%</span>
                </div>
                <div class="h-1.5 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden">
                  <div class="h-full bg-blue-500 transition-all duration-1000" :style="{ width: `${avgSolveRate}%` }" />
                </div>
              </div>

              <div class="space-y-3">
                <span class="text-[10px] font-bold text-slate-400 uppercase block mb-3">全部分类分布</span>
                <div v-for="cat in CATEGORIES.filter(c => c.label !== 'All')" :key="cat.label" class="flex items-center justify-between group cursor-default">
                  <div class="flex items-center gap-2">
                    <div class="w-6 h-6 rounded flex items-center justify-center bg-slate-50 dark:bg-slate-900 border border-slate-100 dark:border-slate-800">
                      <component :is="cat.icon" :size="12" :class="cat.sidebar.inactive.replace('text-[', 'text-').replace(']', '')" />
                    </div>
                    <span class="text-[11px] font-bold text-slate-600 dark:text-slate-400 transition-colors group-hover:text-blue-600">{{ cat.label }}</span>
                  </div>
                  <div class="flex items-center gap-2">
                    <span class="text-[10px] font-mono font-bold text-slate-400">{{ challenges.filter(c => c.category === cat.label).length }}</span>
                    <div class="w-12 h-1 bg-slate-50 dark:bg-slate-900 rounded-full overflow-hidden">
                      <div class="h-full bg-slate-200 dark:bg-slate-800" :style="{ width: `${(challenges.filter(c => c.category === cat.label).length / challenges.length) * 100}%` }" />
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <button class="w-full py-2 bg-slate-50 dark:bg-slate-900 hover:bg-blue-50 dark:hover:bg-blue-900/20 text-[10px] font-black text-slate-400 hover:text-blue-600 uppercase tracking-widest rounded-lg border border-slate-100 dark:border-slate-800 transition-all">
              <PieChart :size="12" class="inline mr-1" /> 查看完整分析报告
            </button>
          </div>
        </div>
      </aside>
    </div>

    <!-- 2. 实例监控 (同步样式) -->
    <div v-if="activeTab === 'instances'" class="space-y-3">
      <div class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl overflow-hidden shadow-sm">
        <table class="w-full text-left border-collapse text-xs">
          <thead>
            <tr class="bg-slate-50 dark:bg-slate-900 text-[10px] font-black text-slate-400 uppercase tracking-widest border-b border-slate-200 dark:border-slate-800">
              <th class="px-4 py-3">
                User / Challenge
              </th>
              <th class="px-4 py-3">
                Endpoint
              </th>
              <th class="px-4 py-3">
                Resources
              </th>
              <th class="px-4 py-3 text-right">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100 dark:divide-slate-800">
            <tr v-for="inst in activeInstances" :key="inst.id" class="hover:bg-slate-50/50 dark:hover:bg-slate-900/30">
              <td class="px-4 py-3">
                <div class="flex items-center gap-3">
                  <img :src="inst.avatar" class="w-7 h-7 rounded-full border border-slate-200 dark:border-slate-700" />
                  <div>
                    <div class="text-xs font-bold text-slate-800 dark:text-slate-100">
                      @{{ inst.user }}
                    </div>
                    <div class="text-[10px] font-bold text-slate-400 uppercase tracking-tighter">
                      {{ inst.challenge }}
                    </div>
                  </div>
                </div>
              </td>
              <td class="px-4 py-3 font-mono text-blue-600 dark:text-blue-400 font-bold">
                {{ inst.ip }}:{{ inst.port }}
              </td>
              <td class="px-4 py-3">
                <div class="flex items-center gap-2">
                  <Cpu :size="12" class="text-slate-400" /> <span class="font-mono">{{ inst.cpu }}%</span>
                  <Layers :size="12" class="text-slate-400 ml-2" /> <span class="font-mono">{{ inst.ram }}M</span>
                </div>
              </td>
              <td class="px-4 py-3 text-right">
                <button class="px-2 py-1 bg-rose-50 dark:bg-rose-900/20 text-rose-600 text-[10px] font-bold rounded border border-rose-100 dark:border-rose-900/30">
                  Destroy
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- 3. 审计日志 (同步样式) -->
    <div v-if="activeTab === 'logs'" class="bg-white dark:bg-slate-950 border border-slate-200 dark:border-slate-800 rounded-xl overflow-hidden shadow-sm">
      <div class="divide-y divide-slate-100 dark:divide-slate-800">
        <div v-for="log in 15" :key="log" class="px-4 py-2.5 flex items-center justify-between hover:bg-slate-50/50 dark:hover:bg-slate-900/30 transition-all text-xs">
          <div class="flex items-center gap-3">
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" class="w-6 h-6 rounded-full" />
            <span class="font-bold text-slate-700 dark:text-slate-200">@rikka 执行了 Flag 提交 "Easy_Heap"</span>
          </div>
          <div :class="['px-2 py-0.5 rounded text-[9px] font-black uppercase', log % 3 === 0 ? 'bg-rose-50 text-rose-500' : 'bg-emerald-50 text-emerald-500']">
            {{ log % 3 === 0 ? 'Failed' : 'Success' }}
          </div>
        </div>
      </div>
    </div>

    <!-- 编辑题目 Modal -->
    <a-modal
      v-model:visible="editModalVisible"
      width="800px"
      title="编辑题目详情"
      :footer="true"
      modal-class="rounded-card overflow-hidden"
    >
      <div v-if="editingChallenge" class="space-y-4 py-2">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase mb-1.5">题目名称</label>
            <input v-model="editingChallenge.title" class="w-full h-10 px-3 bg-app-bg dark:bg-slate-950 border border-app-border rounded-lg text-sm focus:border-blue-500 outline-none transition-all" />
          </div>
          <div>
            <label class="block text-xs font-bold text-slate-400 uppercase mb-1.5">分值</label>
            <input v-model="editingChallenge.points" type="number" class="w-full h-10 px-3 bg-app-bg dark:bg-slate-950 border border-app-border rounded-lg text-sm focus:border-blue-500 outline-none transition-all" />
          </div>
        </div>
        <div>
          <label class="block text-xs font-bold text-slate-400 uppercase mb-1.5">Markdown 详情内容</label>
          <textarea v-model="editingChallenge.description" rows="12" class="w-full p-4 bg-app-bg dark:bg-slate-950 border border-app-border rounded-lg text-sm font-mono focus:border-blue-500 outline-none transition-all leading-relaxed" />
        </div>
      </div>
    </a-modal>
  </div>
</template>