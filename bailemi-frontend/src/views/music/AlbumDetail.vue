<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 py-8">
      <div class="flex flex-col md:flex-row gap-6 mb-8">
        <div class="relative w-48 h-48 md:w-64 md:h-64 flex-shrink-0">
          <img :src="album.cover" class="w-full h-full object-cover rounded-2xl" />
          <div class="absolute inset-0 bg-black/30 rounded-2xl flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
            <button @click="playAlbum" class="w-16 h-16 bg-white/20 backdrop-blur rounded-full flex items-center justify-center text-white text-3xl hover:bg-white/30 transition-colors">
              ▶
            </button>
          </div>
        </div>
        <div class="flex-1">
          <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">{{ album.name }}</h1>
          <p class="text-purple-400 mb-4">{{ album.artist }}</p>
          <div class="flex gap-6 mb-6">
            <span>{{ album.year }}年发行</span>
            <span>{{ album.songs.length }}首歌曲</span>
            <span>{{ album.language }}</span>
          </div>
          <p class="text-slate-300 mb-6">{{ album.description }}</p>
          <div class="flex gap-4">
            <button class="px-6 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:opacity-90">
              ♥ 收藏专辑
            </button>
            <button class="px-6 py-2 bg-white/10 rounded-xl font-medium hover:bg-white/20">
              ➤ 分享
            </button>
          </div>
        </div>
      </div>

      <div class="glass-dark rounded-2xl p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-bold">曲目列表</h2>
          <button @click="playAlbum" class="text-purple-400 hover:text-purple-300">播放全部</button>
        </div>

        <div class="space-y-1">
          <div
            v-for="(song, index) in album.songs"
            :key="song.id"
            @click="playSong(song)"
            :class="[
              'flex items-center gap-4 p-4 rounded-xl cursor-pointer transition-all',
              currentSong?.id === song.id ? 'bg-purple-500/20' : 'hover:bg-white/10'
            ]"
          >
            <div class="w-8 text-center">
              <span v-if="currentSong?.id !== song.id" class="text-slate-400">{{ index + 1 }}</span>
              <span v-else class="text-purple-400 animate-pulse">♪</span>
            </div>
            <img :src="song.cover" class="w-10 h-10 rounded-lg object-cover" />
            <div class="flex-1 min-w-0">
              <h4 class="font-medium truncate">{{ song.title }}</h4>
              <p class="text-sm text-slate-400">{{ song.artist }}</p>
            </div>
            <div class="flex items-center gap-4">
              <span class="text-sm text-slate-500">{{ song.duration }}</span>
              <span v-if="song.isFavorite" class="text-red-500">❤️</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useMusicStore } from '@/stores/music'

const musicStore = useMusicStore()

const currentSong = ref(null)

const album = ref({
  id: 1,
  name: '',
  artist: '',
  cover: 'https://picsum.photos/400/400?random=album1',
  year: '',
  language: '',
  description: '',
  songs: []
})

const playSong = (song) => {
  currentSong.value = song
  musicStore.playSong(song, album.value.songs)
}

const playAlbum = () => {
  if (album.value.songs.length > 0) {
    playSong(album.value.songs[0])
  }
}
</script>
