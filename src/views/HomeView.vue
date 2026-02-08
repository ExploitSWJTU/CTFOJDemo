<script setup lang="ts">
import Banner from '../components/Banner.vue';
import Challenges from '../components/Challenges.vue';
import Announcements from '../components/Announcements.vue';
import ActivityFeed from '../components/ActivityFeed.vue';
import Sidebar from '../components/Sidebar.vue';
</script>

<template>
  <div class="space-y-8">
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
  </div>
</template>
