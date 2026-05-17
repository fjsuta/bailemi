<template>
  <div class="min-h-screen pb-24">
    <Sidebar />
    <div class="pl-16 md:pl-64 transition-all duration-300">
      <div class="container mx-auto px-4 py-6">
        <div class="glass-dark rounded-2xl p-8 text-center">
          <h1 class="text-3xl font-bold gradient-text mb-4">歌手</h1>
          <p class="text-slate-400">歌手功能开发中...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Sidebar from '@/components/Sidebar.vue'
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
