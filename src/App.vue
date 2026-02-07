<script setup lang="ts">
import Header from './components/Header.vue';
import Banner from './components/Banner.vue';
import Challenges from './components/Challenges.vue';
import Announcements from './components/Announcements.vue';
import ActivityFeed from './components/ActivityFeed.vue';
import Sidebar from './components/Sidebar.vue';
</script>

<template>
  <div
    class="min-h-screen bg-slate-50 font-sans text-slate-800 transition-colors duration-300 dark:bg-slate-950 dark:text-slate-200"
  >
    <Header />

    <main class="mx-auto max-w-384 space-y-8 p-6">
      <Banner />

      <section class="grid grid-cols-4 gap-6">
        <div class="col-span-3">
          <Challenges />
        </div>
        <div class="col-span-1">
          <Announcements />
        </div>
      </section>

      <section class="grid grid-cols-4 gap-6 pb-12">
        <ActivityFeed />
        <Sidebar />
      </section>
    </main>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

body {
  font-family: 'Inter', sans-serif;
}

/* 适配 Arco Design 的黑夜模式 */
.dark body {
  --color-bg-1: #0f172a;
  --color-bg-2: #1e293b;
  --color-bg-3: #334155;
  --color-bg-4: #475569;
  --color-bg-5: #64748b;
  --color-bg-white: #0f172a;
  --color-text-1: #f8fafc;
  --color-text-2: #e2e8f0;
  --color-text-3: #cbd5e1;
  --color-text-4: #94a3b8;
  --color-border-1: #1e293b;
  --color-border-2: #334155;
  --color-border-3: #475569;
  --color-border-4: #64748b;
}
</style>
