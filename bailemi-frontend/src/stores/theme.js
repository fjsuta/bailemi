import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  const themeMode = ref(localStorage.getItem('themeMode') || 'system') // 'system' | 'light' | 'dark'
  const backgroundImage = ref(localStorage.getItem('backgroundImage') || '')
  const isDark = ref(false)

  const defaultBackgrounds = [
    {
      id: 'default',
      name: '默认',
      url: '',
      gradient: 'linear-gradient(135deg, #1e1b4b 0%, #0f172a 50%, #1e1b4b 100%)'
    },
    {
      id: 'purple',
      name: '梦幻紫',
      url: '',
      gradient: 'linear-gradient(135deg, #312e81 0%, #4c1d95 50%, #581c87 100%)'
    },
    {
      id: 'ocean',
      name: '深海蓝',
      url: '',
      gradient: 'linear-gradient(135deg, #0c4a6e 0%, #075985 50%, #0369a1 100%)'
    },
    {
      id: 'forest',
      name: '森林绿',
      url: '',
      gradient: 'linear-gradient(135deg, #14532d 0%, #166534 50%, #15803d 100%)'
    },
    {
      id: 'sunset',
      name: '日落橙',
      url: '',
      gradient: 'linear-gradient(135deg, #7c2d12 0%, #9a3412 50%, #c2410c 100%)'
    }
  ]

  const applyTheme = () => {
    if (themeMode.value === 'system') {
      isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    } else {
      isDark.value = themeMode.value === 'dark'
    }
    
    if (isDark.value) {
      document.body.classList.add('dark')
    } else {
      document.body.classList.remove('dark')
    }
  }

  const initTheme = () => {
    applyTheme()
    // 监听系统主题变化
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
      if (themeMode.value === 'system') {
        applyTheme()
      }
    })
  }

  const setThemeMode = (mode) => {
    themeMode.value = mode
    localStorage.setItem('themeMode', mode)
    applyTheme()
  }

  const setBackground = (bg) => {
    backgroundImage.value = bg
    localStorage.setItem('backgroundImage', bg)
    applyBackground()
  }

  const applyBackground = () => {
    const selectedBg = defaultBackgrounds.find(bg => bg.id === backgroundImage.value)
    if (selectedBg) {
      if (selectedBg.url) {
        document.body.style.backgroundImage = `url(${selectedBg.url})`
        document.body.style.backgroundSize = 'cover'
        document.body.style.backgroundPosition = 'center'
        document.body.style.backgroundAttachment = 'fixed'
      } else {
        document.body.style.backgroundImage = selectedBg.gradient
        document.body.style.backgroundSize = 'auto'
        document.body.style.backgroundPosition = 'top left'
        document.body.style.backgroundAttachment = 'fixed'
      }
    } else {
      document.body.style.backgroundImage = defaultBackgrounds[0].gradient
      document.body.style.backgroundSize = 'auto'
      document.body.style.backgroundPosition = 'top left'
      document.body.style.backgroundAttachment = 'fixed'
    }
  }

  return {
    isDark,
    themeMode,
    backgroundImage,
    defaultBackgrounds,
    initTheme,
    setThemeMode,
    setBackground,
    applyBackground
  }
})
