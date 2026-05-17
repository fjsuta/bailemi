<template>
  <div class="min-h-screen">
    <router-view />
    <MusicPlayer />
    <Notification />
    <ThemeToggle />
  </div>
</template>

<script setup>
import MusicPlayer from '@/components/music/MusicPlayer.vue'
import Notification from '@/components/Notification.vue'
import ThemeToggle from '@/components/ThemeToggle.vue'
</script>
