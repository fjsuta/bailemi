<template>
  <div class="min-h-screen pb-24">
    <!-- 侧边栏 -->
    <Sidebar />
    
    <div :style="{ paddingLeft: sidebarStore.width + 'px', transition: 'padding-left 0.15s ease-out' }">
      <div class="container mx-auto px-4 py-6">
        <div class="bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-8 text-center">
          <h1 class="text-3xl font-bold gradient-text mb-4">歌手</h1>
          <p class="text-slate-400">歌手功能开发中...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Sidebar from '@/components/Sidebar.vue'
import { useSidebarStore } from '@/stores/sidebar'

const sidebarStore = useSidebarStore()
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
