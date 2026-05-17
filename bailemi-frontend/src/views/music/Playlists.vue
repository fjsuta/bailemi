<template>
  <div class="min-h-screen pb-24">
    <!-- 顶部导航栏 -->
    <div class="fixed top-0 left-0 right-0 z-50 px-4 py-4">
      <div class="container mx-auto">
        <div class="flex items-center gap-4 mb-4">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
              </svg>
            </div>
            <div class="flex flex-col">
              <h1 class="text-xl font-bold gradient-text">百米乐</h1>
              <span class="text-xs text-slate-400">发现好音乐，感受美好</span>
            </div>
          </div>
          
          <div class="flex-1 flex items-center justify-center">
            <div class="flex items-center gap-3 bg-white/10 backdrop-blur-md rounded-full px-6 py-3 border border-white/20">
              <router-link 
                v-for="item in navItems" 
                :key="item.path"
                :to="item.path"
                class="px-4 py-2 rounded-full text-sm font-medium text-slate-300 hover:text-white hover:bg-white/10 transition-all"
                :class="{ 'bg-purple-600 text-white': $route.path === item.path }"
              >
                {{ item.name }}
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="pt-28">
      <div class="container mx-auto px-4 py-6">
        <div class="glass-dark rounded-2xl p-8 text-center">
          <h1 class="text-3xl font-bold gradient-text mb-4">歌单</h1>
          <p class="text-slate-400">歌单功能开发中...</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const navItems = [
  { name: '首页', path: '/' },
  { name: '排行榜', path: '/charts' },
  { name: '歌手', path: '/artists' },
  { name: '歌单', path: '/playlists' },
  { name: '关于网站', path: '/about' }
]
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
