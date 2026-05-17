<template>
  <div class="min-h-screen pb-24">
    <!-- 侧边栏 -->
    <Sidebar />
    
    <div :style="{ paddingLeft: sidebarStore.width + 'px', transition: 'padding-left 0.15s ease-out' }">
      <div class="container mx-auto px-4 py-6">
        <div class="bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-8">
          <h1 class="text-3xl font-bold gradient-text mb-6 text-center">关于网站</h1>
          <div class="max-w-2xl mx-auto space-y-4">
            <p class="text-slate-300">百米乐是一个专注于音乐分享和发现的平台。</p>
            <p class="text-slate-300">在这里，您可以：</p>
            <ul class="list-disc list-inside text-slate-400 space-y-2">
              <li>发现好听的音乐</li>
              <li>上传和分享您的原创作品</li>
              <li>创建和管理您的歌单</li>
              <li>了解歌手和专辑信息</li>
            </ul>
            <div class="border-t border-white/10 pt-6 mt-6">
              <p class="text-slate-500 text-sm text-center">© 2024 百米乐. 保留所有权利.</p>
            </div>
          </div>
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
