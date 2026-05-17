<template>
  <button
    @click="toggleTheme"
    class="fixed top-4 right-4 z-40 glass rounded-full p-3 hover:bg-white/20 dark:hover:bg-slate-700/50 transition-all"
    :title="themeStore.isDark ? '切换到浅色模式' : '切换到深色模式'"
  >
    <Transition name="theme" mode="out-in">
      <span v-if="themeStore.isDark" key="moon" class="text-2xl">🌙</span>
      <span v-else key="sun" class="text-2xl">☀️</span>
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
