<template>
  <button
    @click="toggleTheme"
    class="fixed top-4 right-4 z-40 glass rounded-full p-3 hover:bg-white/20 dark:hover:bg-slate-700/50 transition-all"
    :title="themeStore.isDark ? '切换到浅色模式' : '切换到深色模式'"
  >
    <Transition name="theme" mode="out-in">
      <svg v-if="themeStore.isDark" key="moon" class="w-6 h-6 text-yellow-300" viewBox="0 0 24 24" fill="currentColor">
        <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"/>
      </svg>
      <svg v-else key="sun" class="w-6 h-6 text-yellow-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <circle cx="12" cy="12" r="5"/>
        <path d="M12 1v2M12 21v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M1 12h2M21 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"/>
      </svg>
    </Transition>
  </button>
</template>

<script setup>
import { useThemeStore } from '@/stores/theme'

const themeStore = useThemeStore()

const toggleTheme = () => {
  themeStore.toggleTheme()
}
</script>

<style scoped>
.theme-enter-active,
.theme-leave-active {
  transition: all 0.3s ease;
}

.theme-enter-from {
  opacity: 0;
  transform: rotate(-90deg) scale(0.5);
}

.theme-leave-to {
  opacity: 0;
  transform: rotate(90deg) scale(0.5);
}
</style>
