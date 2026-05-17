import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useSidebarStore = defineStore('sidebar', () => {
  const isOpen = ref(true)
  
  const width = computed(() => isOpen.value ? 288 : 80)
  
  const toggle = () => {
    isOpen.value = !isOpen.value
  }
  
  const open = () => {
    isOpen.value = true
  }
  
  const close = () => {
    isOpen.value = false
  }
  
  return {
    isOpen,
    width,
    toggle,
    open,
    close
  }
})
