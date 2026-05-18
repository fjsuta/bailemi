<template>
  <div>
    <Teleport to="body">
      <div v-if="showSettings" class="fixed inset-0 z-50 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="showSettings = false"></div>
        <div class="relative w-full max-w-lg mx-4 bg-white/10 backdrop-blur-xl rounded-2xl border border-white/20 p-6 shadow-2xl">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-bold">设置</h2>
            <button @click="showSettings = false" class="p-2 hover:bg-white/10 rounded-xl transition-all">
              <svg class="w-5 h-5 text-slate-400 hover:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M18 6L6 18"/>
                <path d="m6 6 12 12"/>
              </svg>
            </button>
          </div>

          <div class="space-y-6">
            <div class="bg-white/5 rounded-2xl p-4 border border-white/10">
              <h3 class="font-semibold mb-4 flex items-center gap-2">
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="5"/>
                  <path d="M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/>
                </svg>
                主题模式
              </h3>
              <div class="grid grid-cols-3 gap-3">
                <button
                  @click="themeStore.setThemeMode('light')"
                  :class="['p-3 rounded-xl border-2 transition-all', themeStore.themeMode === 'light' ? 'border-purple-500 bg-purple-500/20' : 'border-white/10 hover:border-white/20']"
                >
                  <svg class="w-6 h-6 mx-auto mb-2 text-yellow-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <circle cx="12" cy="12" r="5"/>
                    <path d="M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/>
                  </svg>
                  <span class="text-sm">浅色</span>
                </button>
                <button
                  @click="themeStore.setThemeMode('dark')"
                  :class="['p-3 rounded-xl border-2 transition-all', themeStore.themeMode === 'dark' ? 'border-purple-500 bg-purple-500/20' : 'border-white/10 hover:border-white/20']"
                >
                  <svg class="w-6 h-6 mx-auto mb-2 text-indigo-300" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
                  </svg>
                  <span class="text-sm">深色</span>
                </button>
                <button
                  @click="themeStore.setThemeMode('system')"
                  :class="['p-3 rounded-xl border-2 transition-all', themeStore.themeMode === 'system' ? 'border-purple-500 bg-purple-500/20' : 'border-white/10 hover:border-white/20']"
                >
                  <svg class="w-6 h-6 mx-auto mb-2 text-slate-300" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                  </svg>
                  <span class="text-sm">跟随系统</span>
                </button>
              </div>
            </div>

            <div class="bg-white/5 rounded-2xl p-4 border border-white/10">
              <h3 class="font-semibold mb-4 flex items-center gap-2">
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M4 6a2 2 0 012-2h12a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM4 12a2 2 0 012-2h6a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM4 20a2 2 0 012-2h6a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2z"/>
                </svg>
                界面排版
              </h3>
              <div class="grid grid-cols-2 gap-3">
                <button
                  @click="sidebarStore.setLayout('left')"
                  :class="['p-4 rounded-xl border-2 transition-all flex flex-col items-center gap-2', sidebarStore.layout === 'left' ? 'border-purple-500 bg-purple-500/20' : 'border-white/10 hover:border-white/20']"
                >
                  <svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <rect x="3" y="3" width="4" height="18" rx="1"/>
                    <rect x="9" y="3" width="12" height="18" rx="1" opacity="0.5"/>
                  </svg>
                  <span class="text-sm">侧边栏居左</span>
                </button>
                <button
                  @click="sidebarStore.setLayout('top')"
                  :class="['p-4 rounded-xl border-2 transition-all flex flex-col items-center gap-2', sidebarStore.layout === 'top' ? 'border-purple-500 bg-purple-500/20' : 'border-white/10 hover:border-white/20']"
                >
                  <svg class="w-8 h-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <rect x="3" y="3" width="18" height="4" rx="1"/>
                    <rect x="3" y="9" width="18" height="12" rx="1" opacity="0.5"/>
                  </svg>
                  <span class="text-sm">导航栏置顶</span>
                </button>
              </div>
            </div>

            <div class="bg-white/5 rounded-2xl p-4 border border-white/10">
              <h3 class="font-semibold mb-4 flex items-center gap-2">
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21 15a2 2 0 01-2 2H7l-4 4V5a2 2 0 012-2h14a2 2 0 012 2z"/>
                </svg>
                背景主题
              </h3>
              <div class="grid grid-cols-5 gap-3">
                <button
                  v-for="bg in themeStore.defaultBackgrounds"
                  :key="bg.id"
                  @click="themeStore.setBackground(bg.id)"
                  :style="{ background: bg.gradient }"
                  :class="['w-12 h-12 rounded-xl border-2 transition-all hover:scale-110', themeStore.backgroundImage === bg.id || (!themeStore.backgroundImage && bg.id === 'default') ? 'border-white scale-110' : 'border-white/30']"
                  :title="bg.name"
                >
                </button>
              </div>
              <p class="text-xs text-slate-400 mt-3 text-center">点击选择喜欢的背景主题</p>
            </div>
          </div>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useThemeStore } from '@/stores/theme'
import { useSidebarStore } from '@/stores/sidebar'

const themeStore = useThemeStore()
const sidebarStore = useSidebarStore()
const showSettings = ref(false)

defineExpose({
  showSettings
})
</script>
