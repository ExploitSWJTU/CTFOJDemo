<script setup lang="ts">
import { ref, computed } from 'vue';
import {
  Users,
  Zap,
  Server,
  TrendingUp,
  Clock,
  ShieldCheck,
  ArrowUpRight,
  ArrowDownRight,
  MoreVertical,
} from 'lucide-vue-next';
import VChart from 'vue-echarts';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { LineChart, BarChart } from 'echarts/charts';
import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
} from 'echarts/components';

// Register ECharts components
use([
  CanvasRenderer,
  LineChart,
  BarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
]);

// Mock Dashboard Stats
const stats = ref([
  {
    label: '总注册用户',
    value: '2,845',
    trend: '+12.5%',
    trendUp: true,
    icon: Users,
    color: 'text-blue-600',
    bg: 'bg-blue-50 dark:bg-blue-900/20',
  },
  {
    label: '今日活跃',
    value: '432',
    trend: '+5.2%',
    trendUp: true,
    icon: Zap,
    color: 'text-amber-600',
    bg: 'bg-amber-50 dark:bg-amber-900/20',
  },
  {
    label: '运行中容器',
    value: '86',
    trend: '-2.4%',
    trendUp: false,
    icon: Server,
    color: 'text-emerald-600',
    bg: 'bg-emerald-50 dark:bg-emerald-900/20',
  },
  {
    label: '系统安全评分',
    value: '98',
    trend: '稳定',
    trendUp: true,
    icon: ShieldCheck,
    color: 'text-purple-600',
    bg: 'bg-purple-50 dark:bg-purple-900/20',
  },
]);

// Mock Chart Data
const chartOption = computed(() => ({
  grid: {
    top: 20,
    bottom: 30,
    left: 40,
    right: 10,
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: { type: 'cross' },
  },
  xAxis: {
    type: 'category',
    data: ['00:00', '04:00', '08:00', '12:00', '16:00', '20:00', '23:59'],
    axisLine: { lineStyle: { color: 'rgba(134, 144, 156, 0.1)' } },
    axisLabel: { color: '#86909C', fontSize: 10 },
  },
  yAxis: {
    type: 'value',
    splitLine: { lineStyle: { color: 'rgba(134, 144, 156, 0.1)' } },
    axisLabel: { color: '#86909C', fontSize: 10 },
  },
  series: [
    {
      name: 'CPU 使用率',
      type: 'line',
      smooth: true,
      data: [32, 28, 45, 62, 58, 75, 42],
      itemStyle: { color: '#165DFF' },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            { offset: 0, color: 'rgba(22, 93, 255, 0.2)' },
            { offset: 1, color: 'rgba(22, 93, 255, 0)' },
          ],
        },
      },
    },
  ],
}));

// Recent Activities
const activities = ref([
  { id: 1, user: 'Admin', action: '创建了比赛', target: '春季新秀赛 2026', time: '10 分钟前' },
  { id: 2, user: 'System', action: '自动扩容', target: 'Pwn 容器集群', time: '25 分钟前' },
  { id: 3, user: 'Rikka', action: '提交了题解', target: 'Easy_Heap_Overflow', time: '1 小时前' },
  { id: 4, user: 'Admin', action: '更新了公告', target: '系统维护通知', time: '2 小时前' },
]);
</script>

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-black text-slate-800 dark:text-slate-100 tracking-tight">
          管理概览
        </h2>
        <p class="text-sm text-slate-500 font-medium mt-1">
          欢迎回来，系统管理员。这是今日的运行概况。
        </p>
      </div>
      <div class="flex items-center gap-2">
        <span class="flex items-center gap-1.5 px-3 py-1.5 bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-bold">
          <div class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse" />
          系统运行正常
        </span>
      </div>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <div
        v-for="stat in stats"
        :key="stat.label"
        class="bg-white dark:bg-slate-900 p-5 rounded-2xl border border-slate-100 dark:border-slate-800 shadow-sm group hover:border-blue-500/30 transition-all"
      >
        <div class="flex items-start justify-between mb-4">
          <div :class="['p-2.5 rounded-xl transition-transform group-hover:scale-110', stat.bg]">
            <component :is="stat.icon" :size="20" :class="stat.color" />
          </div>
          <button class="text-slate-300 hover:text-slate-600 dark:hover:text-slate-400">
            <MoreVertical :size="16" />
          </button>
        </div>
        <div class="space-y-1">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-wider">
            {{ stat.label }}
          </p>
          <div class="flex items-end justify-between">
            <h3 class="text-2xl font-black text-slate-800 dark:text-slate-100">
              {{ stat.value }}
            </h3>
            <span
              :class="[
                'flex items-center text-xs font-bold px-1.5 py-0.5 rounded-md',
                stat.trendUp ? 'text-emerald-600 bg-emerald-50 dark:bg-emerald-900/20' : 'text-rose-600 bg-rose-50 dark:bg-rose-900/20'
              ]"
            >
              <component :is="stat.trendUp ? ArrowUpRight : ArrowDownRight" :size="12" class="mr-1" />
              {{ stat.trend }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts and Tables -->
    <div class="grid grid-cols-1 xl:grid-cols-3 gap-6">
      <!-- Main Chart (2/3) -->
      <div class="xl:col-span-2 bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 p-6 shadow-sm">
        <div class="flex items-center justify-between mb-6">
          <div>
            <h3 class="text-base font-bold text-slate-800 dark:text-slate-100">
              服务器负载趋势
            </h3>
            <p class="text-xs text-slate-400 mt-1">
              实时监测 CPU 及内存占用情况
            </p>
          </div>
          <div class="flex items-center gap-2">
            <button class="px-3 py-1.5 text-xs font-bold bg-slate-50 dark:bg-slate-800 text-slate-600 dark:text-slate-300 rounded-lg hover:bg-slate-100 transition-colors">
              导出报告
            </button>
          </div>
        </div>
        <div class="h-72 w-full">
          <v-chart :option="chartOption" autoresize />
        </div>
      </div>

      <!-- Side List (1/3) -->
      <div class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-100 dark:border-slate-800 p-6 shadow-sm">
        <div class="flex items-center justify-between mb-6">
          <h3 class="text-base font-bold text-slate-800 dark:text-slate-100">
            系统最近动态
          </h3>
          <TrendingUp :size="16" class="text-blue-600" />
        </div>
        <div class="space-y-6">
          <div v-for="item in activities" :key="item.id" class="flex gap-4 group">
            <div class="relative">
              <div class="w-8 h-8 rounded-full bg-slate-50 dark:bg-slate-800 flex items-center justify-center border border-slate-100 dark:border-slate-700">
                <Clock :size="14" class="text-slate-400" />
              </div>
              <div v-if="item.id !== activities.length" class="absolute top-8 left-4 bottom-[-24px] w-px bg-slate-100 dark:bg-slate-800" />
            </div>
            <div class="flex-1 pb-2">
              <div class="flex items-center justify-between mb-1">
                <span class="text-xs font-black text-slate-800 dark:text-slate-100">{{ item.user }}</span>
                <span class="text-[10px] font-medium text-slate-400">{{ item.time }}</span>
              </div>
              <p class="text-xs text-slate-500 leading-relaxed">
                {{ item.action }} <span class="text-blue-600 dark:text-blue-400 font-bold">"{{ item.target }}"</span>
              </p>
            </div>
          </div>
        </div>
        <button class="w-full mt-6 py-2.5 rounded-xl border border-slate-100 dark:border-slate-800 text-xs font-bold text-slate-400 hover:bg-slate-50 dark:hover:bg-slate-800 transition-all">
          查看完整日志
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.chart {
  width: 100%;
  height: 100%;
}
</style>