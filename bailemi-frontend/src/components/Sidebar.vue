<template>
  <div class="fixed top-0 left-0 h-full z-40 transition-all duration-300 ease-in-out flex">
    <!-- 侧边栏主体 -->
    <div 
      class="h-full flex flex-col bg-white/10 backdrop-blur-md border-r border-white/20"
      :class="sidebarOpen ? 'w-72' : 'w-20'"
    >
      <!-- 折叠按钮 -->
      <button 
        @click="sidebarOpen = !sidebarOpen"
        class="absolute right-0 top-24 translate-x-1/2 w-8 h-8 bg-white/10 backdrop-blur-md border border-white/20 rounded-full flex items-center justify-center hover:bg-white/20 transition-all z-50"
      >
        <svg 
          class="w-5 h-5 text-slate-300 transition-transform"
          :class="sidebarOpen ? 'rotate-180' : ''"
          viewBox="0 0 24 24" 
          fill="none" 
          stroke="currentColor" 
          stroke-width="2"
        >
          <path d="M15 18l-6-6 6-6"/>
        </svg>
      </button>

      <!-- Logo 区域 -->
      <div class="p-4 border-b border-white/10">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center flex-shrink-0">
            <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
            </svg>
          </div>
          <div v-if="sidebarOpen" class="flex-1 min-w-0">
            <h1 class="text-lg font-bold gradient-text truncate">百米乐</h1>
            <p class="text-xs text-slate-400 truncate">发现好音乐，感受美好</p>
          </div>
        </div>
      </div>

      <!-- 导航项 -->
      <nav class="flex-1 p-3 space-y-2 overflow-y-auto">
        <router-link 
          v-for="item in navItems" 
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-4 py-3 rounded-2xl bg-white/5 hover:bg-white/15 transition-all group"
          :class="{ 'bg-white/20': $route.path === item.path }"
        >
          <span v-html="item.icon" class="w-6 h-6 flex-shrink-0 text-slate-300 group-hover:text-white"></span>
          <span v-if="sidebarOpen" class="text-sm text-slate-300 group-hover:text-white font-medium">
            {{ item.name }}
          </span>
        </router-link>
      </nav>

      <!-- 底部区域 -->
      <div class="p-3 border-t border-white/10">
        <div class="flex items-center gap-3 px-4 py-3">
          <div class="w-10 h-10 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center flex-shrink-0 overflow-hidden">
            <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 7a4 4 0 11-8 0 4 4 0 018 0z"/>
              <path d="M12 14c3.866 0 7-1.892 7-4.231C19 7.423 15.866 5.5 12 5.5S5 7.423 5 9.769C5 12.108 8.134 14 12 14z"/>
            </svg>
          </div>
          <div v-if="sidebarOpen" class="flex-1 min-w-0">
            <p class="text-sm font-medium text-white truncate">访客</p>
            <p class="text-xs text-slate-400 truncate">未登录</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const sidebarOpen = ref(true)

const navItems = computed(() => [
  { name: '首页', path: '/', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M3 9l9-7 9 7"/><path d="M9 22V9"/><path d="M15 22V9"/></svg>' },
  { name: '排行榜', path: '/charts', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/></svg>' },
  { name: '歌手', path: '/artists', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M16 7a4 4 0 11-8 0 4 4 0 018 0z"/><path d="M12 14c3.866 0 7-1.892 7-4.231C19 7.423 15.866 5.5 12 5.5S5 7.423 5 9.769C5 12.108 8.134 14 12 14z"/></svg>' },
  { name: '歌单', path: '/playlists', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4"/><polyline points="7 16 12 21 17 16"/><line x1="12" y1="3" x2="12" y2="21"/></svg>' },
  { name: '关于网站', path: '/about', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/></svg>' }
])
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
