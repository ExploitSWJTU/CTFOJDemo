<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import Header from './components/Header.vue';
import AdminHeader from './components/AdminHeader.vue';

const route = useRoute();
const isAdmin = computed(() => route.path.startsWith('/admin'));
</script>

<template>
  <div
    class="min-h-screen bg-slate-50 font-sans text-slate-800 transition-colors duration-300 dark:bg-slate-950 dark:text-slate-200"
  >
    <AdminHeader v-if="isAdmin" />
    <Header v-else />

    <main class="mx-auto max-w-384 p-6">
      <router-view />
    </main>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

body {
  font-family: 'Inter', sans-serif;
}

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
  --color-bg-popup: transparent;
}
</style>
