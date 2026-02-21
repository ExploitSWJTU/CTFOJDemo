<script setup lang="ts"></script>

<template>
  <div class="group rounded-card relative h-full overflow-hidden bg-slate-800 shadow-lg">
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
          />
          <div
            class="pointer-events-none absolute inset-0 flex flex-col justify-end bg-linear-to-t from-slate-900/80 to-transparent p-10"
          >
            <span class="mb-4 w-fit rounded-full bg-blue-600 px-3 py-1 text-xs font-bold text-white">
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
</template>