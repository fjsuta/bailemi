<template>
  <!-- 左侧布局 -->
  <div 
    v-if="sidebarStore.layout === 'left'"
    class="fixed top-0 left-0 h-full z-40 flex"
    :style="{ width: sidebarStore.width + 'px', transition: 'width 0.15s ease-out' }"
  >
    <!-- 侧边栏主体 -->
    <div class="h-full flex flex-col bg-white/10 backdrop-blur-md border-r border-white/20">
      <!-- 折叠按钮 -->
      <button 
        @click="sidebarStore.toggle"
        class="absolute right-0 top-24 translate-x-1/2 w-8 h-8 bg-white/10 backdrop-blur-md border border-white/20 rounded-full flex items-center justify-center hover:bg-white/20 transition-all z-50"
      >
        <svg 
          class="w-5 h-5 text-slate-300"
          :style="{ transform: sidebarStore.isOpen ? 'rotate(180deg)' : 'rotate(0deg)', transition: 'transform 0.15s ease-out' }"
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
          <div 
            v-show="sidebarStore.isOpen" 
            class="flex-1 min-w-0"
          >
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
          class="flex items-center gap-3 px-4 py-3 rounded-2xl bg-white/5 hover:bg-white/15 transition-all"
          :class="{ 'bg-white/20': $route.path === item.path }"
        >
          <span v-html="item.icon" class="w-6 h-6 flex-shrink-0 text-slate-300"></span>
          <span 
            v-show="sidebarStore.isOpen" 
            class="text-sm text-slate-300 font-medium"
          >
            {{ item.name }}
          </span>
        </router-link>
      </nav>

      <!-- 底部区域 -->
      <div class="p-3 border-t border-white/10">
        <!-- 用户信息 -->
        <div class="flex items-center gap-3 px-4 py-3">
          <div class="w-10 h-10 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center flex-shrink-0 overflow-hidden">
            <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 7a4 4 0 11-8 0 4 4 0 018 0z"/>
              <path d="M12 14c3.866 0 7-1.892 7-4.231C19 7.423 15.866 5.5 12 5.5S5 7.423 5 9.769C5 12.108 8.134 14 12 14z"/>
            </svg>
          </div>
          <div 
            v-show="sidebarStore.isOpen" 
            class="flex-1 min-w-0"
          >
            <p class="text-sm font-medium text-white truncate">访客</p>
            <p class="text-xs text-slate-400 truncate">未登录</p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 顶部布局 -->
  <div 
    v-else
    class="fixed top-0 left-0 right-0 z-40"
    :style="{ height: sidebarStore.height + 'px', transition: 'height 0.15s ease-out' }"
  >
    <!-- 导航栏主体 -->
    <div class="h-full bg-white/10 backdrop-blur-md border-b border-white/20 flex items-center px-4">
      <!-- Logo 区域 -->
      <div class="flex items-center gap-3 mr-8">
        <div class="w-10 h-10 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center flex-shrink-0">
          <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="currentColor">
            <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
          </svg>
        </div>
        <div v-show="sidebarStore.isOpen" class="min-w-0">
          <h1 class="text-lg font-bold gradient-text">百米乐</h1>
        </div>
      </div>

      <!-- 导航项 - 可横向滚动 -->
      <nav class="flex-1 flex items-center gap-2 overflow-x-auto py-2 scrollbar-hide">
        <router-link 
          v-for="item in navItems" 
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-2 px-4 py-2 rounded-2xl bg-white/5 hover:bg-white/15 transition-all flex-shrink-0"
          :class="{ 'bg-white/20': $route.path === item.path }"
        >
          <span v-html="item.icon" class="w-5 h-5 flex-shrink-0 text-slate-300"></span>
          <span 
            v-show="sidebarStore.isOpen" 
            class="text-sm text-slate-300 font-medium whitespace-nowrap"
          >
            {{ item.name }}
          </span>
        </router-link>
      </nav>

      <!-- 右侧按钮区域 -->
      <div class="flex items-center gap-2 ml-4">
        <!-- 折叠按钮 -->
        <button 
          @click="sidebarStore.toggle"
          class="p-3 rounded-2xl bg-white/5 hover:bg-white/15 transition-all"
          title="收起/展开"
        >
          <svg 
            class="w-6 h-6 text-slate-300"
            :style="{ transform: sidebarStore.isOpen ? 'rotate(90deg)' : 'rotate(0deg)', transition: 'transform 0.15s ease-out' }"
            viewBox="0 0 24 24" 
            fill="none" 
            stroke="currentColor" 
            stroke-width="2"
          >
            <path d="M15 18l-6-6 6-6"/>
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useSidebarStore } from '@/stores/sidebar'

const route = useRoute()
const sidebarStore = useSidebarStore()

const navItems = computed(() => [
  { 
    name: '首页', 
    path: '/', 
    icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>' 
  },
  { 
    name: '排行榜', 
    path: '/charts', 
    icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"/><line x1="12" y1="20" x2="12" y2="4"/><line x1="6" y1="20" x2="6" y2="14"/></svg>' 
  },
  { 
    name: '歌手', 
    path: '/artists', 
    icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>' 
  },
  { 
    name: '歌单', 
    path: '/playlists', 
    icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/></svg>' 
  },
  { 
    name: '关于网站', 
    path: '/about', 
    icon: '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/></svg>' 
  }
])
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.scrollbar-hide::-webkit-scrollbar {
  display: none;
}

.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>
