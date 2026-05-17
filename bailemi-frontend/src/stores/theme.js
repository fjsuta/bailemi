import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  const isDark = ref(localStorage.getItem('theme') === 'dark')

  const initTheme = () => {
    const savedTheme = localStorage.getItem('theme')
    if (savedTheme === 'dark') {
      isDark.value = true
      document.body.classList.add('dark')
    } else if (savedTheme === 'light') {
      isDark.value = false
      document.body.classList.remove('dark')
    } else {
      // 检测系统偏好
      if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        isDark.value = true
        document.body.classList.add('dark')
      } else {
        isDark.value = false
        document.body.classList.remove('dark')
      }
    }
  }

  const toggleTheme = () => {
    isDark.value = !isDark.value
    if (isDark.value) {
      document.body.classList.add('dark')
      localStorage.setItem('theme', 'dark')
    } else {
      document.body.classList.remove('dark')
      localStorage.setItem('theme', 'light')
    }
  }

  const setDark = (dark) => {
    isDark.value = dark
    if (dark) {
      document.body.classList.add('dark')
      localStorage.setItem('theme', 'dark')
    } else {
      document.body.classList.remove('dark')
      localStorage.setItem('theme', 'light')
    }
  }

  return {
    isDark,
    initTheme,
    toggleTheme,
    setDark
  }
})
