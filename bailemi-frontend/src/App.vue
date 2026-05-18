<template>
  <div class="min-h-screen">
    <router-view />
    <MusicPlayer />
    <Notification />
    <Settings ref="settingsRef" />
  </div>
</template>

<script setup>
import MusicPlayer from '@/components/music/MusicPlayer.vue'
import Notification from '@/components/Notification.vue'
import Settings from '@/components/Settings.vue'
import { ref } from 'vue'

const settingsRef = ref(null)

// 暴露给全局，供 Sidebar 调用
window.openSettings = () => {
  if (settingsRef.value && settingsRef.value.showSettings !== undefined) {
    settingsRef.value.showSettings.value = true
  }
}
</script>
