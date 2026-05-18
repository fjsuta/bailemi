import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useSidebarStore = defineStore('sidebar', () => {
  const isOpen = ref(true)
  const layout = ref(localStorage.getItem('sidebarLayout') || 'left') // 'left' or 'top'
  
  const width = computed(() => isOpen.value ? 288 : 80)
  const height = computed(() => isOpen.value ? 120 : 72)
  
  const toggle = () => {
    isOpen.value = !isOpen.value
  }
  
  const open = () => {
    isOpen.value = true
  }
  
  const close = () => {
    isOpen.value = false
  }
  
  const setLayout = (newLayout) => {
    layout.value = newLayout
    localStorage.setItem('sidebarLayout', newLayout)
  }
  
  return {
    isOpen,
    layout,
    width,
    height,
    toggle,
    open,
    close,
    setLayout
  }
})
