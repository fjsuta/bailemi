<template>
  <div class="min-h-screen">
    <router-view />
    <MusicPlayer />
    <Notification />
    <Settings />
  </div>
</template>

<script setup>
import MusicPlayer from '@/components/music/MusicPlayer.vue'
import Notification from '@/components/Notification.vue'
import Settings from '@/components/Settings.vue'
</script>
