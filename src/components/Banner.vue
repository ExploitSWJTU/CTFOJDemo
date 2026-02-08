<script setup lang="ts">
import { ref } from 'vue';
import { CalendarCheck2, Sparkles } from 'lucide-vue-next';

const isCheckedIn = ref(false);

const handleCheckIn = () => {
  isCheckedIn.value = true;
};
</script>

<template>
  <section class="flex h-80 items-stretch gap-6">
    <!-- Banner 区域 (左 2/3) -->
    <div class="group rounded-card relative h-full w-2/3 overflow-hidden bg-slate-800 shadow-lg">
      <a-carousel
        class="rounded-card h-full w-full overflow-hidden"
        :auto-play="true"
        indicator-type="line"
        show-arrow="hover"
      >
        <a-carousel-item v-for="i in 3" :key="i">
          <div class="relative h-full w-full">
            <img
              src="https://images.unsplash.com/photo-1550751827-4bd374c3f58b?auto=format&fit=crop&q=80&w=1200"
              class="h-full w-full object-cover opacity-60 transition-transform duration-700 hover:scale-105"
            >
            <div
              class="pointer-events-none absolute inset-0 flex flex-col justify-end bg-linear-to-t from-slate-900/80 to-transparent p-10"
            >
              <span
                class="mb-4 w-fit rounded-full bg-blue-600 px-3 py-1 text-xs font-bold text-white"
              >
                正在进行</span>
              <h2 class="mb-2 text-4xl font-black text-white">
                SWJTU Rookie CTF Game 2026
              </h2>
              <p class="max-w-lg text-slate-200">
                本年CTF新秀杯已经开启，点击进入赛事详情页面报名参赛。
              </p>
            </div>
          </div>
        </a-carousel-item>
      </a-carousel>
    </div>

    <!-- 打卡 & 运势区域 (右 1/3) -->
    <div class="perspective-1000 relative flex w-1/3 flex-col">
      <Transition name="flip" mode="out-in">
        <!-- 未打卡状态 -->
        <div
          v-if="!isCheckedIn"
          class="rounded-card flex h-full w-full flex-col items-center justify-center border border-slate-200 bg-white p-8 shadow-sm transition-all duration-300 dark:border-slate-800 dark:bg-slate-900"
        >
          <div
            class="mb-6 flex h-20 w-20 items-center justify-center rounded-full bg-blue-50 dark:bg-blue-900/20"
          >
            <CalendarCheck2 :size="40" class="text-blue-600 dark:text-blue-400" />
          </div>
          <h3 class="mb-2 text-xl font-bold dark:text-slate-100">
            每日签到
          </h3>
          <p class="mb-8 text-center text-sm text-slate-400 dark:text-slate-500">
            @雷神托尔 jrys
          </p>
          <button
            class="rounded-inner w-full bg-blue-600 py-4 text-lg font-bold text-white transition-all hover:bg-blue-700 hover:shadow-lg active:scale-95"
            @click="handleCheckIn"
          >
            立即打卡
          </button>
        </div>

        <!-- 已打卡状态 -->
        <div
          v-else
          class="rounded-card relative flex h-full w-full flex-col overflow-hidden border border-blue-100 bg-linear-to-br from-white to-blue-50 p-6 shadow-md transition-all duration-300 dark:border-slate-800 dark:from-slate-900 dark:to-slate-950"
        >
          <!-- 背景装饰 -->
          <div class="absolute top-0 right-0 p-4 opacity-10 dark:opacity-5">
            <Sparkles :size="96" class="text-blue-600 dark:text-blue-400" />
          </div>

          <!-- 头部运势信息 -->
          <div class="mb-6 flex items-start justify-between">
            <div>
              <p
                class="text-xs font-bold tracking-widest text-slate-400 uppercase dark:text-slate-500"
              >
                今日运势
              </p>
              <h3 class="mt-1 text-3xl font-black text-blue-600 dark:text-blue-400">
                中吉
              </h3>
            </div>
            <div class="text-right">
              <p
                class="text-xs font-bold tracking-widest text-slate-400 uppercase dark:text-slate-500"
              >
                Streak
              </p>
              <p class="font-mono text-2xl font-bold text-slate-700 dark:text-slate-300">
                7 Days
              </p>
            </div>
          </div>

          <!-- 宜/忌 列表 -->
          <div class="space-y-3">
            <div class="flex items-center gap-3">
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-emerald-100 text-xs font-bold text-emerald-600 dark:bg-emerald-900/30 dark:text-emerald-400"
              >
                宜
              </div>
              <div>
                <p class="text-xs font-bold text-slate-800 dark:text-slate-200">
                  研究 SQL 注入
                </p>
                <p class="text-[10px] leading-tight text-slate-400 dark:text-slate-500">
                  万物皆可 Union，今天灵感爆棚
                </p>
              </div>
            </div>
            <div class="flex items-center gap-3">
              <div
                class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-red-100 text-xs font-bold text-red-600 dark:bg-red-900/30 dark:text-red-400"
              >
                忌
              </div>
              <div>
                <p class="text-xs font-bold text-slate-800 dark:text-slate-200">
                  强制删除容器
                </p>
                <p class="text-[10px] leading-tight text-slate-400 dark:text-slate-500">
                  小心没做完题数据就丢了
                </p>
              </div>
            </div>
          </div>

          <!-- 每日一题卡片 -->
          <div
            class="rounded-inner mt-4 border border-blue-100/50 bg-white/60 p-3 shadow-sm dark:border-slate-800 dark:bg-slate-800/40"
          >
            <div class="mb-1.5 flex items-center justify-between">
              <div class="flex items-center gap-1.5">
                <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-blue-500" />
                <span
                  class="text-[10px] font-bold tracking-wider text-blue-600 uppercase dark:text-blue-400"
                >每日一题</span>
              </div>
              <span class="font-mono text-[10px] text-slate-400 dark:text-slate-500">#PWN-082</span>
            </div>
            <div class="flex items-center justify-between">
              <span class="mr-2 truncate text-sm font-bold text-slate-700 dark:text-slate-200">Easy_Heap_OverFlow</span>
              <button
                class="rounded-button bg-blue-600 px-3 py-1 text-[10px] font-bold text-white shadow-md shadow-blue-200 transition-colors hover:bg-blue-700"
              >
                去挑战
              </button>
            </div>
          </div>

          <!-- 底部格言 -->
          <div class="mt-auto border-t border-blue-50 pt-3 text-center dark:border-slate-800/50">
            <p class="text-[10px] text-blue-400/80 italic dark:text-blue-500/80">
              “ 真正的黑客，在命令行里寻找诗意。”
            </p>
          </div>
        </div>
      </Transition>
    </div>
  </section>
</template>

<style scoped>
.flip-enter-active,
.flip-leave-active {
  transition: all 0.5s ease;
}

.flip-enter-from,
.flip-leave-to {
  transform: rotateY(90deg);
  opacity: 0;
}
</style>
