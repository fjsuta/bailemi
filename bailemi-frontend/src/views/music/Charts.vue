<template>
  <div class="min-h-screen pb-40">
    <!-- 侧边栏 -->
    <Sidebar />

    <div
      :style="sidebarStore.layout === 'left'
        ? { paddingLeft: sidebarStore.width + 'px', paddingTop: '0', transition: 'padding-left 0.15s ease-out, padding-top 0.15s ease-out' }
        : { paddingLeft: '0', paddingTop: sidebarStore.height + 'px', transition: 'padding-left 0.15s ease-out, padding-top 0.15s ease-out' }"
    >
      <div class="container mx-auto px-4 py-6">
        <!-- 页面标题 -->
        <div class="mb-8">
          <h1 class="text-3xl font-bold gradient-text mb-2">排行榜</h1>
          <p class="text-slate-400">发现全站最热门的歌曲</p>
        </div>

        <!-- 使用排行榜组件 -->
        <Chart />
      </div>
    </div>
  </div>
</template>

<script setup>
import { useSidebarStore } from '@/stores/sidebar'
import Sidebar from '@/components/Sidebar.vue'
import Chart from '@/components/music/Chart.vue'

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
